package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"text/template"

	"github.com/gosimple/slug"
	"github.com/qedus/osmpbf"
)

type LatLon struct {
	lat float64
	lon float64
}

func createIndexFile(region string, citySlug string, city string, latLon LatLon, template *template.Template) {
	indexFile := region + "/content/cities/" + citySlug + "/_index.md"
	if _, err := os.Stat(indexFile); os.IsNotExist(err) {
		f, err := os.Create(indexFile)
		if err != nil {
			fmt.Println(err)
			return
		}
		data := map[string]interface{}{
			"title":     city,
			"url":       "/" + citySlug + "/",
			"latitude":  latLon.lat,
			"longitude": latLon.lon,
		}
		if err = template.Execute(f, data); err != nil {
			panic(err)
		}
		f.Close()
	}
}

func createElementFile(region string, citySlug string, nameSlug string, name string, shop string, template *template.Template) {
	elementFile := region + "/content/cities/" + citySlug + "/" + nameSlug + ".md"
	if _, err := os.Stat(elementFile); !os.IsNotExist(err) {
		re := regexp.MustCompile(`\d+$`)
		i := 2
		nameSlug = fmt.Sprintf("%s-%d", nameSlug, i)
		for {
			elementFile = region + "/content/cities/" + citySlug + "/" + nameSlug + ".md"
			if _, err = os.Stat(elementFile); os.IsNotExist(err) {
				break
			}
			i++
			nameSlug = re.ReplaceAllString(nameSlug, strconv.Itoa(i))
		}
	}
	f, err := os.Create(elementFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	data := map[string]interface{}{
		"title": strings.Replace(name, "\"", "", -1),
		"url":   "/" + citySlug + "/" + nameSlug + "/",
		"shop":  shop,
	}
	if err = template.Execute(f, data); err != nil {
		panic(err)
	}
	f.Close()
}

func createDataElementFile(region string, citySlug string, nameSlug string, id int64, elementType string, lat float64, lon float64, tags map[string]string, city string, template *template.Template) {
	elementFile := region + "/data/cities/" + citySlug + "/" + nameSlug + ".yml"
	if _, err := os.Stat(elementFile); !os.IsNotExist(err) {
		re := regexp.MustCompile(`\d+$`)
		i := 2
		nameSlug = fmt.Sprintf("%s-%d", nameSlug, i)
		for {
			elementFile = region + "/data/cities/" + citySlug + "/" + nameSlug + ".yml"
			if _, err = os.Stat(elementFile); os.IsNotExist(err) {
				break
			}
			i++
			nameSlug = re.ReplaceAllString(nameSlug, strconv.Itoa(i))
		}
	}
	f, err := os.Create(elementFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	data := map[string]interface{}{
		"id":            id,
		"type":          elementType,
		"latitude":      lat,
		"longitude":     lon,
		"postcode":      tags["addr:postcode"],
		"city":          city,
		"street":        tags["addr:street"],
		"housenumber":   tags["addr:housenumber"],
		"phone":         strings.Replace(tags["phone"], "\"", "", -1),
		"opening_hours": strings.Replace(tags["opening_hours"], "\"", "", -1),
		"website":       strings.Replace(tags["website"], "\"", "", -1),
	}
	if err = template.Execute(f, data); err != nil {
		panic(err)
	}
	f.Close()
}

func translateShop(shop string) string {
	shopSplit := strings.Split(shop, ";")
	var translatedShop string
	switch shopSplit[0] {
	// Lebensmittel, Getränke
	case "alcohol":
		translatedShop = "Spirituosen"
	case "bakery":
		translatedShop = "Bäckerei"
	case "beverages":
		translatedShop = "Getränke"
	case "brewing_supplies", "brewery":
		translatedShop = "Brauerei"
	case "butcher":
		translatedShop = "Metzgerei"
	case "cheese":
		translatedShop = "Käse"
	case "chocolate":
		translatedShop = "Schokolade"
	case "coffee", "coffee_roasting", "coffeemaker":
		translatedShop = "Kaffee"
	case "confectionery", "sweets", "cookies":
		translatedShop = "Süßwaren"
	case "convenience", "food", "grocery":
		translatedShop = "Lebensmittel"
	case "deli":
		translatedShop = "Feinkost"
	case "dairy":
		translatedShop = "Milch"
	case "farm":
		translatedShop = "Hofladen"
	case "frozen_food":
		translatedShop = "Tiefkühl"
	case "greengrocer", "vegetables":
		translatedShop = "Gemüse & Obst"
	case "health_food", "organic":
		translatedShop = "Bioladen"
	case "ice_cream":
		translatedShop = "Eisprodukte"
	case "pasta":
		translatedShop = "Pasta"
	case "pastry", "cake":
		translatedShop = "Konditorei"
	case "seafood", "fishmonger":
		translatedShop = "Fisch"
	case "spices":
		translatedShop = "Gewürze"
	case "tea":
		translatedShop = "Tee"
	case "wine", "winery":
		translatedShop = "Wein"
	case "water":
		translatedShop = "Wasser"
	// Warenhaus, Kaufhaus, Einkaufszentrum
	case "department_store":
		translatedShop = "Warenhaus"
	case "general", "country_store":
		translatedShop = "Dorfladen"
	case "kiosk":
		translatedShop = "Kiosk"
	case "mall":
		translatedShop = "Einkaufszentrum"
	case "supermarket":
		translatedShop = "Supermarkt"
	case "wholesale":
		translatedShop = "Großhandel"
	// Kleidung, Schuhe, Accessoires
	case "baby_goods":
		translatedShop = "Babysachen"
	case "bag", "suitcases":
		translatedShop = "Taschen & Koffer"
	case "boutique", "fashion", "fashion_accessoires":
		translatedShop = "Modehaus"
	case "clothes":
		translatedShop = "Kleidung"
	case "fabric", "fabrics":
		translatedShop = "Textil"
	case "jewelry", "gemstone", "gold", "gold_buyer":
		translatedShop = "Schmuck"
	case "leather":
		translatedShop = "Leder"
	case "sewing", "sewing_machines":
		translatedShop = "Nähzubehör"
	case "shoes", "shoe_repair", "shoe_repairlocksmith":
		translatedShop = "Schuhe"
	case "tailor":
		translatedShop = "Schneiderei"
	case "watches", "clocks", "watchmaker":
		translatedShop = "Uhren"
	case "wool":
		translatedShop = "Wolle"
	// Discounter, Wohltätigkeit
	case "charity", "second_hand":
		translatedShop = "Gebrauchtwaren"
	case "variety_store":
		translatedShop = "Kramladen"
	// Gesundheit und Schönheitspflege
	case "beauty", "cosmetics", "decorative_cosmetics", "nail_salon", "wellness", "cosmetic":
		translatedShop = "Kosmetik"
	case "chemist":
		translatedShop = "Drogerie"
	case "erotic":
		translatedShop = "Erotik"
	case "hairdresser":
		translatedShop = "Friseur"
	case "hairdresser_supply", "combs":
		translatedShop = "Friseurbedarf"
	case "hearing_aids":
		translatedShop = "Hörgeräte"
	case "herbalist":
		translatedShop = "Kräuter"
	case "massage":
		translatedShop = "Massage"
	case "medical_supply", "medical", "orthopedic", "orthopedics", "sanitary":
		translatedShop = "Sanitätshaus"
	case "nutrition_supplements":
		translatedShop = "Nahrungsergänzung"
	case "optician", "eyeglasses":
		translatedShop = "Optiker"
	case "perfumery":
		translatedShop = "Parfümerie"
	case "tattoo":
		translatedShop = "Tattoo"
	// Do-it-yourself, Haushaltswaren, Baustoffe, Gartenprodukte
	case "agrarian", "agricultural":
		translatedShop = "Landwirtschaftlich"
	case "appliance":
		translatedShop = "Haushaltsgeräte"
	case "bathroom_furnishing", "bathroom", "bath":
		translatedShop = "Badezimmer"
	case "doityourself":
		translatedShop = "Baumarkt"
	case "electrical", "electric":
		translatedShop = "Elektrisch"
	case "energy", "battery":
		translatedShop = "Energie"
	case "fireplace", "furnace", "oven":
		translatedShop = "Kamine & Öfen"
	case "florist":
		translatedShop = "Blumen"
	case "garden_centre":
		translatedShop = "Garten-Center"
	case "garden_furniture":
		translatedShop = "Gartenmöbel"
	case "gas":
		translatedShop = "Gasflaschen"
	case "glaziery":
		translatedShop = "Glaserei"
	case "hardware":
		translatedShop = "Eisenwaren"
	case "houseware", "household", "house":
		translatedShop = "Haushaltsartikel"
	case "locksmith", "keys":
		translatedShop = "Schlüsseldienst"
	case "paint", "paintings", "colors", "painter":
		translatedShop = "Farben"
	case "security":
		translatedShop = "Sicherheit"
	case "trade":
		translatedShop = "Baustoffe"
	// Möbel und Innenausstattung
	case "antiques":
		translatedShop = "Antiquitäten"
	case "bed":
		translatedShop = "Betten"
	case "candles":
		translatedShop = "Kerzen"
	case "carpet":
		translatedShop = "Teppiche"
	case "curtain":
		translatedShop = "Gardinen"
	case "doors":
		translatedShop = "Türen"
	case "flooring", "parquet":
		translatedShop = "Fußböden"
	case "furniture":
		translatedShop = "Möbel"
	case "interior_decoration", "interior", "decoration", "wallpaper", "interior_store":
		translatedShop = "Raumausstattung"
	case "kitchen", "kitchen_appliances", "kitchenware", "kitchen_equipment", "cooking", "crockery", "tableware", "ceramics":
		translatedShop = "Küchen"
	case "lamps", "lighting":
		translatedShop = "Lampen"
	case "tiles", "tile", "tiling":
		translatedShop = "Fliesen"
	case "window_blind", "shutter", "shutters":
		translatedShop = "Jalousien"
	// Elektronik
	case "computer":
		translatedShop = "Computer"
	case "robot":
		translatedShop = "Roboter"
	case "electronics", "electronics_repair", "electro":
		translatedShop = "Elektronik"
	case "hifi":
		translatedShop = "Hifi"
	case "mobile_phone", "telephone", "phone", "communication", "telecommunication":
		translatedShop = "Handy"
	case "radiotechnics":
		translatedShop = "Radiotechnik"
	case "vacuum_cleaner":
		translatedShop = "Staubsauger"
	// Outdoor und Sport, Fahrzeuge
	case "atv":
		translatedShop = "Quad"
	case "bicycle", "bike_repair":
		translatedShop = "Fahrrad"
	case "boat", "yachts":
		translatedShop = "Boot"
	case "car":
		translatedShop = "Autohaus"
	case "car_repair":
		translatedShop = "Autowerkstatt"
	case "car_parts":
		translatedShop = "Autoteile"
	case "caravan", "caravaning":
		translatedShop = "Wohnwagen"
	case "fuel":
		translatedShop = "Treibstoff"
	case "fishing", "fishing_gear":
		translatedShop = "Angeln"
	case "free_flying":
		translatedShop = "Freeflying"
	case "golf":
		translatedShop = "Golf"
	case "hunting":
		translatedShop = "Jagd"
	case "jetski":
		translatedShop = "Jetski"
	case "military_surplus", "military":
		translatedShop = "Militär"
	case "motorcycle", "motorcycle_repair":
		translatedShop = "Motorrad"
	case "outdoor":
		translatedShop = "Outdoor"
	case "scuba_diving":
		translatedShop = "Tauchen"
	case "ski":
		translatedShop = "Ski"
	case "snowmobile":
		translatedShop = "Schneemobil"
	case "sports", "water_sports", "hobby":
		translatedShop = "Sport"
	case "swimming_pool":
		translatedShop = "Pool"
	case "trailer", "car_trailer":
		translatedShop = "Anhänger"
	case "tyres":
		translatedShop = "Reifen"
	// Kunst, Musik, Hobbys
	case "art", "arts", "artwork":
		translatedShop = "Kunst"
	case "collector", "coins", "comics", "stamps":
		translatedShop = "Sammler"
	case "craft":
		translatedShop = "Basteln"
	case "frame", "picture_frames":
		translatedShop = "Rahmen"
	case "games":
		translatedShop = "Spiele"
	case "model", "modelrailway":
		translatedShop = "Modellbau"
	case "music":
		translatedShop = "Musik"
	case "musical_instrument", "woodwind_repair":
		translatedShop = "Instrumente"
	case "photo", "photo_studio", "photographic_studio", "photographer":
		translatedShop = "Foto"
	case "camera":
		translatedShop = "Kamera"
	case "trophy":
		translatedShop = "Pokal"
	case "video":
		translatedShop = "Videothek"
	case "video_games":
		translatedShop = "Videospiele"
	// Schreibwaren, Geschenke, Bücher und Zeitungen
	case "anime", "japan":
		translatedShop = "Anime"
	case "books", "book_restoration":
		translatedShop = "Bücher"
	case "gift":
		translatedShop = "Andenken"
	case "lottery":
		translatedShop = "Lotterie"
	case "newsagent":
		translatedShop = "Zeitungen"
	case "stationary":
		translatedShop = "Schreibwaren"
	case "ticket", "tickets":
		translatedShop = "Tickets"
	// Andere
	case "bookmaker":
		translatedShop = "Wettbüro"
	case "cannabis", "growshop":
		translatedShop = "Hanf"
	case "copyshop", "printing", "print_shop", "printer_ink", "ink_cartridges", "printer", "printers", "paper", "printshop", "printery":
		translatedShop = "Kopieren"
	case "e-cigarette", "vape":
		translatedShop = "E-Zigaretten"
	case "funeral_directors":
		translatedShop = "Bestattungen"
	case "laundry", "rotary_iron", "ironing", "dry_cleaning":
		translatedShop = "Wäscherei"
	case "party":
		translatedShop = "Partyzubehör"
	case "pawnbroker", "money_lender":
		translatedShop = "Leiher"
	case "pet":
		translatedShop = "Tiere"
	case "pet_grooming", "dog_beauty", "dog_hairdresser":
		translatedShop = "Tiersalon"
	case "pest_control":
		translatedShop = "Schädlingsbekämpfung"
	case "pyrotechnics":
		translatedShop = "Pyrotechnik"
	case "religion":
		translatedShop = "Religion"
	case "storage_rental", "rental":
		translatedShop = "Mieten"
	case "tobacco", "smokers":
		translatedShop = "Tabak"
	case "toys":
		translatedShop = "Spielzeug"
	case "travel_agency":
		translatedShop = "Reisebüro"
	case "vacant":
		translatedShop = "Leerstehend"
	case "weapons", "guns", "arms", "knives", "knife":
		translatedShop = "Waffen"
	case "outpost":
		translatedShop = "Außenstelle"
	// Benutzerdefiniert
	case "apiary", "beekeeping", "beekeepers_need", "honey", "beekeeper":
		translatedShop = "Imkerei"
	case "auction_house", "auctioneer", "auction":
		translatedShop = "Auktionshaus"
	case "car_accessories", "child_safety_seats":
		translatedShop = "Autozubehör"
	case "car_service":
		translatedShop = "Autoservice"
	case "caretaker", "building_cleaner":
		translatedShop = "Hausmeister"
	case "carpenter", "cabinet_maker", "carpentry":
		translatedShop = "Schreinerei"
	case "casino", "gambling":
		translatedShop = "Spielkasino"
	case "catering", "catering_supplies":
		translatedShop = "Catering"
	case "equestrian", "horse_equipment", "horse":
		translatedShop = "Pferde"
	case "esoteric":
		translatedShop = "Esoterik"
	case "estate_agent":
		translatedShop = "Immobilien"
	case "event_service":
		translatedShop = "Veranstaltungen"
	case "fanshop":
		translatedShop = "Fanshop"
	case "fitness_equipment":
		translatedShop = "Fitness"
	case "flour":
		translatedShop = "Mehl"
	case "glass":
		translatedShop = "Glas"
	case "garden_service":
		translatedShop = "Gartendienst"
	case "garden_machinery", "gardening_tools", "lawn_mower":
		translatedShop = "Gartenmaschinen"
	case "grill", "bbq":
		translatedShop = "Grillen"
	case "grinding":
		translatedShop = "Schleifen"
	case "hat", "hats":
		translatedShop = "Hüte"
	case "health":
		translatedShop = "Gesundheit"
	case "heating_system", "heater", "heating":
		translatedShop = "Heizungsanlagen"
	case "hookah", "shisha":
		translatedShop = "Wasserpfeife"
	case "hypnotism":
		translatedShop = "Hypnose"
	case "internet_service_provider":
		translatedShop = "Internetanbieter"
	case "joiner":
		translatedShop = "Tischlerei"
	case "kids_furnishing":
		translatedShop = "Kinder"
	case "furs":
		translatedShop = "Pelze"
	case "lettering", "license_plate", "license_plates", "number_plate", "sign_make", "signs":
		translatedShop = "Beschriftungen"
	case "locksmithery", "metalworker", "metalwork", "metalworking":
		translatedShop = "Schlosserei"
	case "machinery", "vehicle", "vehicles", "forklift", "agricultural_machinery", "industrial", "machines":
		translatedShop = "Maschinen"
	case "office_supplies", "office", "stationery":
		translatedShop = "Schreibwaren"
	case "pet_food", "fodder":
		translatedShop = "Tierfutter"
	case "plumber", "plumbing_business":
		translatedShop = "Klempner"
	case "piercing":
		translatedShop = "Piercing"
	case "pottery":
		translatedShop = "Töpferei"
	case "ship_chandler":
		translatedShop = "Schiffe"
	case "software":
		translatedShop = "Software"
	case "solarium", "sunstudio":
		translatedShop = "Solarium"
	case "stones", "tombstone", "tombstones", "gravestones":
		translatedShop = "Steine"
	case "tanning":
		translatedShop = "Gerberei"
	case "tools", "screws":
		translatedShop = "Werkzeuge"
	case "wedding_gown", "wedding":
		translatedShop = "Brautkleider"
	case "whirlpool", "pool":
		translatedShop = "Pool"
	case "wood", "timber", "sawmill":
		translatedShop = "Holz"
	case "worldshop", "one_world", "afro", "oneworld":
		translatedShop = "Weltladen"
	default:
		translatedShop = "Allgemein"
		fmt.Println("unknown shop:", shop)
	}
	return translatedShop
}

func main() {
	region := os.Args[1]
	f, err := os.Open(region + "-latest.osm.pbf")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	d := osmpbf.NewDecoder(f)

	// Cache nodes, ways and cities
	m := make(map[int64]LatLon)
	n := make(map[int64]string)
	p := make(map[string]LatLon)

	// Create templates
	indexTmpl := `---
title: {{ .title }}
url: {{ .url }}
latitude: {{ with .latitude }}{{ . }}{{ end }}
longitude: {{ with .longitude }}{{ . }}{{ end }}
---
`
	indexTemplate := template.Must(template.New("index").Parse(indexTmpl))
	mdTmpl := `---
title: "{{ .title }}"
url: {{ .url }}
shop: {{ .shop }}
---
`
	mdTemplate := template.Must(template.New("markdown").Parse(mdTmpl))
	dataTmpl := `id: {{ .id }}
type: {{ .type }}
latitude: {{ .latitude }}
longitude: {{ .longitude }}
postcode: "{{ .postcode }}"
city: {{ .city }}
street: "{{ .street }}"
housenumber: {{ .housenumber }}
phone: "{{ .phone }}"
opening_hours: "{{ .opening_hours }}"
website: "{{ .website }}"
`
	dataTemplate := template.Must(template.New("data").Parse(dataTmpl))

	// use more memory from the start, it is faster
	d.SetBufferSize(osmpbf.MaxBlobSize)

	// start decoding with several goroutines, it is faster
	err = d.Start(runtime.GOMAXPROCS(-1))
	if err != nil {
		log.Fatal(err)
	}
	var nc, wc, rc uint64
	for {
		if v, err := d.Decode(); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		} else {
			switch v := v.(type) {
			case *osmpbf.Node:
				// Process Node v.
				tags := v.Tags
				city := tags["addr:city"]
				name := tags["name"]
				shop := tags["shop"]
				place := tags["place"]
				if place != "" && name != "" {
					// Cache all cities LatLon
					p[name] = LatLon{v.Lat, v.Lon}
				}
				if city != "" && name != "" && shop != "" {
					citySlug := slug.MakeLang(city, "de")
					nameSlug := slug.MakeLang(name, "de")

					// Exceptions: skip foreign cities
					if citySlug == "s-heerenberg" {
						// 's-Heerenberg in the Netherlands
						continue
					}

					// 1. content
					err = os.MkdirAll(region+"/content/cities/"+citySlug, 0755)
					createIndexFile(region, citySlug, city, p[city], indexTemplate)
					createElementFile(region, citySlug, nameSlug, name, translateShop(shop), mdTemplate)

					// 2. data
					err = os.MkdirAll(region+"/data/cities/"+citySlug, 0755)
					createDataElementFile(region, citySlug, nameSlug, v.ID, "node", v.Lat, v.Lon, tags, city, dataTemplate)
				}
				// Cache all Nodes LatLon
				m[v.ID] = LatLon{v.Lat, v.Lon}
				nc++
			case *osmpbf.Way:
				// Process Way v.
				tags := v.Tags
				city := tags["addr:city"]
				name := tags["name"]
				shop := tags["shop"]
				if city != "" && name != "" && shop != "" && n[v.ID] == "" {
					citySlug := slug.MakeLang(city, "de")
					nameSlug := slug.MakeLang(name, "de")

					// 1. content
					err = os.MkdirAll(region+"/content/cities/"+citySlug, 0755)
					createIndexFile(region, citySlug, city, p[city], indexTemplate)
					createElementFile(region, citySlug, nameSlug, name, translateShop(shop), mdTemplate)

					// 2. data
					err = os.MkdirAll(region+"/data/cities/"+citySlug, 0755)
					node := m[v.NodeIDs[0]] // Lookup coords of first childnode
					createDataElementFile(region, citySlug, nameSlug, v.ID, "way", node.lat, node.lon, tags, city, dataTemplate)

					// Ways might be twice
					n[v.ID] = name
				}
				wc++
			case *osmpbf.Relation:
				// Process Relation v.
				rc++
			default:
				log.Fatalf("unknown type %T\n", v)
			}
		}
	}
	fmt.Printf("Nodes: %d, Ways: %d, Relations: %d\n", nc, wc, rc)
}

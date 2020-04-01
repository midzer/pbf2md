package main

import (
	"fmt"
	"os"
	"io"
	"log"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"text/template"
	"github.com/qedus/osmpbf"
	"github.com/gosimple/slug"
)

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
			return false
	}
	return !info.IsDir()
}

func createIndexFile(citySlug string , city string) {
	indexFile := "content/cities/" + citySlug + "/_index.md"
	if !fileExists(indexFile) {
		f, err := os.Create(indexFile)
		if err != nil {
			fmt.Println(err)
			return
		}
		data := map[string]interface{} {
			"title": city,
			"url":  "/" + citySlug + "/",
		}
		const indexTmpl = `---
title: {{ .title }}
url: {{ .url }}
---
`
	  indexTemplate := template.Must(template.New("index").Parse(indexTmpl))
		if err := indexTemplate.Execute(f, data); err != nil {
				panic(err)
		}
		f.Close()
	}
}

func createElementFile(citySlug string, nameSlug string, name string) {
	elementFile := "content/cities/" + citySlug + "/" + nameSlug + ".md"
	if fileExists(elementFile) {
		var re = regexp.MustCompile(`\d+$`)
		i := 2
		nameSlug = fmt.Sprintf("%s-%d", nameSlug, i)
		for {
			elementFile := "content/cities/" + citySlug + "/" + nameSlug + ".md"
			if !fileExists(elementFile) {
				break
			}
			nameSlug = re.ReplaceAllString(nameSlug, strconv.Itoa(i))
			i++
		}
	}
	f, err := os.Create(elementFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	data := map[string]interface{} {
		"title": strings.Replace(name, "\"", "", -1),
		"url":  "/" + citySlug + "/" + nameSlug + "/",
	}
	const mdTmpl = `---
title: "{{ .title }}"
url: {{ .url }}
---
`
	mdTemplate := template.Must(template.New("markdown").Parse(mdTmpl))
	if err := mdTemplate.Execute(f, data); err != nil {
			panic(err)
	}
	f.Close()
}

func createDataElementFile(citySlug string, nameSlug string, id int64, elementType string, shop string, lat float64, lon float64, tags map[string]string, city string) {
	elementFile := "data/cities/" + citySlug + "/" + nameSlug + ".yml"
	if fileExists(elementFile) {
		var re = regexp.MustCompile(`\d+$`)
		i := 2
		nameSlug = fmt.Sprintf("%s-%d", nameSlug, i)
		for {
			elementFile := "data/cities/" + citySlug + "/" + nameSlug + ".yml"
			if !fileExists(elementFile) {
				break
			}
			nameSlug = re.ReplaceAllString(nameSlug, strconv.Itoa(i))
			i++
		}
	}
	f, err := os.Create(elementFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	data := map[string]interface{} {
		"id":            id,
		"type":          elementType,
		"shop":          shop,
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
	const dataTmpl = `id: {{ .id }}
type: {{ .type }}
shop: {{ .shop }}
latitude: {{ .latitude }}
longitude: {{ .longitude }}
postcode: {{ .postcode }}
city: {{ .city }}
street: {{ .street }}
housenumber: {{ .housenumber }}
phone: "{{ .phone }}"
opening_hours: "{{ .opening_hours }}"
website: "{{ .website }}"
`
  dataTemplate := template.Must(template.New("data").Parse(dataTmpl))
	if err := dataTemplate.Execute(f, data); err != nil {
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
		case "brewing_supplies":
			translatedShop = "Brauerei"
		case "butcher":
			translatedShop = "Metzgerei"
		case "cheese":
			translatedShop = "Käse"
		case "chocolate":
			translatedShop = "Schokolade"
		case "coffee", "coffee_roasting", "coffeemaker":
			translatedShop = "Kaffee"
		case "confectionery", "sweets":
			translatedShop = "Süßwaren"
		case "convenience":
			translatedShop = "Tante Emma"
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
		case "pastry":
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
			translatedShop = "Baby"
		case "bag", "suitcases":
			translatedShop = "Taschen & Koffer"
		case "boutique", "fashion":
			translatedShop = "Modehaus"
		case "clothes":
			translatedShop = "Kleidung"
		case "fabric":
			translatedShop = "Textil"
		case "fashion_accessoires":
			translatedShop = "Modeaccessoires"
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
		case "charity":
			translatedShop = "Gebrauchtwaren"
		case "second_hand":
			translatedShop = "Second Hand"
		case "variety_store":
			translatedShop = "Kramladen"
		// Gesundheit und Schönheitspflege
		case "beauty", "cosmetics", "decorative_cosmetics", "nail_salon", "wellness":
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
		case "medical_supply", "medical", "orthopedics":
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
		case "bathroom_furnishing", "bathroom":
			translatedShop = "Badezimmer"
		case "doityourself":
			translatedShop = "Baumarkt"
		case "electrical":
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
		case "paint", "paintings":
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
		case "tiles", "tile":
			translatedShop = "Fliesen"
		case "window_blind", "shutter":
			translatedShop = "Rollladen"
		// Elektronik
		case "computer":
			translatedShop = "Computer"
		case "robot":
			translatedShop = "Roboter"
		case "electronics", "electronics_repair":
			translatedShop = "Elektronik"
		case "hifi":
			translatedShop = "Hifi"
		case "mobile_phone", "telephone", "phone", "communication", "telecommunication":
			translatedShop = "Handy"
		case "radiotechnics":
			translatedShop = "Elektronisches"
		case "vacuum_cleaner":
			translatedShop = "Staubsauger"
		// Outdoor und Sport, Fahrzeuge
		case "atv":
			translatedShop = "Quad"
		case "bicycle", "bike_repair":
			translatedShop = "Fahrrad"
		case "boat":
			translatedShop = "Boot"
		case "car":
			translatedShop = "Autohaus"
		case "car_repair":
			translatedShop = "Autowerkstatt"
		case "car_parts":
			translatedShop = "Autoteile"
		case "caravan":
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
		case "art":
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
		case "photo", "photo_studio", "photographic_studio":
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
		case "ticket":
			translatedShop = "Tickets"
		// Andere
		case "bookmaker":
			translatedShop = "Wettbüro"
		case "cannabis":
			translatedShop = "Hanf"
		case "copyshop", "printing", "print_shop", "printer_ink", "ink_cartridges":
			translatedShop = "Kopieren"
		case "e-cigarette":
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
		case "apiary", "beekeeping", "beekeepers_need":
			translatedShop = "Imkerei"
		case "auction_house", "auctioneer":
			translatedShop = "Auktionshaus"
		case "car_accessories", "child_safety_seats":
			translatedShop = "Autozubehör"
		case "car_service":
			translatedShop = "Autoservice"
		case "caretaker", "building_cleaner":
			translatedShop = "Hausmeister"
		case "carpenter", "cabinet_maker":
			translatedShop = "Schreinerei"
		case "casino", "gambling":
			translatedShop = "Spielkasino"
		case "catering", "catering_supplies":
			translatedShop = "Catering"
		case "equestrian", "horse_equipment":
			translatedShop = "Reitsport"
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
		case "food", "grocery":
			translatedShop = "Lebensmittel"
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
		case "heating_system":
			translatedShop = "Heizungsanlagen"
		case "hookah":
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
		case "lettering", "license_plate", "license_plates", "number_plate", "sign_make":
			translatedShop = "Beschriftungen"
		case "locksmithery", "metalworker":
			translatedShop = "Schlosserei"
		case "machinery", "vehicle", "forklift", "agricultural_machinery", "industrial":
			translatedShop = "Maschinen"
		case "office_supplies", "office":
			translatedShop = "Büromaterial"
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
		case "tanning":
			translatedShop = "Gerberei"
		case "tools":
			translatedShop = "Werkzeuge"
		case "wedding_gown":
			translatedShop = "Brautkleider"
		case "whirlpool":
			translatedShop = "Whirlpool"
		case "wood":
			translatedShop = "Holz"
		case "worldshop", "one_world":
			translatedShop = "Weltladen"
		case "yes":
			translatedShop = "Allgemeine"
		default:
			translatedShop = "Andere"
			fmt.Println("unknown shop:", shop)
	}
	return translatedShop
}

func main() {
  f, err := os.Open("bayern-latest.osm.pbf")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	d := osmpbf.NewDecoder(f)

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
				if city != "" && name != "" && shop != "" {
					citySlug := slug.MakeLang(city, "de")
					nameSlug := slug.MakeLang(name, "de")
					shop = translateShop(shop)

					// 1. content
					err = os.MkdirAll("content/cities/" + citySlug, 0755)
					createIndexFile(citySlug, city)
					createElementFile(citySlug, nameSlug, name)

					// 2. data
					err = os.MkdirAll("data/cities/" + citySlug, 0755)
					createDataElementFile(citySlug, nameSlug, v.ID, "node", shop, v.Lat, v.Lon, tags, city)
				}
				nc++
			case *osmpbf.Way:
				// Process Way v.
				tags := v.Tags
				city := tags["addr:city"]
				name := tags["name"]
				shop := tags["shop"]
				if city != "" && name != "" && shop != "" {
					citySlug := slug.MakeLang(city, "de")
					nameSlug := slug.MakeLang(name, "de")
					shop = translateShop(shop)

					// 1. content
					err = os.MkdirAll("content/cities/" + citySlug, 0755)
					createIndexFile(citySlug, city)
					createElementFile(citySlug, nameSlug, name)

					// 2. data
					err = os.MkdirAll("data/cities/" + citySlug, 0755)
					var lat, lon float64
					for {
						if w, err := d.Decode(); err == io.EOF {
							break
						} else if err != nil {
							log.Fatal(err)
						} else {
							switch w := w.(type) {
							case *osmpbf.Node:
								if v.NodeIDs[0] == w.ID {
									lat = w.Lat
									lon = w.Lon
									break
								}
							}
						}
					}
					createDataElementFile(citySlug, nameSlug, v.ID, "way", shop, lat, lon, tags, city)
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

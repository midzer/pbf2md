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

type LatLon struct {
	lat float64
	lon float64
}

func createIndexFile(region string, citySlug string , city string, latLon LatLon) {
	indexFile := region + "/content/cities/" + citySlug + "/_index.md"
	if _, err := os.Stat(indexFile); os.IsNotExist(err) {
		f, err := os.Create(indexFile)
		if err != nil {
			fmt.Println(err)
			return
		}
		data := map[string]interface{} {
			"title": city,
			"url":  "/" + citySlug + "/",
			"latitude": latLon.lat,
			"longitude": latLon.lon,
		}
		indexTmpl := `---
title: {{ .title }}
url: {{ .url }}
latitude: {{ with .latitude }}{{ . }}{{ end }}
longitude: {{ with .longitude }}{{ . }}{{ end }}
---
`
	  indexTemplate := template.Must(template.New("index").Parse(indexTmpl))
		if err = indexTemplate.Execute(f, data); err != nil {
			panic(err)
		}
		f.Close()
	}
}

func createElementFile(region string, citySlug string, nameSlug string, name string) {
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
	data := map[string]interface{} {
		"title": strings.Replace(name, "\"", "", -1),
		"url":  "/" + citySlug + "/" + nameSlug + "/",
	}
	mdTmpl := `---
title: "{{ .title }}"
url: {{ .url }}
---
`
	mdTemplate := template.Must(template.New("markdown").Parse(mdTmpl))
	if err = mdTemplate.Execute(f, data); err != nil {
		panic(err)
	}
	f.Close()
}

func createDataElementFile(region string, citySlug string, nameSlug string, id int64, elementType string, shop string, lat float64, lon float64, tags map[string]string, city string) {
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
	data := map[string]interface{} {
		"id":            id,
		"type":          elementType,
		"shop":          shop,
		"latitude":      lat,
		"longitude":     lon,
		"postcode":      tags["addr:postcode"],
		"city":          city,
		"street":        strings.Replace(tags["addr:street"], "\"", "", -1),
		"housenumber":   tags["addr:housenumber"],
		"phone":         strings.Replace(tags["phone"], "\"", "", -1),
		"opening_hours": strings.Replace(tags["opening_hours"], "\"", "", -1),
		"website":       strings.Replace(tags["website"], "\"", "", -1),
	}
	dataTmpl := `id: {{ .id }}
type: {{ .type }}
shop: {{ .shop }}
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
	if err = dataTemplate.Execute(f, data); err != nil {
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
			translatedShop = "alcohol"
		case "bakery":
			translatedShop = "bakery"
		case "beverages":
			translatedShop = "beverages"
		case "brewing_supplies", "brewery":
			translatedShop = "brewery"
		case "butcher":
			translatedShop = "butcher"
		case "cheese":
			translatedShop = "cheese"
		case "chocolate":
			translatedShop = "chocolate"
		case "coffee", "coffee_roasting", "coffeemaker":
			translatedShop = "coffee"
		case "confectionery", "sweets", "cookies":
			translatedShop = "confectionery"
		case "convenience", "food", "grocery":
			translatedShop = "convenience"
		case "deli":
			translatedShop = "deli"
		case "dairy":
			translatedShop = "dairy"
		case "farm":
			translatedShop = "farm"
		case "frozen_food":
			translatedShop = "frozen food"
		case "greengrocer", "vegetables":
			translatedShop = "greengrocer"
		case "health_food", "organic":
			translatedShop = "health food"
		case "ice_cream":
			translatedShop = "ice cream"
		case "pasta":
			translatedShop = "pasta"
		case "pastry", "cake":
			translatedShop = "pastry"
		case "seafood", "fishmonger":
			translatedShop = "seafood"
		case "spices":
			translatedShop = "spices"
		case "tea":
			translatedShop = "tea"
		case "wine", "winery":
			translatedShop = "wine"
		case "water":
			translatedShop = "water"
		// Warenhaus, Kaufhaus, Einkaufszentrum
		case "department_store":
			translatedShop = "department store"
		case "general", "country_store":
			translatedShop = "general"
		case "kiosk":
			translatedShop = "kiosk"
		case "mall":
			translatedShop = "mall"
		case "supermarket":
			translatedShop = "supermarket"
		case "wholesale":
			translatedShop = "wholesale"
		// Kleidung, Schuhe, Accessoires
		case "baby_goods":
			translatedShop = "baby goods"
		case "bag", "suitcases":
			translatedShop = "bag"
		case "boutique", "fashion", "fashion_accessoires":
			translatedShop = "boutique"
		case "clothes":
			translatedShop = "clothes"
		case "fabric", "fabrics":
			translatedShop = "fabric"
		case "jewelry", "gemstone", "gold", "gold_buyer":
			translatedShop = "jewelry"
		case "leather":
			translatedShop = "leather"
		case "sewing", "sewing_machines":
			translatedShop = "sewing"
		case "shoes", "shoe_repair", "shoe_repairlocksmith":
			translatedShop = "shoes"
		case "tailor":
			translatedShop = "tailor"
		case "watches", "clocks", "watchmaker":
			translatedShop = "watches"
		case "wool":
			translatedShop = "wool"
		// Discounter, Wohltätigkeit
		case "charity", "second_hand":
			translatedShop = "charity"
		case "variety_store":
			translatedShop = "variety store"
		// Gesundheit und Schönheitspflege
		case "beauty", "cosmetics", "decorative_cosmetics", "nail_salon", "wellness", "cosmetic":
			translatedShop = "beauty"
		case "chemist":
			translatedShop = "chemist"
		case "erotic":
			translatedShop = "erotic"
		case "hairdresser":
			translatedShop = "hairdresser"
		case "hairdresser_supply", "combs":
			translatedShop = "hairdresser supply"
		case "hearing_aids":
			translatedShop = "hearing aids"
		case "herbalist":
			translatedShop = "herbalist"
		case "massage":
			translatedShop = "massage"
		case "medical_supply", "medical", "orthopedic", "orthopedics", "sanitary":
			translatedShop = "medical supply"
		case "nutrition_supplements":
			translatedShop = "nutrition supplements"
		case "optician", "eyeglasses":
			translatedShop = "optician"
		case "perfumery":
			translatedShop = "perfumery"
		case "tattoo":
			translatedShop = "tattoo"
		// Do-it-yourself, Haushaltswaren, Baustoffe, Gartenprodukte
		case "agrarian", "agricultural":
			translatedShop = "agrarian"
		case "appliance":
			translatedShop = "appliance"
		case "bathroom_furnishing", "bathroom", "bath":
			translatedShop = "bathroom"
		case "doityourself":
			translatedShop = "doityourself"
		case "electrical", "electric":
			translatedShop = "electrical"
		case "energy", "battery":
			translatedShop = "energy"
		case "fireplace", "furnace", "oven":
			translatedShop = "fireplace"
		case "florist":
			translatedShop = "florist"
		case "garden_centre":
			translatedShop = "garden centre"
		case "garden_furniture":
			translatedShop = "garden furniture"
		case "gas":
			translatedShop = "gas"
		case "glaziery":
			translatedShop = "glaziery"
		case "hardware":
			translatedShop = "hardware"
		case "houseware", "household", "house":
			translatedShop = "houseware"
		case "locksmith", "keys":
			translatedShop = "locksmith"
		case "paint", "paintings", "colors", "painter":
			translatedShop = "paint"
		case "security":
			translatedShop = "security"
		case "trade":
			translatedShop = "trade"
		// Möbel und Innenausstattung
		case "antiques":
			translatedShop = "antiques"
		case "bed":
			translatedShop = "bed"
		case "candles":
			translatedShop = "candles"
		case "carpet":
			translatedShop = "carpet"
		case "curtain":
			translatedShop = "curtain"
		case "doors":
			translatedShop = "doors"
		case "flooring", "parquet":
			translatedShop = "flooring"
		case "furniture":
			translatedShop = "furniture"
		case "interior_decoration", "interior", "decoration", "wallpaper", "interior_store":
			translatedShop = "interior decoration"
		case "kitchen", "kitchen_appliances", "kitchenware", "kitchen_equipment", "cooking", "crockery", "tableware", "ceramics":
			translatedShop = "kitchen"
		case "lamps", "lighting":
			translatedShop = "lamps"
		case "tiles", "tile", "tiling":
			translatedShop = "tiles"
		case "window_blind", "shutter", "shutters":
			translatedShop = "window blind"
		// Elektronik
		case "computer":
			translatedShop = "computer"
		case "robot":
			translatedShop = "robot"
		case "electronics", "electronics_repair", "electro":
			translatedShop = "electronics"
		case "hifi":
			translatedShop = "hifi"
		case "mobile_phone", "telephone", "phone", "communication", "telecommunication":
			translatedShop = "mobile phone"
		case "radiotechnics":
			translatedShop = "radiotechnics"
		case "vacuum_cleaner":
			translatedShop = "vacuum cleaner"
		// Outdoor und Sport, Fahrzeuge
		case "atv":
			translatedShop = "atv"
		case "bicycle", "bike_repair":
			translatedShop = "bicycle"
		case "boat", "yachts":
			translatedShop = "boat"
		case "car":
			translatedShop = "car"
		case "car_repair":
			translatedShop = "car repair"
		case "car_parts":
			translatedShop = "car parts"
		case "caravan", "caravaning":
			translatedShop = "caravan"
		case "fuel":
			translatedShop = "fuel"
		case "fishing", "fishing_gear":
			translatedShop = "fishing"
		case "free_flying":
			translatedShop = "free flying"
		case "golf":
			translatedShop = "golf"
		case "hunting":
			translatedShop = "hunting"
		case "jetski":
			translatedShop = "jetski"
		case "military_surplus", "military":
			translatedShop = "military"
		case "motorcycle", "motorcycle_repair":
			translatedShop = "motorcycle"
		case "outdoor":
			translatedShop = "outdoor"
		case "scuba_diving":
			translatedShop = "diving"
		case "ski":
			translatedShop = "ski"
		case "snowmobile":
			translatedShop = "snowmobile"
		case "sports", "water_sports", "hobby":
			translatedShop = "sports"
		case "swimming_pool":
			translatedShop = "swimming pool"
		case "trailer", "car_trailer":
			translatedShop = "trailer"
		case "tyres":
			translatedShop = "tyres"
		// Kunst, Musik, Hobbys
		case "art", "arts", "artwork":
			translatedShop = "art"
		case "collector", "coins", "comics", "stamps":
			translatedShop = "collector"
		case "craft":
			translatedShop = "craft"
		case "frame", "picture_frames":
			translatedShop = "frame"
		case "games":
			translatedShop = "games"
		case "model", "modelrailway":
			translatedShop = "model"
		case "music":
			translatedShop = "music"
		case "musical_instrument", "woodwind_repair":
			translatedShop = "musical instrument"
		case "photo", "photo_studio", "photographic_studio", "photographer":
			translatedShop = "photo"
		case "camera":
			translatedShop = "camera"
		case "trophy":
			translatedShop = "trophy"
		case "video":
			translatedShop = "video"
		case "video_games":
			translatedShop = "video games"
		// Schreibwaren, Geschenke, Bücher und Zeitungen
		case "anime", "japan":
			translatedShop = "anime"
		case "books", "book_restoration":
			translatedShop = "books"
		case "gift":
			translatedShop = "gift"
		case "lottery":
			translatedShop = "lottery"
		case "newsagent":
			translatedShop = "newsagent"
		case "stationary":
			translatedShop = "stationary"
		case "ticket", "tickets":
			translatedShop = "ticket"
		// Andere
		case "bookmaker":
			translatedShop = "bookmaker"
		case "cannabis", "growshop":
			translatedShop = "cannabis"
		case "copyshop", "printing", "print_shop", "printer_ink", "ink_cartridges", "printer", "printers", "paper", "printshop", "printery":
			translatedShop = "copyshop"
		case "e-cigarette", "vape":
			translatedShop = "e-cigarette"
		case "funeral_directors":
			translatedShop = "funeral directors"
		case "laundry", "rotary_iron", "ironing", "dry_cleaning":
			translatedShop = "laundry"
		case "party":
			translatedShop = "party"
		case "pawnbroker", "money_lender":
			translatedShop = "pawnbroker"
		case "pet":
			translatedShop = "pet"
		case "pet_grooming", "dog_beauty", "dog_hairdresser":
			translatedShop = "pet grooming"
		case "pest_control":
			translatedShop = "pest control"
		case "pyrotechnics":
			translatedShop = "pyrotechnics"
		case "religion":
			translatedShop = "religion"
		case "storage_rental", "rental":
			translatedShop = "storage rental"
		case "tobacco", "smokers":
			translatedShop = "tobacco"
		case "toys":
			translatedShop = "toys"
		case "travel_agency":
			translatedShop = "travel agency"
		case "vacant":
			translatedShop = "vacant"
		case "weapons", "guns", "arms", "knives", "knife":
			translatedShop = "weapons"
		case "outpost":
			translatedShop = "outpost"
		// Benutzerdefiniert
		case "apiary", "beekeeping", "beekeepers_need", "honey", "beekeeper":
			translatedShop = "apiary"
		case "auction_house", "auctioneer", "auction":
			translatedShop = "auction house"
		case "car_accessories", "child_safety_seats":
			translatedShop = "car accessories"
		case "car_service":
			translatedShop = "car service"
		case "caretaker", "building_cleaner":
			translatedShop = "caretaker"
		case "carpenter", "cabinet_maker", "carpentry":
			translatedShop = "carpenter"
		case "casino", "gambling":
			translatedShop = "casino"
		case "catering", "catering_supplies":
			translatedShop = "catering"
		case "equestrian", "horse_equipment", "horse":
			translatedShop = "equestrian"
		case "esoteric":
			translatedShop = "esoteric"
		case "estate_agent":
			translatedShop = "estate agent"
		case "event_service":
			translatedShop = "event service"
		case "fanshop":
			translatedShop = "fanshop"
		case "fitness_equipment":
			translatedShop = "fitness equipment"
		case "flour":
			translatedShop = "flour"
		case "glass":
			translatedShop = "glass"
		case "garden_service":
			translatedShop = "garden service"
		case "garden_machinery", "gardening_tools", "lawn_mower":
			translatedShop = "garden machinery"
		case "grill", "bbq":
			translatedShop = "grill"
		case "grinding":
			translatedShop = "grinding"
		case "hat", "hats":
			translatedShop = "hat"
		case "health":
			translatedShop = "health"
		case "heating_system", "heater", "heating":
			translatedShop = "heating system"
		case "hookah", "shisha":
			translatedShop = "hookah"
		case "hypnotism":
			translatedShop = "hypnotism"
		case "internet_service_provider":
			translatedShop = "internet service provider"
	    case "joiner":
			translatedShop = "joiner"
		case "kids_furnishing":
			translatedShop = "kids furnishing"
		case "furs":
			translatedShop = "furs"
		case "lettering", "license_plate", "license_plates", "number_plate", "sign_make", "signs":
			translatedShop = "lettering"
		case "locksmithery", "metalworker", "metalwork", "metalworking":
			translatedShop = "locksmithery"
		case "machinery", "vehicle", "vehicles", "forklift", "agricultural_machinery", "industrial", "machines":
			translatedShop = "machinery"
		case "office_supplies", "office", "stationery":
			translatedShop = "office supplies"
		case "pet_food", "fodder":
			translatedShop = "pet food"
		case "plumber", "plumbing_business":
			translatedShop = "plumber"
		case "piercing":
			translatedShop = "piercing"
		case "pottery":
			translatedShop = "pottery"
		case "ship_chandler":
			translatedShop = "ship chandler"
		case "software":
			translatedShop = "software"
		case "solarium", "sunstudio":
			translatedShop = "solarium"
		case "stones", "tombstone", "tombstones", "gravestones":
			translatedShop = "stones"
		case "tanning":
			translatedShop = "tanning"
		case "tools", "screws":
			translatedShop = "tools"
		case "wedding_gown", "wedding":
			translatedShop = "wedding gown"
		case "whirlpool", "pool":
			translatedShop = "whirlpool"
		case "wood", "timber", "sawmill":
			translatedShop = "wood"
		case "worldshop", "one_world", "afro", "oneworld":
			translatedShop = "worldshop"
		default:
			translatedShop = "shop"
			fmt.Println("unknown shop:", shop)
	}
	return translatedShop
}

func main() {
	regions := []string{"alabama", "alaska", "arizona", "arkansas", "california", "colorado", "connecticut", "delaware", "district-of-columbia", "florida", "georgia", "hawaii", "idaho", "illinois", "indiana", "iowa", "kansas", "kentucky", "louisiana", "maine", "maryland", "massachusetts", "michigan", "minnesota", "mississippi", "missouri", "montana", "nebraska", "nevada", "new-hampshire", "new-jersey", "new-mexico", "new-york", "north-carolina", "north-dakota", "ohio", "oklahoma", "oregon", "pennsylvania", "puerto-rico", "rhode-island", "south-carolina", "south-dakota", "tennessee", "texas", "utah", "vermont", "virginia", "washington", "west-virginia", "wisconsin", "wyoming"}
	for _, r := range regions {
		f, err := os.Open(r + "-latest.osm.pbf")
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		d := osmpbf.NewDecoder(f)

		// Cache nodes, ways and cities
		m := make(map[int64]LatLon)
		n := make(map[int64]string)
		p := make(map[string]LatLon)

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
						citySlug := slug.MakeLang(city, "en")
						nameSlug := slug.MakeLang(name, "en")

						// 1. content
						err = os.MkdirAll(r + "/content/cities/" + citySlug, 0755)
						createIndexFile(r, citySlug, city, p[city])
						createElementFile(r, citySlug, nameSlug, name)

						// 2. data
						err = os.MkdirAll(r + "/data/cities/" + citySlug, 0755)
						createDataElementFile(r, citySlug, nameSlug, v.ID, "node", translateShop(shop), v.Lat, v.Lon, tags, city)
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
						citySlug := slug.MakeLang(city, "en")
						nameSlug := slug.MakeLang(name, "en")

						// 1. content
						err = os.MkdirAll(r + "/content/cities/" + citySlug, 0755)
						createIndexFile(r, citySlug, city, p[city])
						createElementFile(r, citySlug, nameSlug, name)

						// 2. data
						err = os.MkdirAll(r + "/data/cities/" + citySlug, 0755)
						node := m[v.NodeIDs[0]] // Lookup coords of first childnode
						createDataElementFile(r, citySlug, nameSlug, v.ID, "way", translateShop(shop), node.lat, node.lon, tags, city)

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
}

package main

import (
	"fmt"
	"io"
	"log"
	"math"
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

func roundFloat(value float64, decimals float64) float64 {
	factor := math.Pow(10, decimals)
	return math.Round(value*factor)/factor
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
			"latitude":  roundFloat(latLon.lat, 3),
			"longitude": roundFloat(latLon.lon, 3),
		}
		if err = template.Execute(f, data); err != nil {
			panic(err)
		}
		f.Close()
	}
}

func createShopFile(region string, shopSlug string, shop string, icon string, template *template.Template) {
	indexFile := region + "/content/shops/" + shopSlug + "/_index.md"
	if _, err := os.Stat(indexFile); os.IsNotExist(err) {
		f, err := os.Create(indexFile)
		if err != nil {
			fmt.Println(err)
			return
		}
		data := map[string]interface{}{
			"title": shop,
			"url":   "/" + shopSlug + "/",
			"icon":  icon,
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
		"latitude":      roundFloat(lat, 5),
		"longitude":     roundFloat(lon, 5),
		"postcode":      tags["addr:postcode"],
		"city":          city,
		"street":        strings.Replace(tags["addr:street"], "\"", "", -1),
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

func Find(slice []string, val string) (int, bool) {
	for i, item := range slice {
		if item == val {
			return i, true
		}
	}
	return -1, false
}

func getIcon(shop string) string {
	icons := []string{
		"alcohol",
		"art",
		"bag",
		"bakery",
		"beauty",
		"bed",
		"beverages",
		"bicycle",
		"bookmaker",
		"books",
		"butcher",
		"car_parts",
		"carpet",
		"car_repair",
		"car",
		"charity",
		"chemist",
		"clothes",
		"coffee",
		"computer",
		"confectionery",
		"convenience",
		"copyshop",
		"dairy",
		"deli",
		"department_store",
		"diy",
		"doityourself",
		"electronics",
		"fabric",
		"florist",
		"furniture",
		"garden_centre",
		"gift",
		"greengrocer",
		"hairdresser",
		"hifi",
		"houseware",
		"ice_cream",
		"interior_decoration",
		"jewellery",
		"jewelry",
		"kiosk",
		"laundry",
		"marketplace",
		"massage",
		"medical_supply",
		"mobile_phone",
		"motorcycle",
		"musical_instrument",
		"music",
		"newsagent",
		"news",
		"optician",
		"outdoor",
		"paint",
		"perfumery",
		"pet",
		"photo",
		"seafood",
		"second_hand",
		"shoes",
		"sports",
		"stationery",
		"supermarket",
		"tea",
		"ticket",
		"tobacco",
		"toys",
		"trade",
		"travel_agency",
		"tyres",
		"variety_store",
		"video_games",
		"video",
	}
	shopSplit := strings.Split(shop, ";")
	_, found := Find(icons, shopSplit[0])
	if found {
		return shopSplit[0]
	}
	return "other"
}

func translateShop(shop string) string {
	shopSplit := strings.Split(shop, ";")
	var translatedShop string
	switch shopSplit[0] {
	// Lebensmittel, Getränke
	case "alcohol":
		translatedShop = "alcool"
	case "bakery":
		translatedShop = "boulangerie"
	case "beverages":
		translatedShop = "boissons"
	case "brewing_supplies", "brewery":
		translatedShop = "brasserie"
	case "butcher":
		translatedShop = "boucherie"
	case "cheese":
		translatedShop = "fromage"
	case "chocolate":
		translatedShop = "chocolat"
	case "coffee", "coffee_roasting", "coffeemaker":
		translatedShop = "café"
	case "confectionery", "sweets", "cookies":
		translatedShop = "confiserie"
	case "convenience", "food", "grocery":
		translatedShop = "commodité"
	case "deli":
		translatedShop = "charcuterie"
	case "dairy":
		translatedShop = "produits laitiers"
	case "farm":
		translatedShop = "ferme"
	case "frozen_food":
		translatedShop = "aliments surgelés"
	case "greengrocer", "vegetables":
		translatedShop = "légumes"
	case "health_food", "organic":
		translatedShop = "alimentation saine"
	case "ice_cream":
		translatedShop = "crème glacée"
	case "pasta":
		translatedShop = "pâtes"
	case "pastry", "cake":
		translatedShop = "pâtisserie"
	case "seafood", "fishmonger":
		translatedShop = "fruits de mer"
	case "spices":
		translatedShop = "épices"
	case "tea":
		translatedShop = "thé"
	case "wine", "winery":
		translatedShop = "vin"
	case "water":
		translatedShop = "eau"
	// Warenhaus, Kaufhaus, Einkaufszentrum
	case "department_store":
		translatedShop = "grand magasin"
	case "general", "country_store":
		translatedShop = "magasin de campagne"
	case "kiosk":
		translatedShop = "kiosque"
	case "mall":
		translatedShop = "centre commercial"
	case "supermarket":
		translatedShop = "supermarché"
	case "wholesale":
		translatedShop = "vente en gros"
	// Kleidung, Schuhe, Accessoires
	case "baby_goods":
		translatedShop = "produits pour bébés"
	case "bag", "suitcases":
		translatedShop = "sac"
	case "boutique", "fashion", "fashion_accessoires":
		translatedShop = "boutique"
	case "clothes":
		translatedShop = "vêtements"
	case "fabric", "fabrics":
		translatedShop = "tissu"
	case "jewelry", "gemstone", "gold", "gold_buyer":
		translatedShop = "bijoux"
	case "leather":
		translatedShop = "cuir"
	case "sewing", "sewing_machines":
		translatedShop = "couture"
	case "shoes", "shoe_repair", "shoe_repairlocksmith":
		translatedShop = "chaussures"
	case "tailor":
		translatedShop = "tailleur"
	case "watches", "clocks", "watchmaker":
		translatedShop = "montres"
	case "wool":
		translatedShop = "laine"
	// Discounter, Wohltätigkeit
	case "charity", "second_hand":
		translatedShop = "charité"
	case "variety_store":
		translatedShop = "magasin de variétés"
	// Gesundheit und Schönheitspflege
	case "beauty", "cosmetics", "decorative_cosmetics", "nail_salon", "wellness", "cosmetic":
		translatedShop = "beauté"
	case "chemist":
		translatedShop = "chimiste"
	case "erotic":
		translatedShop = "érotique"
	case "hairdresser":
		translatedShop = "coiffeur"
	case "hairdresser_supply", "combs":
		translatedShop = "fournitures pour coiffeurs"
	case "hearing_aids":
		translatedShop = "les appareils auditifs"
	case "herbalist":
		translatedShop = "herboriste"
	case "massage":
		translatedShop = "massage"
	case "medical_supply", "medical", "orthopedic", "orthopedics", "sanitary":
		translatedShop = "approvisionnement médical"
	case "nutrition_supplements":
		translatedShop = "les compléments alimentaires"
	case "optician", "eyeglasses":
		translatedShop = "opticien"
	case "perfumery":
		translatedShop = "parfumerie"
	case "tattoo":
		translatedShop = "tatouage"
	// Do-it-yourself, Haushaltswaren, Baustoffe, Gartenprodukte
	case "agrarian", "agricultural":
		translatedShop = "agraire"
	case "appliance":
		translatedShop = "appareil ménager"
	case "bathroom_furnishing", "bathroom", "bath":
		translatedShop = "salle de bains"
	case "doityourself":
		translatedShop = "à faire soi-même"
	case "electrical", "electric":
		translatedShop = "électrique"
	case "energy", "battery":
		translatedShop = "énergie"
	case "fireplace", "furnace", "oven":
		translatedShop = "cheminée"
	case "florist":
		translatedShop = "fleuriste"
	case "garden_centre":
		translatedShop = "centre de jardinage"
	case "garden_furniture":
		translatedShop = "meubles de jardin"
	case "gas":
		translatedShop = "gaz"
	case "glaziery":
		translatedShop = "vitrerie"
	case "hardware":
		translatedShop = "matériel informatique"
	case "houseware", "household", "house":
		translatedShop = "articles ménagers"
	case "locksmith", "keys":
		translatedShop = "serrurier"
	case "paint", "paintings", "colors", "painter":
		translatedShop = "peinture"
	case "security":
		translatedShop = "sécurité"
	case "trade":
		translatedShop = "commerce"
	// Möbel und Innenausstattung
	case "antiques":
		translatedShop = "antiquités"
	case "bed":
		translatedShop = "lit"
	case "candles":
		translatedShop = "bougies"
	case "carpet":
		translatedShop = "tapis"
	case "curtain":
		translatedShop = "rideau"
	case "doors":
		translatedShop = "portes"
	case "flooring", "parquet":
		translatedShop = "revêtement de sol"
	case "furniture":
		translatedShop = "meubles"
	case "interior_decoration", "interior", "decoration", "wallpaper", "interior_store":
		translatedShop = "décoration intérieure"
	case "kitchen", "kitchen_appliances", "kitchenware", "kitchen_equipment", "cooking", "crockery", "tableware", "ceramics":
		translatedShop = "cuisine"
	case "lamps", "lighting":
		translatedShop = "lampes"
	case "tiles", "tile", "tiling":
		translatedShop = "tuiles"
	case "window_blind", "shutter", "shutters":
		translatedShop = "store de fenêtre"
	// Elektronik
	case "computer":
		translatedShop = "ordinateur"
	case "robot":
		translatedShop = "robot"
	case "electronics", "electronics_repair", "electro":
		translatedShop = "électronique"
	case "hifi":
		translatedShop = "hifi"
	case "mobile_phone", "telephone", "phone", "communication", "telecommunication":
		translatedShop = "téléphone portable"
	case "radiotechnics":
		translatedShop = "radiotechnique"
	case "vacuum_cleaner":
		translatedShop = "aspirateur"
	// Outdoor und Sport, Fahrzeuge
	case "atv":
		translatedShop = "atv"
	case "bicycle", "bike_repair":
		translatedShop = "vélo"
	case "boat", "yachts":
		translatedShop = "bateau"
	case "car":
		translatedShop = "voiture"
	case "car_repair":
		translatedShop = "réparation de voitures"
	case "car_parts":
		translatedShop = "pièces de voitures"
	case "caravan", "caravaning":
		translatedShop = "caravane"
	case "fuel":
		translatedShop = "carburant"
	case "fishing", "fishing_gear":
		translatedShop = "pêche"
	case "free_flying":
		translatedShop = "vol libre"
	case "golf":
		translatedShop = "golf"
	case "hunting":
		translatedShop = "chasse"
	case "jetski":
		translatedShop = "jetski"
	case "military_surplus", "military":
		translatedShop = "militaire"
	case "motorcycle", "motorcycle_repair":
		translatedShop = "moto"
	case "outdoor":
		translatedShop = "extérieur"
	case "scuba_diving":
		translatedShop = "plongée"
	case "ski":
		translatedShop = "ski"
	case "snowmobile":
		translatedShop = "motoneige"
	case "sports", "water_sports", "hobby":
		translatedShop = "sports"
	case "swimming_pool":
		translatedShop = "piscine"
	case "trailer", "car_trailer":
		translatedShop = "remorque"
	case "tyres":
		translatedShop = "pneus"
	// Kunst, Musik, Hobbys
	case "art", "arts", "artwork":
		translatedShop = "art"
	case "collector", "coins", "comics", "stamps":
		translatedShop = "collecteur"
	case "craft":
		translatedShop = "artisanat"
	case "frame", "picture_frames":
		translatedShop = "cadre"
	case "games":
		translatedShop = "jeux"
	case "model", "modelrailway":
		translatedShop = "modèle"
	case "music":
		translatedShop = "musique"
	case "musical_instrument", "woodwind_repair":
		translatedShop = "instrument de musique"
	case "photo", "photo_studio", "photographic_studio", "photographer":
		translatedShop = "photo"
	case "camera":
		translatedShop = "caméra"
	case "trophy":
		translatedShop = "trophée"
	case "video":
		translatedShop = "vidéo"
	case "video_games":
		translatedShop = "jeux vidéo"
	// Schreibwaren, Geschenke, Bücher und Zeitungen
	case "anime", "japan":
		translatedShop = "anime"
	case "books", "book_restoration":
		translatedShop = "livres"
	case "gift":
		translatedShop = "cadeau"
	case "lottery":
		translatedShop = "loterie"
	case "newsagent":
		translatedShop = "marchand de journaux"
	case "stationary":
		translatedShop = "stationnaire"
	case "ticket", "tickets":
		translatedShop = "billet"
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
		translatedShop = "directeurs de funérailles"
	case "laundry", "rotary_iron", "ironing", "dry_cleaning":
		translatedShop = "blanchisserie"
	case "party":
		translatedShop = "fête"
	case "pawnbroker", "money_lender":
		translatedShop = "prêteur sur gages"
	case "pet":
		translatedShop = "animal de compagnie"
	case "pet_grooming", "dog_beauty", "dog_hairdresser":
		translatedShop = "toilettage des animaux"
	case "pest_control":
		translatedShop = "contrôle des nuisibles"
	case "pyrotechnics":
		translatedShop = "pyrotechnie"
	case "religion":
		translatedShop = "religion"
	case "storage_rental", "rental":
		translatedShop = "location de stockage"
	case "tobacco", "smokers":
		translatedShop = "tabac"
	case "toys":
		translatedShop = "jouets"
	case "travel_agency":
		translatedShop = "agence de voyage"
	case "vacant":
		translatedShop = "vacant"
	case "weapons", "guns", "arms", "knives", "knife":
		translatedShop = "armes"
	case "outpost":
		translatedShop = "avant-poste"
	// Benutzerdefiniert
	case "apiary", "beekeeping", "beekeepers_need", "honey", "beekeeper":
		translatedShop = "rucher"
	case "auction_house", "auctioneer", "auction":
		translatedShop = "maison d'enchères"
	case "car_accessories", "child_safety_seats":
		translatedShop = "accessoires automobiles"
	case "car_service":
		translatedShop = "service automobile"
	case "caretaker", "building_cleaner":
		translatedShop = "concierge"
	case "carpenter", "cabinet_maker", "carpentry":
		translatedShop = "charpentier"
	case "casino", "gambling":
		translatedShop = "casino"
	case "catering", "catering_supplies":
		translatedShop = "restauration"
	case "equestrian", "horse_equipment", "horse":
		translatedShop = "équestre"
	case "esoteric":
		translatedShop = "ésotérique"
	case "estate_agent":
		translatedShop = "agent immobilier"
	case "event_service":
		translatedShop = "service événementiel"
	case "fanshop":
		translatedShop = "boutique de fans"
	case "fitness_equipment":
		translatedShop = "matériel de fitness"
	case "flour":
		translatedShop = "farine"
	case "glass":
		translatedShop = "verre"
	case "garden_service":
		translatedShop = "service de jardinage"
	case "garden_machinery", "gardening_tools", "lawn_mower":
		translatedShop = "machines de jardinage"
	case "grill", "bbq":
		translatedShop = "gril"
	case "grinding":
		translatedShop = "broyage"
	case "hat", "hats":
		translatedShop = "chapeau"
	case "health":
		translatedShop = "santé"
	case "heating_system", "heater", "heating":
		translatedShop = "système de chauffage"
	case "hookah", "shisha":
		translatedShop = "narguilé"
	case "hypnotism":
		translatedShop = "hypnotisme"
	case "internet_service_provider":
		translatedShop = "fournisseur de services internet"
	case "joiner":
		translatedShop = "menuisier"
	case "kids_furnishing":
		translatedShop = "mobilier pour enfants"
	case "furs":
		translatedShop = "fourrures"
	case "lettering", "license_plate", "license_plates", "number_plate", "sign_make", "signs":
		translatedShop = "lettrage"
	case "locksmithery", "metalworker", "metalwork", "metalworking":
		translatedShop = "serrurerie"
	case "machinery", "vehicle", "vehicles", "forklift", "agricultural_machinery", "industrial", "machines":
		translatedShop = "machines"
	case "office_supplies", "office", "stationery":
		translatedShop = "fournitures de bureau"
	case "pet_food", "fodder":
		translatedShop = "aliments pour animaux"
	case "plumber", "plumbing_business":
		translatedShop = "plombier"
	case "piercing":
		translatedShop = "piercing"
	case "pottery":
		translatedShop = "poterie"
	case "ship_chandler":
		translatedShop = "manutentionnaire de navires"
	case "software":
		translatedShop = "logiciel"
	case "solarium", "sunstudio":
		translatedShop = "solarium"
	case "stones", "tombstone", "tombstones", "gravestones":
		translatedShop = "pierres"
	case "tanning":
		translatedShop = "bronzage"
	case "tools", "screws":
		translatedShop = "outils"
	case "wedding_gown", "wedding":
		translatedShop = "robe de mariée"
	case "whirlpool", "pool":
		translatedShop = "tourbillon"
	case "wood", "timber", "sawmill":
		translatedShop = "bois"
	case "worldshop", "one_world", "afro", "oneworld":
		translatedShop = "boutique en ligne"
	default:
		translatedShop = "shop"
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
	shopTmpl := `---
title: {{ .title }}
url: {{ .url }}
icon: {{ .icon }}
---
`
	shopTemplate := template.Must(template.New("index").Parse(shopTmpl))
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
				if city != "" && name != "" && shop != "" {
					citySlug := slug.MakeLang(city, "fr")
					nameSlug := slug.MakeLang(name, "fr")
					translatedShop := translateShop(shop)
					shopSlug := slug.MakeLang(translatedShop, "fr")

					// Exceptions: skip foreign cities
					if citySlug == "s-heerenberg" {
						// 's-Heerenberg in the Netherlands
						continue
					}
					// Cache cities LatLon
					_, exists := p[citySlug]
					if !exists {
						p[citySlug] = LatLon{v.Lat, v.Lon}
					}
					// 1. content
					err = os.MkdirAll(region+"/content/cities/"+citySlug, 0755)
					createIndexFile(region, citySlug, city, p[citySlug], indexTemplate)
					err = os.MkdirAll(region+"/content/shops/"+shopSlug, 0755)
					createShopFile(region, shopSlug, translatedShop, getIcon(shop), shopTemplate)
					createElementFile(region, citySlug, nameSlug, name, translatedShop, mdTemplate)

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
					citySlug := slug.MakeLang(city, "fr")
					nameSlug := slug.MakeLang(name, "fr")
					translatedShop := translateShop(shop)
					shopSlug := slug.MakeLang(translatedShop, "fr")
					node := m[v.NodeIDs[0]] // Lookup coords of first childnode

					// Cache cities LatLon
					_, exists := p[citySlug]
					if !exists {
						p[citySlug] = LatLon{node.lat, node.lon}
					}
					// 1. content
					err = os.MkdirAll(region+"/content/cities/"+citySlug, 0755)
					createIndexFile(region, citySlug, city, p[citySlug], indexTemplate)
					err = os.MkdirAll(region+"/content/shops/"+shopSlug, 0755)
					createShopFile(region, shopSlug, translatedShop, getIcon(shop), shopTemplate)
					createElementFile(region, citySlug, nameSlug, name, translatedShop, mdTemplate)

					// 2. data
					err = os.MkdirAll(region+"/data/cities/"+citySlug, 0755)
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

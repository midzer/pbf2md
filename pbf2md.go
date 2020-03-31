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

func createIndexShopFile(citySlug string , shopSlug string, shop string) {
	indexShopFile := "content/cities/" + citySlug + "/" + shopSlug + "/_index.md"
	if !fileExists(indexShopFile) {
		f, err := os.Create(indexShopFile)
		if err != nil {
			fmt.Println(err)
			return
		}
		data := map[string]interface{} {
			"title": shop,
			"url":  "/" + citySlug + "/" + shopSlug + "/",
		}
		const indexShopTmpl = `---
title: {{ .title }}
url: {{ .url }}
---
`
	  indexShopTemplate := template.Must(template.New("indexShop").Parse(indexShopTmpl))
		if err := indexShopTemplate.Execute(f, data); err != nil {
				panic(err)
		}
		f.Close()
	}
}

func createElementFile(citySlug string, shopSlug string, nameSlug string, name string, elementType string, shop string) {
	elementFile := "content/cities/" + citySlug + "/" + shopSlug + "/" + nameSlug + ".md"
	if fileExists(elementFile) {
		var re = regexp.MustCompile(`\d+$`)
		i := 2
		nameSlug = fmt.Sprintf("%s-%d", nameSlug, i)
		for {
			elementFile := "content/cities/" + citySlug + "/" + shopSlug + "/" + nameSlug + ".md"
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
		"url":  "/" + citySlug + "/" + shopSlug + "/" + nameSlug + "/",
		"type": elementType,
		"categories": shop,
	}
	const mdTmpl = `---
title: "{{ .title }}"
url: {{ .url }}
type: {{ .type }}
categories: ['{{ .categories }}']
---
`
	mdTemplate := template.Must(template.New("markdown").Parse(mdTmpl))
	if err := mdTemplate.Execute(f, data); err != nil {
			panic(err)
	}
	f.Close()
}

func createDataElementFile(citySlug string, nameSlug string, shopSlug string, id int64, lat float64, lon float64, tags map[string]string, city string) {
	elementFile := "data/cities/" + citySlug + "/" + nameSlug + "/" + shopSlug + ".yml"
	if fileExists(elementFile) {
		var re = regexp.MustCompile(`\d+$`)
		i := 2
		nameSlug = fmt.Sprintf("%s-%d", nameSlug, i)
		for {
			elementFile := "data/cities/" + citySlug + "/" + nameSlug + "/" + shopSlug + ".yml"
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
		case "alcohol":
			translatedShop = "Alkohol"
		case "art":
			translatedShop = "Kunst"
		case "baby_goods":
			translatedShop = "Babywaren"
		case "bag":
			translatedShop = "Taschen"
		case "bakery":
			translatedShop = "Bäckerei"
		case "beauty":
			translatedShop = "Schönheit"
		case "bed":
			translatedShop = "Betten"
		case "beverages":
			translatedShop = "Getränke"
		case "bicycle":
			translatedShop = "Fahrrad"
		case "bookmaker":
			translatedShop = "Buchbinder"
		case "books":
			translatedShop = "Bücher"
		case "boutique":
			translatedShop = "Laden"
		case "butcher":
			translatedShop = "Metzgerei"
		case "car":
			translatedShop = "Autos"
		case "car_accessories":
			translatedShop = "Autozubehör"
		case "car_parts":
			translatedShop = "Autoteile"
		case "car_repair":
			translatedShop = "Autoreparatur"
		case "carpet":
			translatedShop = "Teppiche"
		case "cheese":
			translatedShop = "Käse"
		case "chemist":
			translatedShop = "Apotheke"
		case "clothes":
			translatedShop = "Kleidung"
		case "coffee":
			translatedShop = "Kaffee"
		case "comics":
			translatedShop = "Comics"
		case "communication":
			translatedShop = "Kommunikation"
		case "computer":
			translatedShop = "Computer"
		case "confectionery":
			translatedShop = "Konditorei"
		case "convenience":
			translatedShop = "Komfort"
		case "copyshop":
			translatedShop = "Copyshop"
		case "cosmetics":
			translatedShop = "Kosmetik"
		case "craft":
			translatedShop = "Handwerk"
		case "curtain":
			translatedShop = "Gardinen"
		case "deli":
			translatedShop = "Feinkostladen"
		case "department_store":
			translatedShop = "Warenhaus"
		case "doityourself":
			translatedShop = "Selbermachen"
		case "dry_cleaning":
			translatedShop = "Trockenreinigung"
		case "e-cigarette":
			translatedShop = "E-Zigaretten"
		case "electrical":
			translatedShop = "Elektrisch"
		case "electronics":
			translatedShop = "Elektrotechnik"
		case "estate_agent":
			translatedShop = "Immobilienmakler"
		case "fabric":
			translatedShop = "Fabrik"
		case "florist":
			translatedShop = "Blumenladen"
		case "funeral_directors":
			translatedShop = "Bestattungsunternehmen"
		case "furniture":
			translatedShop = "Möbel"
		case "gambling":
			translatedShop = "Glücksspiel"
		case "games":
			translatedShop = "Spiele"
		case "garden_centre":
			translatedShop = "Gartencenter"
		case "gift":
			translatedShop = "Geschenke"
		case "glaziery":
			translatedShop = "Glaserei"
		case "greengrocer":
			translatedShop = "Gemüsehändler"
		case "hairdresser":
			translatedShop = "Friseur"
		case "hardware":
			translatedShop = "Hardware"
		case "hearing_aids":
			translatedShop = "Hörgeräte"
		case "hifi":
			translatedShop = "HiFi"
		case "hookah":
			translatedShop = "Wasserpfeife"
		case "houseware":
			translatedShop = "Haushaltswaren"
		case "ice_cream":
			translatedShop = "Eisdiele"
		case "interior_decoration":
			translatedShop = "Innenausstatter"
		case "japan":
			translatedShop = "Japan"
		case "jewelry":
			translatedShop = "Juwelier"
		case "kiosk":
			translatedShop = "Kiosk"
		case "kitchen":
			translatedShop = "Küche"
		case "kitchen_equipment":
			translatedShop = "Küchenzubehör"
		case "laundry":
			translatedShop = "Wäscherei"
		case "license_plate":
			translatedShop = "Kennzeichen"
		case "locksmith":
			translatedShop = "Schlüsseldienst"
		case "mall":
			translatedShop = "Einkaufszentrum"
		case "massage":
			translatedShop = "Massage"
		case "medical_supply":
			translatedShop = "Medizinische Versorgung"
		case "mobile_phone":
			translatedShop = "Mobiltelefone"
		case "motorcycle":
			translatedShop = "Motorräder"
		case "music":
			translatedShop = "Musik"
		case "musical_instrument":
			translatedShop = "Musikinstrumente"
		case "newsagent":
			translatedShop = "Nachrichtenagentur"
		case "optician":
			translatedShop ="Optiker"
		case "opticianhearing_aids":
			translatedShop = "Optiker-Hörhilfen"
		case "outdoor":
			translatedShop = "Außenbereich"
		case "paint":
			translatedShop = "Farbe"
		case "party":
			translatedShop = "Party"
		case "pastry":
			translatedShop = "Backwaren"
		case "pawnbroker":
			translatedShop = "Pfandleiher"
		case "perfumery":
			translatedShop = "Parfümerie"
		case "pet":
			translatedShop = "Tiere"
		case "pet_grooming":
			translatedShop = "Haustier-Pflege"
		case "photo":
			translatedShop = "Foto"
		case "piercing":
			translatedShop = "Piercing"
		case "pottery":
			translatedShop = "Töpferei"
		case "seafood":
			translatedShop = "Meeresfrüchte"
		case "second_hand":
			translatedShop = "Second Hand"
		case "sewing_machines":
			translatedShop = "Nähmaschinen"
		case "ship_chandler":
			translatedShop = "Schiffsausrüster"
		case "shoe_repair":
			translatedShop = "Schuhreparatur"
		case "shoe_repairlocksmith":
			translatedShop = "Schuhreparatur-Schlosserei"
		case "shoes":
			translatedShop = "Schuhe"
		case "sign_make":
			translatedShop = "Schilderherstellung"
		case "solarium":
			translatedShop = "Solarium"
		case "sonderposten":
			translatedShop = "Sonderposten"
		case "spices":
			translatedShop = "Gewürze"
		case "sports":
			translatedShop = "Sportwaren"
		case "stamps":
			translatedShop = "Briefmarken"
		case "stationery":
			translatedShop = "Schreibwaren"
		case "supermarket":
			translatedShop = "Supermarkt"
		case "tailor":
			translatedShop = "Schneiderei"
		case "tattoo":
			translatedShop = "Tattoo"
		case "tea":
			translatedShop = "Teewaren"
		case "ticket":
			translatedShop = "Ticketverkauf"
		case "tobacco":
			translatedShop = "Tabakwaren"
		case "toys":
			translatedShop = "Spielzeug"
		case "trade":
			translatedShop = "Handel"
		case "travel_agency":
			translatedShop = "Reiseagentur"
		case "trophy":
			translatedShop = "Pokale"
		case "tyres":
			translatedShop = "Reifen"
		case "vacuum_cleaner":
			translatedShop = "Staubsauger"
		case "variety_store":
			translatedShop = "Varietéladen"
		case "video":
			translatedShop = "Videothek"
		case "video_games":
			translatedShop = "Videospiele"
		case "weapons":
			translatedShop = "Waffen"
		case "wine":
			translatedShop = "Wein"
		case "wool":
			translatedShop = "Wolle"
		default:
			translatedShop = "Unbekannt"
			fmt.Println("unknown shop:", shop)
	}
	return translatedShop
}

func main() {
  f, err := os.Open("bremen-latest.osm.pbf")
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
					shopSlug := slug.MakeLang(shop, "de")

					// 1. content
					err = os.MkdirAll("content/cities/" + citySlug + "/" + shopSlug, 0755)
					createIndexFile(citySlug, city)
					createIndexShopFile(citySlug, shopSlug, shop)
					createElementFile(citySlug, shopSlug, nameSlug, name, "node", shop)
					

					// 2. data
					err = os.MkdirAll("data/cities/" + citySlug + "/" + shopSlug, 0755)
					createDataElementFile(citySlug, shopSlug, nameSlug, v.ID, v.Lat, v.Lon, tags, city)
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
					shopSlug := slug.MakeLang(shop, "de")

					// 1. content
					err = os.MkdirAll("content/cities/" + citySlug + "/" + shopSlug, 0755)
					createIndexFile(citySlug, city)
					createIndexShopFile(citySlug, shopSlug, shop)
					createElementFile(citySlug, shopSlug, nameSlug, name, "way", shop)

					// 2. data
					err = os.MkdirAll("data/cities/" + citySlug + "/" + shopSlug, 0755)
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
					createDataElementFile(citySlug, shopSlug, nameSlug, v.ID, lat, lon, tags, city)
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

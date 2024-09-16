package main

import (
	"fmt"
	"io"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"text/template"

	"github.com/gosimple/slug"
	"github.com/thomersch/gosmparse"
	"gopkg.in/yaml.v3"
)

type LatLon struct {
	lat float64
	lon float64
}

func roundFloat(value float64, decimals float64) float64 {
	factor := math.Pow(10, decimals)
	return math.Round(value*factor) / factor
}

func createFile(filePath string) *os.File {
	f, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0644)
	if err != nil {
		f.Close()
		return nil
	}

	return f
}

func writeData(file *os.File, template *template.Template, data map[string]interface{}) {
	err := template.Execute(file, data)
	if err != nil {
		panic(err)
	}
	file.Close()
}

func createIndexPath(region string, folder string, slug string) string {
	return region + "/content/" + folder + "/" + slug + "/_index.md"
}

func createElementFile(region string, folder string, citySlug string, nameSlug string, streetSlug string, extension string) (*os.File, string) {
	elementFile := region + "/" + folder + "/cities/" + citySlug + "/" + nameSlug + "." + extension
	f := createFile(elementFile)
	if f == nil {
		if streetSlug != "" {
			nameSlug = fmt.Sprintf("%s-%s", nameSlug, streetSlug)
			elementFile = region + "/" + folder + "/cities/" + citySlug + "/" + nameSlug + "." + extension
		}
		f = createFile(elementFile)
		if f == nil {
			re := regexp.MustCompile(`\d+$`)
			nameSlug = fmt.Sprintf("%s-%d", nameSlug, 2)
			for i := 3; i < 999; i++ {
				elementFile = region + "/" + folder + "/cities/" + citySlug + "/" + nameSlug + "." + extension
				f = createFile(elementFile)
				if f != nil {
					return f, nameSlug
				}
				nameSlug = re.ReplaceAllString(nameSlug, strconv.Itoa(i))
			}
		}
	}

	return f, nameSlug
}

type dataPrepare struct {
	Nodes *uint64
	Ways  *uint64
}

type dataHandler struct {
	Nodes *uint64
	Ways  *uint64
}

func main() {
	region = os.Args[1]
	r, err := os.Open(region + "-latest.osm.pbf")
	if err != nil {
		panic(err)
	}
	dec := gosmparse.NewDecoder(r)
	err = dec.Parse(&dataPrepare{})
	if err != nil {
		panic(err)
	}
	r.Seek(0, io.SeekStart)
	err = dec.Parse(&dataHandler{})
	if err != nil {
		panic(err)
	}
}

// Globals
var (
	region string
	nc     sync.Map
	wc     sync.Map

	// Create templates
	indexTmpl = `---
title: {{ .title }}
url: {{ .url }}
latitude: {{ with .latitude }}{{ . }}{{ end }}
longitude: {{ with .longitude }}{{ . }}{{ end }}
---`
	indexTemplate = template.Must(template.New("index").Parse(indexTmpl))
	shopTmpl      = `---
title: {{ .title }}
url: {{ .url }}
icon: {{ .icon }}
---`
	shopTemplate = template.Must(template.New("shop").Parse(shopTmpl))
	mdTmpl       = `---
title: "{{ .title }}"
url: {{ .url }}
shop: {{ .shop }}
---`
	mdTemplate = template.Must(template.New("markdown").Parse(mdTmpl))
)

type data struct {
	ID           int64
	Type         string
	Latitude     float64
	Longitude    float64
	Postcode     string
	City         string
	Street       string
	Housenumber  string
	Phone        string
	OpeningHours string `yaml:"opening_hours"`
	Website      string
}

func (d *dataPrepare) ReadNode(n gosmparse.Node) {}

func (d *dataHandler) ReadNode(n gosmparse.Node) {
	tags := n.Tags
	city := tags["addr:city"]
	name := tags["name"]
	shop := tags["shop"]
	if city != "" && name != "" && shop != "" {
		citySlug := slug.MakeLang(city, "fr")
		nameSlug := slug.MakeLang(name, "fr")
		translatedShop := translateShop(shop)
		shopSlug := slug.MakeLang(translatedShop, "fr")
		streetSlug := slug.MakeLang(tags["addr:street"], "fr")

		// Exceptions: skip foreign cities
		if citySlug == "s-heerenberg" {
			// 's-Heerenberg in the Netherlands
			return
		}
		// 1. content
		err := os.MkdirAll(region+"/content/cities/"+citySlug, 0755)
		if err != nil && !os.IsExist(err) {
			panic(err)
		}
		f := createFile(createIndexPath(region, "cities", citySlug))
		if f != nil {
			data := map[string]interface{}{
				"title":     city,
				"url":       "/" + citySlug + "/",
				"latitude":  roundFloat(n.Lat, 3),
				"longitude": roundFloat(n.Lon, 3),
			}
			writeData(f, indexTemplate, data)
		}
		err = os.MkdirAll(region+"/content/shops/"+shopSlug, 0755)
		if err != nil && !os.IsExist(err) {
			panic(err)
		}
		f = createFile(createIndexPath(region, "shops", shopSlug))
		if f != nil {
			data := map[string]interface{}{
				"title": translatedShop,
				"url":   "/" + shopSlug + "/",
				"icon":  getIcon(shop),
			}
			writeData(f, shopTemplate, data)
		}
		f, nameSlug = createElementFile(region, "content", citySlug, nameSlug, streetSlug, "md")
		if f != nil {
			data := map[string]interface{}{
				"title": strings.ReplaceAll(name, "\"", ""),
				"url":   "/" + citySlug + "/" + nameSlug + "/",
				"shop":  translatedShop,
			}
			writeData(f, mdTemplate, data)
		}

		// 2. data
		err = os.MkdirAll(region+"/data/cities/"+citySlug, 0755)
		if err != nil && !os.IsExist(err) {
			panic(err)
		}
		f, nameSlug = createElementFile(region, "data", citySlug, nameSlug, streetSlug, "yml")
		if f != nil {
			d := data{
				ID:           n.ID,
				Type:         "node",
				Latitude:     roundFloat(n.Lat, 5),
				Longitude:    roundFloat(n.Lon, 5),
				Postcode:     tags["addr:postcode"],
				City:         city,
				Street:       tags["addr:street"],
				Housenumber:  tags["addr:housenumber"],
				Phone:        tags["phone"],
				OpeningHours: tags["opening_hours"],
				Website:      tags["website"],
			}

			b, _ := yaml.Marshal(&d) // TODO: handle error
			f.Write(b)               // TODO: handle error
			f.Close()                // TODO: handle error
		}

	}
	// Cache necessary Nodes LatLon
	wID, found := wc.Load(n.ID)
	if found {
		nc.Store(wID, LatLon{n.Lat, n.Lon})
	}
}

func (d *dataPrepare) ReadWay(w gosmparse.Way) {
	tags := w.Tags
	city := tags["addr:city"]
	name := tags["name"]
	shop := tags["shop"]
	if city != "" && name != "" && shop != "" {
		// This loop looks ugly
		// but necessary because an individual NodeID might already be taken
		for i := 0; i < 9; i++ {
			_, found := wc.Load(w.NodeIDs[i])
			if found {
				continue
			} else {
				wc.Store(w.NodeIDs[i], w.ID)
				break
			}
		}
	}
}

func (d *dataHandler) ReadWay(w gosmparse.Way) {
	tags := w.Tags
	city := tags["addr:city"]
	name := tags["name"]
	shop := tags["shop"]
	if city != "" && name != "" && shop != "" {
		citySlug := slug.MakeLang(city, "fr")
		nameSlug := slug.MakeLang(name, "fr")
		translatedShop := translateShop(shop)
		shopSlug := slug.MakeLang(translatedShop, "fr")
		streetSlug := slug.MakeLang(tags["addr:street"], "fr")

		// Lookup first childNodes coords
		node, found := nc.Load(w.ID)
		if !found || node == nil {
			// Should never panic, just for safety
			panic(found)
		}
		// 1. content
		os.MkdirAll(region+"/content/cities/"+citySlug, 0755)
		f := createFile(createIndexPath(region, "cities", citySlug))
		if f != nil {
			data := map[string]interface{}{
				"title":     city,
				"url":       "/" + citySlug + "/",
				"latitude":  roundFloat(node.(LatLon).lat, 3),
				"longitude": roundFloat(node.(LatLon).lon, 3),
			}
			writeData(f, indexTemplate, data)
		}
		os.MkdirAll(region+"/content/shops/"+shopSlug, 0755)
		f = createFile(createIndexPath(region, "shops", shopSlug))
		if f != nil {
			data := map[string]interface{}{
				"title": translatedShop,
				"url":   "/" + shopSlug + "/",
				"icon":  getIcon(shop),
			}
			writeData(f, shopTemplate, data)
		}
		f, nameSlug = createElementFile(region, "content", citySlug, nameSlug, streetSlug, "md")
		if f != nil {
			data := map[string]interface{}{
				"title": strings.ReplaceAll(name, "\"", ""),
				"url":   "/" + citySlug + "/" + nameSlug + "/",
				"shop":  translatedShop,
			}
			writeData(f, mdTemplate, data)
		}

		// 2. data
		os.MkdirAll(region+"/data/cities/"+citySlug, 0755)
		f, nameSlug = createElementFile(region, "data", citySlug, nameSlug, streetSlug, "yml")
		if f != nil {
			d := data{
				ID:           w.ID,
				Type:         "node",
				Latitude:     roundFloat(node.(LatLon).lat, 5),
				Longitude:    roundFloat(node.(LatLon).lon, 5),
				Postcode:     tags["addr:postcode"],
				City:         city,
				Street:       tags["addr:street"],
				Housenumber:  tags["addr:housenumber"],
				Phone:        tags["phone"],
				OpeningHours: tags["opening_hours"],
				Website:      tags["website"],
			}

			b, _ := yaml.Marshal(&d) // TODO: handle error
			f.Write(b)               // TODO: handle error
			f.Close()                // TODO: handle error
		}
	}
}

func (d *dataPrepare) ReadRelation(r gosmparse.Relation) {}
func (d *dataHandler) ReadRelation(r gosmparse.Relation) {}

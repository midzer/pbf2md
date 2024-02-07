package main

import (
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
	"sync"
	"text/template"

	"github.com/gosimple/slug"
	"github.com/thomersch/gosmparse"
)

type LatLon struct {
	lat float64
	lon float64
}

func createIndexFile(citySlug string, city string, template *template.Template) {
	indexFile := "content/cities/" + citySlug + "/_index.md"
	if _, err := os.Stat(indexFile); os.IsNotExist(err) {
		f, err := os.Create(indexFile)
		if err != nil {
			fmt.Println(err)
			return
		}
		data := map[string]interface{}{
			"title": city,
			"url":   "/" + citySlug + "/",
		}
		if err = template.Execute(f, data); err != nil {
			panic(err)
		}
		f.Close()
	}
}

func createElementFile(citySlug string, nameSlug string, name string, template *template.Template) {
	elementFile := "content/cities/" + citySlug + "/" + nameSlug + ".md"
	if _, err := os.Stat(elementFile); !os.IsNotExist(err) {
		re := regexp.MustCompile(`\d+$`)
		i := 2
		nameSlug = fmt.Sprintf("%s-%d", nameSlug, i)
		for {
			elementFile = "content/cities/" + citySlug + "/" + nameSlug + ".md"
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
	}
	if err = template.Execute(f, data); err != nil {
		panic(err)
	}
	f.Close()
}

func createDataElementFile(citySlug string, nameSlug string, id int64, lat float64, lon float64, tags map[string]string, city string, template *template.Template) {
	elementFile := "data/cities/" + citySlug + "/" + nameSlug + ".yml"
	if _, err := os.Stat(elementFile); !os.IsNotExist(err) {
		re := regexp.MustCompile(`\d+$`)
		i := 2
		nameSlug = fmt.Sprintf("%s-%d", nameSlug, i)
		for {
			elementFile = "data/cities/" + citySlug + "/" + nameSlug + ".yml"
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

type dataPrepare struct {
	Nodes     *uint64
	Ways      *uint64
}

type dataHandler struct {
	Nodes     *uint64
	Ways      *uint64
}

func main() {
	regions := []string{"austria", "baden-wuerttemberg", "bayern", "brandenburg", "bremen", "hamburg", "hessen", "mecklenburg-vorpommern", "niedersachsen", "nordrhein-westfalen", "rheinland-pfalz", "saarland", "sachsen-anhalt", "sachsen", "schleswig-holstein", "switzerland", "thueringen"}
	for _, region := range regions {
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
}

// Globals
var (
	region string
	nc sync.Map
	wc sync.Map

	// Create templates
	indexTmpl = `---
title: {{ .title }}
url: {{ .url }}
---	`
	indexTemplate = template.Must(template.New("index").Parse(indexTmpl))
	mdTmpl = `---
title: "{{ .title }}"
url: {{ .url }}
---`
	mdTemplate = template.Must(template.New("markdown").Parse(mdTmpl))
	dataTmpl = `id: {{ .id }}
latitude: {{ .latitude }}
longitude: {{ .longitude }}
postcode: "{{ .postcode }}"
city: {{ .city }}
street: "{{ .street }}"
housenumber: {{ .housenumber }}
phone: "{{ .phone }}"
opening_hours: "{{ .opening_hours }}"
website: "{{ .website }}"`
	dataTemplate = template.Must(template.New("data").Parse(dataTmpl))
)

func (d *dataPrepare) ReadNode(n gosmparse.Node) {}

func (d *dataHandler) ReadNode(n gosmparse.Node) {
	tags := n.Tags
	city := tags["addr:city"]
	name := tags["name"]
	cuisine := tags["cuisine"]
	if city != "" && name != "" && (strings.Contains(cuisine, "turkish") || strings.Contains(cuisine, "kebab")) {
		citySlug := slug.MakeLang(city, "de")
		nameSlug := slug.MakeLang(name, "de")

		// 1. content
		os.MkdirAll("content/cities/"+citySlug, 0755)
		createIndexFile(citySlug, city, indexTemplate)
		createElementFile(citySlug, nameSlug, name, mdTemplate)

		// 2. data
		os.MkdirAll("data/cities/"+citySlug, 0755)
		createDataElementFile(citySlug, nameSlug, n.ID, n.Lat, n.Lon, tags, city, dataTemplate)
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
	cuisine := tags["cuisine"]
	if city != "" && name != "" && (strings.Contains(cuisine, "turkish") || strings.Contains(cuisine, "kebab")) {
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
	cuisine := tags["cuisine"]
	if city != "" && name != "" && (strings.Contains(cuisine, "turkish") || strings.Contains(cuisine, "kebab")) {
		citySlug := slug.MakeLang(city, "de")
		nameSlug := slug.MakeLang(name, "de")

		// Lookup first childNodes coords
		node, found := nc.Load(w.ID)
		if !found || node == nil {
			// Should never panic, just for safety
			panic(found)
		}

		// 1. content
		os.MkdirAll("content/cities/"+citySlug, 0755)
		createIndexFile(citySlug, city, indexTemplate)
		createElementFile(citySlug, nameSlug, name, mdTemplate)

		// 2. data
		os.MkdirAll("data/cities/"+citySlug, 0755)
		createDataElementFile(citySlug, nameSlug, w.ID, node.(LatLon).lat, node.(LatLon).lon, tags, city, dataTemplate)
	}
}

func (d *dataPrepare) ReadRelation(r gosmparse.Relation) {}
func (d *dataHandler) ReadRelation(r gosmparse.Relation) {}

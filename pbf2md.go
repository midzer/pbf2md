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

func createIndexFile(citySlug string, city string) {
	indexFile := "content/cities/" + citySlug + "/_index.md"
	if _, err := os.Stat(indexFile); os.IsNotExist(err) {
		f, err := os.Create(indexFile)
		if err != nil {
			fmt.Println(err)
			return
		}
		data := map[string]interface{} {
			"title": city,
			"url":  "/" + citySlug + "/",
		}
		indexTmpl := `---
title: {{ .title }}
url: {{ .url }}
---
`
	  indexTemplate := template.Must(template.New("index").Parse(indexTmpl))
		if err = indexTemplate.Execute(f, data); err != nil {
			panic(err)
		}
		f.Close()
	}
}

func createElementFile(citySlug string, nameSlug string, name string) {
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

func createDataElementFile(citySlug string, nameSlug string, id int64, lat float64, lon float64, tags map[string]string, city string) {
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
	dataTmpl := `id: {{ .id }}
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

func main() {
  f, err := os.Open("germany-latest.osm.pbf")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	d := osmpbf.NewDecoder(f)

	// Cache nodes, ways and cities
	m := make(map[int64]LatLon)
	n := make(map[int64]string)

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
				cuisine := tags["cuisine"]
				if city != "" && name != "" && (cuisine == "turkish" || cuisine == "kebab") {
					citySlug := slug.MakeLang(city, "de")
					nameSlug := slug.MakeLang(name, "de")

					// 1. content
					err = os.MkdirAll("content/cities/" + citySlug, 0755)
					createIndexFile(citySlug, city)
					createElementFile(citySlug, nameSlug, name)

					// 2. data
					err = os.MkdirAll("data/cities/" + citySlug, 0755)
					createDataElementFile(citySlug, nameSlug, v.ID, v.Lat, v.Lon, tags, city)
				}
				// Cache all Nodes LatLon
				m[v.ID] = LatLon{v.Lat, v.Lon}
				nc++
			case *osmpbf.Way:
				// Process Way v.
				tags := v.Tags
				city := tags["addr:city"]
				name := tags["name"]
				cuisine := tags["cuisine"]
				if city != "" && name != "" && n[v.ID] != name && (cuisine == "turkish" || cuisine == "kebab") {
					citySlug := slug.MakeLang(city, "de")
					nameSlug := slug.MakeLang(name, "de")

					// 1. content
					err = os.MkdirAll("content/cities/" + citySlug, 0755)
					createIndexFile(citySlug, city)
					createElementFile(citySlug, nameSlug, name)

					// 2. data
					err = os.MkdirAll("data/cities/" + citySlug, 0755)
					node := m[v.NodeIDs[0]] // Lookup coords of first childnode
					createDataElementFile(citySlug, nameSlug, v.ID, node.lat, node.lon, tags, city)

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

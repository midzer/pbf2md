package main

import (
	"fmt"
	"os"
	"io"
	"log"
	"runtime"
	"text/template"
	"github.com/qedus/osmpbf"
	"github.com/gosimple/slug"
)

func main() {
  f, err := os.Open("germany-latest.osm.pbf")
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
	// Create content and data directory
	err = os.MkdirAll("content", 0755)
	err = os.MkdirAll("content/cities", 0755)
	err = os.MkdirAll("data", 0755)
	err = os.MkdirAll("data/cities", 0755)

	// Index template
	const indexTmpl = `---
title: {{ .city }}
url: "{{ .url }}"
---
`
	indexTemplate := template.Must(template.New("index").Parse(indexTmpl))
	
	// Markdown template
	const mdTmpl = `---
title: {{ .name }}
url: "{{ .url }}"
---
`
	mdTemplate := template.Must(template.New("markdown").Parse(mdTmpl))
	
	// data template
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
				cuisine := tags["cuisine"]
				if cuisine == "turkish" || cuisine == "kebab" {
					city := tags["addr:city"]
					name := tags["name"]
					if city != "" && name != "" {
						citySlug := slug.MakeLang(city, "de")
						nameSlug := slug.MakeLang(name, "de")

						// 1. content
						// Create directory
						err = os.MkdirAll("content/cities/" + citySlug, 0755)

						// Create index file
						f, err := os.Create("content/cities/" + citySlug + "/_index.md")
						if err != nil {
							fmt.Println(err)
							return
						}
						data := map[string]interface{} {
							"city": city,
							"url":  "/" + citySlug + "/",
						}
						if err := indexTemplate.Execute(f, data); err != nil {
								panic(err)
						}
						f.Close()

						// Create element file
						f, err = os.Create("content/cities/" + citySlug + "/" + nameSlug + ".md")
						if err != nil {
							fmt.Println(err)
							return
						}
						data = map[string]interface{} {
							"name": name,
							"url":  "/" + citySlug + "/" + nameSlug + "/",
						}
						if err := mdTemplate.Execute(f, data); err != nil {
								panic(err)
						}
						f.Close()

						// 2. data
						// Create directory
						err = os.MkdirAll("data/cities/" + citySlug, 0755)

						// Create element file
						f, err = os.Create("data/cities/" + citySlug + "/" + nameSlug + ".yml")
						if err != nil {
							fmt.Println(err)
							return
						}
						data = map[string]interface{} {
							"id":            v.ID,
							"latitude":      v.Lat,
							"longitude":     v.Lon,
							"postcode":      tags["addr:postcode"],
							"city":          city,
							"street":        tags["addr:street"],
							"housenumber":   tags["addr:housenumber"],
							"phone":         tags["phone"],
							"opening_hours": tags["opening_hours"],
							"website":       tags["website"],
						}
						if err := dataTemplate.Execute(f, data); err != nil {
							panic(err)
						}
						f.Close()
					}
				}
				nc++
			case *osmpbf.Way:
				// Process Way v.
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

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
	err = os.MkdirAll("data", 0755)

	// Index template
	const indexTmpl = `---
title: {{ .city }}
---
`
	indexTemplate := template.Must(template.New("index").Parse(indexTmpl))
	
	// Markdown template
	const mdTmpl = `---
title: {{ .name }}
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
				if val, ok := tags["cuisine"]; ok {
					if val == "turkish" || val == "kebab" {
						if city, ok := tags["addr:city"]; ok {
							// 1. content
							// Create directory
							err = os.MkdirAll("content/" + slug.MakeLang(city, "de"), 0755)

							// Create index file
							f, err := os.Create("content/" + slug.MakeLang(city, "de") + "/_index.md")
							if err != nil {
								fmt.Println(err)
								return
							}
							data := map[string]interface{} {
								"city": tags["addr:city"],
							}
							if err := indexTemplate.Execute(f, data); err != nil {
									panic(err)
							}
							f.Close()

							// Create element file
							f, err = os.Create("content/" + slug.MakeLang(city, "de") + "/" + slug.MakeLang(tags["name"], "de") + ".md")
							if err != nil {
								fmt.Println(err)
								return
							}
							data = map[string]interface{} {
								"name":          tags["name"],
							}
							if err := mdTemplate.Execute(f, data); err != nil {
									panic(err)
							}
							f.Close()

							// 2. data
							// Create directory
							err = os.MkdirAll("data/" + slug.MakeLang(city, "de"), 0755)

							// Create element file
							f, err = os.Create("data/" + slug.MakeLang(city, "de") + "/" + slug.MakeLang(tags["name"], "de") + ".yml")
							if err != nil {
								fmt.Println(err)
								return
							}
							data = map[string]interface{} {
								"id":            v.ID,
								"latitude":      v.Lat,
								"longitude":     v.Lon,
								"postcode":      tags["addr:postcode"],
								"city":          tags["addr:city"],
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

# pbf2md

Parse OpenStreetMap *.pbf files and convert specific entries to markdown files

## Getting Started
* download and set up Go: see https://golang.org/doc/install
* clone this repository
* download required packages: `go get -v github.com/qedus/osmpbf github.com/gosimple/slug`
* build: `go build`
* download OSM data:
```bash
for prefix in "austria" "germany/baden-wuerttemberg" "germany/bayern" "germany/brandenburg" "germany/bremen" "germany/hamburg" "germany/hessen" "germany/mecklenburg-vorpommern" "germany/niedersachsen" "germany/nordrhein-westfalen" "germany/rheinland-pfalz" "germany/saarland" "germany/sachsen-anhalt" "germany/sachsen" "germany/schleswig-holstein" "switzerland" "germany/thueringen"; do
  wget "https://download.geofabrik.de/europe/$prefix-latest.osm.pbf";
done
```
* for a more simple start, you can also download a single sub-region file from https://download.geofabrik.de/index.html (make sure to get the .osm.pbf format) and then change the `regions` list in `pbf2md.go` source code to your downloaded .osm.pbf file, and re-run `go build`.

* run conversion: `./pbf2md`
  * the converted data will be put into a subdirectory for each region

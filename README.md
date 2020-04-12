# pbf2md

Parse OpenStreetMap *.pbf files and convert specific entries to markdown files

## Getting Started
* download and set up Go: see https://golang.org/doc/install
* clone this repository
* download required packages: `go get -v github.com/qedus/osmpbf github.com/gosimple/slug`
* build: `go build`
* download OSM data:
```bash
for prefix in "alabama" "alaska" "arizona" "arkansas" "california" "colorado" "connecticut" "delaware" "district-of-columbia" "florida" "georgia" "hawaii" "idaho" "illinois" "indiana" "iowa" "kansas" "kentucky" "louisiana" "maine" "maryland" "massachusetts" "michigan" "minnesota" "mississippi" "missouri" "montana" "nebraska" "nevada" "new-hampshire" "new-jersey" "new-mexico" "new-york" "north-carolina" "north-dakota" "ohio" "oklahoma" "oregon" "pennsylvania" "puerto-rico" "rhode-island" "south-carolina" "south-dakota" "tennessee" "texas" "utah" "vermont" "virginia" "washington" "west-virginia" "wisconsin" "wyoming"; do
  wget "https://download.geofabrik.de/north-america/us/$prefix-latest.osm.pbf"; sleep 10;
done
```
* for a more simple start, you can also download a single sub-region file from https://download.geofabrik.de/index.html (make sure to get the .osm.pbf format) and then change the `regions` list in `pbf2md.go` source code to your downloaded .osm.pbf file, and re-run `go build`.

* run conversion: `./pbf2md`
  * the converted data will be put into a subdirectory for each region

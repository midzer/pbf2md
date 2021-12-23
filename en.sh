# Build, move and deploy all english countries
# Copy file to your pbf2md folder
# Usage: ./en.sh DSTFOLDER (e.g. /countries/)

if [ $# -lt 1 ]; then
  echo 1>&2 "$0: not enough arguments"
  exit 2
elif [ $# -gt 1 ]; then
  echo 1>&2 "$0: too many arguments"
  exit 2
fi

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )"
LANG="en"

# Use proper branch
git checkout $LANG
go build

# Download and build single countries in

# Africa
for country in "ghana" "kenya" "liberia" "nigeria" "sierra-leone" "south-africa"; do
  wget "https://download.geofabrik.de/africa/$country-latest.osm.pbf"
  ./pbf2md $country
done

# Asia
for country in "india" "pakistan" "philippines"; do
  wget "https://download.geofabrik.de/asia/$country-latest.osm.pbf"
  ./pbf2md $country
done

# Australia-Oceania
for country in "australia" "new-zealand"; do
  wget "https://download.geofabrik.de/australia-oceania/$country-latest.osm.pbf"
  ./pbf2md $country
done

# Europe
for country in "ireland-and-northern-ireland"; do
  wget "https://download.geofabrik.de/europe/$country-latest.osm.pbf"
  ./pbf2md $country
done

# Update theme, move all countries and commit
for country in "ghana" "kenya" "liberia" "nigeria" "sierra-leone" "south-africa" "india" "pakistan" "philippines" "australia" "new-zealand" "ireland-and-northern-ireland"; do
  cd "$1/$LANG/$country" && git pull origin master && git submodule update --remote --merge && git commit -am "update theme"
  cd $DIR
  rm -rf "$1/$LANG/$country/content/cities"
  rm -rf "$1/$LANG/$country/content/shops"
  mv "$country/content/cities" "$1/$LANG/$country/content/cities"
  mv "$country/content/shops" "$1/$LANG/$country/content/shops"
  rm -rf "$1/$LANG/$country/data/cities"
  mv "$country/data/cities" "$1/$LANG/$country/data/cities"
  cd "$1/$LANG/$country" && git add -A && git commit -m "update content and data" && git push origin master
  cd $DIR
  rm "$country-latest.osm.pbf"
  rm -rf "$country/"
done

# Download, build and move special country regions in

# Great Britain
for region in "england" "scotland" "wales"; do
  wget "https://download.geofabrik.de/europe/great-britain/$region-latest.osm.pbf";
  ./pbf2md $region
  cd "$1/$LANG/great-britain/$region" && git pull origin master && git submodule update --remote --merge && git commit -am "update theme"
  cd $DIR
  rm -rf "$1/$LANG/great-britain/$region/content/cities"
  rm -rf "$1/$LANG/great-britain/$region/content/shops"
  mv "$region/content/cities" "$1/$LANG/great-britain/$region/content/cities"
  mv "$region/content/shops" "$1/$LANG/great-britain/$region/content/shops"
  rm -rf "$1/$LANG/great-britain/$region/data/cities"
  mv "$region/data/cities" "$1/$LANG/great-britain/$region/data/cities"
  cd "$1/$LANG/great-britain/$region" && git add -A && git commit -m "update content and data" && git push origin master
  cd $DIR
  rm "$region-latest.osm.pbf"
  rm -rf "$region/"
done

# Canada
for region in "alberta" "british-columbia" "manitoba" "new-brunswick" "newfoundland-and-labrador" "northwest-territories" "nova-scotia" "nunavut" "ontario" "prince-edward-island" "quebec" "saskatchewan" "yukon"; do
  wget "https://download.geofabrik.de/north-america/canada/$region-latest.osm.pbf"
  ./pbf2md $region
    cd "$1/$LANG/canada/$region" && git pull origin master && git submodule update --remote --merge && git commit -am "update theme"
  cd $DIR
  rm -rf "$1/$LANG/canada/$region/content/cities"
  rm -rf "$1/$LANG/canada/$region/content/shops"
  mv "$region/content/cities" "$1/$LANG/canada/$region/content/cities"
  mv "$region/content/shops" "$1/$LANG/canada/$region/content/shops"
  rm -rf "$1/$LANG/canada/$region/data/cities"
  mv "$region/data/cities" "$1/$LANG/canada/$region/data/cities"
  cd "$1/$LANG/canada/$region" && git add -A && git commit -m "update content and data" && git push origin master
  cd $DIR
  rm "$region-latest.osm.pbf"
  rm -rf "$region/"
done

# United States
for region in "alabama" "alaska" "arizona" "arkansas" "california" "colorado" "connecticut" "delaware" "district-of-columbia" "florida" "georgia" "hawaii" "idaho" "illinois" "indiana" "iowa" "kansas" "kentucky" "louisiana" "maine" "maryland" "massachusetts" "michigan" "minnesota" "mississippi" "missouri" "montana" "nebraska" "nevada" "new-hampshire" "new-jersey" "new-mexico" "new-york" "north-carolina" "north-dakota" "ohio" "oklahoma" "oregon" "pennsylvania" "puerto-rico" "rhode-island" "south-carolina" "south-dakota" "tennessee" "texas" "utah" "vermont" "virginia" "washington" "west-virginia" "wisconsin" "wyoming"; do
  wget "https://download.geofabrik.de/north-america/us/$region-latest.osm.pbf"
  ./pbf2md $region
  cd "$1/$LANG/us/$region" && git pull origin master && git submodule update --remote --merge && git commit -am "update theme"
  cd $DIR
  rm -rf "$1/$LANG/us/$region/content/cities"
  rm -rf "$1/$LANG/us/$region/content/shops"
  mv "$region/content/cities" "$1/$LANG/us/$region/content/cities"
  mv "$region/content/shops" "$1/$LANG/us/$region/content/shops"
  rm -rf "$1/$LANG/us/$region/data/cities"
  mv "$region/data/cities" "$1/$LANG/us/$region/data/cities"
  cd "$1/$LANG/us/$region" && git add -A && git commit -m "update content and data" && git push origin master
  cd $DIR
  rm "$region-latest.osm.pbf"
  rm -rf "$region/"
done

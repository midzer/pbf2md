# Build, move and deploy all french countries
# Copy file to your pbf2md folder
# Usage: ./fr.sh DSTFOLDER (e.g. /countries/)

if [ $# -lt 1 ]; then
  echo 1>&2 "$0: not enough arguments"
  exit 2
elif [ $# -gt 1 ]; then
  echo 1>&2 "$0: too many arguments"
  exit 2
fi

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )"
LANG="fr"

# Use proper branch
git checkout $LANG
go build

# Download and build single countries in

# Africa
for country in "benin" "burkina-faso" "congo-brazzaville" "congo-democratic-republic" "gabon" "guinea" "ivory-coast" "mali" "niger" "senegal-and-gambia" "togo"; do
  wget "https://download.geofabrik.de/africa/$country-latest.osm.pbf"
  ./pbf2md $country
done

# Europe
for country in "monaco"; do
  wget "https://download.geofabrik.de/europe/$country-latest.osm.pbf"
  ./pbf2md $country
done

# Update theme, move all countries and commit
for country in "benin" "burkina-faso" "congo-brazzaville" "congo-democratic-republic" "gabon" "guinea" "ivory-coast" "mali" "niger" "senegal-and-gambia" "togo" "monaco"; do
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

# France
for region in "alsace" "aquitaine" "auvergne" "basse-normandie" "bourgogne" "bretagne" "centre" "champagne-ardenne" "corse" "franche-comte" "guadeloupe" "guyane" "haute-normandie" "ile-de-france" "languedoc-roussillon" "limousin" "lorraine" "martinique" "mayotte" "midi-pyrenees" "nord-pas-de-calais" "pays-de-la-loire" "picardie" "poitou-charentes" "provence-alpes-cote-d-azur" "reunion" "rhone-alpes"; do
  wget "https://download.geofabrik.de/europe/france/$region-latest.osm.pbf";
  ./pbf2md $region
  cd "$1/$LANG/france/$region" && git pull origin master && git submodule update --remote --merge && git commit -am "update theme"
  cd $DIR
  rm -rf "$1/$LANG/france/$region/content/cities"
  rm -rf "$1/$LANG/france/$region/content/shops"
  mv "$region/content/cities" "$1/$LANG/france/$region/content/cities"
  mv "$region/content/shops" "$1/$LANG/france/$region/content/shops"
  rm -rf "$1/$LANG/france/$region/data/cities"
  mv "$region/data/cities" "$1/$LANG/france/$region/data/cities"
  cd "$1/$LANG/france/$region" && git add -A && git commit -m "update content and data" && git push origin master
  cd $DIR
  rm "$region-latest.osm.pbf"
  rm -rf "$region/"
done

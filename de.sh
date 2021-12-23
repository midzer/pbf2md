# Build, move and deploy all german countries
# Copy file to your pbf2md folder
# Usage: ./de.sh DSTFOLDER (e.g. /countries/)

if [ $# -lt 1 ]; then
  echo 1>&2 "$0: not enough arguments"
  exit 2
elif [ $# -gt 1 ]; then
  echo 1>&2 "$0: too many arguments"
  exit 2
fi

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )"
LANG="master"

# Use proper branch
git checkout $LANG
go build

# Download and build single countries in

# Europe
for country in "austria" "switzerland"; do
  wget "https://download.geofabrik.de/europe/$country-latest.osm.pbf"
  ./pbf2md $country
done

# Update theme, move all countries and commit
for country in "austria" "switzerland"; do
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

# Germany
for region in "baden-wuerttemberg" "bayern" "brandenburg" "bremen" "hamburg" "hessen" "mecklenburg-vorpommern" "niedersachsen" "nordrhein-westfalen" "rheinland-pfalz" "saarland" "sachsen-anhalt" "sachsen" "schleswig-holstein" "thueringen"; do
  wget "https://download.geofabrik.de/europe/germany/$region-latest.osm.pbf";
  ./pbf2md $region
  cd "$1/$LANG/germany/$region" && git pull origin master && git submodule update --remote --merge && git commit -am "update theme"
  cd $DIR
  rm -rf "$1/$LANG/germany/$region/content/cities"
  rm -rf "$1/$LANG/germany/$region/content/shops"
  mv "$region/content/cities" "$1/$LANG/germany/$region/content/cities"
  mv "$region/content/shops" "$1/$LANG/germany/$region/content/shops"
  rm -rf "$1/$LANG/germany/$region/data/cities"
  mv "$region/data/cities" "$1/$LANG/germany/$region/data/cities"
  cd "$1/$LANG/germany/$region" && git add -A && git commit -m "update content and data" && git push origin master
  cd $DIR
  rm "$region-latest.osm.pbf"
  rm -rf "$region/"
done

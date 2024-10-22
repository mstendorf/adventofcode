
mkdir "day$2"
cd "day$2"
echo "$1 $2"
curl https://adventofcode.com/$1/day/$2/input --cookie "session=53616c7465645f5f7dec293ad0e6b4079fd4cbe63fdddb16301b4a71f8bf7fc584368547f76e426ad0055cbe84a925f891d13c25d028aec24a13f057816d7f5f" > input
echo "package main\n\nfunc main () {\n}" >main.go

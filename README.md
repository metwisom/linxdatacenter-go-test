# linxdatacenter-go-test

Test for linxdatacenter

## Usage

```sh
cd app
go run . ./data/1.json      #just run for json
go run . ./data/1.csv       #just run for csv
go run . ./data/1.csv ","   #run with specifly comma for csv
go build -o app             #build
./app ./data/1.csv          #run builded app
```
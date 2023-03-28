package main

import (
	"log"

	"go-geo/shapefile/shp"
)

func main() {
	_, err := shp.ParseShapeFile("shapefile/shp/samples/tl_rd22_12001_addrfeat.shp")
	if err != nil {
		log.Fatalf("%v: failed to parse file", err)
	}

	_, err = shp.ParseShapeFile("shapefile/shp/samples/test_polygon_header.shp")
	if err != nil {
		log.Fatalf("%v: failed to parse file", err)
	}
}

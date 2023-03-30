package main

import (
	"log"

	"go-geo/shapefile/shp"
)

func main() {
	shapeFile, err := shp.ParseShapeFile("shapefile/shp/samples/tl_rd22_12001_addrfeat.shp")
	if err != nil {
		log.Fatalf("%v: failed to parse file", err)
	}

	log.Printf("Shapefile parsed in %s", shapeFile.ParseDuration())
	log.Printf("Parsed %d %ss", shapeFile.NumberOfRecords(), shapeFile.ShapeType())

	polygonFile, err := shp.ParseShapeFile("shapefile/shp/samples/test_polygon_header.shp")
	if err != nil {
		log.Fatalf("%v: failed to parse file", err)
	}

	log.Printf("Shapefile parsed in %s", polygonFile.ParseDuration())
	log.Printf("Parsed %d %ss", polygonFile.NumberOfRecords(), polygonFile.ShapeType())
}

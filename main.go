package main

import (
	"fmt"
	"log"
	"os"

	"go-geo/shapefile/shp"
)

func main() {
	s, err := shp.ParseShapeFile("shp/samples/tl_rd22_12001_addrfeat.shp")
	if err != nil {
		log.Fatalf("%v: failed to parse file", err)
	}

	_, err = fmt.Fprint(os.Stdout, s)
	if err != nil {
		log.Fatalf("failed to output parsed shapefile: %v", err)
	}
}

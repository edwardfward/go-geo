package main

import (
	"fmt"
	"go-shp/shapefile/shp"
	"log"
	"os"
)

func main() {
	s, err := shp.ParseShapeFile("shp/samples/tl_rd22_12001_addrfeat.shp")
	if err != nil {
		log.Fatalf("failed to parse shapefile: %v", err)
	}

	_, err = fmt.Fprint(os.Stdout, s)
	if err != nil {
		log.Fatalf("failed to output parsed shapefile: %v", err)
	}
}

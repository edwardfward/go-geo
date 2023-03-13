package main

import (
	"fmt"
	"go-shp/shp"
)

func main() {
	s := shp.ParseShapeFile("shp/samples/tl_rd22_12001_addrfeat.shp")
	fmt.Println(s.ShapeType())
}

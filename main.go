package main

import (
	"fmt"
	"go-shp/shp"
)

func main() {
	s := shp.ParseShapeFile("shp/samples/test_polygon_header.shp")
	fmt.Println(s.ShapeType())
}

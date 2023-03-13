package records

import (
	"go-shp/shp/shapes"
	"os"
)

type Record struct {
	id     int32
	shape  shapes.Shape
	offset int32
	length int32
}

type Records struct {
	records []Record
}

func ParseRecords(f *os.File) (Records, error) {
	return Records{}, nil
}

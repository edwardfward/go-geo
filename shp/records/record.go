package records

import (
	"fmt"
	"go-shp/shp/shapes"
)

type Record struct {
	id     int32
	shape  shapes.Shape
	offset int64
	length int32
}

type Records struct {
	records []Record
	shape   shapes.Shape
}

func (r *Records) Append(new Record) {
	r.records = append(r.records, new)
}

func (r *Records) NumberOfShapes() int32 {
	return int32(len(r.records))
}

func (r *Records) String() string {
	return fmt.Sprintf("Number of Shapes: %d", r.NumberOfShapes())
}

func NewRecord(id int32, shape shapes.Shape, offset int64, length int32) Record {
	return Record{id: id, shape: shape, offset: offset, length: length}
}

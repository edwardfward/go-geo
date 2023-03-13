package records

import (
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
}

func (r *Records) Append(new Record) {
	r.records = append(r.records, new)
}

func NewRecord(id int32, shape shapes.Shape, offset int64, length int32) Record {
	return Record{id: id, shape: shape, offset: offset, length: length}
}

package records

import (
	"fmt"
)

type Record interface {
	Parse(record []byte, header RecordHeader) error
	LengthBytes() int32
	RecordNumber() int32
}

const (
	WORDMULTIPLE = 2 // word represents two bytes in length
)

type ShapeType int32

const (
	NULLSHAPE ShapeType = 0
	POINT     ShapeType = 1
	// POLYLINE    ShapeType = 3
	// POLYGON     ShapeType = 5
	// MULTIPOINT  ShapeType = 8
	// POINTZ      ShapeType = 11
	// POLYLINEZ   ShapeType = 13
	// POLYGONZ    ShapeType = 15
	// MULTIPOINTZ ShapeType = 18
	// POINTM      ShapeType = 21
	// POLYLINEM   ShapeType = 23
	// POLYGONM    ShapeType = 25
	// MULTIPOINTM ShapeType = 28
	// MULTIPATCH  ShapeType = 31
)

// GetShapeType todo documentation.
func GetShapeType(value ShapeType) (Record, error) {
	switch value {
	case NULLSHAPE:
		return &NullShape{header: RecordHeader{recordNumber: 0, contentLength: 0},
			shape: NULLSHAPE}, nil
	case POINT:
		return &Point{header: RecordHeader{recordNumber: 0, contentLength: 0},
			x: 0, y: 0, shape: POINT}, nil
	//case POLYLINE:
	//	return &PolyLine{}, nil
	//case POLYGON:
	//	return &Polygon{}, nil
	//case MULTIPOINT:
	//	return &MultiPoint{}, nil
	//case POINTZ:
	//	return &PointZ{x: 0, y: 0, z: 0, m: 0}, nil
	//case POLYLINEZ:
	//	return &PolyLineZ{}, nil
	//case POLYGONZ:
	//	return &PolygonZ{}, nil
	//case MULTIPOINTZ:
	//	return &MultiPointZ{}, nil
	//case POINTM:
	//	return &PointM{}, nil
	//case POLYLINEM:
	//	return &PolyLineM{}, nil
	//case POLYGONM:
	//	return &PolygonM{}, nil
	//case MULTIPOINTM:
	//	return &MultiPointM{}, nil
	//case MULTIPATCH:
	//	return &MultiPatch{}, nil
	default:
		return &NullShape{header: RecordHeader{recordNumber: 0, contentLength: 0},
				shape: NULLSHAPE},
			fmt.Errorf(": unrecognized shapetype value: %v", value)
	}
}

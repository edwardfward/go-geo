package shapes

import (
	"encoding/binary"
	"fmt"
	"go-shp/utils"
	"log"
)

type ShapeType int32

const (
	NULLSHAPE   ShapeType = 0
	POINT       ShapeType = 1
	POLYLINE    ShapeType = 3
	POLYGON     ShapeType = 5
	MULTIPOINT  ShapeType = 8
	POINTZ      ShapeType = 11
	POLYLINEZ   ShapeType = 13
	POLYGONZ    ShapeType = 15
	MULTIPOINTZ ShapeType = 18
	POINTM      ShapeType = 21
	POLYLINEM   ShapeType = 23
	POLYGONM    ShapeType = 25
	MULTIPOINTM ShapeType = 28
	MULTIPATCH  ShapeType = 31
)

// Shape todo documentation
type Shape interface {
	Parse([]byte)
	Type() ShapeType
	New() Shape
	String() string
}

// BoundaryBox todo documentation
func NewBoundaryBox(box []byte) [2]Point {
	if len(box) != 32 {
		log.Fatalf("shape.go: error parsing boundary box incorrect bytes")
	}
	_xMin := box[0:8]
	_yMin := box[8:16]
	_xMax := box[16:24]
	_yMax := box[24:32]

	var xMin, yMin, xMax, yMax float64
	utils.ReadBinary(_xMin, binary.LittleEndian, &xMin)
	utils.ReadBinary(_yMin, binary.LittleEndian, &yMin)
	utils.ReadBinary(_xMax, binary.LittleEndian, &xMax)
	utils.ReadBinary(_yMax, binary.LittleEndian, &yMax)

	return [2]Point{NewPoint(xMin, yMin), NewPoint(xMax, yMax)}
}

// GetShapeType todo documentation
func GetShapeType(value ShapeType) (Shape, error) {
	switch value {
	case NULLSHAPE:
		return &NullShape{}, nil
	case POINT:
		return &Point{}, nil
	case POLYLINE:
		return &PolyLine{}, nil
	case POLYGON:
		return &Polygon{}, nil
	case MULTIPOINT:
		return &MultiPoint{}, nil
	case POINTZ:
		return &PointZ{}, nil
	case POLYLINEZ:
		return &PolyLineZ{}, nil
	case POLYGONZ:
		return &PolygonZ{}, nil
	case MULTIPOINTZ:
		return &MultiPointZ{}, nil
	case POINTM:
		return &PointM{}, nil
	case POLYLINEM:
		return &PolyLineM{}, nil
	case POLYGONM:
		return &PolygonM{}, nil
	case MULTIPOINTM:
		return &MultiPointM{}, nil
	case MULTIPATCH:
		return &MultiPatch{}, nil
	default:
		return &NullShape{}, fmt.Errorf(": unrecognized shapetype value: %v", value)

	}
}

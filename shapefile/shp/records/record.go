package records

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
)

type Record interface {
	Parse(record []byte, header RecordHeader) error
	LengthBytes() int32
	RecordNumber() int32
}

const (
	WORDMULTIPLE = 2 // word represents two bytes in length
	INT32LENGTH  = 4
)

type ShapeType int32

const (
	NULLSHAPE ShapeType = 0
	POINT     ShapeType = 1
	POLYLINE  ShapeType = 3
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
	// MULTIPATCH  ShapeType = 31.
)

// GetShapeType todo documentation.
func GetShapeType(value ShapeType) (Record, error) {
	var shape Record

	switch value {
	case NULLSHAPE:
		result := EmptyNullShape()
		shape = &result
	case POINT:
		result := EmptyPoint()
		shape = &result
	case POLYLINE:
		result := EmptyPolyLine()
		shape = &result
	// case POLYGON:
	//	return &Polygon{}, nil
	// case MULTIPOINT:
	//	return &MultiPoint{}, nil
	// case POINTZ:
	//	return &PointZ{x: 0, y: 0, z: 0, m: 0}, nil
	// case POLYLINEZ:
	//	return &PolyLineZ{}, nil
	// case POLYGONZ:
	//	return &PolygonZ{}, nil
	// case MULTIPOINTZ:
	//	return &MultiPointZ{}, nil
	// case POINTM:
	//	return &PointM{}, nil
	// case POLYLINEM:
	//	return &PolyLineM{}, nil
	// case POLYGONM:
	//	return &PolygonM{}, nil
	// case MULTIPOINTM:
	//	return &MultiPointM{}, nil
	// case MULTIPATCH:
	//	return &MultiPatch{}, nil
	default:
		result := EmptyNullShape()
		shape = &result

		unrecognizedShapeError := errors.New("unrecognized shape type value")

		return shape, fmt.Errorf("%w: %d", unrecognizedShapeError, value)
	}

	return shape, nil
}

func ParseParts(parts []byte, numParts int32) ([]int32, error) {
	// check right number of bytes received to parts number of parts
	if len(parts) != int(numParts*INT32LENGTH) {
		return nil, errors.New("incorrect number of bytes received to parse " +
			"number of parts requested")
	}

	partsArray := make([]int32, numParts)

	for index := 0; index < int(numParts); index++ {
		var part int32

		err := binary.Read(bytes.NewReader(parts[index*INT32LENGTH:index*INT32LENGTH+INT32LENGTH]),
			binary.LittleEndian, &part)
		if err != nil {
			return nil, err
		}

		partsArray[index] = part
	}

	return partsArray, nil
}

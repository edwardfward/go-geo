package records

import (
	"errors"
	"fmt"
)

type ShapeType int32

const (
	NULLSHAPE ShapeType = 0
	POINT     ShapeType = 1
	POLYLINE  ShapeType = 3
	POLYGON   ShapeType = 5
	// MULTIPOINT ShapeType = 8
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
		result := EmptyPointRecord()
		shape = &result
	case POLYLINE:
		result := EmptyPolyLine()
		shape = &result
	case POLYGON:
		result := EmptyPolygon()
		shape = &result
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

func (s *ShapeType) String() string {
	switch *s {
	case NULLSHAPE:
		return "Nullshape"
	case POINT:
		return "Point"
	case POLYLINE:
		return "PolyLine"
	case POLYGON:
		return "Polygon"
	default:
		return "Nullshape"
	}
}

package shapes

import (
	"bytes"
	"encoding/binary"
	"errors"
	"log"
)

type Shape interface {
	ParseShape([]byte)
	GetShapeType() int32
	String() string
	Copy() Shape
}

func ReadBinary[T int32 | float64](b []byte, order binary.ByteOrder, target *T) {
	err := binary.Read(bytes.NewReader(b), order, target)
	if err != nil {
		log.Fatalf("error reading binary on %v: %v", target, err)
	}
}
func GetShapeType(value int32) (Shape, error) {
	switch value {
	case 0:
		return &NullShape{}, nil
	case 1:
		return &Point{}, nil
	case 3:
		return &PolyLine{}, nil
	case 5:
		return &Polygon{}, nil
	case 8:
		return &MultiPoint{}, nil
	case 11:
		return &PointZ{}, nil
	case 13:
		return &PolyLineZ{}, nil
	case 15:
		return &PolygonZ{}, nil
	case 18:
		return &MultiPointZ{}, nil
	case 21:
		return &PointM{}, nil
	case 23:
		return &PolyLineM{}, nil
	case 25:
		return &PolygonM{}, nil
	case 28:
		return &MultiPointM{}, nil
		// case 31: [todo] implement MultiPatchShape
		// return MultiPatchShape, nil
	}
	return &NullShape{}, errors.New("not a valid shape value")
}

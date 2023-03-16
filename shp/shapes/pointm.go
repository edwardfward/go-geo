package shapes

import (
	"encoding/binary"
	"go-shp/utils"
)

type PointM struct {
	X float64 // x coordinate
	Y float64 // y coordinate
	M float64 // measure
}

func (p *PointM) Parse(r []byte) {
	_shape := r[0:4] // int32 little endian
	_x := r[4:12]    // x float64 little endian
	_y := r[12:20]   // y float little endian
	_m := r[20:28]   // measure float64 little endian

	var shapeType ShapeType
	var x, y, m float64

	utils.ReadBinary(_shape, binary.LittleEndian, &shapeType)
	utils.ReadBinary(_x, binary.LittleEndian, &x)
	utils.ReadBinary(_y, binary.LittleEndian, &y)
	utils.ReadBinary(_m, binary.LittleEndian, &m)

	p.X, p.Y, p.M = x, y, m
}

func (p *PointM) Type() ShapeType {
	return POINTM
}

func (p *PointM) String() string {
	return "PointM"
}

func (p *PointM) New() Shape {
	return new(PointM)
}

func NewPointM(p []byte) PointM {
	return PointM{}
}

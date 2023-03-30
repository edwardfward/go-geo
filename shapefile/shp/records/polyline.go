package records

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type PolyLine struct {
	shape     ShapeType
	box       BoundaryBox
	numParts  int32
	numPoints int32
	parts     []int32
	points    []Point
}

func (p *PolyLine) Parse(record []byte) error {
	// create temporary struct to determine number of parts and points
	// to determine array size before parsing complete point
	initial := struct {
		Shape     ShapeType   // int32 4-bytes
		Box       BoundaryBox // 4 * float64 32-bytes
		NumParts  int32       // 4-bytes
		NumPoints int32       // 4-bytes
	}{}

	// parse shape, box, numparts, and numpoints
	if err := binary.Read(bytes.NewReader(record[:44]),
		binary.LittleEndian, &initial); err != nil {
		return fmt.Errorf("error parsing record information: %w", err)
	}

	// check numparts and numpoints positive
	if initial.NumParts < 0 || initial.NumPoints < 0 {
		return fmt.Errorf("error parsing number of points or parts: cannot have negative " +
			"number of points or parts")
	}

	p.shape, p.box = initial.Shape, initial.Box
	p.numParts, p.numPoints = initial.NumParts, initial.NumPoints

	parts := make([]int32, p.numParts)
	points := make([]Point, p.numPoints)

	// parse polyline part array
	if err := binary.Read(bytes.NewReader(record[44:44+int(p.numParts)*4]),
		binary.LittleEndian, &parts); err != nil {
		return fmt.Errorf("error parsing polyline part array: %w", err)
	}

	// parse polyline point array
	if err := binary.Read(bytes.NewReader(record[44+int(p.numParts)*4:]),
		binary.LittleEndian, &points); err != nil {
		return fmt.Errorf("error parsing polyline point array: %w", err)
	}

	// assign parts and points arrays
	p.parts, p.points = parts, points

	return nil
}

func EmptyPolyLine() PolyLine {
	return PolyLine{shape: POLYLINE,
		box:       EmptyBoundaryBox(),
		numParts:  0,
		numPoints: 0,
		parts:     nil,
		points:    nil}
}

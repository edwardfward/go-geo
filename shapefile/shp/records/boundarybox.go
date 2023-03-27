package records

import (
	"fmt"
)

type BoundaryBox struct {
	points [2]Point
}

const (
	BOUNDARYBOXLENGTH = 32
)

func NewBoundaryBox(box []byte) (BoundaryBox, error) {
	// check to ensure length is 32, four float64
	if len(box) != BOUNDARYBOXLENGTH {
		return BoundaryBox{},
			fmt.Errorf("boundary box parse error: incorrect number of bytes %d",
				len(box))
	}

	boundaryBox := BoundaryBox{points: [2]Point{{x: 0, y: 0}, {x: 0, y: 0}}}

	min, err := ParsePoint(box[0:16])
	if err != nil {
		return boundaryBox, fmt.Errorf("error parsing boundary box minimum point: %w",
			err)
	}

	max, err := ParsePoint(box[16:32])
	if err != nil {
		return boundaryBox, fmt.Errorf("error parsing boundary box maximum point: %w",
			err)
	}

	boundaryBox.points[0] = min
	boundaryBox.points[1] = max

	return boundaryBox, nil
}

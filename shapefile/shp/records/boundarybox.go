package records

import (
	"fmt"
)

// BoundaryBox contains a minimum point (xMin, yMin) and
// maximum point (xMax, yMax) representing the bounds of
// all the shapes contained in a record or shapefile.
type BoundaryBox struct {
	points [2]Point
}

const (
	BOUNDARYBOXLENGTH = 32
)

// ParseBoundaryBox returns a shapefile boundary box. Box must be a 32-byte array
// or slice [xMin float64, yMin float64, xMax float64, yMax float64].
func ParseBoundaryBox(box []byte) (BoundaryBox, error) {
	// check to ensure length is 32, four float64
	if len(box) != BOUNDARYBOXLENGTH {
		return EmptyBoundaryBox(),
			fmt.Errorf("boundary box parse error: incorrect number of bytes %d",
				len(box))
	}

	boundaryBox := EmptyBoundaryBox()

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

// EmptyBoundaryBox returns an empty or default boundary box.
func EmptyBoundaryBox() BoundaryBox {
	return BoundaryBox{points: [2]Point{EmptyPoint(), EmptyPoint()}}
}

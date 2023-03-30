package records

// BoundaryBox contains a minimum point (xMin, yMin) and
// maximum point (xMax, yMax) representing the bounds of
// all the shapes contained in a record or shapefile.
type BoundaryBox struct {
	Points [2]Point
}

// EmptyBoundaryBox returns an empty or default boundary box.
func EmptyBoundaryBox() BoundaryBox {
	return BoundaryBox{Points: [2]Point{EmptyPoint(), EmptyPoint()}}
}

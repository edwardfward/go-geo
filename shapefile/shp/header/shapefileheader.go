package header

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"go-geo/shapefile/shp/records"
)

const (
	FILECODE              int32 = 9994
	VERSION               int32 = 1000
	SHAPEFILEHEADERLENGTH int   = 100
)

type ShapeFileHeader struct {
	fileCode   int32               // 9994 big endian
	fileLength int32               // length in 16-bit (2 byte) words big endian
	version    int32               // 1000 little endian
	shape      records.ShapeType   // int32 little endian
	box        records.BoundaryBox // boundary box (min, max)
	zRange     [2]float64          // [min, max] z
	mRange     [2]float64          // [min, max] measure
}

func (h *ShapeFileHeader) Parse(header []byte) error {
	// check to ensure header is proper length
	if len(header) != SHAPEFILEHEADERLENGTH {
		return fmt.Errorf("error parsing shapefile: shape file headers need to " +
			"be 100 bytes")
	}

	// parse header file code and file length
	fileInfo := struct {
		FileCode   int32
		_          int32
		_          int32
		_          int32
		_          int32
		_          int32
		FileLength int32
	}{}

	if err := binary.Read(bytes.NewReader(header[0:28]), binary.BigEndian, &fileInfo); err != nil {
		return fmt.Errorf("error parsing shapefile file code and file length: %w", err)
	}

	// parse rest of shape file header
	rest := struct {
		Version int32
		Shape   records.ShapeType
		Box     records.BoundaryBox
		ZBox    [2]float64
		MBox    [2]float64
	}{}

	if err := binary.Read(bytes.NewReader(header[28:100]), binary.LittleEndian, &rest); err != nil {
		return fmt.Errorf("error parsing shapefile version, type, bounding box, etc: %w", err)
	}

	h.fileCode, h.fileLength = fileInfo.FileCode, fileInfo.FileLength
	h.version, h.shape, h.box = rest.Version, rest.Shape, rest.Box
	h.zRange, h.mRange = rest.ZBox, rest.MBox

	return nil
}

func (h *ShapeFileHeader) ShapeType() records.ShapeType {
	return h.shape
}

func EmptyShapeFileHeader() ShapeFileHeader {
	return ShapeFileHeader{
		fileCode:   FILECODE,
		fileLength: 0,
		version:    VERSION,
		shape:      records.NULLSHAPE,
		box:        records.EmptyBoundaryBox(),
		zRange:     [2]float64{0.0, 0.0},
		mRange:     [2]float64{0.0, 0.0},
	}
}

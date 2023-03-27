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
		return fmt.Errorf(" expected header of %d bytes, received %d bytes",
			len(header), SHAPEFILEHEADERLENGTH)
	}

	// start parsing the header
	shapeHeader := new(ShapeFileHeader)

	// parse and check file code equals 9994
	err := binary.Read(bytes.NewReader(header[0:4]),
		binary.BigEndian, &shapeHeader.fileCode)
	if err != nil {
		return fmt.Errorf("unable to parse file code")
	}

	if shapeHeader.fileCode != FILECODE {
		return fmt.Errorf("parsed file code (%d) not equal to %d",
			shapeHeader.fileCode, FILECODE)
	}

	// parse file length (16-bit words)
	err = binary.Read(bytes.NewReader(header[24:28]),
		binary.BigEndian, &shapeHeader.fileLength)
	if err != nil {
		return fmt.Errorf("unable to parse file length: %v", err)
	}

	// parse and check version
	err = binary.Read(bytes.NewReader(header[28:32]),
		binary.LittleEndian, &shapeHeader.version)
	if err != nil {
		return fmt.Errorf("unable to parse file version: %v", err)
	}

	if shapeHeader.version != VERSION {
		return fmt.Errorf("parsed version (%d) not equal to %d",
			shapeHeader.version, VERSION)
	}

	// parse and check ShapeType
	err = binary.Read(bytes.NewReader(header[32:36]),
		binary.LittleEndian, &shapeHeader.shape)
	if err != nil {
		return fmt.Errorf("unable to parse shape type: %v",
			err)
	}
	// check if valid shape type
	if _, err := records.GetShapeType(shapeHeader.shape); err != nil {
		return fmt.Errorf("invalid shape type received: %v", err)
	}

	// parse boundary box
	shapeHeader.box, err = records.ParseBoundaryBox(header[36:68])
	if err != nil {
		return fmt.Errorf("%w: unable to parse boundary box", err)
	}

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
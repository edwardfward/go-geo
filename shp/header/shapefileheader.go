package header

import (
	"encoding/binary"
	"go-shp/shp/shapes"
	"go-shp/utils"
	"log"
	"os"
)

type ShapeFileHeader struct {
	fileCode   int32           // 9994
	fileLength int32           // total length of the file in bytes
	version    int32           // 1000
	shape      shapes.Shape    // only one shape type per shapefile
	box        [2]shapes.Point // boundary box (min, max)
	zRange     [2]float64      // [min, max] z
	mRange     [2]float64      // [min, max] measure
}

func (h *ShapeFileHeader) ShapeType() string {
	return h.shape.String()
}

func (h *ShapeFileHeader) NewShape() shapes.Shape {
	return h.shape.New() // [todo] change to New() for Shape interface
}

func Parse(f *os.File) (ShapeFileHeader, error) {
	// shapefile headers are 100 bytes
	headerBytes := make([]byte, 100)
	_, err := f.Read(headerBytes)
	if err != nil {
		log.Fatalf("unable to parse shapefile header: %v", err)
	}

	// create shapefile header
	shapeFileHeader := ShapeFileHeader{}

	// header fields
	_fileCode := headerBytes[0:4]     // int32 big endian
	_fileLength := headerBytes[24:28] // int32 big endian, length is 16-bit words (2-bytes)
	_version := headerBytes[28:32]    // int32 little endian
	_shape := headerBytes[32:36]      // int32 little endian
	_box := headerBytes[36:68]        // boundary box
	_zMin := headerBytes[68:76]       // float64 little endian
	_zMax := headerBytes[76:84]       // float64 little endian
	_mMin := headerBytes[84:92]       // float64 little endian
	_mMax := headerBytes[92:100]      // float64 little endian

	var shapeType shapes.ShapeType
	var zMin, zMax, mMin, mMax float64

	// [todo] try to find a more elegant way to parse the header
	utils.ReadBinary(_fileCode, binary.BigEndian, &shapeFileHeader.fileCode)
	utils.ReadBinary(_fileLength, binary.BigEndian, &shapeFileHeader.fileLength)
	utils.ReadBinary(_version, binary.LittleEndian, &shapeFileHeader.version)
	utils.ReadBinary(_shape, binary.LittleEndian, &shapeType)

	// z and m min and max todo potential to function
	utils.ReadBinary(_zMin, binary.LittleEndian, &zMin)
	utils.ReadBinary(_zMax, binary.LittleEndian, &zMax)
	utils.ReadBinary(_mMin, binary.LittleEndian, &mMin)
	utils.ReadBinary(_mMax, binary.LittleEndian, &mMax)

	shapeFileHeader.shape, err = shapes.GetShapeType(shapeType)
	if err != nil {
		log.Fatalf("unrecognized shape type %d: %v", shapeType, err)
	}

	// construct boundary box
	shapeFileHeader.box = shapes.NewBoundaryBox(_box)
	shapeFileHeader.zRange = [2]float64{zMin, zMax}
	shapeFileHeader.mRange = [2]float64{mMin, mMax}

	return shapeFileHeader, nil
}

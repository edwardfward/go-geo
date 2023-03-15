package header

import (
	"bytes"
	"encoding/binary"
	"go-shp/shp/shapes"
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

func (h *ShapeFileHeader) BoundingBox() (shapes.Point, shapes.Point) {
	return h.box[0], h.box[1]
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
	fileCode := headerBytes[0:4]     // int32 big endian
	fileLength := headerBytes[24:28] // int32 big endian, length is 16-bit words (2-bytes)
	version := headerBytes[28:32]    // int32 little endian
	shapeType := headerBytes[32:36]  // int32 little endian
	xMinSlice := headerBytes[36:44]  // float64 little endian
	yMinSlice := headerBytes[44:52]  // float64 little endian
	xMaxSlice := headerBytes[52:60]  // float 64 little endian
	yMaxSlice := headerBytes[60:68]  // float64 little endian
	zMinSlice := headerBytes[68:76]  // float64 little endian
	zMaxSlice := headerBytes[76:84]  // float64 little endian
	mMinSlice := headerBytes[84:92]  // float64 little endian
	mMaxSlice := headerBytes[92:100] // float64 little endian

	var shapeTypeInt int32
	var xMin, xMax, yMin, yMax, zMin, zMax, mMin, mMax float64

	// [todo] try to find a more elegant way to parse the header
	err = binary.Read(bytes.NewReader(fileCode), binary.BigEndian, &shapeFileHeader.fileCode)
	err = binary.Read(bytes.NewReader(fileLength), binary.BigEndian, &shapeFileHeader.fileLength)
	err = binary.Read(bytes.NewReader(version), binary.LittleEndian, &shapeFileHeader.version)
	err = binary.Read(bytes.NewReader(shapeType), binary.LittleEndian, &shapeTypeInt)
	err = binary.Read(bytes.NewReader(xMinSlice), binary.LittleEndian, &xMin)
	err = binary.Read(bytes.NewReader(yMinSlice), binary.LittleEndian, &yMin)
	err = binary.Read(bytes.NewReader(xMaxSlice), binary.LittleEndian, &xMax)
	err = binary.Read(bytes.NewReader(yMaxSlice), binary.LittleEndian, &yMax)
	err = binary.Read(bytes.NewReader(zMinSlice), binary.LittleEndian, &zMin)
	err = binary.Read(bytes.NewReader(zMaxSlice), binary.LittleEndian, &zMax)
	err = binary.Read(bytes.NewReader(mMinSlice), binary.LittleEndian, &mMin)
	err = binary.Read(bytes.NewReader(mMaxSlice), binary.LittleEndian, &mMax)

	shapeFileHeader.shape, err = shapes.GetShapeType(shapeTypeInt)
	if err != nil {
		log.Fatalf("unrecognized shape type %d: %v", shapeTypeInt, err)
	}

	// construct boundary box
	shapeFileHeader.box[0] = shapes.NewPoint(xMin, yMin)
	shapeFileHeader.box[1] = shapes.NewPoint(xMax, yMax)
	shapeFileHeader.zRange[0], shapeFileHeader.zRange[1] = zMin, zMax
	shapeFileHeader.mRange[0], shapeFileHeader.mRange[0] = mMin, mMax

	return shapeFileHeader, nil
}

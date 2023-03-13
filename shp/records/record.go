package records

import (
	"encoding/binary"
	"go-shp/shp/shapes"
	"go-shp/utils"
)

type Record struct {
	recordNumber  int32
	contentLength int32
	shape         shapes.Shape
	offset        int64
}

func ReadRecordHeader(r []byte) (recordNumber int32, contentLength int32) {
	recordNumberSlice := r[0:4]  // bytes 0-3 of record header
	contentLengthSlice := r[4:8] // bytes 3-7 of record header

	var number, length int32
	utils.ReadBinary(recordNumberSlice, binary.BigEndian, &number)
	utils.ReadBinary(contentLengthSlice, binary.BigEndian, &length)

	return number, length
}

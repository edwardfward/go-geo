package records

import (
	"encoding/binary"
	"fmt"
	"go-shp/shp/shapes"
	"go-shp/utils"
	"log"
	"os"
)

type Records struct {
	records []Record
	shape   shapes.Shape
}

func (r *Records) Append(new Record) {
	r.records = append(r.records, new)
}

func (r *Records) NumberOfShapes() int32 {
	return int32(len(r.records))
}

func (r *Records) String() string {
	return fmt.Sprintf("Number of Shapes: %d", r.NumberOfShapes())
}

func ParseRecords(f *os.File) Records {
	records := Records{}
	for {
		record := Record{}
		// parse record header 8-bytes
		b := make([]byte, 8)
		n, err := f.Read(b)
		if n == 0 { // EOF
			break
		}
		if err != nil {
			log.Fatalf("unable to read record: %v", err)
		}
		number, length := ReadRecordHeader(b)
		record.contentLength = length
		record.recordNumber = number

		// parse shape type
		b = make([]byte, length*2)
		n, err = f.Read(b) // content length is 16-bit words (2 bytes)
		if n == 0 || err != nil {
			log.Fatalf("no shape data in record")
		}
		var shapeType int32
		utils.ReadBinary(b[0:4], binary.LittleEndian, &shapeType)
		shape, err := shapes.GetShapeType(shapeType)
		shape.Parse(b)
		record.shape = shape
		records.Append(record)
	}

	return records
}

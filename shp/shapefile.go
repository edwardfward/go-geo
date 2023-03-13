package shp

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"go-shp/shp/header"
	"go-shp/shp/records"
	"log"
	"os"
	"time"
)

type ShapeFile struct {
	header      header.ShapeFileHeader
	records     records.Records
	sizeBytes   int64
	timeToParse time.Duration
}

func (s *ShapeFile) String() string {
	return fmt.Sprintf("Shapefile contains %d shapes",
		s.records.NumberOfShapes())
}

func ParseShapeFile(filePath string) *ShapeFile {
	// [todo] check for *.shp extension
	// open shapefile
	f, err := os.Open(filePath)
	defer f.Close()

	if err != nil {
		log.Fatalf("unable to open shapefile %s: %v", filePath, err)
	}
	shapeFile := &ShapeFile{}

	shapeFile.header, err = header.ParseHeader(f)
	shapeFile.ParseRecords(f)
	return shapeFile
}

func (s *ShapeFile) ShapeType() string {
	return s.header.ShapeType()
}

func (s *ShapeFile) ParseRecords(f *os.File) {
	var offset int64 = 100 // start reading records at 100 bytes
	s.records = records.Records{}

	// iterate through records
	for {
		// read and parse record header
		b := make([]byte, 8)
		n, err := f.Read(b)
		//
		if n == 0 {
			log.Println("finished parsing records")
			break
		}
		if err != nil {
			log.Fatalf("shapefile.go: unable to parse header for record: %v", err)
		}

		var recordNumber int32
		var recordContentLength int32

		err = binary.Read(bytes.NewReader(b[0:4]), binary.BigEndian, &recordNumber)
		err = binary.Read(bytes.NewReader(b[4:8]), binary.BigEndian, &recordContentLength)

		// read record
		b = make([]byte, recordContentLength*2)
		_, err = f.Read(b) // read record
		if err != nil {
			log.Fatalf("unable to parse record type %v at record %d: %v", s.ShapeType(),
				recordNumber, err)
		}

		// parse record
		shape := s.header.NewShape()
		shape.ParseShape(b)

		// create record
		r := records.NewRecord(recordNumber, shape, offset, recordContentLength)

		// todo need to parse byte slice
		s.records.Append(r)

		offset += int64(8 + recordContentLength*2)
	}
}

package shp

import (
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

	defer func() {
		if err := f.Close(); err != nil {
			log.Fatalf("file failed to close: %v", err)
		}
	}()

	if err != nil {
		log.Fatalf("unable to open shapefile %s: %v", filePath, err)
	}
	shapeFile := &ShapeFile{}
	shapeFile.header, err = header.ParseHeader(f)
	shapeFile.records = records.ParseRecords(f)

	return shapeFile
}

func (s *ShapeFile) ShapeType() string {
	return s.header.ShapeType()
}

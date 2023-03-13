package shp

import (
	"go-shp/shp/header"
	"go-shp/shp/records"
	"log"
	"os"
)

type ShapeFile struct {
	header  header.ShapeFileHeader
	records records.Records
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
	shapeFile.records, err = records.ParseRecords(f)
	return shapeFile
}

func (s *ShapeFile) ShapeType() string {
	return s.header.ShapeType()
}

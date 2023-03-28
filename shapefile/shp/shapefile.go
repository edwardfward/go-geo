package shp

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"go-geo/shapefile/shp/header"
	"go-geo/shapefile/shp/records"
)

// ShapeFile contains information and data from ESRI shapefile.
type ShapeFile struct {
	header      header.ShapeFileHeader
	records     []records.Record
	sizeBytes   int64
	timeToParse time.Duration
}

// ParseShapeFile parses a raw ESRI shapefile.
func ParseShapeFile(filePath string) (*ShapeFile, error) {
	// open shapefile
	file, err := os.Open(filePath)
	if err != nil {
		fileOpenError := errors.New("error opening file")

		return nil, fmt.Errorf("%w: %s", fileOpenError, file.Name())
	}

	defer func() {
		if err := file.Close(); err != nil {
			log.Printf("%s failed to close after parsing: %v", file.Name(), err)
		}
	}()

	shapeFile := EmptyShapeFile()
	err = shapeFile.Parse(file)

	if err != nil {
		return &shapeFile, fmt.Errorf("%w: [shapefile.go] unable to parse file", err)
	}

	return &shapeFile, nil
}

func EmptyShapeFile() ShapeFile {
	return ShapeFile{
		header:      header.EmptyShapeFileHeader(),
		records:     nil,
		sizeBytes:   0,
		timeToParse: time.Nanosecond * 0,
	}
}

// Parse takes a file binary and parses a shapefile.
func (s *ShapeFile) Parse(file *os.File) error {
	// parse shapefile header
	headerBytes := make([]byte, header.SHAPEFILEHEADERLENGTH)

	n, err := file.Read(headerBytes)
	if n != header.SHAPEFILEHEADERLENGTH || err != nil {
		headerParseError := errors.New("unable to parse shapefile header")

		return fmt.Errorf("%w: error reading shapefile header", headerParseError)
	}

	s.header = header.EmptyShapeFileHeader()

	log.Printf("header length: %d\n", len(headerBytes))

	err = s.header.Parse(headerBytes)
	if err != nil {
		return fmt.Errorf("%w: error parsing shapefile header", err)
	}

	// parse shapefile records
	s.records, err = records.ParseRecords(file, s.header.ShapeType())
	if err != nil {
		return fmt.Errorf("%w: unable to parse shapefile records", err)
	}

	return nil
}

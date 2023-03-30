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
	if err := shapeFile.Parse(file); err != nil {
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

	if _, err := file.Read(headerBytes); err != nil {
		return fmt.Errorf("error reading shape file header bytes: %w", err)
	}

	s.header = header.EmptyShapeFileHeader()

	if err := s.header.Parse(headerBytes); err != nil {
		return fmt.Errorf("%w: error parsing shapefile header", err)
	}

	// parse shapefile records
	if shapeRecords, err := records.ParseRecords(file, s.header.ShapeType()); err != nil {
		return fmt.Errorf("unable to parse shapefile records: %w", err)
	} else {
		s.records = shapeRecords
	}

	return nil
}

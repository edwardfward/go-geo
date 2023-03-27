package shp

import (
	"errors"
	"fmt"
	"io"
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

// Parse takes a file binary and parses a shapefile.
func (s *ShapeFile) Parse(file *os.File) error {
	// parse shapefile header
	headerBytes := make([]byte, header.SHAPEFILEHEADERLENGTH)

	n, err := file.Read(headerBytes)
	if n != header.SHAPEFILEHEADERLENGTH || err != nil {
		headerParseError := errors.New("unable to parse shapefile header")

		return fmt.Errorf("%w: error reading shapefile header", headerParseError)
	}

	parsedHeader := header.EmptyShapeFileHeader()

	err = parsedHeader.Parse(headerBytes)
	if err != nil {
		improperHeader := errors.New("unable to parse shapefile header")

		return fmt.Errorf("%w: error parsing shapefile header", improperHeader)
	}

	s.header = parsedHeader // assign shapefile header

	// parse shapefile records
	for {
		// read record header bytes from shapefile
		recordHeaderBytes := make([]byte, records.RECORDHEADERLENGTH)

		// check to ensure correct number of bytes read
		numBytes, err := file.Read(recordHeaderBytes)
		if numBytes != records.RECORDHEADERLENGTH || err != nil {
			// reached the end of the file
			if errors.Is(err, io.EOF) {
				break
			}

			return fmt.Errorf("%w: unable to parse record header", err)
		}

		// create record header
		recordHeader := records.EmptyRecordHeader()

		err = recordHeader.Parse(recordHeaderBytes)
		if err != nil {
			return fmt.Errorf("%w: unable to parse record header", err)
		}

		// create record
		recordBytes := make([]byte, recordHeader.LengthBytes())

		numBytes, err = file.Read(recordBytes)
		if numBytes != recordHeader.LengthBytes() || err != nil {
			return fmt.Errorf("%w: unable to read record bytes for record %d",
				err, recordHeader.RecordNumber())
		}

		record, recordError := records.GetShapeType(s.header.ShapeType())
		if recordError != nil {
			return fmt.Errorf("%w: unable to get record shape type", recordError)
		}

		err = record.Parse(recordBytes, recordHeader)
		if err != nil {
			return fmt.Errorf("%w: error parsing record", err)
		}
	}

	return nil
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

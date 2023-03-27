package shp

import (
	"fmt"
	"go-shp/shapefile/shp/header"
	"go-shp/shapefile/shp/records"
	"log"
	"os"
)

// ShapeFile todo documentation
type ShapeFile struct {
	header  header.ShapeFileHeader
	records records.Records
	// sizeBytes   int64
	// timeToParse time.Duration
}

// ParseShapeFile todo documentation and add meaningful error propagation.
func ParseShapeFile(filePath string) (*ShapeFile, error) {
	// open shapefile
	file, err := os.Open(filePath)

	if err != nil {
		return nil, fmt.Errorf("error opening %s", file.Name())
	}

	defer func() {
		if err := file.Close(); err != nil {
			log.Printf("%s failed to close after parsing: %v", file.Name(), err)
		}
	}()

	shapeFile := &ShapeFile{}

	return shapeFile, nil
}

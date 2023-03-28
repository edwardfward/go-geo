package records

import (
	"bytes"
	"encoding/binary"
	"log"
	"testing"
)

func write(buffer *bytes.Buffer, value any) {
	err := binary.Write(buffer, binary.LittleEndian, value)
	if err != nil {
		log.Panicln("%w: error writing to buffer", err)
	}
}

func TestPolygon_Parse(t *testing.T) {
	goodBytes := new(bytes.Buffer)
	write(goodBytes, int32(POLYGON))
	write(goodBytes, -10.0)    // boundary box xMin
	write(goodBytes, -10.0)    // boundary box yMin
	write(goodBytes, 10.0)     // boundary box xMax
	write(goodBytes, 10.0)     // boundary box yMax
	write(goodBytes, int32(1)) // numParts
	write(goodBytes, int32(5)) // numPoints
	write(goodBytes, int32(0)) // parts
	write(goodBytes, -10.0)    // pt. 1 x start
	write(goodBytes, -10.0)    // pt. 1 y
	write(goodBytes, -10.0)    // pt. 2 x
	write(goodBytes, 10.0)     // pt. 2 y
	write(goodBytes, 10.0)     // pt. 3 x
	write(goodBytes, 10.0)     // pt. 3 y
	write(goodBytes, 10.0)     // pt. 4 x
	write(goodBytes, -10.0)    // pt. 4 y
	write(goodBytes, -10.0)    // pt. 5 x end
	write(goodBytes, -10.0)    // pt. 5 y

	polygon := EmptyPolygon()
	recordHeader := RecordHeader{recordNumber: 1,
		contentLength: int32(len(goodBytes.Bytes()) / WORDMULTIPLE)}

	err := polygon.Parse(goodBytes.Bytes(), recordHeader)

	if err != nil {
		t.Errorf("%v: failed to parse known good polygon byte array", err)
	}

	t.Run("check shape type parse", func(t *testing.T) {
		if polygon.shape != POLYGON {
			t.Errorf("parsing failed to produce the correct shape type (%d): %d",
				POLYGON, polygon.shape)
		}
	})

	t.Run("check boundary box", func(t *testing.T) {
		if polygon.box.points[1].x != 10.0 || polygon.box.points[1].y != 10.0 {
			t.Errorf("error parsing polygon boundary box")
		}
	})
}

func TestEmptyPolygon(t *testing.T) {
	got := EmptyPolygon()
	if got.shape != POLYGON {
		t.Errorf("invalid shape received: %d", got.shape)
	}

	if got.points != nil {
		t.Errorf("points array not set to nil")
	}

	if got.parts != nil {
		t.Errorf("parts array not set to nil")
	}

	if got.numPoints != 0 {
		t.Errorf("number of points not set to 0")
	}

	if got.numParts != 0 {
		t.Errorf("number of parts not set to 0")
	}
}

package records_test

import (
	"bytes"
	"encoding/binary"
	"go-geo/shapefile/shp/records"
	"testing"
)

func TestParseParts(t *testing.T) {

	// test good byte sequence with two parts
	t.Run("Good byte sequence", func(t *testing.T) {
		buf := new(bytes.Buffer)

		var firstPoint, secondPoint int32 = 0, 5

		var numPoints int32 = 2

		err := binary.Write(buf, binary.LittleEndian, firstPoint)
		if err != nil {
			t.Fatalf("binary.Write failed: %v", err)
		}

		err = binary.Write(buf, binary.LittleEndian, secondPoint)
		if err != nil {
			t.Fatalf("binary.Write failed: %v\n", err)
		}

		parts, er := records.ParseParts(buf.Bytes(), numPoints)
		if er != nil {
			t.Fatalf("ParseParts failed")
		}

		if parts[0] != firstPoint || parts[1] != secondPoint {
			t.Errorf("expecting [%d, %d], received [%d, %d]",
				firstPoint, secondPoint, parts[0], parts[1])
		}

		t.Logf("sent parts bytes for [%d, %d] and parsed [%d, %d]",
			firstPoint, secondPoint, parts[0], parts[1])
	})

	t.Run("Bad byte sequence", func(t *testing.T) {
		buf := new(bytes.Buffer)

		var firstPoint, secondPoint int32 = 0, 5
		var numPoints int32 = 3

		err := binary.Write(buf, binary.LittleEndian, firstPoint)
		if err != nil {
			t.Fatalf("binary.Write failed: %v", err)
		}

		err = binary.Write(buf, binary.LittleEndian, secondPoint)
		if err != nil {
			t.Fatalf("binary.Write failed: %v", err)
		}

		_, err = records.ParseParts(buf.Bytes(), numPoints)
		if err == nil {
			t.Fatalf("parsed a bad parts byte sequence")
		}
	})
}

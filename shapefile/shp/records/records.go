package records

import "fmt"

type Records struct {
	records []Record
}

func (r *Records) Parse(records []byte, shapeType ShapeType) error {
	cursor := 0 // records array cursor

	for {
		// check to make sure there's enough room in the remaining bytes
		// to not panic with an out-of-range error
		if cursor >= len(records)-1 {
			break
		}

		shape, err := GetShapeType(shapeType)
		if err != nil {
			return fmt.Errorf("error creating new shapetype: %v", err)
		}
		// parse record
		header, e := ParseRecordHeader(records[cursor : cursor+8])
		if e != nil {
			return fmt.Errorf("error parsing record header: %v", err)
		}

		cursor += 8 // advance cursor 8 bytes
		endCursor := cursor + int(header.contentLength)*WORDMULTIPLE

		err = shape.Parse(records[cursor:endCursor], header)
		if err != nil {
			return fmt.Errorf("error parsing record: %v", err)
		}

		cursor = endCursor

		// add shape
		r.records = append(r.records, shape)
	}

	return nil
}

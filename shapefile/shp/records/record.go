package records

type Record interface {
	Parse(record []byte, header RecordHeader) error
	LengthBytes() int32
	RecordNumber() int32
}

const (
	WORDMULTIPLE = 2 // word represents two bytes in length
	INT32LENGTH  = 4
)

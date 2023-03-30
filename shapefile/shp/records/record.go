package records

type Record interface {
	Parse(record []byte) error
}

const (
	WORDMULTIPLE = 2 // word represents two bytes in length
)

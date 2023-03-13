package utils

import (
	"bytes"
	"encoding/binary"
	"log"
)

func ReadBinary[T int32 | float64](b []byte, order binary.ByteOrder, target *T) {
	err := binary.Read(bytes.NewReader(b), order, target)
	if err != nil {
		log.Fatalf("error reading binary on %v: %v", target, err)
	}
}

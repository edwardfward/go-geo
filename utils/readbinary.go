package utils

import (
	"bytes"
	"encoding/binary"
	"log"
)

func ReadBinary[T any](b []byte, order binary.ByteOrder, target *T) {
	err := binary.Read(bytes.NewReader(b), order, target)
	if err != nil {
		log.Fatalf("error reading binary on %v: %v", target, err)
	}
}

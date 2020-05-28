package emoji

import (
	"bytes"
	"unicode/utf8"
)

func Fuzz(data []byte) int {
	// This Fuzz function allocates heavily.
	// To avoid OOM, cap data at 16k.
	// This should be plenty to catch genuine bugs.
	if len(data) > 16384 {
		return -1
	}
	prev := utf8.RuneCount(data)

	g, ok, n := Decode(data)
	if len(g) != n {
		panic("len(g) != n")
	}
	decoded := append(g, data[n:]...)
	if !bytes.Equal(decoded, data) {
		panic("total != data")
	}
	after := utf8.RuneCount(g) + utf8.RuneCount(data[n:])
	if after != prev {
		panic("unexpected split")
	}

	if !ok {
		return 0
	}
	if !utf8.Valid(g) {
		panic("invalid g")
	}
	return 1
}

package emoji

import (
	"testing"
	"unicode"
)

func Test_Tag(t *testing.T) {
	for _, c := range "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ" {
		r := c + 0xE0000
		if !unicode.Is(Tag, r) {
			t.Errorf("%q code %q + 0xE0000 is not counted as a tag", r, c)
		}
	}

	for _, c := range "abcxyz世界\nş123#. 😇🔁" {
		if unicode.Is(Tag, c) {
			t.Errorf("%q code %X is counted as a regional indicator", c, c)
		}
	}
}

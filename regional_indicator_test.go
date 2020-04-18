package emoji

import (
	"testing"
	"unicode"
	"unicode/utf8"

	"github.com/Succo/countrymoji"
)

func Test_RegionalIndicator_Test(t *testing.T) {
	for _, c := range "abcdefghijklmnopqrstuvwxyz" {
		r, _ := utf8.DecodeRuneInString(countrymoji.Alpha2ToFlag(string(c)))
		if !unicode.Is(RegionalIndicator, r) {
			t.Errorf("%q code %X is not counted as a regional indicator", r, r)
		}
	}

	for _, c := range "abcxyz世界\nş123#. 😇🔁" {
		if unicode.Is(RegionalIndicator, c) {
			t.Errorf("%q code %X is not counted as a regional indicator", c, c)
		}
	}
}

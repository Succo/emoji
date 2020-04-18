package main

import (
	"fmt"
	"sort"
	"testing"
	"unicode"

	"github.com/Succo/emoji"
)

func Test_emojiTable_is_sorted(t *testing.T) {
	for _, table := range []*unicode.RangeTable{emoji.Emoji, emoji.EmojiPresentation, emoji.EmojiModifier, emoji.EmojiModifierBase, emoji.EmojiComponent, emoji.ExtendedPictographic} {
		if !sort.SliceIsSorted(table.R16, func(i, j int) bool { return table.R16[i].Lo < table.R16[i].Lo }) {
			t.Errorf("table.R16 not sorted for Lo")
		}
		if !sort.SliceIsSorted(table.R16, func(i, j int) bool { return table.R16[i].Hi < table.R16[i].Hi }) {
			t.Errorf("table.R16 not sorted for Hi")
		}
		if !sort.SliceIsSorted(table.R32, func(i, j int) bool { return table.R32[i].Lo < table.R32[i].Lo }) {
			t.Errorf("table.R32 not sorted for Lo")
		}
		if !sort.SliceIsSorted(table.R32, func(i, j int) bool { return table.R32[i].Hi < table.R32[i].Hi }) {
			t.Errorf("table.R32 not sorted for Hi")
		}
	}
}

func Test_emojiTable(t *testing.T) {
	var tests = []struct {
		in  rune
		out bool
	}{
		{'r', false},
		{'2', true}, // numbers are part of emoji 🤔
		{' ', false},
		{'\n', false},
		{'{', false},
		{'ç', false},
		{'ğ', false},
		{'ş', false},

		{'😀', true},
		{'😇', true},
		{'😜', true},
		{'😔', true},
		{'🥶', true},
		{'😨', true},
		{'🤡', true},
		{'😿', true},
		{'💙', true},
		{'✋', true},
		{'🤝', true},
		{'🫀', true},
		{'🧑', true},
		{'🧝', true},
		{'🚵', true},
		{'🐘', true},
		{'🌸', true},
		{'🥔', true},
		{'🍗', true},
		{'🥫', true},
		{'🦑', true},
		{'🏪', true},
		{'🚄', true},
		{'🛬', true},
		{'🕛', true},
		{'🌘', true},
		{'🌪', true},
		{'🧨', true},
		{'🥇', true},
		{'🎱', true},
		{'👕', true},
		{'🥿', true},
		{'💄', true},
		{'🔕', true},
		{'🎸', true},
		{'📟', true},
		{'📸', true},
		{'🗞', true},
		{'📇', true},
		{'🔑', true},
		{'🏹', true},
		{'🧰', true},
		{'🧬', true},
		{'🚪', true},
		{'🚭', true},
		{'⤵', true},
		{'✡', true},
		{'♊', true},
		{'🔁', true},
		{'📴', true},
		{'⚧', true},
		{'❓', true},
		{'🔱', true},
		{'❇', true},
		{'🆎', true},
		{'🆚', true},
		{'🈸', true},
		{'🔵', true},
		{'🔺', true},
		{'🏳', true},
	}

	for _, tt := range tests {
		if unicode.Is(emoji.Emoji, tt.in) != tt.out {
			Is(emoji.Emoji, tt.in)
			t.Errorf("got %t for %q code %X", !tt.out, tt.in, tt.in)
		}
	}

}

//Copy pasted debug code from stdlib

const linearMax = 18

// Is reports whether the rune is in the specified table of ranges.
func Is(rangeTab *unicode.RangeTable, r rune) bool {
	fmt.Printf("is %q\n", r)
	r16 := rangeTab.R16
	if len(r16) > 0 && r <= rune(r16[len(r16)-1].Hi) {
		return is16(r16, uint16(r))
	}
	r32 := rangeTab.R32
	if len(r32) > 0 && r >= rune(r32[0].Lo) {
		return is32(r32, uint32(r))
	}
	return false
}

// is16 reports whether r is in the sorted slice of 16-bit ranges.
func is16(ranges []unicode.Range16, r uint16) bool {
	fmt.Printf("is16 %q %X\n", r, r)
	if len(ranges) <= linearMax || r <= unicode.MaxLatin1 {
		fmt.Printf("is16 linear %q\n", r)
		for i := range ranges {
			range_ := &ranges[i]
			fmt.Println(range_)
			if r < range_.Lo {
				return false
			}
			fmt.Printf("is16 linear %q Stride %d\n", r, range_.Stride)
			if r <= range_.Hi {
				return range_.Stride == 1 || (r-range_.Lo)%range_.Stride == 0
			}
		}
		return false
	}

	// binary search over ranges
	lo := 0
	hi := len(ranges)
	for lo < hi {
		m := lo + (hi-lo)/2
		range_ := &ranges[m]
		if range_.Lo <= r && r <= range_.Hi {
			return range_.Stride == 1 || (r-range_.Lo)%range_.Stride == 0
		}
		if r < range_.Lo {
			fmt.Printf("%X, %X <\n", range_.Lo, range_.Hi)
			hi = m
		} else {
			fmt.Printf("%X, %X >\n", range_.Lo, range_.Hi)
			lo = m + 1
		}
	}
	return false
}

// is32 reports whether r is in the sorted slice of 32-bit ranges.
func is32(ranges []unicode.Range32, r uint32) bool {
	if len(ranges) <= linearMax {
		for i := range ranges {
			range_ := &ranges[i]
			if r < range_.Lo {
				return false
			}
			if r <= range_.Hi {
				return range_.Stride == 1 || (r-range_.Lo)%range_.Stride == 0
			}
		}
		return false
	}

	// binary search over ranges
	lo := 0
	hi := len(ranges)
	for lo < hi {
		m := lo + (hi-lo)/2
		range_ := ranges[m]
		if range_.Lo <= r && r <= range_.Hi {
			return range_.Stride == 1 || (r-range_.Lo)%range_.Stride == 0
		}
		if r < range_.Lo {
			hi = m
		} else {
			lo = m + 1
		}
	}
	return false
}

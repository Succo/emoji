package main

import (
	"sort"
	"testing"
	"unicode"

	"github.com/Succo/emoji"
)

func Test_emojiTable_is_sorted(t *testing.T) {
	if !sort.SliceIsSorted(emoji.Table.R16, func(i, j int) bool { return emoji.Table.R16[i].Lo < emoji.Table.R16[i].Lo }) {
		t.Errorf("emoji.Table.R16 not sorted for Lo")
	}
	if !sort.SliceIsSorted(emoji.Table.R16, func(i, j int) bool { return emoji.Table.R16[i].Hi < emoji.Table.R16[i].Hi }) {
		t.Errorf("emoji.Table.R16 not sorted for Hi")
	}
	if !sort.SliceIsSorted(emoji.Table.R32, func(i, j int) bool { return emoji.Table.R32[i].Lo < emoji.Table.R32[i].Lo }) {
		t.Errorf("emoji.Table.R32 not sorted for Lo")
	}
	if !sort.SliceIsSorted(emoji.Table.R32, func(i, j int) bool { return emoji.Table.R32[i].Hi < emoji.Table.R32[i].Hi }) {
		t.Errorf("emoji.Table.R32 not sorted for Hi")
	}
}

func Test_emojiTable(t *testing.T) {
	var tests = []struct {
		in  rune
		out bool
	}{
		{'r', false},
		{'2', false},
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
		if unicode.Is(emoji.Table, tt.in) != tt.out {
			t.Errorf("got %t for %q code %X", !tt.out, tt.in, tt.in)
		}
	}

}

package main

import (
	"testing"
	"unicode"

	"github.com/Succo/emoji"
)

func Test_emojiTable_is_sorted(t *testing.T) {
	for _, table := range []*unicode.RangeTable{emoji.Emoji, emoji.EmojiPresentation, emoji.EmojiModifier, emoji.EmojiModifierBase, emoji.EmojiComponent, emoji.ExtendedPictographic} {
		for i, r := range table.R16 {
			if r.Lo > r.Hi {
				t.Errorf("table.R16 wrong range for table")
			}
			if i+1 < len(table.R16) && table.R16[i+1].Lo <= r.Hi {
				t.Errorf("table.R16 overlap")
			}
		}
		for i, r := range table.R32 {
			if r.Lo > r.Hi {
				t.Errorf("table.R32 wrong range for table")
			}
			if i+1 < len(table.R32) && table.R32[i+1].Lo <= r.Hi {
				t.Errorf("table.R32 overlap")
			}
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
		{'#', true}, // # is part of emoji 🤔
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
		{'🇦', true},
	}

	for _, tt := range tests {
		if unicode.Is(emoji.Emoji, tt.in) != tt.out {
			t.Errorf("got %t for %q code %X", !tt.out, tt.in, tt.in)
		}
	}

}

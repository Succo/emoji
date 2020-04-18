package emoji

import (
	"testing"
	"unicode"
)

func Test_emojiTable_is_sorted(t *testing.T) {
	for _, table := range []*unicode.RangeTable{Emoji, EmojiPresentation, EmojiModifier, EmojiModifierBase, EmojiComponent, ExtendedPictographic} {
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

var nonEmoji = []rune{
	'r',
	' ',
	'\n',
	'{',
	'ç',
	'ğ',
	'ş',
}

var emojiNonPictographic = []rune{
	'2',
	'#',
	'*',
	'🇦',
}

var emojiPictographic = []rune{
	'😀',
	'😇',
	'😜',
	'😔',
	'🥶',
	'😨',
	'🤡',
	'😿',
	'💙',
	'✋',
	'🤝',
	'🫀',
	'🧑',
	'🧝',
	'🚵',
	'🐘',
	'🌸',
	'🥔',
	'🍗',
	'🥫',
	'🦑',
	'🏪',
	'🚄',
	'🛬',
	'🕛',
	'🌘',
	'🌪',
	'🧨',
	'🥇',
	'🎱',
	'👕',
	'🥿',
	'💄',
	'🔕',
	'🎸',
	'📟',
	'📸',
	'🗞',
	'📇',
	'🔑',
	'🏹',
	'🧰',
	'🧬',
	'🚪',
	'🚭',
	'⤵',
	'✡',
	'♊',
	'🔁',
	'📴',
	'⚧',
	'❓',
	'🔱',
	'❇',
	'🆎',
	'🆚',
	'🈸',
	'🔵',
	'🔺',
	'🏳',
}

func Test_emojiTable(t *testing.T) {
	for _, r := range nonEmoji {
		if unicode.Is(Emoji, r) {
			t.Errorf("%q code %X is counted as an emoji", r, r)
		}
	}

	for _, r := range emojiNonPictographic {
		if !unicode.Is(Emoji, r) {
			t.Errorf("%q code %X is not counted as an emoji", r, r)
		}
		if unicode.Is(ExtendedPictographic, r) {
			t.Errorf("%q code %X is counted as pictographic", r, r)
		}
	}

	for _, r := range emojiPictographic {
		if !unicode.Is(Emoji, r) {
			t.Errorf("%q code %X is not counted as an emoji", r, r)
		}
		if !unicode.Is(ExtendedPictographic, r) {
			t.Errorf("%q code %X is not counted as pictographic", r, r)
		}
	}
}

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

var nonEmojiTest = []rune{
	'r',
	' ',
	'\n',
	'{',
	'ç',
	'ğ',
	'ş',
}

var emojiNonPictographicTest = []rune{
	'2',
	'#',
	'*',
	'🇦',
}

var emojiPictographicTest = []rune{
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

var emojiModifierTest = []rune{
	'🏼',
}

var emojiModifierBaseTest = []rune{
	'👰',
	'🤡',
	'😀',
	'😇',
	'😜',
}

var emojiComponentTest = []rune{
	'🦰',
	'6',
	'*',
	0xFE0F, // combining enclosing keycap
	0x200D, // VARIATION SELECTOR-16
	'🇦',
	'🏼',
}

func Test_emojiTable(t *testing.T) {
	for _, r := range nonEmojiTest {
		if unicode.Is(Emoji, r) {
			t.Errorf("%q code %X is counted as an emoji", r, r)
		}
	}

	for _, r := range emojiNonPictographicTest {
		if !unicode.Is(Emoji, r) {
			t.Errorf("%q code %X is not counted as an emoji", r, r)
		}
		if unicode.Is(ExtendedPictographic, r) {
			t.Errorf("%q code %X is counted as pictographic", r, r)
		}
	}

	for _, r := range emojiPictographicTest {
		if !unicode.Is(Emoji, r) {
			t.Errorf("%q code %X is not counted as an emoji", r, r)
		}
		if !unicode.Is(ExtendedPictographic, r) {
			t.Errorf("%q code %X is not counted as pictographic", r, r)
		}
	}

	for _, r := range emojiModifierTest {
		if !unicode.Is(Emoji, r) {
			t.Errorf("%q code %X is not counted as an emoji", r, r)
		}
		if !unicode.Is(EmojiModifier, r) {
			t.Errorf("%q code %X is not counted as an emoji modifier", r, r)
		}
		if unicode.Is(ExtendedPictographic, r) {
			t.Errorf("%q code %X is counted as pictographic", r, r)
		}
	}
	for _, r := range emojiComponentTest {
		if !unicode.Is(EmojiComponent, r) {
			t.Errorf("%q code %X is not counted as an emoji component", r, r)
		}
	}
}

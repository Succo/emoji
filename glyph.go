package emoji

import (
	"unicode"
	"unicode/utf8"
)

const zeroWidthJoiner = 0x200D
const emojiVS = 0xFE0F
const enclosingKeycap = 0x20E3
const termTag = 0xE007F

// PossibleGlyph checks is the given string might be an emoji
// based on the EBNF from https://www.unicode.org/reports/tr51/#EBNF_and_Regex
//
// possible_emoji :=
//  flag_sequence
//  | zwj_element (\x{200D} zwj_element)*
//
//
// flag_sequence :=
//   \p{RI}\p{RI}
//
// zwj_element :=
//   \p{Emoji} emoji_modification?
//
// emoji_modification :=
//   \p{EMod}
// | \x{FE0F} \x{20E3}?
//
// tag_modifier :=
//   [\x{E0020}-\x{E007E}]+ \x{E007F}
//
// There should not be false negative (glyph wrongly detected as an emoji)
// There is false positive such inexistant flag indicator
func PossibleGlyph(s string) bool {
	r1, n1 := utf8.DecodeRuneInString(s)
	if n1 == 0 {
		return false
	}
	if unicode.Is(RegionalIndicator, r1) {
		r2, n2 := utf8.DecodeRuneInString(s[n1:])
		if n2 == 0 {
			return false
		}
		return unicode.Is(RegionalIndicator, r2) && len(s) == n1+n2
	}
	n := n1
	for unicode.Is(Emoji, r1) {
		r2, n2 := utf8.DecodeRuneInString(s[n:])
		if n2 == 0 {
			return unicode.Is(ExtendedPictographic, r1)
		}
		n += n2

		if r2 == emojiVS {
			r3, n3 := utf8.DecodeRuneInString(s[n:])
			if n3 == 0 {
				return true
			}
			n += n3
			if r3 == enclosingKeycap {
				r2, n2 = utf8.DecodeRuneInString(s[n:])
				if n2 == 0 {
					return true
				}
				n += n2
			} else {
				r2, n2 = r3, n3
			}
		} else if unicode.Is(EmojiModifier, r2) && unicode.Is(EmojiModifierBase, r1) {
			r2, n2 = utf8.DecodeRuneInString(s[n:])
			if n2 == 0 {
				return true
			}
			n += n2
		} else if unicode.Is(Tag, r2) && r1 == '🏴' {
			for unicode.Is(Tag, r2) {
				r2, n2 = utf8.DecodeRuneInString(s[n:])
				n += n2
			}
			if r2 != termTag {
				return false
			}
			r2, n2 = utf8.DecodeRuneInString(s[n:])
			if n2 == 0 {
				return true
			}
			n += n2
		}

		if r2 != zeroWidthJoiner {
			return false
		}

		r1, n1 = utf8.DecodeRuneInString(s[n:])
		n += n1
	}
	return false
}

// Decode returns
// - the first complete emoji, true, and it's width in bytes is available
// - the full non emoji sequence, false and it's width in bytes (might be a rune or multiples in case of malformed emoji)
func Decode(s string) (string, bool, int) {
	r1, n1 := utf8.DecodeRuneInString(s)
	if n1 == 0 {
		return "", false, 0
	}
	if unicode.Is(RegionalIndicator, r1) {
		r2, n2 := utf8.DecodeRuneInString(s[n1:])
		if !unicode.Is(RegionalIndicator, r2) {
			return string(r1), false, n1
		}
		return string(r1) + string(r2), true, n1 + n2
	}
	n := n1
	for unicode.Is(Emoji, r1) {
		r2, n2 := utf8.DecodeRuneInString(s[n:])
		if n2 == 0 {
			return s[:n], unicode.Is(ExtendedPictographic, r1), n
		}

		if r2 == emojiVS {
			n += n2
			r3, n3 := utf8.DecodeRuneInString(s[n:])
			if n3 == 0 {
				return s[:n], true, n
			}
			if r3 == enclosingKeycap {
				n += n3
				r2, n2 = utf8.DecodeRuneInString(s[n:])
				if n2 == 0 {
					return s[:n], true, n
				}
			} else {
				r2, n2 = r3, n3
			}
		} else if unicode.Is(EmojiModifier, r2) && unicode.Is(EmojiModifierBase, r1) {
			n += n2
			r2, n2 = utf8.DecodeRuneInString(s[n:])
			if n2 == 0 {
				return s[:n], true, n
			}
		} else if unicode.Is(Tag, r2) && r1 == '🏴' {
			for unicode.Is(Tag, r2) {
				r2, n2 = utf8.DecodeRuneInString(s[n:])
				n += n2
			}
			if r2 != termTag {
				return s[:n], false, n
			}
			r2, n2 = utf8.DecodeRuneInString(s[n:])
			if n2 == 0 {
				return s[:n], true, n
			}
		}

		if r2 != zeroWidthJoiner {
			return s[:n], true, n
		}
		n += n2

		r1, n1 = utf8.DecodeRuneInString(s[n:])
		n += n1
	}
	return s[:n], false, n
}

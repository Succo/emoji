package emoji

import (
	"bytes"
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
func PossibleGlyph(b []byte) bool {
	_, ok, n := Decode(b)
	return ok && n == len(b)
}

// Decode returns
// - the first complete emoji, true, and it's width in bytes is available
// - the full non emoji sequence, false and it's width in bytes (might be a rune or multiples in case of malformed emoji)
func Decode(b []byte) ([]byte, bool, int) {
	r1, n1 := utf8.DecodeRune(b)
	if n1 == 0 {
		return nil, false, 0
	}
	if unicode.Is(RegionalIndicator, r1) {
		r2, n2 := utf8.DecodeRune(b[n1:])
		if !unicode.Is(RegionalIndicator, r2) {
			return []byte(string(r1)), false, n1
		}
		return []byte(string(r1) + string(r2)), true, n1 + n2
	}
	n := n1
	for unicode.Is(Emoji, r1) {
		r2, n2 := utf8.DecodeRune(b[n:])
		if n2 == 0 {
			return b[:n], unicode.Is(ExtendedPictographic, r1), n
		}

		if r2 == emojiVS {
			n += n2
			r3, n3 := utf8.DecodeRune(b[n:])
			if n3 == 0 {
				return b[:n], true, n
			}
			if r3 == enclosingKeycap {
				n += n3
				r2, n2 = utf8.DecodeRune(b[n:])
				if n2 == 0 {
					return b[:n], true, n
				}
			} else {
				r2, n2 = r3, n3
			}
		} else if unicode.Is(EmojiModifier, r2) && unicode.Is(EmojiModifierBase, r1) {
			n += n2
			r2, n2 = utf8.DecodeRune(b[n:])
			if n2 == 0 {
				return b[:n], true, n
			}
		} else if r1 == '🏴' && unicode.Is(Tag, r2) {
			for unicode.Is(Tag, r2) {
				r2, n2 = utf8.DecodeRune(b[n:])
				n += n2
			}
			if r2 != termTag {
				return b[:n], false, n
			}
			r2, n2 = utf8.DecodeRune(b[n:])
			if n2 == 0 {
				return b[:n], true, n
			}
		}

		if r2 != zeroWidthJoiner {
			return b[:n], true, n
		}
		n += n2

		r1, n1 = utf8.DecodeRune(b[n:])
		n += n1
	}
	return b[:n], false, n
}

// Find returns the n first emoji in b
// of all of thems if max == -1
func Find(b []byte, max int) [][]byte {
	emojis := [][]byte{}
	if max == 0 {
		return emojis
	}
	for {
		g, ok, n := Decode(b)
		if n == 0 {
			break
		}
		if ok {
			emojis = append(emojis, g)
			if len(emojis) == max {
				return emojis
			}
		}
		b = b[n:]
	}
	return emojis
}

// Replace replace the n first all emoji with f(b)
// of all of thems if max == -1
func Replace(b []byte, max int, f func([]byte) []byte) []byte {
	if max == 0 {
		return b
	}
	if len(Find(b, max)) == 0 {
		return b
	}
	var buf bytes.Buffer
	var count int
	for {
		g, ok, n := Decode(b)
		b = b[n:]
		if n == 0 {
			break
		}
		if ok {
			buf.Write(f(g))
			count++
			if count == max {
				buf.Write(b)
				return buf.Bytes()
			}
		} else {
			buf.Write(g)
		}
	}
	return buf.Bytes()
}

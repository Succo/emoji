package emoji

import (
	"bytes"
	"unicode"
	"unicode/utf8"
)

const zeroWidthJoiner = rune(0x200D)
const emojiVS = rune(0xFE0F)
const enclosingKeycap = rune(0x20E3)
const termTag = rune(0xE007F)

var enclosingKeycapS = string(enclosingKeycap)
var enclosingKeycapB = []byte(enclosingKeycapS)

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
	u1 := uint32(r1)
	if RegionalIndicator.R32[0].Lo <= u1 && u1 <= RegionalIndicator.R32[0].Hi {
		r2, n2 := utf8.DecodeRune(b[n1:])
		u2 := uint32(r2)
		if RegionalIndicator.R32[0].Lo > u2 || u2 > RegionalIndicator.R32[0].Hi {
			return []byte(string(r1)), false, n1
		}
		return b[:n1+n2], true, n1 + n2
	}
	n := n1
	for unicode.Is(Emoji, r1) {
		r2, n2 := utf8.DecodeRune(b[n:])
		if n2 == 0 {
			return b[:n], unicode.Is(ExtendedPictographic, r1), n
		}

		if r2 == emojiVS {
			n += n2
			if bytes.HasPrefix(b[n:], enclosingKeycapB) {
				n += len(enclosingKeycapB)
			}
			r2, n2 = utf8.DecodeRune(b[n:])
			if n2 == 0 {
				return b[:n], true, n
			}
		} else if isEmod(r2) && unicode.Is(EmojiModifierBase, r1) {
			n += n2
			r2, n2 = utf8.DecodeRune(b[n:])
			if n2 == 0 {
				return b[:n], true, n
			}
		} else if r1 == '🏴' && isTag(r2) {
			for isTag(r2) {
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

func isEmod(r rune) bool {
	u := uint32(r)
	return EmojiModifier.R32[0].Lo <= u && u <= EmojiModifier.R32[0].Hi
}

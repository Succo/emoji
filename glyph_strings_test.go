package emoji

import (
	"strings"
	"testing"
	"unicode"

	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

var notEmojiTest = []string{
	"",
	"a",
	"😀b",
	"b😀",
	"\n",
	"test",
	".",
	"😀👍",
	"🇧",
	"😀🏴󠁵󠁳󠁴󠁸󠁿",
	"A🏴󠁵󠁳󠁴󠁸󠁿",
	"🏴" + string(rune(0xE0031)) + string(rune(0xE004F)),
	"⛰️🏼",
	"🏥🏼",
	"🏼",
	"2",
	"#",
	string(rune(0x200D)),
}

var emojiTest = []string{
	"©️",
	"⏏️",
	"😀",
	"👍",
	"⛰️",
	"🏕️",
	"🏥",
	"🛢️",
	"💈",
	"⛱️",
	"🪁",
	"☎️",
	"💡",
	"💳",
	"🖌️",
	"🔓",
	"⛓️",
	"🧴",
	"🍌",
	"🍓",
	"🥔",
	"🧅",
	"🥐",
	"🥞",
	"🍳",
	"🍠",
	"🥃",
	"🍽️",
	"😆",
	"😍",
	"😚",
	"😬",
	"🤯",
	"💀",
	"😫",
	"✌️",
	"🧠",
	"🧙",
	"💇",
	"🎒",
	"⛑️",
	"🥺",
	"😂",
	"😊",
	"🔥",

	"0️⃣",

	"🙍‍♂️",
	"👩‍🎓",
	"🧑‍🏫",
	"🧑‍⚕️",
	"🧑‍🍳",
	"👨‍🍳",
	"👮‍♀️",
	"🧙‍♂️",
	"🧟‍♂️",
	"👩‍🦯",
	"👨‍👩‍👦‍👦",
	"👩‍👩‍👧‍👧",
	"👨‍👧‍👦",

	"👋🏼",
	"🖖🏿",
	"🦻🏻",
	"👨‍🦰",
	"👩🏼‍🦰",
	"🧏🏼‍♀️",
	"🧜🏼‍♀️",
	"🧝🏼‍♀️",
	"👯🏼‍♀️",
	"👩‍❤️‍👩",
	"👩🏾‍👨🏾‍👦🏾",

	"🏳️",
	"🎌",
	"🇧🇳",
	"🏳️‍⚧️",
	"🇧🇸",
	"🇪🇨",
	"🇭🇲",
	"🇱🇮",
	"🇳🇬",
	"🇸🇬",
	"🇺🇬",
	"🏴󠁵󠁳󠁴󠁸󠁿",
	"🏴󠁧󠁢󠁳󠁣󠁴󠁿",
}

func Test_PossibleGlyphString(t *testing.T) {
	for _, s := range notEmojiTest {
		if PossibleGlyphString(s) {
			t.Errorf("%q returned positive", s)
		}
	}

	for _, s := range emojiTest {
		if !PossibleGlyphString(s) {
			t.Errorf("%q returned negative", s)
		}
		if PossibleGlyphString(s + "a") {
			t.Errorf("%q returned positive", s+"a")
		}
		if PossibleGlyphString(s + s) {
			t.Errorf("%q returned positive", s+s)
		}
	}
}

func Test_unicodeTransform(t *testing.T) {
	isMn := func(r rune) bool {
		return unicode.Is(unicode.Mn, r) && r != emojiVS
	}
	tr := transform.Chain(norm.NFD, transform.RemoveFunc(isMn), norm.NFC)
	for _, s := range emojiTest {
		b := make([]byte, len(s))
		n, _, _ := tr.Transform(b, []byte(s), true)
		o := string(b[:n])
		if o != s {
			t.Errorf("%q != %q %X != %X", o, s, o, s)
		}
		if !PossibleGlyphString(o) {
			t.Errorf("%q returned negative", o)
		}
	}
}

func Test_DecodeString(t *testing.T) {
	for _, s := range notEmojiTest {
		g, ok, n := DecodeString(s)
		if ok && !PossibleGlyphString(g) {
			t.Errorf("DecodeString(%q) returned positive %q", s, g)
		}
		if len(g) != n {
			t.Errorf("DecodeString(%q) returned incoherent len", s)
		}
	}

	for _, s := range emojiTest {
		g, ok, n := DecodeString(s)
		if !ok {
			t.Errorf("DecodeString(%q) returned negative %q ", s, g)
		}
		if g != s {
			t.Errorf("DecodeString(%q) returned not the full string but %q (%X != %X)", s, g, s, g)
		}
		if len(g) != n {
			t.Errorf("DecodeString(%q) returned incoherent len", s)
		}
	}
	for _, s1 := range emojiTest {
		for _, s2 := range append([]string{"aaa", "bbb"}, emojiTest...) {
			s := s1 + s2
			g, ok, n := DecodeString(s)
			if !ok {
				t.Errorf("DecodeString(%q) returned negative %q ", s, g)
			}
			if g != s1 {
				t.Errorf("DecodeString(%q) returned not the full string but %q (%X != %X)", s, g, s1, g)
			}
			if len(g) != n {
				t.Errorf("DecodeString(%q) returned incoherent len", s)
			}
			if s[n:] != s2 {
				t.Errorf("DecodeString(%q) returned incoherent len full string but %q (%X != %X)", s2, s[n:], s2, s[n:])
			}
		}
	}
}

func Test_DecodeStringCanReadAText(t *testing.T) {
	text := strings.Join(emojiTest, "test phrase")
	var b strings.Builder
	for {
		g, ok, n := DecodeString(text)
		if n == 0 {
			break
		}
		if ok {
			b.WriteString(" " + g + " ")
		} else {
			b.WriteString(g)
		}
		text = text[n:]
	}
	emojiWithSpace := make([]string, len(emojiTest))
	for i, e := range emojiTest {
		emojiWithSpace[i] = " " + e + " "
	}
	if b.String() != strings.Join(emojiWithSpace, "test phrase") {
		t.Errorf("Got :\n%q\nExpected :\n%q", b.String(), strings.Join(emojiWithSpace, "test phrase"))
	}
}

func Test_FindString(t *testing.T) {
	text := strings.Join(emojiTest, "test phrase")
	for n := range emojiTest {
		found := FindString(text, n)
		if len(found) != n {
			t.Errorf("FindString wrong len %d not %d", len(found), n)
		}
		for i, s := range emojiTest[:n] {
			if s != found[i] {
				t.Errorf("FindString wrong %d result, %q not  %q", i, found[i], s)
			}
		}
	}
	found := FindString(text, -1)
	if len(found) != len(emojiTest) {
		t.Errorf("FindString wrong len %d not %d", len(found), len(emojiTest))
	}
	for i, s := range emojiTest {
		if s != found[i] {
			t.Errorf("FindString wrong %d result, %q not  %q", i, found[i], s)
		}
	}
}

func Test_ReplaceString(t *testing.T) {
	text := strings.Join(emojiTest, "test phrase")
	for n := range emojiTest {
		replaced := ReplaceString(text, n, func(s string) string { return ";" + s + "|" })
		emojiEdited := make([]string, len(emojiTest))
		copy(emojiEdited, emojiTest)
		for i, e := range emojiTest[:n] {
			emojiEdited[i] = ";" + e + "|"
		}
		expected := strings.Join(emojiEdited, "test phrase")
		if replaced != expected {
			t.Errorf("ReplaceString error %q not %q", replaced, expected)
		}
	}
	emojiEdited := make([]string, len(emojiTest))
	for i, e := range emojiTest {
		emojiEdited[i] = ";" + e + "|"
	}
	replaced := ReplaceString(text, -1, func(s string) string { return ";" + s + "|" })
	expected := strings.Join(emojiEdited, "test phrase")
	if replaced != expected {
		t.Errorf("ReplaceString error %q not %q", replaced, expected)
	}
}

func Benchmark_FindString(b *testing.B) {
	var n int
	for i := 0; i < b.N; i++ {
		l := FindString("0⛱️1☎️2🙍‍♂️3👩🏾‍👨🏾‍👦🏾4🇭🇲5🏴󠁧󠁢󠁳󠁣󠁴󠁿6789", -1)
		n += len(l)
	}
}

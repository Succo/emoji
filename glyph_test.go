package emoji

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

func Test_PossibleGlyph(t *testing.T) {
	for _, s := range notEmojiTest {
		if PossibleGlyph([]byte(s)) {
			t.Errorf("%q returned positive", s)
		}
	}

	for _, s := range emojiTest {
		if !PossibleGlyph([]byte(s)) {
			t.Errorf("%q returned negative", s)
		}
		if PossibleGlyph([]byte(s + "a")) {
			t.Errorf("%q returned positive", s+"a")
		}
		if PossibleGlyph([]byte(s + s)) {
			t.Errorf("%q returned positive", s+s)
		}
	}
}

func Test_Decode(t *testing.T) {
	for _, s := range notEmojiTest {
		g, ok, n := Decode([]byte(s))
		if ok && !PossibleGlyph([]byte(g)) {
			t.Errorf("DecodeString(%q) returned positive %q", s, g)
		}
		if len(g) != n {
			t.Errorf("DecodeString(%q) returned incoherent len", s)
		}
	}

	for _, s := range emojiTest {
		g, ok, n := Decode([]byte(s))
		if !ok {
			t.Errorf("DecodeString(%q) returned negative %q ", s, g)
		}
		if string(g) != s {
			t.Errorf("DecodeString(%q) returned not the full string but %q (%X != %X)", s, g, s, g)
		}
		if len(g) != n {
			t.Errorf("DecodeString(%q) returned incoherent len", s)
		}
	}
	for _, s1 := range emojiTest {
		for _, s2 := range append([]string{"aaa", "bbb"}, emojiTest...) {
			s := s1 + s2
			g, ok, n := Decode([]byte(s))
			if !ok {
				t.Errorf("DecodeString(%q) returned negative %q ", s, g)
			}
			if string(g) != s1 {
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

func Test_DecodeCanReadAText(t *testing.T) {
	text := strings.Join(emojiTest, "test phrase")
	var b bytes.Buffer
	for {
		g, ok, n := Decode([]byte(text))
		if n == 0 {
			break
		}
		if ok {
			b.WriteRune(' ')
			b.Write(g)
			b.WriteRune(' ')
		} else {
			b.Write(g)
		}
		text = text[n:]
	}
	emojiWithSpace := make([][]byte, len(emojiTest))
	for i, e := range emojiTest {
		emojiWithSpace[i] = append(append([]byte(" "), e...), ' ')
	}
	if bytes.Compare(b.Bytes(), bytes.Join(emojiWithSpace, []byte("test phrase"))) != 0 {
		t.Errorf("Got :\n%q\nExpected :\n%q", string(b.Bytes()), string(bytes.Join(emojiWithSpace, []byte("test phrase"))))
	}
}

func Test_Find(t *testing.T) {
	text := strings.Join(emojiTest, "test phrase")
	for n := range emojiTest {
		found := Find([]byte(text), n)
		if len(found) != n {
			t.Errorf("FindString wrong len %d not %d", len(found), n)
		}
		for i, s := range emojiTest[:n] {
			if s != string(found[i]) {
				t.Errorf("FindString wrong %d result, %q not  %q", i, found[i], s)
			}
		}
	}
	found := Find([]byte(text), -1)
	if len(found) != len(emojiTest) {
		t.Errorf("FindString wrong len %d not %d", len(found), len(emojiTest))
	}
	for i, s := range emojiTest {
		if s != string(found[i]) {
			t.Errorf("FindString wrong %d result, %q not  %q", i, found[i], s)
		}
	}
}

func Test_Replace(t *testing.T) {
	text := []byte(strings.Join(emojiTest, "test phrase"))
	for n := range emojiTest {
		replaced := Replace(text, n, func(b []byte) []byte { return append(append([]byte(";"), b...), '|') })
		emojiEdited := make([]string, len(emojiTest))
		copy(emojiEdited, emojiTest)
		for i, e := range emojiTest[:n] {
			emojiEdited[i] = ";" + e + "|"
		}
		expected := strings.Join(emojiEdited, "test phrase")
		if string(replaced) != expected {
			t.Errorf("ReplaceString error %q not %q", replaced, expected)
		}
	}
	emojiEdited := make([]string, len(emojiTest))
	for i, e := range emojiTest {
		emojiEdited[i] = ";" + e + "|"
	}
	replaced := Replace(text, -1, func(b []byte) []byte { return append(append([]byte(";"), b...), '|') })
	expected := strings.Join(emojiEdited, "test phrase")
	if string(replaced) != expected {
		t.Errorf("ReplaceString error %q not %q", replaced, expected)
	}
}

func Benchmark_Find(b *testing.B) {
	var n int
	s := []byte("0⛱️1☎️2🙍‍♂️3👩🏾‍👨🏾‍👦🏾4🇭🇲5🏴󠁧󠁢󠁳󠁣󠁴󠁿6789")
	for i := 0; i < b.N; i++ {
		l := Find(s, -1)
		n += len(l)
	}
	fmt.Println(n)
}

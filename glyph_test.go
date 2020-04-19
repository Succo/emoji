package emoji

import "testing"

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
	"🏴" + string(0xE0031) + string(0xE004F),
	"⛰️🏼",
	"🏥🏼",
	"🏼",
	"2",
	"#",
	string(0x200D),
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

func Test_PossibleGlyph(t *testing.T) {
	for _, s := range notEmojiTest {
		if PossibleGlyph(s) {
			t.Errorf("%q returned positive", s)
		}
	}

	for _, s := range emojiTest {
		if !PossibleGlyph(s) {
			t.Errorf("%q returned negative", s)
		}
		if PossibleGlyph(s + "a") {
			t.Errorf("%q returned positive", s+"a")
		}
		if PossibleGlyph(s + s) {
			t.Errorf("%q returned positive", s+s)
		}
	}
}

func Test_Decode(t *testing.T) {
	for _, s := range notEmojiTest {
		g, ok, n := Decode(s)
		if ok && !PossibleGlyph(g) {
			t.Errorf("Decode(%q) returned positive %q", s, g)
		}
		if len(g) != n {
			t.Errorf("Decode(%q) returned incoherent len", s)
		}
	}

	for _, s := range emojiTest {
		if !PossibleGlyph(s) {
			t.Errorf("%q returned negative", s)
		}
		if PossibleGlyph(s + "a") {
			t.Errorf("%q returned positive", s+"a")
		}
		if PossibleGlyph(s + s) {
			t.Errorf("%q returned positive", s+s)
		}
	}
}

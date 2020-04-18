This library provides tools to extract and characterize emoji character as defined per https://www.unicode.org/reports/tr51/

It builds unicode.RangeTable from https://www.unicode.org/Public/13.0.0/ucd/emoji/emoji-data.txt

The provided tables are

* `Emoji` all character that are emoji (includes emoji component such as #)
* `EmojiPresentation` all characters that have emoji presentation by default
* `EmojiModifier` all characters that are emoji modifiers
* `EmojiModifierBase` all for characters that can serve as a base for emoji modifiers (i.e person or gestures which might be combined with skin color modifiers)
* `EmojiComponent` all for characters used in emoji sequences that normally do not appear on emoji keyboards as separate choices, such as keycap base characters or RegionalIndicator characters. All characters in emoji sequences are either Emoji or EmojiComponent. Implementations must not, however, assume that all EmojiComponent characters are also Emoji. There are some non-emoji characters that are used in various emoji sequences, such as tag characters and ZWJ.
* `ExtendedPictographic` all characters that are used to future-proof segmentation. The ExtendedPictographic characters contain all the Emoji characters except for some EmojiComponent characters.
* `RegionalIndicator` all base letter for regional indicator flag
* `Tag` all possible tag character

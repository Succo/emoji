package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"unicode"
)

const maxR16 = 1 << 16

func main() {
	data, err := os.Open("emoji-data.txt")
	if err != nil {
		log.Fatalf("open emoji-data.txt %v", err)
	}
	reader := bufio.NewReader(data)

	emoji := &unicode.RangeTable{}
	emojiPresentation := &unicode.RangeTable{}
	emojiModifier := &unicode.RangeTable{}
	emojiModifierBase := &unicode.RangeTable{}
	emojiComponent := &unicode.RangeTable{}
	extendedPictographic := &unicode.RangeTable{}

	for {
		l, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("readline %v", err)
		}
		l = strings.TrimSpace(l)
		if strings.HasPrefix(l, "#") {
			continue
		}
		if len(l) == 0 {
			continue
		}
		split := strings.IndexRune(l, ';')
		var property string
		_, err = fmt.Sscanf(l[split+1:], "%s", &property)
		if err != nil {
			log.Fatalf("Sscanf property %q %v", l, err)
		}
		l = l[:split]
		var lo uint32
		var hi uint32
		if strings.IndexRune(l, '.') != -1 {
			n, err := fmt.Sscanf(l, "%X..%X", &lo, &hi)
			if err != nil || n != 2 {
				log.Fatalf("Sscanf %q %v", l, err)
			}
		} else {
			var codepoint uint32
			n, err := fmt.Sscanf(l, "%X", &codepoint)
			if err != nil || n != 1 {
				log.Fatalf("Sscanf %q %v", l, err)
			}
			lo = codepoint
			hi = codepoint
		}
		var table *unicode.RangeTable
		switch property {
		case "Emoji":
			table = emoji
		case "Emoji_Presentation":
			table = emojiPresentation
		case "Emoji_Modifier":
			table = emojiModifier
		case "Emoji_Modifier_Base":
			table = emojiModifierBase
		case "Emoji_Component":
			table = emojiComponent
		case "Extended_Pictographic#": // mising space in file
			table = extendedPictographic
		default:
			log.Fatalf("unknown table %s", property)
		}

		if lo < maxR16 && hi > maxR16 {
			log.Fatal("pair on the border")
		}
		if hi < maxR16 {
			table.R16 = append(table.R16, unicode.Range16{uint16(lo), uint16(hi), 1})
			if hi < unicode.MaxLatin1 {
				table.LatinOffset += 1 + int(hi) - int(lo)
			}
		} else {
			table.R32 = append(table.R32, unicode.Range32{lo, hi, 1})
		}
	}

	res, err := os.Create("emoji.go")
	if err != nil {
		log.Fatalf("create emoji.go %v", err)
	}

	_, err = res.Write([]byte(`package emoji

import "unicode"

`))
	if err != nil {
		log.Fatalf("Write %v", err)
	}

	_, err = res.Write([]byte("\nvar Emoji = "))
	if err != nil {
		log.Fatalf("Write %v", err)
	}
	_, err = fmt.Fprintf(res, "%#v\n", emoji)
	if err != nil {
		log.Fatalf("Fprintf %v", err)
	}

	_, err = res.Write([]byte("\nvar EmojiPresentation = "))
	if err != nil {
		log.Fatalf("Write %v", err)
	}
	_, err = fmt.Fprintf(res, "%#v\n", emojiPresentation)
	if err != nil {
		log.Fatalf("Fprintf %v", err)
	}

	_, err = res.Write([]byte("\nvar EmojiModifier = "))
	if err != nil {
		log.Fatalf("Write %v", err)
	}
	_, err = fmt.Fprintf(res, "%#v\n", emojiModifier)
	if err != nil {
		log.Fatalf("Fprintf %v", err)
	}

	_, err = res.Write([]byte("\nvar EmojiModifierBase = "))
	if err != nil {
		log.Fatalf("Write %v", err)
	}
	_, err = fmt.Fprintf(res, "%#v\n", emojiModifierBase)
	if err != nil {
		log.Fatalf("Fprintf %v", err)
	}

	_, err = res.Write([]byte("\nvar EmojiComponent = "))
	if err != nil {
		log.Fatalf("Write %v", err)
	}
	_, err = fmt.Fprintf(res, "%#v\n", emojiComponent)
	if err != nil {
		log.Fatalf("Fprintf %v", err)
	}

	_, err = res.Write([]byte("\nvar ExtendedPictographic = "))
	if err != nil {
		log.Fatalf("Write %v", err)
	}
	_, err = fmt.Fprintf(res, "%#v\n", extendedPictographic)
	if err != nil {
		log.Fatalf("Fprintf %v", err)
	}
}

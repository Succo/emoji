package main

import (
	"fmt"
	"unicode"

	"github.com/Succo/emoji"
)

func main() {
	fmt.Println(unicode.Is(emoji.Emoji, 'r'))
	fmt.Println(unicode.Is(emoji.Emoji, '😀'))
}

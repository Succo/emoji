package main

import (
	"fmt"
	"unicode"

	"github.com/Succo/emoji"
)

func main() {
	fmt.Println(unicode.Is(emoji.Table, 'r'))
	fmt.Println(unicode.Is(emoji.Table, '😀'))
}

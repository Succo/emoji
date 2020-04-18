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
	var table unicode.RangeTable
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

		l = l[:strings.IndexRune(l, ';')]
		var lo uint32
		var hi uint32
		if strings.IndexRune(l, '.') != -1 {
			n, err := fmt.Sscanf(l, "%X..%X", &lo, &hi)
			if err != nil {
				log.Fatalf("Sscanf %q %v", l, err)
			}
			fmt.Println(l, lo, hi, n)
		} else {
			var codepoint uint32
			n, err := fmt.Sscanf(l, "%X", &codepoint)
			if err != nil {
				log.Fatalf("Sscanf %q %v", l, err)
			}
			fmt.Println(l, codepoint, n)
			lo = codepoint
			hi = codepoint
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

var table = `))
	if err != nil {
		log.Fatalf("Write %v", err)
	}

	_, err = fmt.Fprintf(res, "%#v\n", table)
	if err != nil {
		log.Fatalf("Fprintf %v", err)
	}
}

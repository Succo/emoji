package emoji

import "unicode"

var Tag = &unicode.RangeTable{
	R32: []unicode.Range32{
		{Lo: uint32(0xE0030), Hi: uint32(0xE007E), Stride: 1},
	},
}

func isTag(r rune) bool {
	u := uint32(r)
	return Tag.R32[0].Lo <= u && u <= Tag.R32[0].Hi
}

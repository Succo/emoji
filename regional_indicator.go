package emoji

import "unicode"

var RegionalIndicator = &unicode.RangeTable{
	R32: []unicode.Range32{
		{Lo: uint32('🇦'), Hi: uint32('🇿'), Stride: 1},
	},
}

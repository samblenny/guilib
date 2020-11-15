// Copyright (c) 2020 Sam Blenny
// SPDX-License-Identifier: Apache-2.0 OR MIT
//
package font

import (
	"fmt"
	"strconv"
	"strings"
)

// Holds specification for a Unicode block (codepoint bounds and identifying string)
type UBlock struct {
	Low  uint32
	High uint32
	Name string
}

// Unicode blocks containing leading Unicode Scalars of grapheme clusters used
// in currently supported fonts. For full list of Unicode blocks, see
// https://www.unicode.org/Public/UCD/latest/ucd/Blocks.txt
func KnownBlocks() []UBlock {
	return []UBlock{
		UBlock{0x0000, 0x007F, "BASIC_LATIN"},                             // Latin, Emoji
		UBlock{0x0080, 0x00FF, "LATIN_1_SUPPLEMENT"},                      // Latin, Emoji
		UBlock{0x0100, 0x017F, "LATIN_EXTENDED_A"},                        // Latin
		UBlock{0x2000, 0x206F, "GENERAL_PUNCTUATION"},                     // Latin, Emoji
		UBlock{0x20A0, 0x20CF, "CURRENCY_SYMBOLS"},                        // Latin
		UBlock{0x2100, 0x214F, "LETTERLIKE_SYMBOLS"},                      // Emoji
		UBlock{0x2190, 0x21FF, "ARROWS"},                                  // Emoji
		UBlock{0x2300, 0x23FF, "MISCELLANEOUS_TECHNICAL"},                 // Emoji
		UBlock{0x2460, 0x24FF, "ENCLOSED_ALPHANUMERICS"},                  // Emoji
		UBlock{0x25A0, 0x25FF, "GEOMETRIC_SHAPES"},                        // Emoji
		UBlock{0x2600, 0x26FF, "MISCELLANEOUS_SYMBOLS"},                   // Emoji
		UBlock{0x2700, 0x27BF, "DINGBATS"},                                // Emoji
		UBlock{0x2900, 0x297F, "SUPPLEMENTAL_ARROWS_B"},                   // Emoji
		UBlock{0x2B00, 0x2BFF, "MISCELLANEOUS_SYMBOLS_AND_ARROWS"},        // Emoji
		UBlock{0x3000, 0x303F, "CJK_SYMBOLS_AND_PUNCTUATION"},             // Emoji
		UBlock{0x3200, 0x32FF, "ENCLOSED_CJK_LETTERS_AND_MONTHS"},         // Emoji
		UBlock{0xE000, 0xF8FF, "PRIVATE_USE_AREA"},                        // Latin, Emoji (UI sprites, 109)
		UBlock{0xFFF0, 0xFFFF, "SPECIALS"},                                // Latin (replacement char)
		UBlock{0x1F000, 0x1F02F, "MAHJONG_TILES"},                         // Emoji
		UBlock{0x1F0A0, 0x1F0FF, "PLAYING_CARDS"},                         // Emoji
		UBlock{0x1F100, 0x1F1FF, "ENCLOSED_ALPHANUMERIC_SUPPLEMENT"},      // Emoji
		UBlock{0x1F200, 0x1F2FF, "ENCLOSED_IDEOGRAPHIC_SUPPLEMENT"},       // Emoji
		UBlock{0x1F300, 0x1F5FF, "MISCELLANEOUS_SYMBOLS_AND_PICTOGRAPHS"}, // Emoji
		UBlock{0x1F600, 0x1F64F, "EMOTICONS"},                             // Emoji
		UBlock{0x1F680, 0x1F6FF, "TRANSPORT_AND_MAP_SYMBOLS"},             // Emoji
		UBlock{0x1F780, 0x1F7FF, "GEOMETRIC_SHAPES_EXTENDED"},             // Emoji
		UBlock{0x1F900, 0x1F9FF, "SUPPLEMENTAL_SYMBOLS_AND_PICTOGRAPHS"},  // Emoji
		UBlock{0x1FA70, 0x1FAFF, "SYMBOLS_AND_PICTOGRAPHS_EXTENDED_A"},    // Emoji
	}
}

// Given a Unicode codepoint, return the Unicode block it belongs to
func Block(c uint32) UBlock {
	for _, b := range KnownBlocks() {
		if b.Low <= c && c <= b.High {
			return b
		}
	}
	panic(fmt.Errorf("Codepoint %X belongs to an unknown Unicode block", c))
}

// Holds mappings from extended grapheme clusters to sprite sheet glyph grid coordinates
type CharSpec struct {
	HexCluster string
	Row        int
	Col        int
}

// Parse and return the first codepoint of a grapheme cluster string.
// For example, "1f3c4-200d-2640-fe0f" -> 0x1F3C4
func (cs CharSpec) FirstCodepoint() uint32 {
	base := 16
	bits := 32
	hexCodepoints := strings.Split(cs.HexCluster, "-")
	n, err := strconv.ParseUint(hexCodepoints[0], base, bits)
	if len(hexCodepoints) < 1 || err != nil {
		panic(fmt.Errorf("unexpected value for CharSpec.HexCluster: %q", cs))
	}
	return uint32(n)
}

// Convert a grapheme cluster string from hexidecimal codepoints to a string.
// For example, "1f3c4-200d-2640-fe0f" -> "\U0001F3C4\u200d\u2640\ufe0f"
func (cs CharSpec) GraphemeCluster() string {
	base := 16
	bits := 32
	cluster := ""
	hexCodepoints := strings.Split(cs.HexCluster, "-")
	if len(hexCodepoints) < 1 {
		panic(fmt.Errorf("unexpected value for CharSpec.HexCluster: %q", cs))
	}
	for _, hc := range hexCodepoints {
		n, err := strconv.ParseUint(hc, base, bits)
		if err != nil {
			panic(fmt.Errorf("unexpected value for CharSpec.HexCluster: %q", cs))
		}
		cluster += string(rune(n))
	}
	return cluster
}

func SysLatinMap() []CharSpec {
	return []CharSpec{
		// Unicode Basic Latin block
		CharSpec{"20", 0, 2},  // " "
		CharSpec{"21", 1, 2},  // "!"
		CharSpec{"22", 2, 2},  // "\""
		CharSpec{"23", 3, 2},  // "#"
		CharSpec{"24", 4, 2},  // "$"
		CharSpec{"25", 5, 2},  // "%"
		CharSpec{"26", 6, 2},  // "&"
		CharSpec{"27", 7, 2},  // "'"
		CharSpec{"28", 8, 2},  // "("
		CharSpec{"29", 9, 2},  // ")"
		CharSpec{"2A", 10, 2}, // "*"
		CharSpec{"2B", 11, 2}, // "+"
		CharSpec{"2C", 12, 2}, // ","
		CharSpec{"2D", 13, 2}, // "-"
		CharSpec{"2E", 14, 2}, // "."
		CharSpec{"2F", 15, 2}, // "/"
		CharSpec{"30", 0, 3},  // "0"
		CharSpec{"31", 1, 3},  // "1"
		CharSpec{"32", 2, 3},  // "2"
		CharSpec{"33", 3, 3},  // "3"
		CharSpec{"34", 4, 3},  // "4"
		CharSpec{"35", 5, 3},  // "5"
		CharSpec{"36", 6, 3},  // "6"
		CharSpec{"37", 7, 3},  // "7"
		CharSpec{"38", 8, 3},  // "8"
		CharSpec{"39", 9, 3},  // "9"
		CharSpec{"3A", 10, 3}, // ":"
		CharSpec{"3B", 11, 3}, // ";"
		CharSpec{"3C", 12, 3}, // "<"
		CharSpec{"3D", 13, 3}, // "="
		CharSpec{"3E", 14, 3}, // ">"
		CharSpec{"3F", 15, 3}, // "?"
		CharSpec{"40", 0, 4},  // "@"
		CharSpec{"41", 1, 4},  // "A"
		CharSpec{"42", 2, 4},  // "B"
		CharSpec{"43", 3, 4},  // "C"
		CharSpec{"44", 4, 4},  // "D"
		CharSpec{"45", 5, 4},  // "E"
		CharSpec{"46", 6, 4},  // "F"
		CharSpec{"47", 7, 4},  // "G"
		CharSpec{"48", 8, 4},  // "H"
		CharSpec{"49", 9, 4},  // "I"
		CharSpec{"4A", 10, 4}, // "J"
		CharSpec{"4B", 11, 4}, // "K"
		CharSpec{"4C", 12, 4}, // "L"
		CharSpec{"4D", 13, 4}, // "M"
		CharSpec{"4E", 14, 4}, // "N"
		CharSpec{"4F", 15, 4}, // "O"
		CharSpec{"50", 0, 5},  // "P"
		CharSpec{"51", 1, 5},  // "Q"
		CharSpec{"52", 2, 5},  // "R"
		CharSpec{"53", 3, 5},  // "S"
		CharSpec{"54", 4, 5},  // "T"
		CharSpec{"55", 5, 5},  // "U"
		CharSpec{"56", 6, 5},  // "V"
		CharSpec{"57", 7, 5},  // "W"
		CharSpec{"58", 8, 5},  // "X"
		CharSpec{"59", 9, 5},  // "Y"
		CharSpec{"5A", 10, 5}, // "Z"
		CharSpec{"5B", 11, 5}, // "["
		CharSpec{"5C", 12, 5}, // "\\"
		CharSpec{"5D", 13, 5}, // "]"
		CharSpec{"5E", 14, 5}, // "^"
		CharSpec{"5F", 15, 5}, // "_"
		CharSpec{"60", 0, 6},  // "`"
		CharSpec{"61", 1, 6},  // "a"
		CharSpec{"62", 2, 6},  // "b"
		CharSpec{"63", 3, 6},  // "c"
		CharSpec{"64", 4, 6},  // "d"
		CharSpec{"65", 5, 6},  // "e"
		CharSpec{"66", 6, 6},  // "f"
		CharSpec{"67", 7, 6},  // "g"
		CharSpec{"68", 8, 6},  // "h"
		CharSpec{"69", 9, 6},  // "i"
		CharSpec{"6A", 10, 6}, // "j"
		CharSpec{"6B", 11, 6}, // "k"
		CharSpec{"6C", 12, 6}, // "l"
		CharSpec{"6D", 13, 6}, // "m"
		CharSpec{"6E", 14, 6}, // "n"
		CharSpec{"6F", 15, 6}, // "o"
		CharSpec{"70", 0, 7},  // "p"
		CharSpec{"71", 1, 7},  // "q"
		CharSpec{"72", 2, 7},  // "r"
		CharSpec{"73", 3, 7},  // "s"
		CharSpec{"74", 4, 7},  // "t"
		CharSpec{"75", 5, 7},  // "u"
		CharSpec{"76", 6, 7},  // "v"
		CharSpec{"77", 7, 7},  // "w"
		CharSpec{"78", 8, 7},  // "x"
		CharSpec{"79", 9, 7},  // "y"
		CharSpec{"7A", 10, 7}, // "z"
		CharSpec{"7B", 11, 7}, // "{"
		CharSpec{"7C", 12, 7}, // "|"
		CharSpec{"7D", 13, 7}, // "}"
		CharSpec{"7E", 14, 7}, // "~"

		// Unicode Latin 1 block
		CharSpec{"A0", 0, 2},   // No-Break Space
		CharSpec{"A1", 1, 12},  // "¡"
		CharSpec{"A2", 2, 10},  // "¢"
		CharSpec{"A3", 3, 10},  // "£"
		CharSpec{"A4", 15, 1},  // "¤"
		CharSpec{"A5", 4, 11},  // "¥"
		CharSpec{"A6", 15, 7},  // "¦"
		CharSpec{"A7", 4, 10},  // "§"
		CharSpec{"A8", 12, 10}, // "¨"
		CharSpec{"A9", 9, 10},  // "©"
		CharSpec{"AA", 11, 11}, // "ª"
		CharSpec{"AB", 7, 12},  // "«"
		CharSpec{"AC", 2, 12},  // "¬"
		CharSpec{"AD", 13, 2},  // Soft Hyphen
		CharSpec{"AE", 8, 10},  // "®"
		CharSpec{"AF", 8, 15},  // "¯" Macron
		CharSpec{"B0", 1, 10},  // "°" Degree Sign
		CharSpec{"B1", 1, 11},  // "±"
		CharSpec{"B2", 3, 1},   // "²"
		CharSpec{"B3", 4, 1},   // "³"
		CharSpec{"B4", 11, 10}, // "´"
		CharSpec{"B5", 5, 11},  // "µ"
		CharSpec{"B6", 6, 10},  // "¶"
		CharSpec{"B7", 1, 14},  // "·"
		CharSpec{"B8", 12, 15}, // "¸" Cedillia
		CharSpec{"B9", 2, 1},   // "¹"
		CharSpec{"BA", 12, 11}, // "º"
		CharSpec{"BB", 8, 12},  // "»"
		CharSpec{"BC", 5, 1},   // "¼"
		CharSpec{"BD", 6, 1},   // "½"
		CharSpec{"BE", 7, 1},   // "¾"
		CharSpec{"BF", 0, 12},  // "¿"
		CharSpec{"C0", 11, 12}, // "À"
		CharSpec{"C1", 7, 14},  // "Á"
		CharSpec{"C2", 5, 14},  // "Â"
		CharSpec{"C3", 12, 12}, // "Ã"
		CharSpec{"C4", 0, 8},   // "Ä"
		CharSpec{"C5", 1, 8},   // "Å"
		CharSpec{"C6", 14, 10}, // "Æ"
		CharSpec{"C7", 2, 8},   // "Ç"
		CharSpec{"C8", 9, 14},  // "È"
		CharSpec{"C9", 3, 8},   // "É"
		CharSpec{"CA", 6, 14},  // "Ê"
		CharSpec{"CB", 8, 14},  // "Ë"
		CharSpec{"CC", 13, 14}, // "Ì"
		CharSpec{"CD", 10, 14}, // "Í"
		CharSpec{"CE", 11, 14}, // "Î"
		CharSpec{"CF", 12, 14}, // "Ï"
		CharSpec{"D0", 8, 1},   // "Ð"
		CharSpec{"D1", 4, 8},   // "Ñ"
		CharSpec{"D2", 1, 15},  // "Ò"
		CharSpec{"D3", 14, 14}, // "Ó"
		CharSpec{"D4", 15, 14}, // "Ô"
		CharSpec{"D5", 13, 12}, // "Õ"
		CharSpec{"D6", 5, 8},   // "Ö"
		CharSpec{"D7", 9, 1},   // "×" Multiplication Sign
		CharSpec{"D8", 15, 10}, // "Ø"
		CharSpec{"D9", 4, 15},  // "Ù"
		CharSpec{"DA", 2, 15},  // "Ú"
		CharSpec{"DB", 3, 15},  // "Û"
		CharSpec{"DC", 6, 8},   // "Ü"
		CharSpec{"DD", 10, 1},  // "Ý"
		CharSpec{"DE", 11, 1},  // "Þ"
		CharSpec{"DF", 7, 10},  // "ß"
		CharSpec{"E0", 8, 8},   // "à"
		CharSpec{"E1", 7, 8},   // "á"
		CharSpec{"E2", 9, 8},   // "â"
		CharSpec{"E3", 11, 8},  // "ã"
		CharSpec{"E4", 10, 8},  // "ä"
		CharSpec{"E5", 12, 8},  // "å"
		CharSpec{"E6", 14, 11}, // "æ"
		CharSpec{"E7", 13, 8},  // "ç"
		CharSpec{"E8", 15, 8},  // "è"
		CharSpec{"E9", 14, 8},  // "é"
		CharSpec{"EA", 0, 9},   // "ê"
		CharSpec{"EB", 1, 9},   // "ë"
		CharSpec{"EC", 3, 9},   // "ì"
		CharSpec{"ED", 2, 9},   // "í"
		CharSpec{"EE", 4, 9},   // "î"
		CharSpec{"EF", 5, 9},   // "ï"
		CharSpec{"F0", 12, 1},  // "ð"
		CharSpec{"F1", 6, 9},   // "ñ"
		CharSpec{"F2", 8, 9},   // "ò"
		CharSpec{"F3", 7, 9},   // "ó"
		CharSpec{"F4", 9, 9},   // "ô"
		CharSpec{"F5", 11, 9},  // "õ"
		CharSpec{"F6", 10, 9},  // "ö"
		CharSpec{"F7", 6, 13},  // "÷"
		CharSpec{"F8", 15, 11}, // "ø"
		CharSpec{"F9", 13, 9},  // "ù"
		CharSpec{"FA", 12, 9},  // "ú"
		CharSpec{"FB", 14, 9},  // "û"
		CharSpec{"FC", 15, 9},  // "ü"
		CharSpec{"FD", 13, 1},  // "ý"
		CharSpec{"FE", 14, 1},  // "þ"
		CharSpec{"FF", 8, 13},  // "ÿ"

		// Unicode Latin Extended A block
		CharSpec{"152", 14, 12}, // "Œ"
		CharSpec{"153", 15, 12}, // "œ"

		// Unicode General Punctuation block
		CharSpec{"2018", 4, 13}, // "‘" Left Single Quotation Mark
		CharSpec{"2019", 5, 13}, // "’" Right Single Quotation Mark
		CharSpec{"201A", 2, 14}, // "‚" Single Low-9 Quotation Mark
		CharSpec{"201B", 7, 11}, // "‛" Single High-Reversed-9 Quotation Mark
		CharSpec{"201C", 2, 13}, // "“" Left Double Quotation Mark
		CharSpec{"201D", 3, 13}, // "”" Right Double Quotation Mark
		CharSpec{"201E", 3, 14}, // "„" Double Low-9 Quotation Mark
		CharSpec{"201F", 8, 11}, // "‟" Double High-Reversed-9 Quotation Mark
		CharSpec{"2020", 0, 10}, // "†"
		CharSpec{"2021", 0, 14}, // "‡"
		CharSpec{"2022", 5, 10}, // "•"

		// Unicode Currency Symbols block
		CharSpec{"20AC", 11, 13}, // "€"

		// Unicode Private Use Area assignments for UI sprites
		CharSpec{"E700", 0, 0},  // Battery_05
		CharSpec{"E701", 1, 0},  // Battery_25
		CharSpec{"E702", 2, 0},  // Battery_50
		CharSpec{"E703", 3, 0},  // Battery_75
		CharSpec{"E704", 4, 0},  // Battery_99
		CharSpec{"E705", 5, 0},  // Radio_3
		CharSpec{"E706", 6, 0},  // Radio_2
		CharSpec{"E707", 7, 0},  // Radio_1
		CharSpec{"E708", 8, 0},  // Radio_0
		CharSpec{"E709", 9, 0},  // Radio_Off
		CharSpec{"E70A", 13, 0}, // Shift_Arrow
		CharSpec{"E70B", 14, 0}, // Backspace_Symbol
		CharSpec{"E70C", 15, 0}, // Enter_Symbol

		// Unicode Specials Block
		CharSpec{"FFFD", 0, 15}, // "�"
	}
}

// Copyright (c) 2020 Sam Blenny
// SPDX-License-Identifier: Apache-2.0 OR MIT
//
package font

import (
	"fmt"
	"strconv"
	"strings"
)

type UnicodeBlock int

// Unicode blocks containing leading Unicode Scalar of grapheme clusters used
// in currently supported fonts. For full list of Unicode blocks, see
// https://www.unicode.org/Public/UCD/latest/ucd/Blocks.txt
const (
	BASIC_LATIN                           UnicodeBlock = iota // 0000..007F: Latin, Emoji
	LATIN_1_SUPPLEMENT                                        // 0080..00FF: Latin, Emoji
	LATIN_EXTENDED_A                                          // 0100..017F: Latin
	GENERAL_PUNCTUATION                                       // 2000..206F: Latin, Emoji
	CURRENCY_SYMBOLS                                          // 20A0..20CF: Latin
	LETTERLIKE_SYMBOLS                                        // 2100..214F: Emoji
	ARROWS                                                    // 2190..21FF: Emoji
	MISCELLANEOUS_TECHNICAL                                   // 2300..23FF: Emoji
	ENCLOSED_ALPHANUMERICS                                    // 2460..24FF: Emoji
	GEOMETRIC_SHAPES                                          // 25A0..25FF: Emoji
	MISCELLANEOUS_SYMBOLS                                     // 2600..26FF: Emoji
	DINGBATS                                                  // 2700..27BF: Emoji
	SUPPLEMENTAL_ARROWS_B                                     // 2900..297F: Emoji
	MISCELLANEOUS_SYMBOLS_AND_ARROWS                          // 2B00..2BFF: Emoji
	CJK_SYMBOLS_AND_PUNCTUATION                               // 3000..303F: Emoji
	ENCLOSED_CJK_LETTERS_AND_MONTHS                           // 3200..32FF: Emoji
	PRIVATE_USE_AREA                                          // E000..F8FF: Latin, Emoji
	SPECIALS                                                  // FFF0..FFFF: Latin (replacement symbol)
	MAHJONG_TILES                                             // 1F000..1F02F: Emoji
	PLAYING_CARDS                                             // 1F0A0..1F0FF: Emoji
	ENCLOSED_ALPHANUMERIC_SUPPLEMENT                          // 1F100..1F1FF: Emoji (many)
	ENCLOSED_IDEOGRAPHIC_SUPPLEMENT                           // 1F200..1F2FF: Emoji
	MISCELLANEOUS_SYMBOLS_AND_PICTOGRAPHS                     // 1F300..1F5FF: Emoji (very many)
	EMOTICONS                                                 // 1F600..1F64F: Emoji (many)
	TRANSPORT_AND_MAP_SYMBOLS                                 // 1F680..1F6FF: Emoji (many)
	GEOMETRIC_SHAPES_EXTENDED                                 // 1F780..1F7FF: Emoji
	SUPPLEMENTAL_SYMBOLS_AND_PICTOGRAPHS                      // 1F900..1F9FF: Emoji (very many)
	SYMBOLS_AND_PICTOGRAPHS_EXTENDED_A                        // 1FA70..1FAFF: Emoji (many)
	UNKNOWN
)

type CharSpec struct {
	HexCluster      string
	Row             int
	Col             int
	GraphemeCluster string
}

func (cs CharSpec) FirstCodepoint() uint32 {
	base := 16
	bits := 32
	hexCodepoints := strings.Split(cs.HexCluster, "-")
	n, err := strconv.ParseUint(hexCodepoints[0], base, bits)
	if len(hexCodepoints) < 1 || err != nil {
		fmt.Println(cs)
		panic("unexpected value for CharSpec.HexCluster")
	}
	return uint32(n)
}

// Convert Unicode codepoint to its block
func Block(c uint32) UnicodeBlock {
	switch {
	case 0x0000 <= c && c <= 0x007F:
		return BASIC_LATIN
	case 0x0080 <= c && c <= 0x00FF:
		return LATIN_1_SUPPLEMENT
	case 0x0100 <= c && c <= 0x017F:
		return LATIN_EXTENDED_A
	case 0x2000 <= c && c <= 0x206F:
		return GENERAL_PUNCTUATION
	case 0x20A0 <= c && c <= 0x20CF:
		return CURRENCY_SYMBOLS
	case 0x2100 <= c && c <= 0x214F:
		return LETTERLIKE_SYMBOLS
	case 0x2190 <= c && c <= 0x21FF:
		return ARROWS
	case 0x2300 <= c && c <= 0x23FF:
		return MISCELLANEOUS_TECHNICAL
	case 0x2460 <= c && c <= 0x24FF:
		return ENCLOSED_ALPHANUMERICS
	case 0x25A0 <= c && c <= 0x25FF:
		return GEOMETRIC_SHAPES
	case 0x2600 <= c && c <= 0x26FF:
		return MISCELLANEOUS_SYMBOLS
	case 0x2700 <= c && c <= 0x27BF:
		return DINGBATS
	case 0x2900 <= c && c <= 0x297F:
		return SUPPLEMENTAL_ARROWS_B
	case 0x2B00 <= c && c <= 0x2BFF:
		return MISCELLANEOUS_SYMBOLS_AND_ARROWS
	case 0x3000 <= c && c <= 0x303F:
		return CJK_SYMBOLS_AND_PUNCTUATION
	case 0x3200 <= c && c <= 0x32FF:
		return ENCLOSED_CJK_LETTERS_AND_MONTHS
	case 0xE000 <= c && c <= 0xF8FF:
		return PRIVATE_USE_AREA
	case 0xFFF0 <= c && c <= 0xFFFF:
		return SPECIALS
	case 0x1F000 <= c && c <= 0x1F02F:
		return MAHJONG_TILES
	case 0x1F0A0 <= c && c <= 0x1F0FF:
		return PLAYING_CARDS
	case 0x1F100 <= c && c <= 0x1F1FF:
		return ENCLOSED_ALPHANUMERIC_SUPPLEMENT
	case 0x1F200 <= c && c <= 0x1F2FF:
		return ENCLOSED_IDEOGRAPHIC_SUPPLEMENT
	case 0x1F300 <= c && c <= 0x1F5FF:
		return MISCELLANEOUS_SYMBOLS_AND_PICTOGRAPHS
	case 0x1F600 <= c && c <= 0x1F64F:
		return EMOTICONS
	case 0x1F680 <= c && c <= 0x1F6FF:
		return TRANSPORT_AND_MAP_SYMBOLS
	case 0x1F780 <= c && c <= 0x1F7FF:
		return GEOMETRIC_SHAPES_EXTENDED
	case 0x1F900 <= c && c <= 0x1F9FF:
		return SUPPLEMENTAL_SYMBOLS_AND_PICTOGRAPHS
	case 0x1FA70 <= c && c <= 0x1FAFF:
		return SYMBOLS_AND_PICTOGRAPHS_EXTENDED_A
	default:
		return UNKNOWN
	}
}

func SysLatinMap() []CharSpec {
	return []CharSpec{
		// Unicode Basic Latin block
		CharSpec{"20", 0, 2, " "},
		CharSpec{"21", 1, 2, "!"},
		CharSpec{"22", 2, 2, "\""},
		CharSpec{"23", 3, 2, "#"},
		CharSpec{"24", 4, 2, "$"},
		CharSpec{"25", 5, 2, "%"},
		CharSpec{"26", 6, 2, "&"},
		CharSpec{"27", 7, 2, "\\'"},
		CharSpec{"28", 8, 2, "("},
		CharSpec{"29", 9, 2, ")"},
		CharSpec{"2A", 10, 2, "*"},
		CharSpec{"2B", 11, 2, "+"},
		CharSpec{"2C", 12, 2, ","},
		CharSpec{"2D", 13, 2, "-"},
		CharSpec{"2E", 14, 2, "."},
		CharSpec{"2F", 15, 2, "/"},
		CharSpec{"30", 0, 3, "0"},
		CharSpec{"31", 1, 3, "1"},
		CharSpec{"32", 2, 3, "2"},
		CharSpec{"33", 3, 3, "3"},
		CharSpec{"34", 4, 3, "4"},
		CharSpec{"35", 5, 3, "5"},
		CharSpec{"36", 6, 3, "6"},
		CharSpec{"37", 7, 3, "7"},
		CharSpec{"38", 8, 3, "8"},
		CharSpec{"39", 9, 3, "9"},
		CharSpec{"3A", 10, 3, ":"},
		CharSpec{"3B", 11, 3, ";"},
		CharSpec{"3C", 12, 3, "<"},
		CharSpec{"3D", 13, 3, "="},
		CharSpec{"3E", 14, 3, ">"},
		CharSpec{"3F", 15, 3, "?"},
		CharSpec{"40", 0, 4, "@"},
		CharSpec{"41", 1, 4, "A"},
		CharSpec{"42", 2, 4, "B"},
		CharSpec{"43", 3, 4, "C"},
		CharSpec{"44", 4, 4, "D"},
		CharSpec{"45", 5, 4, "E"},
		CharSpec{"46", 6, 4, "F"},
		CharSpec{"47", 7, 4, "G"},
		CharSpec{"48", 8, 4, "H"},
		CharSpec{"49", 9, 4, "I"},
		CharSpec{"4A", 10, 4, "J"},
		CharSpec{"4B", 11, 4, "K"},
		CharSpec{"4C", 12, 4, "L"},
		CharSpec{"4D", 13, 4, "M"},
		CharSpec{"4E", 14, 4, "N"},
		CharSpec{"4F", 15, 4, "O"},
		CharSpec{"50", 0, 5, "P"},
		CharSpec{"51", 1, 5, "Q"},
		CharSpec{"52", 2, 5, "R"},
		CharSpec{"53", 3, 5, "S"},
		CharSpec{"54", 4, 5, "T"},
		CharSpec{"55", 5, 5, "U"},
		CharSpec{"56", 6, 5, "V"},
		CharSpec{"57", 7, 5, "W"},
		CharSpec{"58", 8, 5, "X"},
		CharSpec{"59", 9, 5, "Y"},
		CharSpec{"5A", 10, 5, "Z"},
		CharSpec{"5B", 11, 5, "["},
		CharSpec{"5C", 12, 5, "\\\\"},
		CharSpec{"5D", 13, 5, "]"},
		CharSpec{"5E", 14, 5, "^"},
		CharSpec{"5F", 15, 5, "_"},
		CharSpec{"60", 0, 6, "`"},
		CharSpec{"61", 1, 6, "a"},
		CharSpec{"62", 2, 6, "b"},
		CharSpec{"63", 3, 6, "c"},
		CharSpec{"64", 4, 6, "d"},
		CharSpec{"65", 5, 6, "e"},
		CharSpec{"66", 6, 6, "f"},
		CharSpec{"67", 7, 6, "g"},
		CharSpec{"68", 8, 6, "h"},
		CharSpec{"69", 9, 6, "i"},
		CharSpec{"6A", 10, 6, "j"},
		CharSpec{"6B", 11, 6, "k"},
		CharSpec{"6C", 12, 6, "l"},
		CharSpec{"6D", 13, 6, "m"},
		CharSpec{"6E", 14, 6, "n"},
		CharSpec{"6F", 15, 6, "o"},
		CharSpec{"70", 0, 7, "p"},
		CharSpec{"71", 1, 7, "q"},
		CharSpec{"72", 2, 7, "r"},
		CharSpec{"73", 3, 7, "s"},
		CharSpec{"74", 4, 7, "t"},
		CharSpec{"75", 5, 7, "u"},
		CharSpec{"76", 6, 7, "v"},
		CharSpec{"77", 7, 7, "w"},
		CharSpec{"78", 8, 7, "x"},
		CharSpec{"79", 9, 7, "y"},
		CharSpec{"7A", 10, 7, "z"},
		CharSpec{"7B", 11, 7, "{"},
		CharSpec{"7C", 12, 7, "|"},
		CharSpec{"7D", 13, 7, "}"},
		CharSpec{"7E", 14, 7, "~"},

		// Unicode Latin 1 block
		CharSpec{"A0", 0, 2, "\u00A0"}, // No-Break Space
		CharSpec{"A1", 1, 12, "¡"},
		CharSpec{"A2", 2, 10, "¢"},
		CharSpec{"A3", 3, 10, "£"},
		CharSpec{"A4", 15, 1, "¤"},
		CharSpec{"A5", 4, 11, "¥"},
		CharSpec{"A6", 15, 7, "¦"},
		CharSpec{"A7", 4, 10, "§"},
		CharSpec{"A8", 12, 10, "¨"},
		CharSpec{"A9", 9, 10, "©"},
		CharSpec{"AA", 11, 11, "ª"},
		CharSpec{"AB", 7, 12, "«"},
		CharSpec{"AC", 2, 12, "¬"},
		CharSpec{"AD", 13, 2, "\u00AD"}, // Soft Hyphen
		CharSpec{"AE", 8, 10, "®"},
		CharSpec{"AF", 8, 15, "¯"}, // Macron
		CharSpec{"B0", 1, 10, "°"}, // Degree Sign
		CharSpec{"B1", 1, 11, "±"},
		CharSpec{"B2", 3, 1, "²"},
		CharSpec{"B3", 4, 1, "³"},
		CharSpec{"B4", 11, 10, "´"},
		CharSpec{"B5", 5, 11, "µ"},
		CharSpec{"B6", 6, 10, "¶"},
		CharSpec{"B7", 1, 14, "·"},
		CharSpec{"B8", 12, 15, "¸"}, // Cedillia
		CharSpec{"B9", 2, 1, "¹"},
		CharSpec{"BA", 12, 11, "º"},
		CharSpec{"BB", 8, 12, "»"},
		CharSpec{"BC", 5, 1, "¼"},
		CharSpec{"BD", 6, 1, "½"},
		CharSpec{"BE", 7, 1, "¾"},
		CharSpec{"BF", 0, 12, "¿"},
		CharSpec{"C0", 11, 12, "À"},
		CharSpec{"C1", 7, 14, "Á"},
		CharSpec{"C2", 5, 14, "Â"},
		CharSpec{"C3", 12, 12, "Ã"},
		CharSpec{"C4", 0, 8, "Ä"},
		CharSpec{"C5", 1, 8, "Å"},
		CharSpec{"C6", 14, 10, "Æ"},
		CharSpec{"C7", 2, 8, "Ç"},
		CharSpec{"C8", 9, 14, "È"},
		CharSpec{"C9", 3, 8, "É"},
		CharSpec{"CA", 6, 14, "Ê"},
		CharSpec{"CB", 8, 14, "Ë"},
		CharSpec{"CC", 13, 14, "Ì"},
		CharSpec{"CD", 10, 14, "Í"},
		CharSpec{"CE", 11, 14, "Î"},
		CharSpec{"CF", 12, 14, "Ï"},
		CharSpec{"D0", 8, 1, "Ð"},
		CharSpec{"D1", 4, 8, "Ñ"},
		CharSpec{"D2", 1, 15, "Ò"},
		CharSpec{"D3", 14, 14, "Ó"},
		CharSpec{"D4", 15, 14, "Ô"},
		CharSpec{"D5", 13, 12, "Õ"},
		CharSpec{"D6", 5, 8, "Ö"},
		CharSpec{"D7", 9, 1, "×"}, // Multiplication Sign
		CharSpec{"D8", 15, 10, "Ø"},
		CharSpec{"D9", 4, 15, "Ù"},
		CharSpec{"DA", 2, 15, "Ú"},
		CharSpec{"DB", 3, 15, "Û"},
		CharSpec{"DC", 6, 8, "Ü"},
		CharSpec{"DD", 10, 1, "Ý"},
		CharSpec{"DE", 11, 1, "Þ"},
		CharSpec{"DF", 7, 10, "ß"},
		CharSpec{"E0", 8, 8, "à"},
		CharSpec{"E1", 7, 8, "á"},
		CharSpec{"E2", 9, 8, "â"},
		CharSpec{"E3", 11, 8, "ã"},
		CharSpec{"E4", 10, 8, "ä"},
		CharSpec{"E5", 12, 8, "å"},
		CharSpec{"E6", 14, 11, "æ"},
		CharSpec{"E7", 13, 8, "ç"},
		CharSpec{"E8", 15, 8, "è"},
		CharSpec{"E9", 14, 8, "é"},
		CharSpec{"EA", 0, 9, "ê"},
		CharSpec{"EB", 1, 9, "ë"},
		CharSpec{"EC", 3, 9, "ì"},
		CharSpec{"ED", 2, 9, "í"},
		CharSpec{"EE", 4, 9, "î"},
		CharSpec{"EF", 5, 9, "ï"},
		CharSpec{"F0", 12, 1, "ð"},
		CharSpec{"F1", 6, 9, "ñ"},
		CharSpec{"F2", 8, 9, "ò"},
		CharSpec{"F3", 7, 9, "ó"},
		CharSpec{"F4", 9, 9, "ô"},
		CharSpec{"F5", 11, 9, "õ"},
		CharSpec{"F6", 10, 9, "ö"},
		CharSpec{"F7", 6, 13, "÷"},
		CharSpec{"F8", 15, 11, "ø"},
		CharSpec{"F9", 13, 9, "ù"},
		CharSpec{"FA", 12, 9, "ú"},
		CharSpec{"FB", 14, 9, "û"},
		CharSpec{"FC", 15, 9, "ü"},
		CharSpec{"FD", 13, 1, "ý"},
		CharSpec{"FE", 14, 1, "þ"},
		CharSpec{"FF", 8, 13, "ÿ"},

		// Unicode Latin Extended A block
		CharSpec{"152", 14, 12, "Œ"},
		CharSpec{"153", 15, 12, "œ"},

		// Unicode General Punctuation block
		CharSpec{"2018", 4, 13, "‘"}, // Left Single Quotation Mark
		CharSpec{"2019", 5, 13, "’"}, // Right Single Quotation Mark
		CharSpec{"201A", 2, 14, "‚"}, // Single Low-9 Quotation Mark
		CharSpec{"201B", 7, 11, "‛"}, // Single High-Reversed-9 Quotation Mark
		CharSpec{"201C", 2, 13, "“"}, // Left Double Quotation Mark
		CharSpec{"201D", 3, 13, "”"}, // Right Double Quotation Mark
		CharSpec{"201E", 3, 14, "„"}, // Double Low-9 Quotation Mark
		CharSpec{"201F", 8, 11, "‟"}, // Double High-Reversed-9 Quotation Mark
		CharSpec{"2020", 0, 10, "†"}, // Dagger
		CharSpec{"2021", 0, 14, "‡"}, // Double Dagger
		CharSpec{"2022", 5, 10, "•"}, // Bullet

		// Unicode Currency Symbols block
		CharSpec{"20AC", 11, 13, "€"},

		// Unicode Private Use Area assignments for UI sprites
		CharSpec{"E700", 0, 0, "\uE700"},  // Battery_05
		CharSpec{"E701", 1, 0, "\uE701"},  // Battery_25
		CharSpec{"E702", 2, 0, "\uE702"},  // Battery_50
		CharSpec{"E703", 3, 0, "\uE703"},  // Battery_75
		CharSpec{"E704", 4, 0, "\uE704"},  // Battery_99
		CharSpec{"E705", 5, 0, "\uE705"},  // Radio_3
		CharSpec{"E706", 6, 0, "\uE706"},  // Radio_2
		CharSpec{"E707", 7, 0, "\uE707"},  // Radio_1
		CharSpec{"E708", 8, 0, "\uE708"},  // Radio_0
		CharSpec{"E709", 9, 0, "\uE709"},  // Radio_Off
		CharSpec{"E70A", 13, 0, "\uE70A"}, // Shift_Arrow
		CharSpec{"E70B", 14, 0, "\uE70B"}, // Backspace_Symbol
		CharSpec{"E70C", 15, 0, "\uE70C"}, // Enter_Symbol

		// Unicode Specials Block
		CharSpec{"FFFD", 0, 15, "�"},
	}
}

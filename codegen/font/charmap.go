package font

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
	FirstCodepoint  uint32
	Row             int
	Col             int
	GraphemeCluster string
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
		CharSpec{0x20, 0, 2, " "},
		CharSpec{0x21, 1, 2, "!"},
		CharSpec{0x22, 2, 2, "\""},
		CharSpec{0x23, 3, 2, "#"},
		CharSpec{0x24, 4, 2, "$"},
		CharSpec{0x25, 5, 2, "%"},
		CharSpec{0x26, 6, 2, "&"},
		CharSpec{0x27, 7, 2, "\\'"},
		CharSpec{0x28, 8, 2, "("},
		CharSpec{0x29, 9, 2, ")"},
		CharSpec{0x2A, 10, 2, "*"},
		CharSpec{0x2B, 11, 2, "+"},
		CharSpec{0x2C, 12, 2, ","},
		CharSpec{0x2D, 13, 2, "-"},
		CharSpec{0x2E, 14, 2, "."},
		CharSpec{0x2F, 15, 2, "/"},
		CharSpec{0x30, 0, 3, "0"},
		CharSpec{0x31, 1, 3, "1"},
		CharSpec{0x32, 2, 3, "2"},
		CharSpec{0x33, 3, 3, "3"},
		CharSpec{0x34, 4, 3, "4"},
		CharSpec{0x35, 5, 3, "5"},
		CharSpec{0x36, 6, 3, "6"},
		CharSpec{0x37, 7, 3, "7"},
		CharSpec{0x38, 8, 3, "8"},
		CharSpec{0x39, 9, 3, "9"},
		CharSpec{0x3A, 10, 3, ":"},
		CharSpec{0x3B, 11, 3, ";"},
		CharSpec{0x3C, 12, 3, "<"},
		CharSpec{0x3D, 13, 3, "="},
		CharSpec{0x3E, 14, 3, ">"},
		CharSpec{0x3F, 15, 3, "?"},
		CharSpec{0x40, 0, 4, "@"},
		CharSpec{0x41, 1, 4, "A"},
		CharSpec{0x42, 2, 4, "B"},
		CharSpec{0x43, 3, 4, "C"},
		CharSpec{0x44, 4, 4, "D"},
		CharSpec{0x45, 5, 4, "E"},
		CharSpec{0x46, 6, 4, "F"},
		CharSpec{0x47, 7, 4, "G"},
		CharSpec{0x48, 8, 4, "H"},
		CharSpec{0x49, 9, 4, "I"},
		CharSpec{0x4A, 10, 4, "J"},
		CharSpec{0x4B, 11, 4, "K"},
		CharSpec{0x4C, 12, 4, "L"},
		CharSpec{0x4D, 13, 4, "M"},
		CharSpec{0x4E, 14, 4, "N"},
		CharSpec{0x4F, 15, 4, "O"},
		CharSpec{0x50, 0, 5, "P"},
		CharSpec{0x51, 1, 5, "Q"},
		CharSpec{0x52, 2, 5, "R"},
		CharSpec{0x53, 3, 5, "S"},
		CharSpec{0x54, 4, 5, "T"},
		CharSpec{0x55, 5, 5, "U"},
		CharSpec{0x56, 6, 5, "V"},
		CharSpec{0x57, 7, 5, "W"},
		CharSpec{0x58, 8, 5, "X"},
		CharSpec{0x59, 9, 5, "Y"},
		CharSpec{0x5A, 10, 5, "Z"},
		CharSpec{0x5B, 11, 5, "["},
		CharSpec{0x5C, 12, 5, "\\\\"},
		CharSpec{0x5D, 13, 5, "]"},
		CharSpec{0x5E, 14, 5, "^"},
		CharSpec{0x5F, 15, 5, "_"},
		CharSpec{0x60, 0, 6, "`"},
		CharSpec{0x61, 1, 6, "a"},
		CharSpec{0x62, 2, 6, "b"},
		CharSpec{0x63, 3, 6, "c"},
		CharSpec{0x64, 4, 6, "d"},
		CharSpec{0x65, 5, 6, "e"},
		CharSpec{0x66, 6, 6, "f"},
		CharSpec{0x67, 7, 6, "g"},
		CharSpec{0x68, 8, 6, "h"},
		CharSpec{0x69, 9, 6, "i"},
		CharSpec{0x6A, 10, 6, "j"},
		CharSpec{0x6B, 11, 6, "k"},
		CharSpec{0x6C, 12, 6, "l"},
		CharSpec{0x6D, 13, 6, "m"},
		CharSpec{0x6E, 14, 6, "n"},
		CharSpec{0x6F, 15, 6, "o"},
		CharSpec{0x70, 0, 7, "p"},
		CharSpec{0x71, 1, 7, "q"},
		CharSpec{0x72, 2, 7, "r"},
		CharSpec{0x73, 3, 7, "s"},
		CharSpec{0x74, 4, 7, "t"},
		CharSpec{0x75, 5, 7, "u"},
		CharSpec{0x76, 6, 7, "v"},
		CharSpec{0x77, 7, 7, "w"},
		CharSpec{0x78, 8, 7, "x"},
		CharSpec{0x79, 9, 7, "y"},
		CharSpec{0x7A, 10, 7, "z"},
		CharSpec{0x7B, 11, 7, "{"},
		CharSpec{0x7C, 12, 7, "|"},
		CharSpec{0x7D, 13, 7, "}"},
		CharSpec{0x7E, 14, 7, "~"},

		// Unicode Latin 1 block
		CharSpec{0xA0, 0, 2, "\u00A0"}, // No-Break Space
		CharSpec{0xA1, 1, 12, "¡"},
		CharSpec{0xA2, 2, 10, "¢"},
		CharSpec{0xA3, 3, 10, "£"},
		CharSpec{0xA4, 15, 1, "¤"},
		CharSpec{0xA5, 4, 11, "¥"},
		CharSpec{0xA6, 15, 7, "¦"},
		CharSpec{0xA7, 4, 10, "§"},
		CharSpec{0xA8, 12, 10, "¨"},
		CharSpec{0xA9, 9, 10, "©"},
		CharSpec{0xAA, 11, 11, "ª"},
		CharSpec{0xAB, 7, 12, "«"},
		CharSpec{0xAC, 2, 12, "¬"},
		CharSpec{0xAD, 13, 2, "\u00AD"}, // Soft Hyphen
		CharSpec{0xAE, 8, 10, "®"},
		CharSpec{0xAF, 8, 15, "¯"}, // Macron
		CharSpec{0xB0, 1, 10, "°"}, // Degree Sign
		CharSpec{0xB1, 1, 11, "±"},
		CharSpec{0xB2, 3, 1, "²"},
		CharSpec{0xB3, 4, 1, "³"},
		CharSpec{0xB4, 11, 10, "´"},
		CharSpec{0xB5, 5, 11, "µ"},
		CharSpec{0xB6, 6, 10, "¶"},
		CharSpec{0xB7, 1, 14, "·"},
		CharSpec{0xB8, 12, 15, "¸"}, // Cedillia
		CharSpec{0xB9, 2, 1, "¹"},
		CharSpec{0xBA, 12, 11, "º"},
		CharSpec{0xBB, 8, 12, "»"},
		CharSpec{0xBC, 5, 1, "¼"},
		CharSpec{0xBD, 6, 1, "½"},
		CharSpec{0xBE, 7, 1, "¾"},
		CharSpec{0xBF, 0, 12, "¿"},
		CharSpec{0xC0, 11, 12, "À"},
		CharSpec{0xC1, 7, 14, "Á"},
		CharSpec{0xC2, 5, 14, "Â"},
		CharSpec{0xC3, 12, 12, "Ã"},
		CharSpec{0xC4, 0, 8, "Ä"},
		CharSpec{0xC5, 1, 8, "Å"},
		CharSpec{0xC6, 14, 10, "Æ"},
		CharSpec{0xC7, 2, 8, "Ç"},
		CharSpec{0xC8, 9, 14, "È"},
		CharSpec{0xC9, 3, 8, "É"},
		CharSpec{0xCA, 6, 14, "Ê"},
		CharSpec{0xCB, 8, 14, "Ë"},
		CharSpec{0xCC, 13, 14, "Ì"},
		CharSpec{0xCD, 10, 14, "Í"},
		CharSpec{0xCE, 11, 14, "Î"},
		CharSpec{0xCF, 12, 14, "Ï"},
		CharSpec{0xD0, 8, 1, "Ð"},
		CharSpec{0xD1, 4, 8, "Ñ"},
		CharSpec{0xD2, 1, 15, "Ò"},
		CharSpec{0xD3, 14, 14, "Ó"},
		CharSpec{0xD4, 15, 14, "Ô"},
		CharSpec{0xD5, 13, 12, "Õ"},
		CharSpec{0xD6, 5, 8, "Ö"},
		CharSpec{0xD7, 9, 1, "×"}, // Multiplication Sign
		CharSpec{0xD8, 15, 10, "Ø"},
		CharSpec{0xD9, 4, 15, "Ù"},
		CharSpec{0xDA, 2, 15, "Ú"},
		CharSpec{0xDB, 3, 15, "Û"},
		CharSpec{0xDC, 6, 8, "Ü"},
		CharSpec{0xDD, 10, 1, "Ý"},
		CharSpec{0xDE, 11, 1, "Þ"},
		CharSpec{0xDF, 7, 10, "ß"},
		CharSpec{0xE0, 8, 8, "à"},
		CharSpec{0xE1, 7, 8, "á"},
		CharSpec{0xE2, 9, 8, "â"},
		CharSpec{0xE3, 11, 8, "ã"},
		CharSpec{0xE4, 10, 8, "ä"},
		CharSpec{0xE5, 12, 8, "å"},
		CharSpec{0xE6, 14, 11, "æ"},
		CharSpec{0xE7, 13, 8, "ç"},
		CharSpec{0xE8, 15, 8, "è"},
		CharSpec{0xE9, 14, 8, "é"},
		CharSpec{0xEA, 0, 9, "ê"},
		CharSpec{0xEB, 1, 9, "ë"},
		CharSpec{0xEC, 3, 9, "ì"},
		CharSpec{0xED, 2, 9, "í"},
		CharSpec{0xEE, 4, 9, "î"},
		CharSpec{0xEF, 5, 9, "ï"},
		CharSpec{0xF0, 12, 1, "ð"},
		CharSpec{0xF1, 6, 9, "ñ"},
		CharSpec{0xF2, 8, 9, "ò"},
		CharSpec{0xF3, 7, 9, "ó"},
		CharSpec{0xF4, 9, 9, "ô"},
		CharSpec{0xF5, 11, 9, "õ"},
		CharSpec{0xF6, 10, 9, "ö"},
		CharSpec{0xF7, 6, 13, "÷"},
		CharSpec{0xF8, 15, 11, "ø"},
		CharSpec{0xF9, 13, 9, "ù"},
		CharSpec{0xFA, 12, 9, "ú"},
		CharSpec{0xFB, 14, 9, "û"},
		CharSpec{0xFC, 15, 9, "ü"},
		CharSpec{0xFD, 13, 1, "ý"},
		CharSpec{0xFE, 14, 1, "þ"},
		CharSpec{0xFF, 8, 13, "ÿ"},

		// Unicode Latin Extended A block
		CharSpec{0x152, 14, 12, "Œ"},
		CharSpec{0x153, 15, 12, "œ"},

		// Unicode General Punctuation block
		CharSpec{0x2018, 4, 13, "‘"}, // Left Single Quotation Mark
		CharSpec{0x2019, 5, 13, "’"}, // Right Single Quotation Mark
		CharSpec{0x201A, 2, 14, "‚"}, // Single Low-9 Quotation Mark
		CharSpec{0x201B, 7, 11, "‛"}, // Single High-Reversed-9 Quotation Mark
		CharSpec{0x201C, 2, 13, "“"}, // Left Double Quotation Mark
		CharSpec{0x201D, 3, 13, "”"}, // Right Double Quotation Mark
		CharSpec{0x201E, 3, 14, "„"}, // Double Low-9 Quotation Mark
		CharSpec{0x201F, 8, 11, "‟"}, // Double High-Reversed-9 Quotation Mark
		CharSpec{0x2020, 0, 10, "†"}, // Dagger
		CharSpec{0x2021, 0, 14, "‡"}, // Double Dagger
		CharSpec{0x2022, 5, 10, "•"}, // Bullet

		// Unicode Currency Symbols block
		CharSpec{0x20AC, 11, 13, "€"},

		// Unicode Private Use Area assignments for UI sprites
		CharSpec{0xE700, 0, 0, "Battery_05"},
		CharSpec{0xE701, 1, 0, "Battery_25"},
		CharSpec{0xE702, 2, 0, "Battery_50"},
		CharSpec{0xE703, 3, 0, "Battery_75"},
		CharSpec{0xE704, 4, 0, "Battery_99"},
		CharSpec{0xE705, 5, 0, "Radio_3"},
		CharSpec{0xE706, 6, 0, "Radio_2"},
		CharSpec{0xE707, 7, 0, "Radio_1"},
		CharSpec{0xE708, 8, 0, "Radio_0"},
		CharSpec{0xE709, 9, 0, "Radio_Off"},
		CharSpec{0xE70A, 13, 0, "Shift_Arrow"},
		CharSpec{0xE70B, 14, 0, "Backspace_Symbol"},
		CharSpec{0xE70C, 15, 0, "Enter_Symbol"},

		// Unicode Specials Block
		CharSpec{0xFFFD, 0, 15, "�"},
	}
}

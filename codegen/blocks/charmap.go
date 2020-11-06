package blocks

func DoNothing() {
	return
}

type UnicodeBlock int

// Unicode blocks
const (
	BASIC_LATIN         UnicodeBlock = iota // Block:     00..7E; Subset:     20..7E
	LATIN_1                                 // Block:     80..FF; Subset:     A0..FF
	LATIN_EXTENDED_A                        // Block:   100..17F; Subset:   152..153
	GENERAL_PUNCTUATION                     // Block: 2000..206F; Subset: 2018..2022
	CURRENCY_SYMBOLS                        // Block: 20A0..20CF; Subset: 20AC..20AC
	PRIVATE_USE_AREA                        // Block: E000..F8FF; Subset: E700..E70C
	SPECIALS                                // Block: FFF0..FFFF; Subset: FFFD..FFFD
	UNKNOWN
)

type CharSpec struct {
	row int
	col int
	hex string
	chr string
}

// Convert Unicode codepoint to its block and offset from the start of that block
func BlockForRune(c rune) (UnicodeBlock, int) {
	if 0x20 <= c && c <= 0x7E {
		return BASIC_LATIN, int(c - 0x20)
	} else if 0xA0 <= c && c <= 0xFF {
		return LATIN_1, int(c - 0xA0)
	} else if 0x152 <= c && c <= 0x153 {
		return LATIN_EXTENDED_A, int(c - 0x152)
	} else if 0x2018 <= c && c <= 0x2022 {
		return GENERAL_PUNCTUATION, int(c - 0x2018)
	} else if 0x20AC <= c && c <= 0x20AC {
		return CURRENCY_SYMBOLS, int(c - 0x20AC)
	} else if 0xE700 <= c && c <= 0xE70C {
		return PRIVATE_USE_AREA, int(c - 0xE700)
	} else if 0xFFFD <= c && c <= 0xFFFD {
		return SPECIALS, int(c - 0xFFFD)
	} else {
		return UNKNOWN, 0
	}
}

func SysLatinMap() map[int]CharSpec {
	return map[int]CharSpec{
		// Unicode Basic Latin block
		32:  CharSpec{0, 2, "20", " "},
		33:  CharSpec{1, 2, "21", "!"},
		34:  CharSpec{2, 2, "22", "\""},
		35:  CharSpec{3, 2, "23", "#"},
		36:  CharSpec{4, 2, "24", "$"},
		37:  CharSpec{5, 2, "25", "%"},
		38:  CharSpec{6, 2, "26", "&"},
		39:  CharSpec{7, 2, "27", "'"},
		40:  CharSpec{8, 2, "28", "("},
		41:  CharSpec{9, 2, "29", ")"},
		42:  CharSpec{10, 2, "2A", "*"},
		43:  CharSpec{11, 2, "2B", "+"},
		44:  CharSpec{12, 2, "2C", ","},
		45:  CharSpec{13, 2, "2D", "-"},
		46:  CharSpec{14, 2, "2E", "."},
		47:  CharSpec{15, 2, "2F", "/"},
		48:  CharSpec{0, 3, "30", "0"},
		49:  CharSpec{1, 3, "31", "1"},
		50:  CharSpec{2, 3, "32", "2"},
		51:  CharSpec{3, 3, "33", "3"},
		52:  CharSpec{4, 3, "34", "4"},
		53:  CharSpec{5, 3, "35", "5"},
		54:  CharSpec{6, 3, "36", "6"},
		55:  CharSpec{7, 3, "37", "7"},
		56:  CharSpec{8, 3, "38", "8"},
		57:  CharSpec{9, 3, "39", "9"},
		58:  CharSpec{10, 3, "3A", ":"},
		59:  CharSpec{11, 3, "3B", ";"},
		60:  CharSpec{12, 3, "3C", "<"},
		61:  CharSpec{13, 3, "3D", "="},
		62:  CharSpec{14, 3, "3E", ">"},
		63:  CharSpec{15, 3, "3F", "?"},
		64:  CharSpec{0, 4, "40", "@"},
		65:  CharSpec{1, 4, "41", "A"},
		66:  CharSpec{2, 4, "42", "B"},
		67:  CharSpec{3, 4, "43", "C"},
		68:  CharSpec{4, 4, "44", "D"},
		69:  CharSpec{5, 4, "45", "E"},
		70:  CharSpec{6, 4, "46", "F"},
		71:  CharSpec{7, 4, "47", "G"},
		72:  CharSpec{8, 4, "48", "H"},
		73:  CharSpec{9, 4, "49", "I"},
		74:  CharSpec{10, 4, "4A", "J"},
		75:  CharSpec{11, 4, "4B", "K"},
		76:  CharSpec{12, 4, "4C", "L"},
		77:  CharSpec{13, 4, "4D", "M"},
		78:  CharSpec{14, 4, "4E", "N"},
		79:  CharSpec{15, 4, "4F", "O"},
		80:  CharSpec{0, 5, "50", "P"},
		81:  CharSpec{1, 5, "51", "Q"},
		82:  CharSpec{2, 5, "52", "R"},
		83:  CharSpec{3, 5, "53", "S"},
		84:  CharSpec{4, 5, "54", "T"},
		85:  CharSpec{5, 5, "55", "U"},
		86:  CharSpec{6, 5, "56", "V"},
		87:  CharSpec{7, 5, "57", "W"},
		88:  CharSpec{8, 5, "58", "X"},
		89:  CharSpec{9, 5, "59", "Y"},
		90:  CharSpec{10, 5, "5A", "Z"},
		91:  CharSpec{11, 5, "5B", "["},
		92:  CharSpec{12, 5, "5C", "\\"},
		93:  CharSpec{13, 5, "5D", "]"},
		94:  CharSpec{14, 5, "5E", "^"},
		95:  CharSpec{15, 5, "5F", "_"},
		96:  CharSpec{0, 6, "60", "`"},
		97:  CharSpec{1, 6, "61", "a"},
		98:  CharSpec{2, 6, "62", "b"},
		99:  CharSpec{3, 6, "63", "c"},
		100: CharSpec{4, 6, "64", "d"},
		101: CharSpec{5, 6, "65", "e"},
		102: CharSpec{6, 6, "66", "f"},
		103: CharSpec{7, 6, "67", "g"},
		104: CharSpec{8, 6, "68", "h"},
		105: CharSpec{9, 6, "69", "i"},
		106: CharSpec{10, 6, "6A", "j"},
		107: CharSpec{11, 6, "6B", "k"},
		108: CharSpec{12, 6, "6C", "l"},
		109: CharSpec{13, 6, "6D", "m"},
		110: CharSpec{14, 6, "6E", "n"},
		111: CharSpec{15, 6, "6F", "o"},
		112: CharSpec{0, 7, "70", "p"},
		113: CharSpec{1, 7, "71", "q"},
		114: CharSpec{2, 7, "72", "r"},
		115: CharSpec{3, 7, "73", "s"},
		116: CharSpec{4, 7, "74", "t"},
		117: CharSpec{5, 7, "75", "u"},
		118: CharSpec{6, 7, "76", "v"},
		119: CharSpec{7, 7, "77", "w"},
		120: CharSpec{8, 7, "78", "x"},
		121: CharSpec{9, 7, "79", "y"},
		122: CharSpec{10, 7, "7A", "z"},
		123: CharSpec{11, 7, "7B", "{"},
		124: CharSpec{12, 7, "7C", "|"},
		125: CharSpec{13, 7, "7D", "}"},
		126: CharSpec{14, 7, "7E", "~"},

		// Unicode Latin 1 block
		160: CharSpec{0, 2, "A0", "\xA0"}, // No-Break Space
		161: CharSpec{1, 12, "A1", "¡"},
		162: CharSpec{2, 10, "A2", "¢"},
		163: CharSpec{3, 10, "A3", "£"},
		164: CharSpec{15, 1, "A4", "¤"},
		165: CharSpec{4, 11, "A5", "¥"},
		166: CharSpec{15, 7, "A6", "¦"},
		167: CharSpec{4, 10, "A7", "§"},
		168: CharSpec{12, 10, "A8", "¨"},
		169: CharSpec{9, 10, "A9", "©"},
		170: CharSpec{11, 11, "AA", "ª"},
		171: CharSpec{7, 12, "AB", "«"},
		172: CharSpec{2, 12, "AC", "¬"},
		173: CharSpec{13, 2, "AD", "\xAD"}, // Soft Hyphen
		174: CharSpec{8, 10, "AE", "®"},
		175: CharSpec{8, 15, "AF", "¯"}, // Macron
		176: CharSpec{1, 10, "B0", "°"}, // Degree Sign
		177: CharSpec{1, 11, "B1", "±"},
		178: CharSpec{3, 1, "B2", "²"},
		179: CharSpec{4, 1, "B3", "³"},
		180: CharSpec{11, 10, "B4", "´"},
		181: CharSpec{5, 11, "B5", "µ"},
		182: CharSpec{6, 10, "B6", "¶"},
		183: CharSpec{1, 14, "B7", "·"},
		184: CharSpec{12, 15, "B8", "¸"}, // Cedillia
		185: CharSpec{2, 1, "B9", "¹"},
		186: CharSpec{12, 11, "BA", "º"},
		187: CharSpec{8, 12, "BB", "»"},
		188: CharSpec{5, 1, "BC", "¼"},
		189: CharSpec{6, 1, "BD", "½"},
		190: CharSpec{7, 1, "BE", "¾"},
		191: CharSpec{0, 12, "BF", "¿"},
		192: CharSpec{11, 12, "C0", "À"},
		193: CharSpec{7, 14, "C1", "Á"},
		194: CharSpec{5, 14, "C2", "Â"},
		195: CharSpec{12, 12, "C3", "Ã"},
		196: CharSpec{0, 8, "C4", "Ä"},
		197: CharSpec{1, 8, "C5", "Å"},
		198: CharSpec{14, 10, "C6", "Æ"},
		199: CharSpec{2, 8, "C7", "Ç"},
		200: CharSpec{9, 14, "C8", "È"},
		201: CharSpec{3, 8, "C9", "É"},
		202: CharSpec{6, 14, "CA", "Ê"},
		203: CharSpec{8, 14, "CB", "Ë"},
		204: CharSpec{13, 14, "CC", "Ì"},
		205: CharSpec{10, 14, "CD", "Í"},
		206: CharSpec{11, 14, "CE", "Î"},
		207: CharSpec{12, 14, "CF", "Ï"},
		208: CharSpec{8, 1, "D0", "Ð"},
		209: CharSpec{4, 8, "D1", "Ñ"},
		210: CharSpec{1, 15, "D2", "Ò"},
		211: CharSpec{14, 14, "D3", "Ó"},
		212: CharSpec{15, 14, "D4", "Ô"},
		213: CharSpec{13, 12, "D5", "Õ"},
		214: CharSpec{5, 8, "D6", "Ö"},
		215: CharSpec{9, 1, "D7", "×"}, // Multiplication Sign
		216: CharSpec{15, 10, "D8", "Ø"},
		217: CharSpec{4, 15, "D9", "Ù"},
		218: CharSpec{2, 15, "DA", "Ú"},
		219: CharSpec{3, 15, "DB", "Û"},
		220: CharSpec{6, 8, "DC", "Ü"},
		221: CharSpec{10, 1, "DD", "Ý"},
		222: CharSpec{11, 1, "DE", "Þ"},
		223: CharSpec{7, 10, "DF", "ß"},
		224: CharSpec{8, 8, "E0", "à"},
		225: CharSpec{7, 8, "E1", "á"},
		226: CharSpec{9, 8, "E2", "â"},
		227: CharSpec{11, 8, "E3", "ã"},
		228: CharSpec{10, 8, "E4", "ä"},
		229: CharSpec{12, 8, "E5", "å"},
		230: CharSpec{14, 11, "E6", "æ"},
		231: CharSpec{13, 8, "E7", "ç"},
		232: CharSpec{15, 8, "E8", "è"},
		233: CharSpec{14, 8, "E9", "é"},
		234: CharSpec{0, 9, "EA", "ê"},
		235: CharSpec{1, 9, "EB", "ë"},
		236: CharSpec{3, 9, "EC", "ì"},
		237: CharSpec{2, 9, "ED", "í"},
		238: CharSpec{4, 9, "EE", "î"},
		239: CharSpec{5, 9, "EF", "ï"},
		240: CharSpec{12, 1, "F0", "ð"},
		241: CharSpec{6, 9, "F1", "ñ"},
		242: CharSpec{8, 9, "F2", "ò"},
		243: CharSpec{7, 9, "F3", "ó"},
		244: CharSpec{9, 9, "F4", "ô"},
		245: CharSpec{11, 9, "F5", "õ"},
		246: CharSpec{10, 9, "F6", "ö"},
		247: CharSpec{6, 13, "F7", "÷"},
		248: CharSpec{15, 11, "F8", "ø"},
		249: CharSpec{13, 9, "F9", "ù"},
		250: CharSpec{12, 9, "FA", "ú"},
		251: CharSpec{14, 9, "FB", "û"},
		252: CharSpec{15, 9, "FC", "ü"},
		253: CharSpec{13, 1, "FD", "ý"},
		254: CharSpec{14, 1, "FE", "þ"},
		255: CharSpec{8, 13, "FF", "ÿ"},

		// Unicode Latin Extended A block
		338: CharSpec{14, 12, "152", "Œ"},
		339: CharSpec{15, 12, "153", "œ"},

		// Unicode General Punctuation block
		8216: CharSpec{4, 13, "2018", "‘"}, // Left Single Quotation Mark
		8217: CharSpec{5, 13, "2019", "’"}, // Right Single Quotation Mark
		8218: CharSpec{2, 14, "201A", "‚"}, // Single Low-9 Quotation Mark
		8219: CharSpec{7, 11, "201B", "‛"}, // Single High-Reversed-9 Quotation Mark
		8220: CharSpec{2, 13, "201C", "“"}, // Left Double Quotation Mark
		8221: CharSpec{3, 13, "201D", "”"}, // Right Double Quotation Mark
		8222: CharSpec{3, 14, "201E", "„"}, // Double Low-9 Quotation Mark
		8223: CharSpec{8, 11, "201F", "‟"}, // Double High-Reversed-9 Quotation Mark
		8224: CharSpec{0, 10, "2020", "†"}, // Dagger
		8225: CharSpec{0, 14, "2021", "‡"}, // Double Dagger
		8226: CharSpec{5, 10, "2022", "•"}, // Bullet

		// Unicode Currency Symbols block
		8364: CharSpec{11, 13, "20AC", "€"},

		// Unicode Private Use Area assignments for UI sprites
		59136: CharSpec{0, 0, "E700", "Battery_05"},
		59137: CharSpec{1, 0, "E701", "Battery_25"},
		59138: CharSpec{2, 0, "E702", "Battery_50"},
		59139: CharSpec{3, 0, "E703", "Battery_75"},
		59140: CharSpec{4, 0, "E704", "Battery_99"},
		59141: CharSpec{5, 0, "E705", "Radio_3"},
		59142: CharSpec{6, 0, "E706", "Radio_2"},
		59143: CharSpec{7, 0, "E707", "Radio_1"},
		59144: CharSpec{8, 0, "E708", "Radio_0"},
		59145: CharSpec{9, 0, "E709", "Radio_Off"},
		59146: CharSpec{13, 0, "E70A", "Shift_Arrow"},
		59147: CharSpec{14, 0, "E70B", "Backspace_Symbol"},
		59148: CharSpec{15, 0, "E70C", "Enter_Symbol"},

		// Unicode Specials Block
		65533: CharSpec{0, 15, "FFFD", "�"},
	}
}

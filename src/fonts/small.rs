#![allow(dead_code)]
//! Small Font
//
// This code incorporates glyphs from the Geneva bitmap typeface which was
// designed by Susan Kare and released by Apple in 1984. Geneva is a registered
// trademark of Apple Inc.

/// Return offset into DATA[] for start of pattern depicting glyph for character c
pub fn get_glyph_pattern_offset(c: char) -> usize {
    match c as u32 {
        0x20..=0x7E => BASIC_LATIN[(c as usize) - 0x20] as usize,
        0xA0..=0xFF => LATIN_1[(c as usize) - 0xA0] as usize,
        0x152..=0x153 => LATIN_EXTENDED_A[(c as usize) - 0x152] as usize,
        0x2018..=0x2022 => GENERAL_PUNCTUATION[(c as usize) - 0x2018] as usize,
        0x20AC..=0x20AC => CURRENCY_SYMBOLS[(c as usize) - 0x20AC] as usize,
        _ => SPECIALS[(0xFFFD as usize) - 0xFFFD] as usize,
    }
}

// Index to Unicode Basic Latin block glyph patterns
const BASIC_LATIN: [u16; 95] = [
    0, // ' '
    2, // '!'
    4, // '"'
    6, // '#'
    11, // '$'
    18, // '%'
    26, // '&'
    34, // '\''
    36, // '('
    41, // ')'
    46, // '*'
    51, // '+'
    56, // ','
    58, // '-'
    60, // '.'
    62, // '/'
    67, // '0'
    73, // '1'
    76, // '2'
    82, // '3'
    88, // '4'
    94, // '5'
    100, // '6'
    106, // '7'
    112, // '8'
    118, // '9'
    124, // ':'
    126, // ';'
    129, // '<'
    132, // '='
    135, // '>'
    138, // '?'
    143, // '@'
    151, // 'A'
    157, // 'B'
    163, // 'C'
    169, // 'D'
    175, // 'E'
    180, // 'F'
    185, // 'G'
    191, // 'H'
    197, // 'I'
    199, // 'J'
    205, // 'K'
    211, // 'L'
    216, // 'M'
    224, // 'N'
    230, // 'O'
    236, // 'P'
    242, // 'Q'
    248, // 'R'
    254, // 'S'
    260, // 'T'
    266, // 'U'
    272, // 'V'
    278, // 'W'
    286, // 'X'
    292, // 'Y'
    298, // 'Z'
    303, // '['
    307, // '\\'
    312, // ']'
    316, // '^'
    318, // '_'
    320, // '`'
    322, // 'a'
    326, // 'b'
    331, // 'c'
    335, // 'd'
    340, // 'e'
    344, // 'f'
    349, // 'g'
    354, // 'h'
    359, // 'i'
    361, // 'j'
    366, // 'k'
    371, // 'l'
    373, // 'm'
    379, // 'n'
    383, // 'o'
    387, // 'p'
    392, // 'q'
    397, // 'r'
    401, // 's'
    405, // 't'
    409, // 'u'
    413, // 'v'
    418, // 'w'
    424, // 'x'
    429, // 'y'
    435, // 'z'
    439, // '{'
    444, // '|'
    447, // '}'
    452, // '~'
];

// Index to Unicode Latin 1 block glyph patterns
const LATIN_1: [u16; 96] = [
    455, // ' '
    457, // '¡'
    459, // '¢'
    465, // '£'
    471, // '¤'
    477, // '¥'
    483, // '¦'
    486, // '§'
    493, // '¨'
    495, // '©'
    504, // 'ª'
    508, // '«'
    513, // '¬'
    516, // '­'
    518, // '®'
    527, // '¯'
    529, // '°'
    532, // '±'
    537, // '²'
    540, // '³'
    543, // '´'
    545, // 'µ'
    552, // '¶'
    560, // '·'
    562, // '¸'
    564, // '¹'
    567, // 'º'
    571, // '»'
    576, // '¼'
    589, // '½'
    602, // '¾'
    615, // '¿'
    620, // 'À'
    628, // 'Á'
    636, // 'Â'
    644, // 'Ã'
    652, // 'Ä'
    659, // 'Å'
    667, // 'Æ'
    675, // 'Ç'
    682, // 'È'
    688, // 'É'
    694, // 'Ê'
    700, // 'Ë'
    706, // 'Ì'
    710, // 'Í'
    714, // 'Î'
    719, // 'Ï'
    724, // 'Ð'
    731, // 'Ñ'
    739, // 'Ò'
    747, // 'Ó'
    755, // 'Ô'
    763, // 'Õ'
    771, // 'Ö'
    778, // '×'
    783, // 'Ø'
    791, // 'Ù'
    799, // 'Ú'
    807, // 'Û'
    815, // 'Ü'
    822, // 'Ý'
    830, // 'Þ'
    835, // 'ß'
    841, // 'à'
    846, // 'á'
    851, // 'â'
    856, // 'ã'
    861, // 'ä'
    866, // 'å'
    872, // 'æ'
    878, // 'ç'
    883, // 'è'
    888, // 'é'
    893, // 'ê'
    898, // 'ë'
    903, // 'ì'
    906, // 'í'
    909, // 'î'
    913, // 'ï'
    917, // 'ð'
    922, // 'ñ'
    927, // 'ò'
    932, // 'ó'
    937, // 'ô'
    942, // 'õ'
    947, // 'ö'
    952, // '÷'
    957, // 'ø'
    963, // 'ù'
    968, // 'ú'
    973, // 'û'
    978, // 'ü'
    983, // 'ý'
    991, // 'þ'
    997, // 'ÿ'
];

// Index to Unicode Latin Extended A block glyph patterns
const LATIN_EXTENDED_A: [u16; 2] = [
    1004, // 'Œ'
    1012, // 'œ'
];

// Index to General Punctuation block glyph patterns
const GENERAL_PUNCTUATION: [u16; 11] = [
    1018, // '‘'
    1020, // '’'
    1022, // '‚'
    1024, // '‛'
    1026, // '“'
    1029, // '”'
    1032, // '„'
    1035, // '‟'
    1038, // '†'
    1041, // '‡'
    1046, // '•'
];

// Index to Unicode Currency Symbols block glyph patterns
const CURRENCY_SYMBOLS: [u16; 1] = [
    1051, // '€'
];

// Index to Unicode Specials block glyph patterns
const SPECIALS: [u16; 1] = [
    1058, // '�'
];

/// Maximum height of glyph patterns in this bitmap typeface.
/// This will be true: h + yOffset <= MAX_HEIGHT
pub const MAX_HEIGHT: u8 = 24;

/// Packed glyph pattern data.
/// Record format:
///  [offset+0]: ((w as u8) << 16) | ((h as u8) << 8) | (yOffset as u8)
///  [offset+1..=ceil(w*h/32)]: packed 1-bit pixels; 0=clear, 1=set
/// Pixels are packed in top to bottom, left to right order with MSB of first
/// pixel word containing the top left pixel.
///  w: Width of pattern in pixels
///  h: Height of pattern in pixels
///  yOffset: Vertical offset (pixels downward from top of line) to position
///     glyph pattern properly relative to text baseline
pub const DATA: [u32; 1071] = [
    // [0]: 20 ' '
    0x0004020b, 0x00000000,
    // [2]: 21 '!'
    0x00020e06, 0xfffff0f0,
    // [4]: 22 '"'
    0x00060406, 0xcf3cf300,
    // [6]: 23 '#'
    0x000a0a04, 0xcc330fff, 0xff330ccf, 0xffff0cc3, 0x30000000,
    // [11]: 24 '$'
    0x000a1204, 0x0c0303f0, 0xfcccf330, 0xcc333f0f, 0xccc330cc, 0xf333f0fc, 0x0c030000,
    // [18]: 25 '%'
    0x00100e06, 0xfffcfffc, 0x30c330c3, 0x0cc30cc3, 0x3f3c3f3c, 0xc3c0c3c0, 0xc330c330, 0x3c0c3c0c,
    // [26]: 26 '&'
    0x000e1004, 0x03c00f00, 0xc3030c03, 0x300cc00c, 0x00303330, 0xccc0c0c3, 0x03330ccc, 0x3c0f303c,
    // [34]: 27 '\''
    0x00020406, 0xff000000,
    // [36]: 28 '('
    0x00061204, 0xc3030c0c, 0x30c30c30, 0xc30c330c, 0xc3000000,
    // [41]: 29 ')'
    0x00061204, 0x0c330cc3, 0x0c30c30c, 0x30c3030c, 0x0c300000,
    // [46]: 2A '*'
    0x000a0a06, 0x0c030ccf, 0x333f0fc3, 0x30ccc0f0, 0x30000000,
    // [51]: 2B '+'
    0x000a0a08, 0x0c0300c0, 0x30fffff0, 0xc0300c03, 0x00000000,
    // [56]: 2C ','
    0x00040612, 0xcccc3300,
    // [58]: 2D '-'
    0x0008020c, 0xffff0000,
    // [60]: 2E '.'
    0x00020212, 0xf0000000,
    // [62]: 2F '/'
    0x00081004, 0xc0c0c0c0, 0x30303030, 0x0c0c0c0c, 0x03030303,
    // [67]: 30 '0'
    0x000a0e06, 0x3f0fcc0f, 0x03c0f03c, 0x0f03c0f0, 0x3c0f033f, 0x0fc00000,
    // [73]: 31 '1'
    0x00040e06, 0xccffcccc, 0xcccccc00,
    // [76]: 32 '2'
    0x000a0e06, 0x3f0fcc0f, 0x03c03003, 0x00c00c03, 0x00300cff, 0xfff00000,
    // [82]: 33 '3'
    0x000a0e06, 0xfffff300, 0xc00c0303, 0xf0fcc030, 0x0c0f033f, 0x0fc00000,
    // [88]: 34 '4'
    0x000a0e06, 0x300c03c0, 0xf0330cc3, 0x0cc3ffff, 0xf300c030, 0x0c000000,
    // [94]: 35 '5'
    0x000a0e06, 0xfffff00c, 0x033fcffc, 0x0300c030, 0x0c0f033f, 0x0fc00000,
    // [100]: 36 '6'
    0x000a0e06, 0x3c0f0030, 0x0c00c033, 0xfcffc0f0, 0x3c0f033f, 0x0fc00000,
    // [106]: 37 '7'
    0x000a0e06, 0xfffffc03, 0x00300c03, 0x00c00c03, 0x00c0300c, 0x03000000,
    // [112]: 38 '8'
    0x000a0e06, 0x3f0fcc0f, 0x03c0f033, 0xf0fcc0f0, 0x3c0f033f, 0x0fc00000,
    // [118]: 39 '9'
    0x000a0e06, 0x3f0fcc0f, 0x03c0f03f, 0xf3fcc030, 0x0300c00f, 0x03c00000,
    // [124]: 3A ':'
    0x00020a0a, 0xf000f000,
    // [126]: 3B ';'
    0x00040e0a, 0xcc000000, 0xcccc3300,
    // [129]: 3C '<'
    0x00060a08, 0xc3030c0c, 0x330cc300,
    // [132]: 3D '='
    0x000a060a, 0xfffff000, 0x00fffff0,
    // [135]: 3E '>'
    0x00060a08, 0x0c330cc3, 0x030c0c30,
    // [138]: 3F '?'
    0x00080e06, 0x3c3cc3c3, 0xc0c03030, 0x0c0c0000, 0x0c0c0000,
    // [143]: 40 '@'
    0x000e1006, 0x0fc03f03, 0x030c0ccf, 0x0f3c3ccc, 0xf333cccf, 0x3333f0cf, 0xc3003000, 0xc0fc03f0,
    // [151]: 41 'A'
    0x000a0e06, 0x0c0300c0, 0x30330cc3, 0x30ccffff, 0xfc0f03c0, 0xf0300000,
    // [157]: 42 'B'
    0x000a0e06, 0x3fcffc0f, 0x03c0f033, 0xfcffc0f0, 0x3c0f033f, 0xcff00000,
    // [163]: 43 'C'
    0x000a0e06, 0x3f0fcc0f, 0x0300c030, 0x0c0300c0, 0x3c0f033f, 0x0fc00000,
    // [169]: 44 'D'
    0x000a0e06, 0x0fc3f30c, 0xc3c0f03c, 0x0f03c0f0, 0x330cc30f, 0xc3f00000,
    // [175]: 45 'E'
    0x00080e06, 0xffff0303, 0x03033f3f, 0x03030303, 0xffff0000,
    // [180]: 46 'F'
    0x00080e06, 0xffff0303, 0x03033f3f, 0x03030303, 0x03030000,
    // [185]: 47 'G'
    0x000a0e06, 0x3f0fcc0f, 0x0300c03f, 0x0fc3c0f0, 0x3c0f033f, 0x0fc00000,
    // [191]: 48 'H'
    0x000a0e06, 0xc0f03c0f, 0x03c0f03f, 0xffffc0f0, 0x3c0f03c0, 0xf0300000,
    // [197]: 49 'I'
    0x00020e06, 0xfffffff0,
    // [199]: 4A 'J'
    0x000a0e06, 0xc0300c03, 0x00c0300c, 0x0300c0f0, 0x3c0f033f, 0x0fc00000,
    // [205]: 4B 'K'
    0x000a0e06, 0xc0f0330c, 0xc30cc330, 0x3c0f0cc3, 0x330cc3c0, 0xf0300000,
    // [211]: 4C 'L'
    0x00080e06, 0x03030303, 0x03030303, 0x03030303, 0xffff0000,
    // [216]: 4D 'M'
    0x000e0e06, 0xc00f003f, 0x03fc0fcc, 0xcf333c30, 0xf0c3c00f, 0x003c00f0, 0x03c00f00, 0x30000000,
    // [224]: 4E 'N'
    0x000a0e06, 0xc3f0fc3f, 0x0fccf33c, 0xcf33f0fc, 0x3f0fc3c0, 0xf0300000,
    // [230]: 4F 'O'
    0x000a0e06, 0x3f0fcc0f, 0x03c0f03c, 0x0f03c0f0, 0x3c0f033f, 0x0fc00000,
    // [236]: 50 'P'
    0x000a0e06, 0x3fcffc0f, 0x03c0f033, 0xfcff00c0, 0x300c0300, 0xc0300000,
    // [242]: 51 'Q'
    0x000a1006, 0x3f0fcc0f, 0x03c0f03c, 0x0f03c0f0, 0x3ccf333f, 0x0fc300c0,
    // [248]: 52 'R'
    0x000a0e06, 0x3fcffc0f, 0x03c0f033, 0xfcff0cc3, 0x330cc3c0, 0xf0300000,
    // [254]: 53 'S'
    0x000a0e06, 0x3f0fcc0f, 0x0300c033, 0xf0fcc030, 0x0c0f033f, 0x0fc00000,
    // [260]: 54 'T'
    0x000a0e06, 0xfffff0c0, 0x300c0300, 0xc0300c03, 0x00c0300c, 0x03000000,
    // [266]: 55 'U'
    0x000a0e06, 0xc0f03c0f, 0x03c0f03c, 0x0f03c0f0, 0x3c0f033f, 0x0fc00000,
    // [272]: 56 'V'
    0x000a0e06, 0xc0f03c0f, 0x03c0f033, 0x30cc330c, 0xc0c0300c, 0x03000000,
    // [278]: 57 'W'
    0x000e0e06, 0xc00f003c, 0x00f00333, 0x30ccc333, 0x0ccc0cc0, 0x3300cc03, 0x300cc033, 0x00000000,
    // [286]: 58 'X'
    0x000a0e06, 0xc0f03c0f, 0x03330cc0, 0xc030330c, 0xcc0f03c0, 0xf0300000,
    // [292]: 59 'Y'
    0x000a0e06, 0xc0f03c0f, 0x03330cc0, 0xc0300c03, 0x00c0300c, 0x03000000,
    // [298]: 5A 'Z'
    0x00080e06, 0xffffc0c0, 0x30300c0c, 0x03030303, 0xffff0000,
    // [303]: 5B '['
    0x00041204, 0xff333333, 0x33333333, 0xff000000,
    // [307]: 5C '\\'
    0x00081004, 0x03030303, 0x0c0c0c0c, 0x30303030, 0xc0c0c0c0,
    // [312]: 5D ']'
    0x00041204, 0xffcccccc, 0xcccccccc, 0xff000000,
    // [316]: 5E '^'
    0x00060406, 0x30ccf300,
    // [318]: 5F '_'
    0x000c0212, 0xffffff00,
    // [320]: 60 '`'
    0x00040406, 0x33cc0000,
    // [322]: 61 'a'
    0x00080a0a, 0x3c3cc0c0, 0xfcfcc3c3, 0xfcfc0000,
    // [326]: 62 'b'
    0x00080e06, 0x03030303, 0x3f3fc3c3, 0xc3c3c3c3, 0x3f3f0000,
    // [331]: 63 'c'
    0x00080a0a, 0x3c3cc3c3, 0x0303c3c3, 0x3c3c0000,
    // [335]: 64 'd'
    0x00080e06, 0xc0c0c0c0, 0xfcfcc3c3, 0xc3c3c3c3, 0xfcfc0000,
    // [340]: 65 'e'
    0x00080a0a, 0x3c3cc3c3, 0xffff0303, 0x3c3c0000,
    // [344]: 66 'f'
    0x00080e06, 0xf0f00c0c, 0x3f3f0c0c, 0x0c0c0c0c, 0x0c0c0000,
    // [349]: 67 'g'
    0x00080e0a, 0xfcfcc3c3, 0xc3c3c3c3, 0xfcfcc0c0, 0x3c3c0000,
    // [354]: 68 'h'
    0x00080e06, 0x03030303, 0x3f3fc3c3, 0xc3c3c3c3, 0xc3c30000,
    // [359]: 69 'i'
    0x00020e06, 0xf0fffff0,
    // [361]: 6A 'j'
    0x00061206, 0xc30000c3, 0x0c30c30c, 0x30c30c30, 0x3cf00000,
    // [366]: 6B 'k'
    0x00080e06, 0x03030303, 0xc3c33333, 0x0f0f3333, 0xc3c30000,
    // [371]: 6C 'l'
    0x00020e06, 0xfffffff0,
    // [373]: 6D 'm'
    0x000e0a0a, 0x3cfcf3fc, 0x30f0c3c3, 0x0f0c3c30, 0xf0c3c30f, 0x0c300000,
    // [379]: 6E 'n'
    0x00080a0a, 0x3f3fc3c3, 0xc3c3c3c3, 0xc3c30000,
    // [383]: 6F 'o'
    0x00080a0a, 0x3c3cc3c3, 0xc3c3c3c3, 0x3c3c0000,
    // [387]: 70 'p'
    0x00080e0a, 0x3f3fc3c3, 0xc3c3c3c3, 0x3f3f0303, 0x03030000,
    // [392]: 71 'q'
    0x00080e0a, 0xfcfcc3c3, 0xc3c3c3c3, 0xfcfcc0c0, 0xc0c00000,
    // [397]: 72 'r'
    0x00080a0a, 0xf3f30f0f, 0x03030303, 0x03030000,
    // [401]: 73 's'
    0x00080a0a, 0xfcfc0303, 0x3c3cc0c0, 0x3f3f0000,
    // [405]: 74 't'
    0x00060e06, 0x30c30cff, 0xf30c30c3, 0x0cc30000,
    // [409]: 75 'u'
    0x00080a0a, 0xc3c3c3c3, 0xc3c3c3c3, 0xfcfc0000,
    // [413]: 76 'v'
    0x000a0a0a, 0xc0f03330, 0xcc330cc0, 0xc0300c03, 0x00000000,
    // [418]: 77 'w'
    0x000e0a0a, 0xc00f0033, 0x330ccc33, 0x30ccc0cc, 0x03300cc0, 0x33000000,
    // [424]: 78 'x'
    0x000a0a0a, 0xc0f03330, 0xcc0c0303, 0x30ccc0f0, 0x30000000,
    // [429]: 79 'y'
    0x000a0e0a, 0xc0f03c0f, 0x03330cc3, 0x30cc0c03, 0x00c03003, 0xc0f00000,
    // [435]: 7A 'z'
    0x00080a0a, 0xffff3030, 0x0c0c0303, 0xffff0000,
    // [439]: 7B '{'
    0x00061204, 0xc3030c30, 0xc30c0c33, 0x0c30c30c, 0xc3000000,
    // [444]: 7C '|'
    0x00021204, 0xffffffff, 0xf0000000,
    // [447]: 7D '}'
    0x00061204, 0x0c330c30, 0xc30cc303, 0x0c30c30c, 0x0c300000,
    // [452]: 7E '~'
    0x000a0406, 0xcf33c3cc, 0xf3000000,
    // [455]: A0 ' '
    0x0004020b, 0x00000000,
    // [457]: A1 '¡'
    0x00020e06, 0xf0fffff0,
    // [459]: A2 '¢'
    0x000a0e06, 0x0c0303f0, 0xfcccf330, 0xcc33ccf3, 0x33f0fc0c, 0x03000000,
    // [465]: A3 '£'
    0x000a0e06, 0x3c0f0030, 0x0c0300c0, 0xfc3f0300, 0xcc330c3f, 0xcff00000,
    // [471]: A4 '¤'
    0x000a0e06, 0xc0f033f0, 0xfcc0f03c, 0x0f03c0f0, 0x33f0fcc0, 0xf0300000,
    // [477]: A5 '¥'
    0x000a0e06, 0xc0f03330, 0xccfffff0, 0xc030ffff, 0xf0c0300c, 0x03000000,
    // [483]: A6 '¦'
    0x00021204, 0xffff0fff, 0xf0000000,
    // [486]: A7 '§'
    0x000a1204, 0x3f0fcc0f, 0x0300c033, 0xf0fcc0f0, 0x33f0fcc0, 0x300c0f03, 0x3f0fc000,
    // [493]: A8 '¨'
    0x00060206, 0xcf300000,
    // [495]: A9 '©'
    0x00101004, 0x0ff00ff0, 0x300c300c, 0xc3c3c3c3, 0xc033c033, 0xc033c033, 0xc3c3c3c3, 0x300c300c,
    0x0ff00ff0,
    // [504]: AA 'ª'
    0x00060e04, 0x30ccf3f3, 0xccf3f3c0, 0x00fff000,
    // [508]: AB '«'
    0x000c0a0a, 0xc30c3030, 0xc30c0c30, 0xc330c30c, 0xc30c3000,
    // [513]: AC '¬'
    0x0008060c, 0xffffc0c0, 0xc0c00000,
    // [516]: AD '­'
    0x0008020c, 0xffff0000,
    // [518]: AE '®'
    0x00101004, 0x0ff00ff0, 0x300c300c, 0xc3f3c3f3, 0xcc33cc33, 0xc3f3c3f3, 0xcc33cc33, 0x300c300c,
    0x0ff00ff0,
    // [527]: AF '¯'
    0x00060206, 0xfff00000,
    // [529]: B0 '°'
    0x00080804, 0x3c3cc3c3, 0xc3c33c3c,
    // [532]: B1 '±'
    0x000a0c08, 0x0c0300c0, 0x30fffff0, 0xc0300c03, 0x0fffff00,
    // [537]: B2 '²'
    0x00060a02, 0xfffc30ff, 0xf0c3fff0,
    // [540]: B3 '³'
    0x00060a02, 0xfffc30ff, 0xfc30fff0,
    // [543]: B4 '´'
    0x00040406, 0xcc330000,
    // [545]: B5 'µ'
    0x000c0e0a, 0xc30c30c3, 0x0c3030c3, 0x0c30c30c, 0xcfccfc00, 0x30030030, 0x03000000,
    // [552]: B6 '¶'
    0x000c1206, 0xffcffc33, 0xf33f33f3, 0x3f33c33c, 0x33033033, 0x03303303, 0x30330330, 0x33033000,
    // [560]: B7 '·'
    0x0002020e, 0xf0000000,
    // [562]: B8 '¸'
    0x00040414, 0xcc330000,
    // [564]: B9 '¹'
    0x00060a02, 0x30c3cf30, 0xc30cfff0,
    // [567]: BA 'º'
    0x00080c06, 0x3c3cc3c3, 0xc3c33c3c, 0x0000ffff,
    // [571]: BB '»'
    0x000c0a0a, 0x0c30c330, 0xc30cc30c, 0x3030c30c, 0x0c30c300,
    // [576]: BC '¼'
    0x00121400, 0x0c030300, 0xc0c03c30, 0x0f030300, 0xc0c03030, 0x0c0c00cf, 0xc033fccc, 0x033300cc,
    0x30330c0f, 0xc303f0c0, 0xc00c3003, 0x0c00c300, 0x30000000,
    // [589]: BD '½'
    0x00121400, 0x0c030300, 0xc0c03c30, 0x0f030300, 0xc0c03030, 0x0c0c00cf, 0xc033ffcc, 0x03f300c0,
    0x30300c0f, 0xc303f0c0, 0x0c0c0303, 0x0fc0c3f0, 0x30000000,
    // [602]: BE '¾'
    0x00121400, 0x0c0fc303, 0xf0c0c030, 0x30030fc0, 0xc3f030c0, 0x0c3000cf, 0xc033fccc, 0x033300cc,
    0x30330c0f, 0xc303f0c0, 0xc00c3003, 0x0c00c300, 0x30000000,
    // [615]: BF '¿'
    0x00080e06, 0x30300000, 0x30300c0c, 0x0303c3c3, 0x3c3c0000,
    // [620]: C0 'À'
    0x000a1400, 0x0300c0c0, 0x30000000, 0xc0300c03, 0x0330cc33, 0x0ccfffff, 0xc0f03c0f, 0x03000000,
    // [628]: C1 'Á'
    0x000a1400, 0x300c00c0, 0x30000000, 0xc0300c03, 0x0330cc33, 0x0ccfffff, 0xc0f03c0f, 0x03000000,
    // [636]: C2 'Â'
    0x000a1400, 0x0c030330, 0xcc000000, 0xc0300c03, 0x0330cc33, 0x0ccfffff, 0xc0f03c0f, 0x03000000,
    // [644]: C3 'Ã'
    0x000a1400, 0xcf33c3cc, 0xf3000000, 0xc0300c03, 0x0330cc33, 0x0ccfffff, 0xc0f03c0f, 0x03000000,
    // [652]: C4 'Ä'
    0x000a1202, 0x330cc000, 0x000c0300, 0xc030330c, 0xc330ccff, 0xfffc0f03, 0xc0f03000,
    // [659]: C5 'Å'
    0x000a1400, 0x3f0fcc0f, 0x033f0fc0, 0xc0300c03, 0x0330cc33, 0x0ccfffff, 0xc0f03c0f, 0x03000000,
    // [667]: C6 'Æ'
    0x00100e06, 0xffc0ffc0, 0x03300330, 0x03300330, 0x3ffc3ffc, 0x030c030c, 0x03030303, 0xff03ff03,
    // [675]: C7 'Ç'
    0x000a1206, 0x3f0fcc0f, 0x0300c030, 0x0c0300c0, 0x3c0f033f, 0x0fc0c030, 0x0300c000,
    // [682]: C8 'È'
    0x00081400, 0x0c0c3030, 0x0000ffff, 0x03030303, 0x3f3f0303, 0x0303ffff,
    // [688]: C9 'É'
    0x00081400, 0x30300c0c, 0x0000ffff, 0x03030303, 0x3f3f0303, 0x0303ffff,
    // [694]: CA 'Ê'
    0x00081400, 0x0c0c3333, 0x0000ffff, 0x03030303, 0x3f3f0303, 0x0303ffff,
    // [700]: CB 'Ë'
    0x00081202, 0x33330000, 0xffff0303, 0x03033f3f, 0x03030303, 0xffff0000,
    // [706]: CC 'Ì'
    0x00041400, 0x33cc00cc, 0xcccccccc, 0xcccc0000,
    // [710]: CD 'Í'
    0x00041400, 0xcc330033, 0x33333333, 0x33330000,
    // [714]: CE 'Î'
    0x00061400, 0x30ccf300, 0x030c30c3, 0x0c30c30c, 0x30c30c00,
    // [719]: CF 'Ï'
    0x00061202, 0xcf300030, 0xc30c30c3, 0x0c30c30c, 0x30c00000,
    // [724]: D0 'Ð'
    0x000c0e06, 0x0fc0fc30, 0xc30cc0cc, 0x0cc3fc3f, 0xc0cc0c30, 0xc30c0fc0, 0xfc000000,
    // [731]: D1 'Ñ'
    0x000a1400, 0xcf33c3cc, 0xf300000c, 0x3f0fc3f0, 0xfccf33cc, 0xf33f0fc3, 0xf0fc3c0f, 0x03000000,
    // [739]: D2 'Ò'
    0x000a1400, 0x0300c0c0, 0x30000003, 0xf0fcc0f0, 0x3c0f03c0, 0xf03c0f03, 0xc0f033f0, 0xfc000000,
    // [747]: D3 'Ó'
    0x000a1400, 0x300c00c0, 0x30000003, 0xf0fcc0f0, 0x3c0f03c0, 0xf03c0f03, 0xc0f033f0, 0xfc000000,
    // [755]: D4 'Ô'
    0x000a1400, 0x0c030330, 0xcc000003, 0xf0fcc0f0, 0x3c0f03c0, 0xf03c0f03, 0xc0f033f0, 0xfc000000,
    // [763]: D5 'Õ'
    0x000a1400, 0xcf33c3cc, 0xf3000003, 0xf0fcc0f0, 0x3c0f03c0, 0xf03c0f03, 0xc0f033f0, 0xfc000000,
    // [771]: D6 'Ö'
    0x000a1202, 0x330cc000, 0x003f0fcc, 0x0f03c0f0, 0x3c0f03c0, 0xf03c0f03, 0x3f0fc000,
    // [778]: D7 '×'
    0x000a0a0a, 0xc0f03330, 0xcc0c0303, 0x30ccc0f0, 0x30000000,
    // [783]: D8 'Ø'
    0x000e0e06, 0xcfc33f03, 0x030c0c3c, 0x30f0c333, 0x0ccc30f0, 0xc3c3030c, 0x0c0fcc3f, 0x30000000,
    // [791]: D9 'Ù'
    0x000a1400, 0x0300c0c0, 0x3000000c, 0x0f03c0f0, 0x3c0f03c0, 0xf03c0f03, 0xc0f033f0, 0xfc000000,
    // [799]: DA 'Ú'
    0x000a1400, 0x300c00c0, 0x3000000c, 0x0f03c0f0, 0x3c0f03c0, 0xf03c0f03, 0xc0f033f0, 0xfc000000,
    // [807]: DB 'Û'
    0x000a1400, 0x0c030330, 0xcc00000c, 0x0f03c0f0, 0x3c0f03c0, 0xf03c0f03, 0xc0f033f0, 0xfc000000,
    // [815]: DC 'Ü'
    0x000a1202, 0x330cc000, 0x00c0f03c, 0x0f03c0f0, 0x3c0f03c0, 0xf03c0f03, 0x3f0fc000,
    // [822]: DD 'Ý'
    0x000a1400, 0x300c00c0, 0x3000000c, 0x0f03c0f0, 0x3330cc0c, 0x0300c030, 0x0c0300c0, 0x30000000,
    // [830]: DE 'Þ'
    0x00080e06, 0x0303033f, 0x3fc3c3c3, 0xc33f3f03, 0x03030000,
    // [835]: DF 'ß'
    0x000a0e06, 0x0f03c30c, 0xc30cc330, 0xcc3330cc, 0x3c0f033c, 0xcf300000,
    // [841]: E0 'à'
    0x00081004, 0x0c0c3030, 0x00003c3c, 0xc3c3fcfc, 0xc3c3fcfc,
    // [846]: E1 'á'
    0x00081004, 0x30300c0c, 0x00003c3c, 0xc3c3fcfc, 0xc3c3fcfc,
    // [851]: E2 'â'
    0x00081004, 0x3030cccc, 0x00003c3c, 0xc3c3fcfc, 0xc3c3fcfc,
    // [856]: E3 'ã'
    0x00081004, 0xcccc3333, 0x00003c3c, 0xc3c3fcfc, 0xc3c3fcfc,
    // [861]: E4 'ä'
    0x00080e06, 0xcccc0000, 0x3c3cc3c3, 0xfcfcc3c3, 0xfcfc0000,
    // [866]: E5 'å'
    0x00081202, 0x3c3cc3c3, 0x3c3c0000, 0x3c3cc3c3, 0xfcfcc3c3, 0xfcfc0000,
    // [872]: E6 'æ'
    0x000e0a0a, 0x3cf0f3cc, 0x3030c0ff, 0xf3ffc030, 0xc0c33cf0, 0xf3c00000,
    // [878]: E7 'ç'
    0x00080e0a, 0x3c3cc3c3, 0x0303c3c3, 0x3c3c3030, 0x0c0c0000,
    // [883]: E8 'è'
    0x00081004, 0x0c0c3030, 0x00003c3c, 0xc3c3ffff, 0x03033c3c,
    // [888]: E9 'é'
    0x00081004, 0x30300c0c, 0x00003c3c, 0xc3c3ffff, 0x0303fcfc,
    // [893]: EA 'ê'
    0x00081004, 0x0c0c3333, 0x00003c3c, 0xc3c3ffff, 0x03033c3c,
    // [898]: EB 'ë'
    0x00080e06, 0xc3c30000, 0x3c3cc3c3, 0xffff0303, 0xfcfc0000,
    // [903]: EC 'ì'
    0x00041004, 0x33cc00cc, 0xcccccccc,
    // [906]: ED 'í'
    0x00041004, 0xcc330033, 0x33333333,
    // [909]: EE 'î'
    0x00061004, 0x30ccf300, 0x030c30c3, 0x0c30c30c,
    // [913]: EF 'ï'
    0x00060e06, 0xcf300030, 0xc30c30c3, 0x0c30c000,
    // [917]: F0 'ð'
    0x00081004, 0x03333c0f, 0x3330cccc, 0xc3c3c3c3, 0xc3c33c3c,
    // [922]: F1 'ñ'
    0x00081004, 0xcccc3333, 0x00003f3f, 0xc3c3c3c3, 0xc3c3c3c3,
    // [927]: F2 'ò'
    0x00081004, 0x0c0c3030, 0x00003c3c, 0xc3c3c3c3, 0xc3c33c3c,
    // [932]: F3 'ó'
    0x00081004, 0x30300c0c, 0x00003c3c, 0xc3c3c3c3, 0xc3c33c3c,
    // [937]: F4 'ô'
    0x00081004, 0x0c0c3333, 0x00003c3c, 0xc3c3c3c3, 0xc3c33c3c,
    // [942]: F5 'õ'
    0x00081004, 0xcccc3333, 0x00003c3c, 0xc3c3c3c3, 0xc3c33c3c,
    // [947]: F6 'ö'
    0x00080e06, 0xc3c30000, 0x3c3cc3c3, 0xc3c3c3c3, 0x3c3c0000,
    // [952]: F7 '÷'
    0x000a0a08, 0x0c030000, 0x00fffff0, 0x00000c03, 0x00000000,
    // [957]: F8 'ø'
    0x000c0c08, 0xcf0cf030, 0xc30c3cc3, 0xcc33c33c, 0x30c30c0f, 0x30f30000,
    // [963]: F9 'ù'
    0x00081004, 0x0c0c3030, 0x0000c3c3, 0xc3c3c3c3, 0xc3c3fcfc,
    // [968]: FA 'ú'
    0x00081004, 0x30300c0c, 0x0000c3c3, 0xc3c3c3c3, 0xc3c3fcfc,
    // [973]: FB 'û'
    0x00081004, 0x0c0c3333, 0x0000c3c3, 0xc3c3c3c3, 0xc3c3fcfc,
    // [978]: FC 'ü'
    0x00080e06, 0xc3c30000, 0xc3c3c3c3, 0xc3c3c3c3, 0xfcfc0000,
    // [983]: FD 'ý'
    0x000a1404, 0x300c00c0, 0x3000000c, 0x0f03c0f0, 0x3330cc33, 0x0cc0c030, 0x0c03003c, 0x0f000000,
    // [991]: FE 'þ'
    0x00081206, 0x03030303, 0x3f3fc3c3, 0xc3c3c3c3, 0x3f3f0303, 0x03030000,
    // [997]: FF 'ÿ'
    0x000a1206, 0x330cc000, 0x00c0f03c, 0x0f03330c, 0xc330cc0c, 0x0300c030, 0x03c0f000,
    // [1004]: 152 'Œ'
    0x00100e06, 0xfffcfffc, 0x03030303, 0x03030303, 0x3f033f03, 0x03030303, 0x03030303, 0xfffcfffc,
    // [1012]: 153 'œ'
    0x000e0a0a, 0x3cf0f3cc, 0x30f0c3ff, 0x0ffc3030, 0xc0c33ff0, 0xffc00000,
    // [1018]: 2018 '‘'
    0x00040604, 0xcc333300,
    // [1020]: 2019 '’'
    0x00040604, 0xcccc3300,
    // [1022]: 201A '‚'
    0x00040612, 0xcccc3300,
    // [1024]: 201B '‛'
    0x00040604, 0x3333cc00,
    // [1026]: 201C '“'
    0x00080604, 0xcccc3333, 0x33330000,
    // [1029]: 201D '”'
    0x00080604, 0xcccccccc, 0x33330000,
    // [1032]: 201E '„'
    0x00080612, 0xcccccccc, 0x33330000,
    // [1035]: 201F '‟'
    0x00080604, 0x33333333, 0xcccc0000,
    // [1038]: 2020 '†'
    0x00060a04, 0x30cfff30, 0xc30c30c0,
    // [1041]: 2021 '‡'
    0x00061206, 0x30c30cff, 0xf30c30c3, 0x0cfff30c, 0x30c00000,
    // [1046]: 2022 '•'
    0x000a0a08, 0x3f0fcfff, 0xffffffff, 0xffff3f0f, 0xc0000000,
    // [1051]: 20AC '€'
    0x000c0e04, 0x3f03f0c0, 0xcc0c0ff0, 0xff00c00c, 0x0ff0ffc0, 0xcc0c3f03, 0xf0000000,
    // [1058]: FFFD '�'
    0x00121402, 0x00c00030, 0x003f000f, 0xc00f3c03, 0xcf03ccf0, 0xf33cfcff, 0xff3ffff3, 0xfffcff3f,
    0xff0fffc0, 0xf3c03cf0, 0x03f000fc, 0x000c0003, 0x00000000,
];

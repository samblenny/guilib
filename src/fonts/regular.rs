#![allow(dead_code)]
//! Regular Font
//
// This code incorporates glyphs from the Geneva bitmap typeface which was
// designed by Susan Kare and released by Apple in 1984. Geneva is a registered
// trademark of Apple Inc.
//

/// Maximum height of glyph patterns in this bitmap typeface.
/// This will be true: h + y_offset <= MAX_HEIGHT
pub const MAX_HEIGHT: u8 = 30;

/// Return Okay(offset into DATA[]) for start of blit pattern for grapheme cluster.
///
/// Before doing an expensive lookup for the whole cluster, this does a pre-filter
/// check to see whether the first character falls into one of the codepoint ranges
/// for Unicode blocks included in this font.
///
pub fn get_blit_pattern_offset(cluster: &str) -> Result<usize, super::GlyphNotFound> {
    let first_char: u32;
    match cluster.chars().next() {
        Some(c) => first_char = c as u32,
        None => return Err(super::GlyphNotFound),
    }
    return match first_char {
        0x00..=0x7E        // Basic Latin
        | 0x80..=0xFF      // Latin 1
        | 0x100..=0x17F    // Latin Extended A
        | 0x2000..=0x206F  // General Punctuation
        | 0x20A0..=0x20CF  // Currency Symbols
        | 0xE700..=0xE70C  // Subset of Private Use Area (UI Icons)
        | 0xFFFD => {      // Subset of Specials (replacement character)
            find_pattern(cluster)
        }
        _ => Err(super::GlyphNotFound),
    };
}

/// Use binary search on table of grapheme cluster hashes to find blit pattern for grapheme cluster
fn find_pattern(cluster: &str) -> Result<usize, super::GlyphNotFound> {
    let seed = 0;
    let key = super::murmur3(cluster, seed);
    match HASHED_CLUSTERS.binary_search(&key) {
        Ok(index) => return Ok(PATTERN_OFFSETS[index]),
        _ => Err(super::GlyphNotFound),
    }
}

// Index of murmur3(grapheme cluster) with sort order matching PATTERN_OFFSETS
const HASHED_CLUSTERS: [u32; 219] = [
    0x00A5B2D1, // Battery_75
    0x00EAC56E, // '°'
    0x0254FD66, // '®'
    0x049BF55E, // '}'
    0x056DE9E6, // 'Î'
    0x06FE368D, // 'Ö'
    0x0B9EA876, // '8'
    0x0D0042B1, // '•'
    0x111FB3D8, // Shift_Arrow
    0x13156136, // 'F'
    0x131AB870, // 'O'
    0x137E3259, // 'Õ'
    0x145967B1, // Radio_Off
    0x1586C0AA, // 'j'
    0x1672836B, // '`'
    0x16F9A05F, // '#'
    0x18D53BB6, // '('
    0x18E68A7D, // '‚'
    0x18E6F281, // 'à'
    0x1A01594C, // 'Œ'
    0x1A5A66CD, // 'ø'
    0x1ACA36BB, // '€'
    0x1B0DE252, // '¢'
    0x1B43C661, // '¤'
    0x1BF82D35, // 'y'
    0x1C505947, // 'í'
    0x1CFCD293, // ')'
    0x1F79A4D0, // '$'
    0x21840E11, // 'q'
    0x22ABE92F, // 'K'
    0x23BDB32A, // Battery_25
    0x252158D3, // '¥'
    0x25872B22, // 'ë'
    0x259437AC, // '§'
    0x259E68D3, // '6'
    0x268B196D, // 'ä'
    0x26FF6E36, // '3'
    0x275E5021, // '¨'
    0x2766361E, // 'Ú'
    0x278A377A, // 'Ì'
    0x28CA6AD7, // 'Ø'
    0x291E971E, // 'Ï'
    0x2A3184F1, // 'U'
    0x2B038801, // 'a'
    0x2C17F13B, // 'r'
    0x2D2269F8, // 'm'
    0x31099644, // 't'
    0x315F5687, // 'I'
    0x3280724D, // '»'
    0x3673CB35, // '{'
    0x36953E7B, // 'õ'
    0x372ED469, // '“'
    0x388FB155, // 'ï'
    0x39056A43, // 'É'
    0x39BC06BB, // 'L'
    0x3A073FBE, // 'È'
    0x3A3C25C2, // '”'
    0x3AE00F65, // 'ü'
    0x3CA7FBAF, // Radio_0
    0x3E168B4D, // 'ó'
    0x3E82329E, // 'û'
    0x402DCFF6, // '×'
    0x4236CDD1, // '¾'
    0x43603571, // 'C'
    0x44C216BA, // 'R'
    0x44D0C2C4, // 'ê'
    0x450DB83D, // 'Û'
    0x45EF573F, // ';'
    0x4938EA00, // '7'
    0x4AB0FC13, // '4'
    0x4B9C97D9, // Radio_1
    0x4BCD3197, // 'A'
    0x4CC5A899, // '|'
    0x4F92356C, // 'µ'
    0x511BCA1D, // 'á'
    0x524C4F7E, // 'B'
    0x52D3CB36, // '%'
    0x554A5349, // 'Æ'
    0x56DEE618, // 'v'
    0x58A5DA35, // '�'
    0x5A352E5E, // 'c'
    0x5B06F60A, // '1'
    0x5C904E6C, // '_'
    0x5CA0F32F, // '&'
    0x5E39BCEE, // '´'
    0x5EB9D2A0, // 'ç'
    0x5EE10367, // 'ò'
    0x5FF80125, // 'S'
    0x61901224, // '5'
    0x62D494F4, // 'Ý'
    0x68E75BE7, // '@'
    0x6984AA2D, // 'Í'
    0x6BA96CC3, // 'Ü'
    0x6E3C05CF, // 'ý'
    0x6E85A379, // 'x'
    0x6EF8ED06, // '­'
    0x6F6FD0E9, // Radio_2
    0x6FA3C127, // 'ã'
    0x71EE15F2, // 'Ò'
    0x72DE1EFB, // '*'
    0x73688484, // 'p'
    0x7619D892, // 'ª'
    0x7627CA6C, // 'D'
    0x76554D10, // 'Ç'
    0x779CDA2E, // 'Y'
    0x77CD98B4, // Battery_50
    0x78FD8C32, // 'u'
    0x79026F8E, // 'Å'
    0x7B6B1369, // '9'
    0x7BA23029, // 'J'
    0x7E2A203C, // '÷'
    0x7FDE2AE0, // '\\'
    0x803AF153, // '['
    0x804744C4, // 'e'
    0x80E4E277, // '‡'
    0x81BAB05B, // '\''
    0x839D40CB, // '£'
    0x84510F7D, // 'o'
    0x84FF8F94, // 'Ô'
    0x86A8B043, // ','
    0x885E576B, // '‟'
    0x8AE0A2B9, // 'å'
    0x8B80C01D, // '‛'
    0x8B86BA97, // '='
    0x8C1A54EE, // 'g'
    0x8C60DA30, // 'œ'
    0x8C69C315, // ' '
    0x8D741045, // 'Þ'
    0x8DB20A6B, // '!'
    0x8EA3F31F, // '³'
    0x8F0BE1DB, // '„'
    0x8F56D335, // '†'
    0x90AE32C9, // Backspace_Symbol
    0x9194CCB5, // '¬'
    0x92979158, // 'ô'
    0x9500C437, // '<'
    0x9582C02E, // 'G'
    0x95C4DC63, // 'Ã'
    0x97A9C1DA, // '¶'
    0x97C63B05, // ' '
    0x98453FD6, // '?'
    0x99933C47, // 'w'
    0x9AF9DB7B, // Radio_3
    0xA10C8120, // '·'
    0xA13B361A, // '/'
    0xA1A410C7, // '2'
    0xA5872342, // 'k'
    0xA5F3368B, // '¹'
    0xA6C64F29, // '‘'
    0xA7883CA5, // '¸'
    0xA7B0FE42, // '©'
    0xA891E88A, // '0'
    0xABE208D0, // 'P'
    0xABF561B2, // '«'
    0xAD97621D, // 'Ð'
    0xADABEA12, // 'd'
    0xAF41DB71, // ']'
    0xB12BF2EE, // 'V'
    0xB1A64F1E, // 'Â'
    0xB1BF38B8, // 'º'
    0xB35F1351, // 'E'
    0xB55DA0FA, // 'î'
    0xB5CA05F5, // '¦'
    0xB7001103, // '¡'
    0xB8ED6BE2, // 'X'
    0xB9229669, // '¿'
    0xBA8E9829, // 'Ä'
    0xBB608964, // 'ú'
    0xBFE47D0D, // 'Q'
    0xC06E932F, // 'T'
    0xC0837C42, // '¼'
    0xC0E7AB63, // '¯'
    0xC1209A73, // 'f'
    0xC20FE7B7, // 'þ'
    0xC2DDC575, // '"'
    0xC3A1EBC2, // Enter_Symbol
    0xC6EF143E, // 'ß'
    0xC890E3DD, // '>'
    0xC98BE718, // 'Á'
    0xC99309A2, // '-'
    0xC9BEA311, // 'Z'
    0xCAD0511F, // 'é'
    0xCAECCF17, // 's'
    0xCAF24984, // 'Ù'
    0xCAF83468, // 'H'
    0xCB7242E0, // '+'
    0xCD3FDBE3, // 'M'
    0xCE94AE25, // 'b'
    0xD040D3E3, // 'è'
    0xD4D6097A, // 'Ë'
    0xD4FFA898, // ':'
    0xD7B03F23, // 'z'
    0xD83534E2, // 'Ñ'
    0xDB01AF22, // 'n'
    0xDC135ABD, // 'N'
    0xDD061C7A, // '²'
    0xDE073178, // Battery_05
    0xDE3FB757, // 'ì'
    0xE14F2E81, // 'â'
    0xE2532EDC, // 'ð'
    0xE29813B0, // '’'
    0xE2AA1EBB, // '.'
    0xE51FF77C, // '½'
    0xE5CA55BF, // 'i'
    0xE5DA3FED, // '±'
    0xE6556D01, // 'h'
    0xE9131AA1, // 'ù'
    0xEBC8143D, // 'ö'
    0xEC99A516, // 'À'
    0xEF026B52, // 'l'
    0xF0F19A38, // '~'
    0xF29AB82D, // 'W'
    0xF446CD1A, // 'Ó'
    0xF520EF41, // 'æ'
    0xF64ED921, // 'ñ'
    0xFB89D3F4, // 'ÿ'
    0xFBFECC1C, // 'Ê'
    0xFC2C2430, // '^'
    0xFEE9E453, // Battery_99
];

// Lookup table from hashed cluster to blit pattern offset (sort order matches HASHED_CLUSTERS)
const PATTERN_OFFSETS: [usize; 219] = [
    1429, // Battery_75
    690,  // '°'
    679,  // '®'
    597,  // '}'
    927,  // 'Î'
    998,  // 'Ö'
    139,  // '8'
    1384, // '•'
    1514, // Shift_Arrow
    231,  // 'F'
    297,  // 'O'
    988,  // 'Õ'
    1501, // Radio_Off
    477,  // 'j'
    419,  // '`'
    7,    // '#'
    42,   // '('
    1353, // '‚'
    1087, // 'à'
    1327, // 'Œ'
    1259, // 'ø'
    1390, // '€'
    611,  // '¢'
    626,  // '¤'
    573,  // 'y'
    1186, // 'í'
    48,   // ')'
    13,   // '$'
    521,  // 'q'
    265,  // 'K'
    1409, // Battery_25
    633,  // '¥'
    1175, // 'ë'
    643,  // '§'
    123,  // '6'
    1119, // 'ä'
    98,   // '3'
    653,  // '¨'
    1033, // 'Ú'
    919,  // 'Ì'
    1013, // 'Ø'
    933,  // 'Ï'
    348,  // 'U'
    421,  // 'a'
    528,  // 'r'
    493,  // 'm'
    540,  // 't'
    254,  // 'I'
    740,  // '»'
    588,  // '{'
    1239, // 'õ'
    1357, // '“'
    1195, // 'ï'
    893,  // 'É'
    273,  // 'L'
    884,  // 'È'
    1360, // '”'
    1291, // 'ü'
    1488, // Radio_0
    1223, // 'ó'
    1283, // 'û'
    1008, // '×'
    771,  // '¾'
    208,  // 'C'
    323,  // 'R'
    1167, // 'ê'
    1043, // 'Û'
    157,  // ';'
    131,  // '7'
    106,  // '4'
    1475, // Radio_1
    191,  // 'A'
    594,  // '|'
    707,  // 'µ'
    1095, // 'á'
    200,  // 'B'
    21,   // '%'
    861,  // 'Æ'
    552,  // 'v'
    1550, // '�'
    434,  // 'c'
    86,   // '1'
    417,  // '_'
    30,   // '&'
    705,  // '´'
    1144, // 'ç'
    1215, // 'ò'
    331,  // 'S'
    115,  // '5'
    1063, // 'Ý'
    182,  // '@'
    923,  // 'Í'
    1053, // 'Ü'
    1298, // 'ý'
    567,  // 'x'
    677,  // '­'
    1462, // Radio_2
    1111, // 'ã'
    958,  // 'Ò'
    54,   // '*'
    514,  // 'p'
    664,  // 'ª'
    216,  // 'D'
    874,  // 'Ç'
    384,  // 'Y'
    1419, // Battery_50
    546,  // 'u'
    850,  // 'Å'
    147,  // '9'
    257,  // 'J'
    1254, // '÷'
    402,  // '\\'
    398,  // '['
    447,  // 'e'
    1376, // '‡'
    40,   // '\''
    618,  // '£'
    508,  // 'o'
    978,  // 'Ô'
    64,   // ','
    1366, // '‟'
    1126, // 'å'
    1355, // '‛'
    165,  // '='
    459,  // 'g'
    1340, // 'œ'
    0,    // ' '
    1072, // 'Þ'
    2,    // '!'
    702,  // '³'
    1363, // '„'
    1369, // '†'
    1522, // Backspace_Symbol
    674,  // '¬'
    1231, // 'ô'
    160,  // '<'
    238,  // 'G'
    827,  // 'Ã'
    717,  // '¶'
    606,  // ' '
    174,  // '?'
    558,  // 'w'
    1449, // Radio_3
    728,  // '·'
    70,   // '/'
    90,   // '2'
    483,  // 'k'
    732,  // '¹'
    1349, // '‘'
    730,  // '¸'
    655,  // '©'
    78,   // '0'
    305,  // 'P'
    669,  // '«'
    939,  // 'Ð'
    440,  // 'd'
    410,  // ']'
    356,  // 'V'
    815,  // 'Â'
    735,  // 'º'
    224,  // 'E'
    1190, // 'î'
    640,  // '¦'
    608,  // '¡'
    377,  // 'X'
    784,  // '¿'
    839,  // 'Ä'
    1275, // 'ú'
    313,  // 'Q'
    339,  // 'T'
    745,  // '¼'
    688,  // '¯'
    453,  // 'f'
    1309, // 'þ'
    5,    // '"'
    1538, // Enter_Symbol
    1079, // 'ß'
    169,  // '>'
    803,  // 'Á'
    66,   // '-'
    391,  // 'Z'
    1159, // 'é'
    534,  // 's'
    1023, // 'Ù'
    246,  // 'H'
    59,   // '+'
    280,  // 'M'
    427,  // 'b'
    1151, // 'è'
    911,  // 'Ë'
    155,  // ':'
    582,  // 'z'
    948,  // 'Ñ'
    502,  // 'n'
    289,  // 'N'
    699,  // '²'
    1399, // Battery_05
    1182, // 'ì'
    1103, // 'â'
    1200, // 'ð'
    1351, // '’'
    68,   // '.'
    758,  // '½'
    474,  // 'i'
    693,  // '±'
    467,  // 'h'
    1267, // 'ù'
    1247, // 'ö'
    791,  // 'À'
    490,  // 'l'
    603,  // '~'
    365,  // 'W'
    968,  // 'Ó'
    1135, // 'æ'
    1207, // 'ñ'
    1317, // 'ÿ'
    902,  // 'Ê'
    414,  // '^'
    1439, // Battery_99
];

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
pub const DATA: [u32; 1563] = [
    // [0]: 20 ' '
    0x0004020e, 0x00000000,
    // [2]: 21 '!'
    0x00021206, 0xfffffff0, 0xf0000000,
    // [5]: 22 '"'
    0x00060406, 0xcf3cf300,
    // [7]: 23 '#'
    0x000e0a06, 0x3300cc0f, 0xff3ffc0c, 0xc03303ff, 0xcfff0330, 0x0cc00000,
    // [13]: 24 '$'
    0x000a1604, 0x0c0303f0, 0xfcccf330, 0xcc330f03, 0xc3c0f0cc, 0x330cc330, 0xccf333f0, 0xfc0c0300,
    // [21]: 25 '%'
    0x000e1206, 0xfff3ffcc, 0x30f0c333, 0x0ccc30cf, 0x033c0300, 0x0c03cc0f, 0x30c3330c, 0xcc30f0c3,
    0x3c0cf030,
    // [30]: 26 '&'
    0x00101206, 0x00f000f0, 0x030c030c, 0x00cc00cc, 0x00300030, 0x00cc00cc, 0x33033303, 0x0c030c03,
    0x33033303, 0xc0fcc0fc,
    // [40]: 27 '\''
    0x00020406, 0xff000000,
    // [42]: 28 '('
    0x00061604, 0xc3030c30, 0xc0c30c30, 0xc30c30c3, 0x30c30cc3, 0x00000000,
    // [48]: 29 ')'
    0x00061604, 0x0c330c30, 0xcc30c30c, 0x30c30c30, 0x30c30c0c, 0x30000000,
    // [54]: 2A '*'
    0x000a0a08, 0x0c030ccf, 0x333f0fc3, 0x30ccc0f0, 0x30000000,
    // [59]: 2B '+'
    0x000a0a0a, 0x0c0300c0, 0x30fffff0, 0xc0300c03, 0x00000000,
    // [64]: 2C ','
    0x00040616, 0xcccc3300,
    // [66]: 2D '-'
    0x000a020e, 0xfffff000,
    // [68]: 2E '.'
    0x00020216, 0xf0000000,
    // [70]: 2F '/'
    0x000a1404, 0xc0300c03, 0x00300c03, 0x00c00c03, 0x00c03003, 0x00c0300c, 0x00c0300c, 0x03000000,
    // [78]: 30 '0'
    0x000c1206, 0x0f00f030, 0xc30cc03c, 0x03c03c03, 0xc03c03c0, 0x3c03c03c, 0x0330c30c, 0x0f00f000,
    // [86]: 31 '1'
    0x00041206, 0xccffcccc, 0xcccccccc, 0xcc000000,
    // [90]: 32 '2'
    0x000c1206, 0x3f03f0c0, 0xcc0cc03c, 0x03c00c00, 0x3003000c, 0x00c00300, 0x3000c00c, 0xffffff00,
    // [98]: 33 '3'
    0x000c1206, 0xffffff30, 0x03000c00, 0xc00f00f0, 0x300300c0, 0x0c00c00c, 0x00303303, 0x0fc0fc00,
    // [106]: 34 '4'
    0x000e1206, 0x3000c003, 0xc00f0033, 0x00cc030c, 0x0c303030, 0xc0c300cc, 0x03ffffff, 0xf3000c00,
    0x3000c000,
    // [115]: 35 '5'
    0x000c1206, 0xffffff00, 0x30030030, 0x033ff3ff, 0xc00c00c0, 0x0c00c00c, 0x00c03c03, 0x3fc3fc00,
    // [123]: 36 '6'
    0x000c1206, 0x3f03f000, 0xc00c0030, 0x033ff3ff, 0xc03c03c0, 0x3c03c03c, 0x03c03c03, 0x3fc3fc00,
    // [131]: 37 '7'
    0x000c1206, 0xffffffc0, 0x0c003003, 0x00300300, 0x0c00c00c, 0x00c00300, 0x30030030, 0x03003000,
    // [139]: 38 '8'
    0x000c1206, 0x3fc3fcc0, 0x3c03c03c, 0x033fc3fc, 0xc03c03c0, 0x3c03c03c, 0x03c03c03, 0x3fc3fc00,
    // [147]: 39 '9'
    0x000c1206, 0x3fc3fcc0, 0x3c03c03c, 0x03c03c03, 0xc03c03ff, 0xcffcc00c, 0x00300300, 0x0fc0fc00,
    // [155]: 3A ':'
    0x00020c0c, 0xf0000f00,
    // [157]: 3B ';'
    0x0004100c, 0xcc000000, 0x00cccc33,
    // [160]: 3C '<'
    0x00080e0a, 0xc0c03030, 0x0c0c0303, 0x0c0c3030, 0xc0c00000,
    // [165]: 3D '='
    0x000a080e, 0xfffff000, 0x0000000f, 0xffff0000,
    // [169]: 3E '>'
    0x00080e0a, 0x03030c0c, 0x3030c0c0, 0x30300c0c, 0x03030000,
    // [174]: 3F '?'
    0x000c1206, 0x3fc3fcc0, 0x3c03c03c, 0x03300300, 0x0c00c003, 0x00300300, 0x30000000, 0x03003000,
    // [182]: 40 '@'
    0x00101008, 0x0ff00ff0, 0x300c300c, 0xcfc3cfc3, 0xcc33cc33, 0xcc33cc33, 0x3fc33fc3, 0x000c000c,
    0x03f003f0,
    // [191]: 41 'A'
    0x000e1206, 0x03000c00, 0x3000c00c, 0xc03300cc, 0x03303030, 0xc0c3030c, 0x0cffffff, 0xfc00f003,
    0xc00f0030,
    // [200]: 42 'B'
    0x000c1206, 0x3ff3ffc0, 0x3c03c03c, 0x033ff3ff, 0xc03c03c0, 0x3c03c03c, 0x03c03c03, 0x3ff3ff00,
    // [208]: 43 'C'
    0x000c1206, 0x3fc3fcc0, 0x3c030030, 0x03003003, 0x00300300, 0x30030030, 0x03c03c03, 0x3fc3fc00,
    // [216]: 44 'D'
    0x000c1206, 0x0ff0ff30, 0x3303c03c, 0x03c03c03, 0xc03c03c0, 0x3c03c03c, 0x03303303, 0x0ff0ff00,
    // [224]: 45 'E'
    0x000a1206, 0xfffff00c, 0x0300c030, 0x0c033fcf, 0xf00c0300, 0xc0300c03, 0xfffff000,
    // [231]: 46 'F'
    0x000a1206, 0xfffff00c, 0x0300c030, 0x0c033fcf, 0xf00c0300, 0xc0300c03, 0x00c03000,
    // [238]: 47 'G'
    0x000c1206, 0x3fc3fcc0, 0x3c030030, 0x03003003, 0xfc3fc3c0, 0x3c03c03c, 0x03c03c03, 0x3fc3fc00,
    // [246]: 48 'H'
    0x000c1206, 0xc03c03c0, 0x3c03c03c, 0x03c03c03, 0xffffffc0, 0x3c03c03c, 0x03c03c03, 0xc03c0300,
    // [254]: 49 'I'
    0x00021206, 0xffffffff, 0xf0000000,
    // [257]: 4A 'J'
    0x000c1206, 0xc00c00c0, 0x0c00c00c, 0x00c00c00, 0xc00c00c0, 0x0c00c03c, 0x03c03c03, 0x3fc3fc00,
    // [265]: 4B 'K'
    0x000c1206, 0xc03c0330, 0x33030c30, 0xc3033033, 0x00f00f03, 0x30330c30, 0xc3303303, 0xc03c0300,
    // [273]: 4C 'L'
    0x000a1206, 0x00c0300c, 0x0300c030, 0x0c0300c0, 0x300c0300, 0xc0300c03, 0xfffff000,
    // [280]: 4D 'M'
    0x000e1206, 0xc00f003f, 0x03fc0fcc, 0xcf333c30, 0xf0c3c00f, 0x003c00f0, 0x03c00f00, 0x3c00f003,
    0xc00f0030,
    // [289]: 4E 'N'
    0x000c1206, 0xc0fc0fc0, 0xfc0fc33c, 0x33c33c33, 0xcc3cc3cc, 0x3cc3f03f, 0x03f03f03, 0xc03c0300,
    // [297]: 4F 'O'
    0x000c1206, 0x3fc3fcc0, 0x3c03c03c, 0x03c03c03, 0xc03c03c0, 0x3c03c03c, 0x03c03c03, 0x3fc3fc00,
    // [305]: 50 'P'
    0x000c1206, 0x3ff3ffc0, 0x3c03c03c, 0x03c03c03, 0x3ff3ff00, 0x30030030, 0x03003003, 0x00300300,
    // [313]: 51 'Q'
    0x000c1606, 0x3fc3fcc0, 0x3c03c03c, 0x03c03c03, 0xc03c03c0, 0x3c03c03c, 0x03c03c03, 0x3fc3fc0c,
    0x00c0f00f, 0x00000000,
    // [323]: 52 'R'
    0x000c1206, 0x3ff3ffc0, 0x3c03c03c, 0x03c03c03, 0x3ff3ff0c, 0x30c33033, 0x03c03c03, 0xc03c0300,
    // [331]: 53 'S'
    0x000c1206, 0x3fc3fcc0, 0x3c030030, 0x03003003, 0x3fc3fcc0, 0x0c00c00c, 0x00c03c03, 0x3fc3fc00,
    // [339]: 54 'T'
    0x000e1206, 0xfffffff0, 0x3000c003, 0x000c0030, 0x00c00300, 0x0c003000, 0xc003000c, 0x003000c0,
    0x03000c00,
    // [348]: 55 'U'
    0x000c1206, 0xc03c03c0, 0x3c03c03c, 0x03c03c03, 0xc03c03c0, 0x3c03c03c, 0x03c03c03, 0x3fc3fc00,
    // [356]: 56 'V'
    0x000e1206, 0xc00f003c, 0x00f003c0, 0x0f003303, 0x0c0c3030, 0xc0c0cc03, 0x300cc033, 0x003000c0,
    0x03000c00,
    // [365]: 57 'W'
    0x00121206, 0xc000f000, 0x3c000f00, 0x03c0c0f0, 0x30330c30, 0xc30c3333, 0x0cccc333, 0x30cccc0c,
    0x0c030300, 0xc0c03030, 0x0c0c0303, 0x00000000,
    // [377]: 58 'X'
    0x000a1206, 0xc0f03c0f, 0x03330cc3, 0x30cc0c03, 0x0330cc33, 0x0ccc0f03, 0xc0f03000,
    // [384]: 59 'Y'
    0x000a1206, 0xc0f03c0f, 0x03c0f033, 0x30cc330c, 0xc0c0300c, 0x0300c030, 0x0c030000,
    // [391]: 5A 'Z'
    0x000a1206, 0xfffffc03, 0x00300c03, 0x00c00c03, 0x00300c03, 0x00c00c03, 0xfffff000,
    // [398]: 5B '['
    0x00041604, 0xff333333, 0x33333333, 0x3333ff00,
    // [402]: 5C '\\'
    0x000a1404, 0x00c0300c, 0x030300c0, 0x300c0c03, 0x00c03030, 0x0c0300c0, 0xc0300c03, 0x00000000,
    // [410]: 5D ']'
    0x00041604, 0xffcccccc, 0xcccccccc, 0xccccff00,
    // [414]: 5E '^'
    0x00060806, 0x30c30ccf, 0x3cf30000,
    // [417]: 5F '_'
    0x00100216, 0xffffffff,
    // [419]: 60 '`'
    0x00040406, 0x33cc0000,
    // [421]: 61 'a'
    0x000a0e0a, 0x3f0fcc0f, 0x03c0300f, 0xf3fcc0f0, 0x3c0f03ff, 0x3fc00000,
    // [427]: 62 'b'
    0x000a1206, 0x00c0300c, 0x033fcffc, 0x0f03c0f0, 0x3c0f03c0, 0xf03c0f03, 0x3fcff000,
    // [434]: 63 'c'
    0x000a0e0a, 0x3f0fcc0f, 0x0300c030, 0x0c0300c0, 0x3c0f033f, 0x0fc00000,
    // [440]: 64 'd'
    0x000a1206, 0xc0300c03, 0x00ff3fcc, 0x0f03c0f0, 0x3c0f03c0, 0xf03c0f03, 0xff3fc000,
    // [447]: 65 'e'
    0x000a0e0a, 0x3f0fcc0f, 0x03c0f03f, 0xffff00c0, 0x3c0f033f, 0x0fc00000,
    // [453]: 66 'f'
    0x00081206, 0xf0f00c0c, 0x3f3f0c0c, 0x0c0c0c0c, 0x0c0c0c0c, 0x0c0c0000,
    // [459]: 67 'g'
    0x000a140a, 0xff3fcc0f, 0x03c0f03c, 0x0f03c0f0, 0x3c0f03ff, 0x3fcc0300, 0xc0f033f0, 0xfc000000,
    // [467]: 68 'h'
    0x000a1206, 0x00c0300c, 0x033ccf3c, 0x3f0fc0f0, 0x3c0f03c0, 0xf03c0f03, 0xc0f03000,
    // [474]: 69 'i'
    0x00021206, 0xf0ffffff, 0xf0000000,
    // [477]: 6A 'j'
    0x00061806, 0xc30000c3, 0x0c30c30c, 0x30c30c30, 0xc30c30c3, 0x03cf0000,
    // [483]: 6B 'k'
    0x000a1206, 0x00c0300c, 0x03c0f033, 0x0cc30cc3, 0x303c0f0c, 0xc3330cc3, 0xc0f03000,
    // [490]: 6C 'l'
    0x00021206, 0xffffffff, 0xf0000000,
    // [493]: 6D 'm'
    0x00120e0a, 0x3c3ccf0f, 0x3c3c3f0f, 0x0fc0c0f0, 0x303c0c0f, 0x0303c0c0, 0xf0303c0c, 0x0f0303c0,
    0xc0f03030,
    // [502]: 6E 'n'
    0x000a0e0a, 0x3ccf3c3f, 0x0fc0f03c, 0x0f03c0f0, 0x3c0f03c0, 0xf0300000,
    // [508]: 6F 'o'
    0x000a0e0a, 0x3f0fcc0f, 0x03c0f03c, 0x0f03c0f0, 0x3c0f033f, 0x0fc00000,
    // [514]: 70 'p'
    0x000a120a, 0x3fcffc0f, 0x03c0f03c, 0x0f03c0f0, 0x3c0f033f, 0xcff00c03, 0x00c03000,
    // [521]: 71 'q'
    0x000a120a, 0xff3fcc0f, 0x03c0f03c, 0x0f03c0f0, 0x3c0f03ff, 0x3fcc0300, 0xc0300000,
    // [528]: 72 'r'
    0x000a0e0a, 0xfcff303c, 0x0f00c030, 0x0c0300c0, 0x300c0300, 0xc0300000,
    // [534]: 73 's'
    0x000a0e0a, 0x3f0fcc0f, 0x0300c033, 0xf0fcc030, 0x0c0f033f, 0x0fc00000,
    // [540]: 74 't'
    0x00081206, 0x0c0c0c0c, 0x3f3f0c0c, 0x0c0c0c0c, 0x0c0c0c0c, 0xf0f00000,
    // [546]: 75 'u'
    0x000a0e0a, 0xc0f03c0f, 0x03c0f03c, 0x0f03c0f0, 0x3f0fc3cf, 0x33c00000,
    // [552]: 76 'v'
    0x000a0e0a, 0xc0f03c0f, 0x03c0f033, 0x30cc330c, 0xc0c0300c, 0x03000000,
    // [558]: 77 'w'
    0x00120e0a, 0xc0c0f030, 0x3c0c0f03, 0x0333330c, 0xccc33330, 0xcccc0c0c, 0x030300c0, 0xc030300c,
    0x0c030300,
    // [567]: 78 'x'
    0x000a0e0a, 0xc0f03c0f, 0x03330cc0, 0xc030330c, 0xcc0f03c0, 0xf0300000,
    // [573]: 79 'y'
    0x000c140a, 0xc0cc0cc0, 0xcc0cc0cc, 0x0c330330, 0x3303300c, 0x00c00c00, 0xc0030030, 0x03003000,
    0xf00f0000,
    // [582]: 7A 'z'
    0x000a0e0a, 0xfffffc03, 0x00300c00, 0xc0300300, 0xc00c03ff, 0xfff00000,
    // [588]: 7B '{'
    0x00061604, 0xc3030c30, 0xc30c30c0, 0xc330c30c, 0x30c30cc3, 0x00000000,
    // [594]: 7C '|'
    0x00021604, 0xffffffff, 0xfff00000,
    // [597]: 7D '}'
    0x00061604, 0x0c330c30, 0xc30c30cc, 0x3030c30c, 0x30c30c0c, 0x30000000,
    // [603]: 7E '~'
    0x000c0406, 0xc3cc3c3c, 0x33c30000,
    // [606]: A0 ' '
    0x0004020e, 0x00000000,
    // [608]: A1 '¡'
    0x00021206, 0xf0ffffff, 0xf0000000,
    // [611]: A2 '¢'
    0x000a1208, 0x0c0303f0, 0xfcccf330, 0xcc330cc3, 0x30cc33cc, 0xf333f0fc, 0x0c030000,
    // [618]: A3 '£'
    0x000c1206, 0x3c03c003, 0x00300300, 0x30030030, 0x0ff0ff00, 0xc00c00c0, 0x0c00c00c, 0xffffff00,
    // [626]: A4 '¤'
    0x000c1006, 0x30c30c0f, 0x00f030c3, 0x0cc03c03, 0xc03c0330, 0xc30c0f00, 0xf030c30c,
    // [633]: A5 '¥'
    0x000a1206, 0xc0f03330, 0xccfffff0, 0xc030ffff, 0xf0c0300c, 0x0300c030, 0x0c030000,
    // [640]: A6 '¦'
    0x00021604, 0xfffff0ff, 0xfff00000,
    // [643]: A7 '§'
    0x000c1606, 0x3fc3fcc0, 0x3c030030, 0x0303c03c, 0x3c33c3c0, 0x3c03c3cc, 0x3c3c03c0, 0xc00c00c0,
    0x3c033fc3, 0xfc000000,
    // [653]: A8 '¨'
    0x00060206, 0xcf300000,
    // [655]: A9 '©'
    0x00101008, 0x0ff00ff0, 0x300c300c, 0xc3c3c3c3, 0xc033c033, 0xc033c033, 0xc3c3c3c3, 0x300c300c,
    0x0ff00ff0,
    // [664]: AA 'ª'
    0x00080e06, 0x3c3cc0c0, 0xfcfcc3c3, 0xfcfc0000, 0xffff0000,
    // [669]: AB '«'
    0x000c0a0e, 0xc30c3030, 0xc30c0c30, 0xc330c30c, 0xc30c3000,
    // [674]: AC '¬'
    0x000a060e, 0xfffffc03, 0x00c03000,
    // [677]: AD '­'
    0x000a020e, 0xfffff000,
    // [679]: AE '®'
    0x00101008, 0x0ff00ff0, 0x300c300c, 0xc3f3c3f3, 0xcc33cc33, 0xc3f3c3f3, 0xcc33cc33, 0x300c300c,
    0x0ff00ff0,
    // [688]: AF '¯'
    0x00060206, 0xfff00000,
    // [690]: B0 '°'
    0x00080806, 0x3c3cc3c3, 0xc3c33c3c,
    // [693]: B1 '±'
    0x000a0e0a, 0x0c0300c0, 0x30fffff0, 0xc0300c03, 0x000000ff, 0xfff00000,
    // [699]: B2 '²'
    0x00060a02, 0xfffc30ff, 0xf0c3fff0,
    // [702]: B3 '³'
    0x00060a02, 0xfffc30ff, 0xfc30fff0,
    // [705]: B4 '´'
    0x00040406, 0xcc330000,
    // [707]: B5 'µ'
    0x0010120a, 0xc0c0c0c0, 0xc0c0c0c0, 0x30303030, 0x30303030, 0x30303030, 0x0c0c0c0c, 0x33fc33fc,
    0x00030003, 0x00030003,
    // [717]: B6 '¶'
    0x000e1606, 0xfff3ffc3, 0x3fccff33, 0xfccff33f, 0xccff33f0, 0xcfc3300c, 0xc03300cc, 0x03300cc0,
    0x3300cc03, 0x300cc033, 0x00cc0000,
    // [728]: B7 '·'
    0x00020210, 0xf0000000,
    // [730]: B8 '¸'
    0x00040618, 0x33ccff00,
    // [732]: B9 '¹'
    0x00060a02, 0x30c3cf30, 0xc30cfff0,
    // [735]: BA 'º'
    0x00080e06, 0x3c3cc3c3, 0xc3c3c3c3, 0x3c3c0000, 0xffff0000,
    // [740]: BB '»'
    0x000c0a0e, 0x0c30c330, 0xc30cc30c, 0x3030c30c, 0x0c30c300,
    // [745]: BC '¼'
    0x00121402, 0x0c030300, 0xc0c03c30, 0x0f030300, 0xc0c03030, 0x0c0c00cf, 0xc033fccc, 0x033300cc,
    0x30330c0f, 0xc303f0c0, 0xc00c3003, 0x0c00c300, 0x30000000,
    // [758]: BD '½'
    0x00121402, 0x0c030300, 0xc0c03c30, 0x0f030300, 0xc0c03030, 0x0c0c00cf, 0xc033ffcc, 0x03f300c0,
    0x30300c0f, 0xc303f0c0, 0x0c0c0303, 0x0fc0c3f0, 0x30000000,
    // [771]: BE '¾'
    0x00121402, 0x0c0fc303, 0xf0c0c030, 0x30030fc0, 0xc3f030c0, 0x0c3000cf, 0xc033fccc, 0x033300cc,
    0x30330c0f, 0xc303f0c0, 0xc00c3003, 0x0c00c300, 0x30000000,
    // [784]: BF '¿'
    0x000a1206, 0x0c030000, 0x000c0300, 0xc0300300, 0xc00c03c0, 0xf03c0f03, 0x3f0fc000,
    // [791]: C0 'À'
    0x000e1800, 0x00c00300, 0x3000c000, 0x00000030, 0x00c00300, 0x0c00cc03, 0x300cc033, 0x03030c0c,
    0x3030c0cf, 0xffffffc0, 0x0f003c00, 0xf0030000,
    // [803]: C1 'Á'
    0x000e1800, 0x0c003000, 0x3000c000, 0x00000030, 0x00c00300, 0x0c00cc03, 0x300cc033, 0x03030c0c,
    0x3030c0cf, 0xffffffc0, 0x0f003c00, 0xf0030000,
    // [815]: C2 'Â'
    0x000e1800, 0x03000c00, 0xcc033000, 0x00000030, 0x00c00300, 0x0c00cc03, 0x300cc033, 0x03030c0c,
    0x3030c0cf, 0xffffffc0, 0x0f003c00, 0xf0030000,
    // [827]: C3 'Ã'
    0x000e1800, 0xc3c30f03, 0xc30f0c00, 0x00000030, 0x00c00300, 0x0c00cc03, 0x300cc033, 0x03030c0c,
    0x3030c0cf, 0xffffffc0, 0x0f003c00, 0xf0030000,
    // [839]: C4 'Ä'
    0x000e1602, 0x0cc03300, 0x00000003, 0x000c0030, 0x00c00cc0, 0x3300cc03, 0x303030c0, 0xc3030c0c,
    0xfffffffc, 0x00f003c0, 0x0f003000,
    // [850]: C5 'Å'
    0x000e1602, 0x03000c00, 0xcc033003, 0x000c0030, 0x00c00cc0, 0x3300cc03, 0x303030c0, 0xc3030c0c,
    0xfffffffc, 0x00f003c0, 0x0f003000,
    // [861]: C6 'Æ'
    0x00141206, 0xfff00fff, 0x0000cc00, 0x0cc000cc, 0x000cc000, 0xc3000c30, 0x3fff03ff, 0xf000c0c0,
    0x0c0c00c0, 0xc00c0c00, 0xc0300c03, 0xffc03ffc, 0x03000000,
    // [874]: C7 'Ç'
    0x000c1806, 0x3fc3fcc0, 0x3c030030, 0x03003003, 0x00300300, 0x30030030, 0x03c03c03, 0x3fc3fc0c,
    0x00c03003, 0x000c00c0,
    // [884]: C8 'È'
    0x000a1800, 0x0300c0c0, 0x3000000f, 0xffff00c0, 0x300c0300, 0xc033fcff, 0x00c0300c, 0x0300c03f,
    0xffff0000,
    // [893]: C9 'É'
    0x000a1800, 0x300c00c0, 0x3000000f, 0xffff00c0, 0x300c0300, 0xc033fcff, 0x00c0300c, 0x0300c03f,
    0xffff0000,
    // [902]: CA 'Ê'
    0x000a1800, 0x0c030330, 0xcc00000f, 0xffff00c0, 0x300c0300, 0xc033fcff, 0x00c0300c, 0x0300c03f,
    0xffff0000,
    // [911]: CB 'Ë'
    0x000a1602, 0x330cc000, 0x00fffff0, 0x0c0300c0, 0x300c033f, 0xcff00c03, 0x00c0300c, 0x03fffff0,
    // [919]: CC 'Ì'
    0x00041800, 0x33cc00cc, 0xcccccccc, 0xcccccccc,
    // [923]: CD 'Í'
    0x00041800, 0xcc330033, 0x33333333, 0x33333333,
    // [927]: CE 'Î'
    0x00061800, 0x30ccf300, 0x030c30c3, 0x0c30c30c, 0x30c30c30, 0xc30c0000,
    // [933]: CF 'Ï'
    0x00061602, 0xcf300030, 0xc30c30c3, 0x0c30c30c, 0x30c30c30, 0xc0000000,
    // [939]: D0 'Ð'
    0x000e1206, 0x0ff03fc3, 0x030c0cc0, 0x3300cc03, 0x300cc0ff, 0x03fc0330, 0x0cc03300, 0xc3030c0c,
    0x0ff03fc0,
    // [948]: D1 'Ñ'
    0x000c1800, 0xc3cc3c3c, 0x33c30000, 0x00c0fc0f, 0xc0fc0fc3, 0x3c33c33c, 0x33cc3cc3, 0xcc3cc3f0,
    0x3f03f03f, 0x03c03c03,
    // [958]: D2 'Ò'
    0x000c1800, 0x0300300c, 0x00c00000, 0x003fc3fc, 0xc03c03c0, 0x3c03c03c, 0x03c03c03, 0xc03c03c0,
    0x3c03c03c, 0x033fc3fc,
    // [968]: D3 'Ó'
    0x000c1800, 0x0c00c003, 0x00300000, 0x003fc3fc, 0xc03c03c0, 0x3c03c03c, 0x03c03c03, 0xc03c03c0,
    0x3c03c03c, 0x033fc3fc,
    // [978]: D4 'Ô'
    0x000c1800, 0x0300300c, 0xc0cc0000, 0x003fc3fc, 0xc03c03c0, 0x3c03c03c, 0x03c03c03, 0xc03c03c0,
    0x3c03c03c, 0x033fc3fc,
    // [988]: D5 'Õ'
    0x000c1800, 0xc3cc3c3c, 0x33c30000, 0x003fc3fc, 0xc03c03c0, 0x3c03c03c, 0x03c03c03, 0xc03c03c0,
    0x3c03c03c, 0x033fc3fc,
    // [998]: D6 'Ö'
    0x000c1602, 0x30c30c00, 0x00003fc3, 0xfcc03c03, 0xc03c03c0, 0x3c03c03c, 0x03c03c03, 0xc03c03c0,
    0x3c033fc3, 0xfc000000,
    // [1008]: D7 '×'
    0x000a0a0e, 0xc0f03330, 0xcc0c0303, 0x30ccc0f0, 0x30000000,
    // [1013]: D8 'Ø'
    0x00101206, 0xcff0cff0, 0x300c300c, 0x3c0c3c0c, 0x330c330c, 0x30cc30cc, 0x303c303c, 0x300c300c,
    0x300f300f, 0x0ff00ff0,
    // [1023]: D9 'Ù'
    0x000c1800, 0x0300300c, 0x00c00000, 0x00c03c03, 0xc03c03c0, 0x3c03c03c, 0x03c03c03, 0xc03c03c0,
    0x3c03c03c, 0x033fc3fc,
    // [1033]: DA 'Ú'
    0x000c1800, 0x0c00c003, 0x00300000, 0x00c03c03, 0xc03c03c0, 0x3c03c03c, 0x03c03c03, 0xc03c03c0,
    0x3c03c03c, 0x033fc3fc,
    // [1043]: DB 'Û'
    0x000c1800, 0x0300300c, 0xc0cc0000, 0x00c03c03, 0xc03c03c0, 0x3c03c03c, 0x03c03c03, 0xc03c03c0,
    0x3c03c03c, 0x033fc3fc,
    // [1053]: DC 'Ü'
    0x000c1602, 0x30c30c00, 0x0000c03c, 0x03c03c03, 0xc03c03c0, 0x3c03c03c, 0x03c03c03, 0xc03c03c0,
    0x3c033fc3, 0xfc000000,
    // [1063]: DD 'Ý'
    0x000a1800, 0x18060060, 0x1800000c, 0x0f03c0f0, 0x3c0f0333, 0x0cc330cc, 0x0c0300c0, 0x300c0300,
    0xc0300000,
    // [1072]: DE 'Þ'
    0x000a1206, 0x00c0300c, 0x033fcffc, 0x0f03c0f0, 0x3c0f033f, 0xcff00c03, 0x00c03000,
    // [1079]: DF 'ß'
    0x000c1206, 0x0fc0fc30, 0x33030c30, 0xc3033033, 0x0330333c, 0x33c3c03c, 0x03c03c03, 0x3f33f300,
    // [1087]: E0 'à'
    0x000a1404, 0x0c030300, 0xc0000003, 0xf0fcc0f0, 0x3ff3fcc0, 0xf03c0f03, 0xc0f03ff3, 0xfc000000,
    // [1095]: E1 'á'
    0x000a1404, 0x0c030030, 0x0c000003, 0xf0fcc0f0, 0x3ff3fcc0, 0xf03c0f03, 0xc0f03ff3, 0xfc000000,
    // [1103]: E2 'â'
    0x000a1404, 0x0c030330, 0xcc000003, 0xf0fcc0f0, 0x3ff3fcc0, 0xf03c0f03, 0xc0f03ff3, 0xfc000000,
    // [1111]: E3 'ã'
    0x000a1404, 0xcf33c3cc, 0xf3000003, 0xf0fcc0f0, 0x3ff3fcc0, 0xf03c0f03, 0xc0f03ff3, 0xfc000000,
    // [1119]: E4 'ä'
    0x000a1206, 0x330cc000, 0x003f0fcc, 0x0f03ff3f, 0xcc0f03c0, 0xf03c0f03, 0xff3fc000,
    // [1126]: E5 'å'
    0x000a1800, 0x0f03c30c, 0xc330cc30, 0xf03c0000, 0x03f0fcc0, 0xf03ff3fc, 0xc0f03c0f, 0x03c0f03f,
    0xf3fc0000,
    // [1135]: E6 'æ'
    0x00120e0a, 0x3f3f0fcf, 0xcc0c0f03, 0x03c0c030, 0x300ffff3, 0xfffc00c0, 0xc0303c0c, 0x0f03033f,
    0x3f0fcfc0,
    // [1144]: E7 'ç'
    0x000a120a, 0x3f0fcc0f, 0x0300c030, 0x0c0300c0, 0x3c0f033f, 0x0fc0c030, 0x0300c000,
    // [1151]: E8 'è'
    0x000a1404, 0x0300c0c0, 0x30000003, 0xf0fcc0f0, 0x3c0f03ff, 0xfff00c03, 0xc0f033f0, 0xfc000000,
    // [1159]: E9 'é'
    0x000a1404, 0x0c030030, 0x0c000003, 0xf0fcc0f0, 0x3c0f03ff, 0xfff00c03, 0xc0f033f0, 0xfc000000,
    // [1167]: EA 'ê'
    0x000a1404, 0x0c030330, 0xcc000003, 0xf0fcc0f0, 0x3c0f03ff, 0xfff00c03, 0xc0f033f0, 0xfc000000,
    // [1175]: EB 'ë'
    0x000a1206, 0x330cc000, 0x003f0fcc, 0x0f03c0f0, 0x3fffff00, 0xc03c0f03, 0x3f0fc000,
    // [1182]: EC 'ì'
    0x00041404, 0x33cc00cc, 0xcccccccc, 0xcccc0000,
    // [1186]: ED 'í'
    0x00041404, 0xcc330033, 0x33333333, 0x33330000,
    // [1190]: EE 'î'
    0x00061404, 0x30ccf300, 0x030c30c3, 0x0c30c30c, 0x30c30c00,
    // [1195]: EF 'ï'
    0x00061206, 0xcf300030, 0xc30c30c3, 0x0c30c30c, 0x30c00000,
    // [1200]: F0 'ð'
    0x000a1305, 0x00c0f330, 0xfc0fcf33, 0xc30cc330, 0x3c0f03c0, 0xf03c0cc3, 0x30c3c0f0,
    // [1207]: F1 'ñ'
    0x000a1404, 0xcf33c3cc, 0xf3000003, 0xccf3c3f0, 0xfc0f03c0, 0xf03c0f03, 0xc0f03c0f, 0x03000000,
    // [1215]: F2 'ò'
    0x000a1404, 0x0300c0c0, 0x30000003, 0xf0fcc0f0, 0x3c0f03c0, 0xf03c0f03, 0xc0f033f0, 0xfc000000,
    // [1223]: F3 'ó'
    0x000a1404, 0x300c00c0, 0x30000003, 0xf0fcc0f0, 0x3c0f03c0, 0xf03c0f03, 0xc0f033f0, 0xfc000000,
    // [1231]: F4 'ô'
    0x000a1404, 0x0c030330, 0xcc000003, 0xf0fcc0f0, 0x3c0f03c0, 0xf03c0f03, 0xc0f033f0, 0xfc000000,
    // [1239]: F5 'õ'
    0x000a1404, 0xcf33c3cc, 0xf3000003, 0xf0fcc0f0, 0x3c0f03c0, 0xf03c0f03, 0xc0f033f0, 0xfc000000,
    // [1247]: F6 'ö'
    0x000a1206, 0x330cc000, 0x003f0fcc, 0x0f03c0f0, 0x3c0f03c0, 0xf03c0f03, 0x3f0fc000,
    // [1254]: F7 '÷'
    0x000a0a0a, 0x0c030000, 0x00fffff0, 0x00000c03, 0x00000000,
    // [1259]: F8 'ø'
    0x000e0e0a, 0xcfc33f03, 0x030c0c3c, 0x30f0c333, 0x0ccc30f0, 0xc3c3030c, 0x0c0fcc3f, 0x30000000,
    // [1267]: F9 'ù'
    0x000a1404, 0x0300c0c0, 0x3000000c, 0x0f03c0f0, 0x3c0f03c0, 0xf03c0f03, 0xf0fc3cf3, 0x3c000000,
    // [1275]: FA 'ú'
    0x000a1404, 0x300c00c0, 0x3000000c, 0x0f03c0f0, 0x3c0f03c0, 0xf03c0f03, 0xf0fc3cf3, 0x3c000000,
    // [1283]: FB 'û'
    0x000a1404, 0x0c030330, 0xcc00000c, 0x0f03c0f0, 0x3c0f03c0, 0xf03c0f03, 0xf0fc3cf3, 0x3c000000,
    // [1291]: FC 'ü'
    0x000a1206, 0x330cc000, 0x00c0f03c, 0x0f03c0f0, 0x3c0f03c0, 0xf03f0fc3, 0xcf33c000,
    // [1298]: FD 'ý'
    0x000c1a04, 0x0c00c003, 0x00300000, 0x00c0cc0c, 0xc0cc0cc0, 0xcc0c3303, 0x30330330, 0x0c00c00c,
    0x00c00300, 0x30030030, 0x00f00f00,
    // [1309]: FE 'þ'
    0x000a1606, 0x00c0300c, 0x033fcffc, 0x0f03c0f0, 0x3c0f03c0, 0xf03c0f03, 0x3fcff00c, 0x0300c030,
    // [1317]: FF 'ÿ'
    0x000c1806, 0x33033000, 0x0000c0cc, 0x0cc0cc0c, 0xc0cc0c33, 0x03303303, 0x300c00c0, 0x0c00c003,
    0x00300300, 0x3000f00f,
    // [1327]: 152 'Œ'
    0x00141206, 0xff3fcff3, 0xfc00c030, 0x0c0300c0, 0x300c0300, 0xc0300c03, 0x3fc033fc, 0x0300c030,
    0x0c0300c0, 0x300c0300, 0xc0300c03, 0xff3fcff3, 0xfc000000,
    // [1340]: 153 'œ'
    0x00120e0a, 0x3f3f0fcf, 0xcc0c0f03, 0x03c0c0f0, 0x303ffc0f, 0xff0300c0, 0xc0303c0c, 0x0f03033f,
    0x3f0fcfc0,
    // [1349]: 2018 '‘'
    0x00040604, 0xcc333300,
    // [1351]: 2019 '’'
    0x00040604, 0xcccc3300,
    // [1353]: 201A '‚'
    0x00040616, 0xcccc3300,
    // [1355]: 201B '‛'
    0x00040604, 0x3333cc00,
    // [1357]: 201C '“'
    0x000a0604, 0xc330c30c, 0xc330cc30,
    // [1360]: 201D '”'
    0x000a0604, 0xc330cc33, 0x0c30cc30,
    // [1363]: 201E '„'
    0x000a0616, 0xc330cc33, 0x0c30cc30,
    // [1366]: 201F '‟'
    0x000a0604, 0x30cc330c, 0xc3c330c0,
    // [1369]: 2020 '†'
    0x000a1206, 0x0c0300c0, 0x30fffff0, 0xc0300c03, 0x00c0300c, 0x0300c030, 0x0c030000,
    // [1376]: 2021 '‡'
    0x000a1606, 0x0c0300c0, 0x30fffff0, 0xc0300c03, 0x00c0300c, 0x0300c030, 0xfffff0c0, 0x300c0300,
    // [1384]: 2022 '•'
    0x000c0c0a, 0x3fc3fcff, 0xffffffff, 0xffffffff, 0xffffff3f, 0xc3fc0000,
    // [1390]: 20AC '€'
    0x00101008, 0x3fc03fc0, 0xc030c030, 0x000c000c, 0x0fff0fff, 0x000c000c, 0x0fff0fff, 0xc030c030,
    0x3fc03fc0,
    // [1399]: E700 Battery_05
    0x00180c0c, 0x3ffffe40, 0x00014000, 0x0dc0000d, 0xc0000dc0, 0x000dc000, 0x0dc0000d, 0xc0000d40,
    0x000d4000, 0x013ffffe,
    // [1409]: E701 Battery_25
    0x00180c0c, 0x3ffffe40, 0x00014000, 0x7dc0007d, 0xc0007dc0, 0x007dc000, 0x7dc0007d, 0xc0007d40,
    0x007d4000, 0x013ffffe,
    // [1419]: E702 Battery_50
    0x00180c0c, 0x3ffffe40, 0x0001400f, 0xfdc00ffd, 0xc00ffdc0, 0x0ffdc00f, 0xfdc00ffd, 0xc00ffd40,
    0x0ffd4000, 0x013ffffe,
    // [1429]: E703 Battery_75
    0x00180c0c, 0x3ffffe40, 0x000141ff, 0xfdc1fffd, 0xc1fffdc1, 0xfffdc1ff, 0xfdc1fffd, 0xc1fffd41,
    0xfffd4000, 0x013ffffe,
    // [1439]: E704 Battery_99
    0x00180c0c, 0x3ffffe40, 0x00015fff, 0xfddffffd, 0xdffffddf, 0xfffddfff, 0xfddffffd, 0xdffffd5f,
    0xfffd4000, 0x013ffffe,
    // [1449]: E705 Radio_3
    0x00151107, 0x00f8001f, 0xf003e3e0, 0x3c078380, 0x0e387c3b, 0x8ff8e8f1, 0xe20e0380, 0xe10e023e,
    0x2003f800, 0x38e00082, 0x00008000, 0x0e000020, 0x00000000,
    // [1462]: E706 Radio_2
    0x00151107, 0x00000000, 0x00000000, 0x00000000, 0x00007c00, 0x0ff800f1, 0xe00e0380, 0xe10e023e,
    0x2003f800, 0x38e00082, 0x00008000, 0x0e000020, 0x00000000,
    // [1475]: E707 Radio_1
    0x00151107, 0x00000000, 0x00000000, 0x00000000, 0x00000000, 0x00000000, 0x00000000, 0x0100003e,
    0x0003f800, 0x38e00082, 0x00008000, 0x0e000020, 0x00000000,
    // [1488]: E708 Radio_0
    0x00151107, 0x00000000, 0x00000000, 0x00000000, 0x00000000, 0x00000000, 0x00000000, 0x00000000,
    0x00000000, 0x00000000, 0x00008000, 0x0e000020, 0x00000000,
    // [1501]: E709 Radio_Off
    0x00151107, 0x00f80018, 0x30030060, 0x20008200, 0x0220000a, 0x00002800, 0x02200020, 0x80020200,
    0x20080200, 0x20200082, 0x00022000, 0x0a000020, 0x00000000,
    // [1514]: E70A Shift_Arrow
    0x000a1406, 0x0c0783f1, 0xfefffff0, 0xc0300c03, 0x00c0300c, 0x0300c030, 0x0c0300c0, 0x30000000,
    // [1522]: E70B Backspace_Symbol
    0x001a1206, 0xffffc03f, 0xfff80c00, 0x07030000, 0xe0c6061c, 0x31c3838c, 0x39c07307, 0xe00ec0f0,
    0x01f03c00, 0x7c1f803b, 0x0e701cc7, 0x0e0e3181, 0x870c0003, 0x830001c0, 0xffffe03f, 0xfff00000,
    // [1538]: E70C Enter_Symbol
    0x00180e08, 0xc00000c0, 0x0000c000, 0x00c00000, 0xc00030c0, 0x0038c000, 0x3cc0003e, 0xffffffff,
    0xffff0000, 0x3e00003c, 0x00003800, 0x00300000,
    // [1550]: FFFD '�'
    0x00121404, 0x00c00030, 0x003f000f, 0xc00f3c03, 0xcf03ccf0, 0xf33cfcff, 0xff3ffff3, 0xfffcff3f,
    0xff0fffc0, 0xf3c03cf0, 0x03f000fc, 0x000c0003, 0x00000000,
];

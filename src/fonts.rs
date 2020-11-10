/// This code includes an adaptation of the the murmur3 hash algorithm.
/// The murmur3 public domain notice, as retrieved on August 3, 2020 from
/// https://github.com/aappleby/smhasher/blob/master/src/MurmurHash3.cpp, states:
///
/// > MurmurHash3 was written by Austin Appleby, and is placed in the public
/// > domain. The author hereby disclaims copyright to this source code.
///
pub mod bold;
pub mod regular;
pub mod small;

use core::fmt;

/// Strings with Unicode Private Use Area characters for UI Sprites
pub mod pua {
    pub const BATTERY_05: &str = &"\u{E700}";
    pub const BATTERY_25: &str = &"\u{E701}";
    pub const BATTERY_50: &str = &"\u{E702}";
    pub const BATTERY_75: &str = &"\u{E703}";
    pub const BATTERY_99: &str = &"\u{E704}";
    pub const RADIO_3: &str = &"\u{E705}";
    pub const RADIO_2: &str = &"\u{E706}";
    pub const RADIO_1: &str = &"\u{E707}";
    pub const RADIO_0: &str = &"\u{E708}";
    pub const RADIO_OFF: &str = &"\u{E709}";
    pub const SHIFT_ARROW: &str = &"\u{E70A}";
    pub const BACKSPACE_SYMBOL: &str = &"\u{E70B}";
    pub const ENTER_SYMBOL: &str = &"\u{E70C}";
}

/// Holds header data for a font glyph
pub struct GlyphHeader {
    pub w: usize,
    pub h: usize,
    pub y_offset: usize,
}
impl GlyphHeader {
    /// Unpack glyph header of format: (w:u8)<<16 | (h:u8)<<8 | yOffset:u8
    pub fn new(header: u32) -> GlyphHeader {
        let w = ((header << 8) >> 24) as usize;
        let h = ((header << 16) >> 24) as usize;
        let y_offset = (header & 0x000000ff) as usize;
        GlyphHeader { w, h, y_offset }
    }
}

/// Available typeface glyph sets
pub enum GlyphSet {
    Bold,
    Regular,
    Small,
}

/// Error type for when a font has no glyph to match a grapheme cluster query
#[derive(Debug, Clone)]
pub struct GlyphNotFound;
impl fmt::Display for GlyphNotFound {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
        write!(f, "Font has no glyph for requested grapheme cluster")
    }
}

/// Abstraction for working with typeface glyph sets
#[derive(Copy, Clone)]
pub struct Font {
    pub glyph_pattern_offset: GlyphPatternOffsetFnPtr,
    pub glyph_data: GlyphDataFnPtr,
}
pub type GlyphPatternOffsetFnPtr = fn(&str) -> Result<usize, GlyphNotFound>;
pub type GlyphDataFnPtr = fn(usize) -> u32;
impl Font {
    pub fn new(gs: GlyphSet) -> Font {
        match gs {
            GlyphSet::Bold => Font {
                glyph_pattern_offset: bold::get_blit_pattern_offset,
                glyph_data: bold_data,
            },
            GlyphSet::Regular => Font {
                glyph_pattern_offset: regular::get_blit_pattern_offset,
                glyph_data: regular_data,
            },
            GlyphSet::Small => Font {
                glyph_pattern_offset: small::get_blit_pattern_offset,
                glyph_data: small_data,
            },
        }
    }
}

/// Get word of packed glyph data for bold
pub fn bold_data(index: usize) -> u32 {
    bold::DATA[index]
}

/// Get word of packed glyph data for regular
pub fn regular_data(index: usize) -> u32 {
    regular::DATA[index]
}

/// Get word of packed glyph data for small
pub fn small_data(index: usize) -> u32 {
    small::DATA[index]
}

/// Return Murmur3 hash function of a string using each char as a u32 block.
/// Fonts use this to hash grapheme clusters from &str into u32.
pub fn murmur3(key: &str, seed: u32) -> u32 {
    let mut h = seed;
    let mut k;
    // Hash each character as its own u32 block
    for c in key.chars() {
        k = c as u32;
        k = k.wrapping_mul(0xcc9e2d51);
        k = k.rotate_left(15);
        k = k.wrapping_mul(0x1b873593);
        h ^= k;
        h = h.rotate_left(13);
        h = h.wrapping_mul(5);
        h = h.wrapping_add(0xe6546b64);
    }
    h ^= key.bytes().count() as u32;
    // Finalize with avalanche
    h ^= h >> 16;
    h = h.wrapping_mul(0x85ebca6b);
    h ^= h >> 13;
    h = h.wrapping_mul(0xc2b2ae35);
    h ^ (h >> 16)
}

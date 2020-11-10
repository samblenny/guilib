#![allow(dead_code)]
//! Emoji Font
//
// This code includes modified versions of graphics from the twemoji project.
// The modified emoji PNG files were converted from color PNG format to monochrome
// PNG with dithered grayscale shading.
//
// - Twemoji License Notice
//   > Copyright 2019 Twitter, Inc and other contributors
//   >
//   > Code licensed under the MIT License: http://opensource.org/licenses/MIT
//   >
//   > Graphics licensed under CC-BY 4.0: https://creativecommons.org/licenses/by/4.0/
//
// - Twemoji Source Code Link:
//   https://github.com/twitter/twemoji
//

/// Maximum height of glyph patterns in this bitmap typeface.
/// This will be true: h + y_offset <= MAX_HEIGHT
pub const MAX_HEIGHT: u8 = 48;

/* TODO: Emoji data */

// Copyright (c) 2020 Sam Blenny
// SPDX-License-Identifier: Apache-2.0 OR MIT
//
package main

import (
	"bytes"
	"fmt"
	"guilib/codegen/font"
	"image"
	"image/png"
	"io/ioutil"
	"math/bits"
	"os"
	"path"
	"sort"
	"strings"
	"text/template"
)

// Command line switch to confirm intent of writing output files
const confirm = "--replace-font-files"

// Main: check for confirmation switch before writing files
func main() {
	if len(os.Args) == 2 && os.Args[1] == confirm {
		codegen()
	} else {
		usage()
	}
}

// Change this to control the visibility of debug messages
const enableDebug = false

// Path for output files with generated font code
const outPath = "../src/fonts"

// Index and alias files for grapheme clusters that go with img/emoji48x48_o3x3.png
const emojiIndexPath = "img/emoji_13_0_index.txt"
const emojiAliasPath = "img/emoji_13_0_variation_sequence_aliases.txt"

// Spec for how to generate font source code files from glyph grid sprite sheets
func fonts() []font.FontSpec {
	return []font.FontSpec{
		font.FontSpec{"Emoji", "img/emoji48x48_o3x0.png", 48, 16, 0, 0, twemoji, "emoji.rs"},
		font.FontSpec{"Bold", "img/bold.png", 30, 16, 2, 2, chicago, "bold.rs"},
		font.FontSpec{"Regular", "img/regular.png", 30, 16, 2, 2, geneva, "regular.rs"},
	}
}

// Generate rust source code files for fonts
func codegen() {
	t := template.Must(template.New("usage").Parse(fontFileTemplate))
	for _, f := range fonts() {
		var data string
		switch f.Name {
		case "Emoji":
			data = emojiData(f)
		case "Bold":
			data = sysLatinData(f)
		case "Regular":
			data = sysLatinData(f)
		default:
			panic("unexpected FontSpec.Name")
		}
		context := struct {
			Font    font.FontSpec
			OutPath string
			Data    string
		}{f, outPath, data}
		var buf bytes.Buffer
		err := t.Execute(&buf, context)
		if err != nil {
			panic(err)
		}
		// Write the file
		op := path.Join(outPath, f.RustOut)
		fmt.Println("Writing to", op)
		ioutil.WriteFile(op, buf.Bytes(), 0644)
	}
}

// Generate rust code for emoji glyph blit pattern and grapheme cluster index data
func emojiData(font font.FontSpec) string {
	data := "/* TODO: Emoji data */"
	return data
}

// Generate rust code for sysLatin glyph blit pattern and grapheme cluster index data
func sysLatinData(fs font.FontSpec) string {
	// Find all the glyphs and pack them into a list of blit pattern objects
	pl := patternListFromSpriteSheet(fs, font.SysLatinMap())
	// Make rust code for the blit pattern DATA array, plus an index list
	rb := rustyBlitsFromPatternList(pl)
	// Render data template
	dataT := template.Must(template.New("dataBuf").Parse(dataTemplate))
	dataContext := struct{ RB RustyBlits }{rb}
	var dataStrBuf bytes.Buffer
	err := dataT.Execute(&dataStrBuf, dataContext)
	if err != nil {
		panic(err)
	}
	return dataStrBuf.String()
}

// Extract glyph sprites from a PNG grid and pack them into a list of blit pattern objects
func patternListFromSpriteSheet(fs font.FontSpec, csList []font.CharSpec) []font.BlitPattern {
	// Read glyphs from png file
	img := readPNGFile(fs.Sprites)
	var patternList []font.BlitPattern
	for _, cs := range csList {
		blitPattern := font.ConvertGlyphToBlitPattern(img, fs, cs, enableDebug)
		patternList = append(patternList, blitPattern)
	}
	return patternList
}

// Make rust source code and an index list from a list of glyph blit patterns.
// When this finishes, rust source code for the `DATA: [u32; n] = [...];` array
// of concatenated blit patterns is in the return values's .Code. The length (n)
// of the `DATA: [u32; n]...` blit pattern array is in .DataLen, and the
// ClusterOffsetEntry{...} index entries are in .Index.
func rustyBlitsFromPatternList(pl []font.BlitPattern) RustyBlits {
	rb := RustyBlits{"", 0, ClusterOffsetIndex{}}
	for _, p := range pl {
		label := labelForCluster(p.CS.GraphemeCluster)
		comment := fmt.Sprintf("[%d]: %X %s", rb.DataLen, p.CS.FirstCodepoint, label)
		rb.Code += font.ConvertPatternToRust(p, comment)
		// Update the index with the correct offset (DATA[n]) for pattern header
		rb.Index = append(rb.Index, ClusterOffsetEntry{
			murmur3(p.CS.GraphemeCluster, 0),
			p.CS.GraphemeCluster,
			rb.DataLen,
		})
		rb.DataLen += len(p.Bytes)
	}
	// Sort the index
	sort.Slice(rb.Index, func(i, j int) bool { return rb.Index[i].M3Hash < rb.Index[j].M3Hash })
	return rb
}

// Read the specified PNG file and convert its data into an image object
func readPNGFile(name string) image.Image {
	pngFile, err := os.Open(name)
	if err != nil {
		panic("unable to open png file")
	}
	img, err := png.Decode(pngFile)
	if err != nil {
		panic("unable to decode png file")
	}
	pngFile.Close()
	return img
}

// Make label for grapheme cluster with special handling for UI sprites in PUA block
func labelForCluster(c string) string {
	switch c {
	case "\uE700":
		return "Battery_05"
	case "\uE701":
		return "Battery_25"
	case "\uE702":
		return "Battery_50"
	case "\uE703":
		return "Battery_75"
	case "\uE704":
		return "Battery_99"
	case "\uE705":
		return "Radio_3"
	case "\uE706":
		return "Radio_2"
	case "\uE707":
		return "Radio_1"
	case "\uE708":
		return "Radio_0"
	case "\uE709":
		return "Radio_Off"
	case "\uE70A":
		return "Shift_Arrow"
	case "\uE70B":
		return "Backspace_Symbol"
	case "\uE70C":
		return "Enter_Symbol"
	default:
		return fmt.Sprintf("'%s'", c)
	}
}

// Holds an index list and rust source code for a font's worth of blit patterns
type RustyBlits struct {
	Code    string
	DataLen int
	Index   ClusterOffsetIndex
}

// Type for the index so it can be used with .RustCodeFor* methods in templates
type ClusterOffsetIndex []ClusterOffsetEntry

// An index entry for translating from grapheme cluster to blit pattern
type ClusterOffsetEntry struct {
	M3Hash     uint32
	Cluster    string
	DataOffset int
}

// Return Murmur3 hash function of a string using each character as a u32 block
func murmur3(key string, seed uint32) uint32 {
	h := seed
	k := uint32(0)
	// Hash each codepoint in the string as its own uint32 block
	for _, c := range key {
		k = uint32(c)
		k *= 0xcc9e2d51
		k = bits.RotateLeft32(k, 15)
		k *= 0x1b873593
		h ^= k
		h = bits.RotateLeft32(h, 13)
		h *= 5
		h += 0xe6546b64
	}
	h ^= uint32(len(key))
	// Finalize with avalanche
	h ^= h >> 16
	h *= 0x85ebca6b
	h ^= h >> 13
	h *= 0xc2b2ae35
	return h ^ (h >> 16)
}

// Format the inner elements of a [u32; n] cluster hash index table
func (coIndex ClusterOffsetIndex) RustCodeForClusterHashes() string {
	var rustCode []string
	for _, entry := range coIndex {
		hash := fmt.Sprintf("0x%08X", entry.M3Hash)
		label := labelForCluster(entry.Cluster)
		rustCode = append(rustCode, fmt.Sprintf("%s, // %s", hash, label))
	}
	return strings.Join(rustCode, "\n    ")
}

// Format the inner elements of a [u32; n] blit pattern offset table
func (coIndex ClusterOffsetIndex) RustCodeForOffsets() string {
	var rustCode []string
	for _, entry := range coIndex {
		offset := fmt.Sprintf("%d,", entry.DataOffset)
		label := labelForCluster(entry.Cluster)
		rustCode = append(rustCode, fmt.Sprintf("%-5s // %s", offset, label))
	}
	return strings.Join(rustCode, "\n    ")
}

// Print usage message
func usage() {
	templateString := `
This tool generates fonts in the form of rust source code.
To confirm that you want to write the files, use the {{.Confirm}} switch.

Font files that will be generated:{{range $f := .Fonts}}
  {{$.OutPath}}/{{$f.RustOut}}{{end}}

Usage:
    go run main.go {{.Confirm}}
`
	t := template.Must(template.New("usage").Parse(templateString))
	context := struct {
		Confirm string
		OutPath string
		Fonts   []font.FontSpec
	}{confirm, outPath, fonts()}
	err := t.Execute(os.Stdout, context)
	if err != nil {
		panic(err)
	}
}

// Template with rust source code for a outer structure of a font file
const fontFileTemplate = `#![forbid(unsafe_code)]
#![allow(dead_code)]
//! {{.Font.Name}} Font
// DO NOT MAKE EDITS HERE because this file is automatically generated.
// To make changes, see guilib/codegen/main.go
//
// Copyright (c) 2020 Sam Blenny
// SPDX-License-Identifier: Apache-2.0 OR MIT
//
// NOTE: The copyright notice above applies to the rust source code in this
// file, but not to the bitmap graphics encoded in the DATA array.
//
// CREDITS:
{{.Font.Legal}}

/// Maximum height of glyph patterns in this bitmap typeface.
/// This will be true: h + y_offset <= MAX_HEIGHT
pub const MAX_HEIGHT: u8 = {{.Font.Size}};

{{.Data}}
`

// Template with rust source code for the data and index portion of a font file
const dataTemplate = `/// Return Okay(offset into DATA[]) for start of blit pattern for grapheme cluster.
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
const HASHED_CLUSTERS: [u32; {{len .RB.Index}}] = [
    {{.RB.Index.RustCodeForClusterHashes}}
];

// Lookup table from hashed cluster to blit pattern offset (sort order matches HASHED_CLUSTERS)
const PATTERN_OFFSETS: [usize; {{len .RB.Index}}] = [
    {{.RB.Index.RustCodeForOffsets}}
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
pub const DATA: [u32; {{.RB.DataLen}}] = [
{{.RB.Code}}];`

// Emoji graphics legal notice
const twemoji = `// This code includes encoded bitmaps with modified versions of graphics from
// the twemoji project. The modified emoji PNG files were converted from color
// PNG format to monochrome PNG with dithered grayscale shading.
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
//`

// Bold font legal notice
const chicago = `// This code includes encoded bitmaps of glyphs from the Chicago typeface which
// was designed by Susan Kare and released by Apple in 1984. Chicago is a
// registered trademark of Apple Inc.
//`

// Regular font legal notice
const geneva = `// This code includes encoded bitmaps of glyphs from the Geneva typeface which
// was designed by Susan Kare and released by Apple in 1984. Geneva is a
// registered trademark of Apple Inc.
//`

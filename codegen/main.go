package main

import (
	"bytes"
	"fmt"
	"guilib/codegen/font"
	"image/png"
	"io/ioutil"
	"math/bits"
	"os"
	"path"
	"sort"
	"strings"
	"text/template"
)

// Path for output files with generated font code
const outPath = "../src/fonts"

// Spec for how to generate font source code files from glyph grid sprite sheets
func fonts() []font.FontSpec {
	return []font.FontSpec{
		font.FontSpec{"Emoji", "img/emoji48x48_o3x0.png", 48, 16, 0, 0, twemoji, "emoji.rs"},
		font.FontSpec{"Bold", "img/bold.png", 30, 16, 2, 2, chicago, "bold.rs"},
		font.FontSpec{"Regular", "img/regular.png", 30, 16, 2, 2, geneva, "regular.rs"},
		font.FontSpec{"Small", "img/small.png", 24, 16, 2, 2, geneva, "small.rs"},
	}
}

// Generate rust source code files for fonts
func codegen() {
	templateString := strings.TrimLeft(`
#![allow(dead_code)]
//! {{.Font.Name}} Font
//
{{.Font.Legal}}
/// Maximum height of glyph patterns in this bitmap typeface.
/// This will be true: h + y_offset <= MAX_HEIGHT
pub const MAX_HEIGHT: u8 = {{.Font.Size}};

{{.Data}}
`, "\n")
	t := template.Must(template.New("usage").Parse(templateString))
	for _, f := range fonts() {
		var data string
		switch f.Name {
		case "Emoji":
			data = emojiData(f)
		case "Bold":
			data = sysLatinData(f, true)
		case "Regular":
			data = sysLatinData(f, true)
		case "Small":
			data = sysLatinData(f, false)
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
func sysLatinData(fs font.FontSpec, hasPUA bool) string {
	dataTemplate := strings.TrimSpace(`
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
        | 0x20A0..=0x20CF  // Currency Symbols{{if .PrivateUseArea}}
        | 0xE700..=0xE70C  // Subset of Private Use Area (UI Icons){{end}}
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
const HASHED_CLUSTERS: [u32; {{len .COIndex}}] = [
    {{.COIndex.RustCodeForClusterHashes}}
];

// Lookup table from hashed cluster to blit pattern offset (sort order matches HASHED_CLUSTERS)
const PATTERN_OFFSETS: [usize; {{len .COIndex}}] = [
    {{.COIndex.RustCodeForOffsets}}
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
pub const DATA: [u32; {{.DataBufLen}}] = [
{{.DataBufRust}}];
`)
	// Read glyphs from png file
	pngFile, err := os.Open(fs.Sprites)
	if err != nil {
		panic("unable to open png file")
	}
	img, err := png.Decode(pngFile)
	if err != nil {
		panic("unable to decode png file")
	}
	pngFile.Close()
	var dataBuf []uint32
	var dataBufRust string
	var coIndex []ClusterOffsetEntry
	puaCount := 0
	// Loop through the big list of []CharSpec structs that map from
	// grapheme clusters to row and column grid coordinates in the sprite
	// sheet of glyphs. For each grapheme cluster:
	// 1. Read the glyph pixels
	// 2. Pack them into the blit pattern data array
	// 3. Add an entry to the index
	for _, cs := range font.SysLatinMap() {
		// Identify the Unicode block for this grapheme cluster
		block := font.Block(cs.FirstCodepoint)
		if block == font.PRIVATE_USE_AREA && fs.Name != "Bold" && fs.Name != "Regular" {
			// Skip PUA block glyphs for fonts that don't include them
			continue
		} else if block == font.UNKNOWN {
            fmt.Printf("0x%08X\n", cs.FirstCodepoint)
			panic("character map includes cluster from unknown Unicode block")
		}
		// Find the glyph and pack it into a [u32] blit pattern for the data buffer
		matrix, yOffset := font.ConvertGlyphBoxToMatrix(img, fs, cs.Row, cs.Col)
		// Uncomment the two `fmt...` lines below if you want an ASCII art debug dump of
		// the blit patterns on stdout. The debug dump can help to find and fix data
		// entry problems in the []CharSpec character map.
		//
		// fmt.Printf("%X: %s\n", cs.Codepoint, cs.Chr)
		// fmt.Println(font.ConvertMatrixToText(matrix, yOffset))
		blitPattern := font.ConvertMatrixToPattern(matrix, yOffset)
		dataOffset := len(dataBuf)
		dataBuf = append(dataBuf, blitPattern...)
		// Subtle point: distinction between cluster vs. clusterLabel
		// has to do with the way ./font/charmap.go::SysLatinMap()
		// encodes UI sprites in the PRIVATE USE AREA block. Normally,
		// CharSpec.GraphemeCluster holds an actual grapheme cluster.
		// But, the CharSpec structs for UI sprites abuse the field to
		// store a descriptive label instead. The workaround below is
		// a way to avoid adding another field to CharSpec. The goal
		// is to annotate the big const [u32; n] arrays with readable
		// comments about what all the mysterious int literals mean.
		// TODO: maybe find a less confusing solution?
		cluster := cs.GraphemeCluster
		clusterLabel := cs.GraphemeCluster
		comment := fmt.Sprintf("[%d]: %X '%s'", dataOffset, cs.FirstCodepoint, cluster)
		if block == font.PRIVATE_USE_AREA {
			// For UI sprites, adjust the grapheme cluster string and comment because
			// cs.GraphemeCluster holds a descriptive label instead of the cluster string
			cluster = string(rune(cs.FirstCodepoint))
			comment = fmt.Sprintf("[%d]: %X %s", dataOffset, cs.FirstCodepoint, clusterLabel)
			puaCount += 1
		}
		rustCode := font.ConvertPatternToRust(blitPattern, comment)
		dataBufRust += rustCode
		// Update the index with the correct offset (DATA[n]) for pattern header
		seed := uint32(0)
		coIndex = append(coIndex, ClusterOffsetEntry{
			murmur3(cluster, seed),
			clusterLabel,
			dataOffset,
			block == font.PRIVATE_USE_AREA,
		})
	}
	// Sort the index
	sort.Slice(coIndex, func(i, j int) bool { return coIndex[i].M3Hash < coIndex[j].M3Hash })
	// Render data template
	dataT := template.Must(template.New("dataBuf").Parse(dataTemplate))
	dataContext := struct {
		PrivateUseArea bool
		DataBufRust    string
		DataBufLen     int
		COIndex        ClusterOffsetIndex
	}{puaCount > 0, dataBufRust, len(dataBuf), coIndex}
	var dataStrBuf bytes.Buffer
	err = dataT.Execute(&dataStrBuf, dataContext)
	if err != nil {
		panic(err)
	}
	return dataStrBuf.String()
}

// An index entry for translating from grapheme cluster to blit pattern
type ClusterOffsetEntry struct {
	M3Hash         uint32
	ClusterLabel   string
	DataOffset     int
	PrivateUseArea bool
}

// Type for the index so it can be used with .RustCodeFor* methods in templates
type ClusterOffsetIndex []ClusterOffsetEntry

// Format the inner elements of a [u32; n] cluster hash index table
func (coIndex ClusterOffsetIndex) RustCodeForClusterHashes() string {
	var rustCode []string
	for _, entry := range coIndex {
		hash := fmt.Sprintf("0x%08X", entry.M3Hash)
		clusterLabel := entry.ClusterLabel
		if !entry.PrivateUseArea {
			rustCode = append(rustCode, fmt.Sprintf("%s, // '%s'", hash, clusterLabel))
		} else {
			// PUA block chars have description strings, so format comment without quotes
			rustCode = append(rustCode, fmt.Sprintf("%s, // %s", hash, clusterLabel))
		}
	}
	return strings.Join(rustCode, "\n    ")
}

// Format the inner elements of a [u32; n] blit pattern offset table
func (coIndex ClusterOffsetIndex) RustCodeForOffsets() string {
	var rustCode []string
	for _, entry := range coIndex {
		offset := fmt.Sprintf("%d,", entry.DataOffset)
		clusterLabel := entry.ClusterLabel
		if !entry.PrivateUseArea {
			rustCode = append(rustCode, fmt.Sprintf("%-5s // '%s'", offset, clusterLabel))
		} else {
			// PUA block chars have description strings, so format comment without quotes
			rustCode = append(rustCode, fmt.Sprintf("%-5s // %s", offset, clusterLabel))
		}
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

// Main: check for confirmation switch before writing files
func main() {
	if len(os.Args) == 2 && os.Args[1] == confirm {
		codegen()
	} else {
		usage()
	}
}

// Command line switch to confirm intent of writing output files
const confirm = "--replace-font-files"

// Legal Notices
const twemoji = `// This code includes modified versions of graphics from the twemoji project.
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
`

const chicago = `// This code incorporates glyphs from the Chicago bitmap typeface which was
// designed by Susan Kare and released by Apple in 1984. Chicago is a registered
// trademark of Apple Inc.
//
`

const geneva = `// This code incorporates glyphs from the Geneva bitmap typeface which was
// designed by Susan Kare and released by Apple in 1984. Geneva is a registered
// trademark of Apple Inc.
//
`

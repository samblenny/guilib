package main

import (
	"bytes"
	"codegen/blocks"
	"fmt"
	"os"
	"strings"
	"text/template"
)

// Path for output files with generated font code
const outPath = "../src/fonts"

// Path for glyph grid sprite sheet input files
const inPath = "img"

// Spec for how to generate font source code files from glyph grid sprite sheets
func fonts() []FontSpec {
	return []FontSpec{
		FontSpec{"Emoji", "emoji48x48_o3x0.png", 48, 16, 0, 0, "emoji_index.txt", twemoji, "emoji.rs"},
		FontSpec{"Bold", "bold.png", 30, 16, 2, 2, "", chicago, "bold.rs"},
		FontSpec{"Regular", "regular.png", 30, 16, 2, 2, "", geneva, "regular.rs"},
		FontSpec{"Small", "small.png", 24, 16, 2, 2, "", geneva, "small.rs"},
	}
}

// Holds description of sprite sheet and character map for generating a font
type FontSpec struct {
	Name    string // Name of font
	Sprites string // Which file holds the sprite sheet image with the grid of glyphs?
	Size    int    // How many pixels on a side is each glyph (precondition: square glyphs)
	Cols    int    // How many glyphs wide is the grid?
	Gutter  int    // How many px between glyphs?
	Border  int    // How many px wide are top and left borders?
	Index   string // Which file holds index of unicode for sprites?
	Legal   string // What credits or license notices need to be included in font file comments?
	RustOut string // Where should the generated source code go?
}

// Generate rust source code files for fonts
func codegen() {
	blocks.DoNothing()
	templateString := strings.TrimLeft(`
#![allow(dead_code)]
//! {{.Font.Name}} Font
//
{{.Font.Legal}}

{{.CharMapIndex}}

/// Maximum height of glyph patterns in this bitmap typeface.
/// This will be true: h + yOffset <= MAX_HEIGHT
pub const MAX_HEIGHT: u8 = {{.Font.Size}};

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
{{.DataBuf}}
];
`, "\n")
	t := template.Must(template.New("usage").Parse(templateString))
	for _, f := range fonts() {
		fmt.Println("=======================================================")
		var index string
		var dataBuf string
		var dataBufLen int
		switch f.Name {
		case "Emoji":
			index = "/* TODO: Emoji char index */"
			dataBuf, dataBufLen = emojiData(f)
		case "Bold":
			index = sysLatinIndex(true)
			dataBuf, dataBufLen = sysLatinData(f, true)
		case "Regular":
			index = sysLatinIndex(true)
			dataBuf, dataBufLen = sysLatinData(f, true)
		case "Small":
			index = sysLatinIndex(false)
			dataBuf, dataBufLen = sysLatinData(f, false)
		default:
			panic("unexpected FontSpec.Name")
		}
		context := struct {
			Font         FontSpec
			OutPath      string
			CharMapIndex string
			DataBufLen   int
			DataBuf      string
		}{f, outPath, index, dataBufLen, dataBuf}
		err := t.Execute(os.Stdout, context)
		if err != nil {
			panic(err)
		}
	}
}

// Generate rust code (comments and u32 literals) for emoji glyphs packed as [u32]
func emojiData(font FontSpec) (dataBuf string, dataBufLen int) {
	dataBuf = "/* TODO */"
	dataBufLen = -1
	return
}

// Generate rust code (comments and u32 literals) for sysLatin glyphs packed as [u32]
func sysLatinData(font FontSpec, privateUseArea bool) (dataBuf string, dataBufLen int) {
	dataBuf = "/* TODO */"
	dataBufLen = -1
	return
}

// Generate charmap index source code string for sysLatin style glyph sheet (16x16 grid)
func sysLatinIndex(privateUseArea bool) string {
	templateString := strings.TrimSpace(`
/// Return offset into DATA[] for start of pattern depicting glyph for character c
pub fn get_glyph_pattern_offset(c: char) -> usize {
    match c as u32 {
        0x20..=0x7E => BASIC_LATIN[(c as usize) - 0x20] as usize,
        0xA0..=0xFF => LATIN_1[(c as usize) - 0xA0] as usize,
        0x152..=0x153 => LATIN_EXTENDED_A[(c as usize) - 0x152] as usize,
        0x2018..=0x2022 => GENERAL_PUNCTUATION[(c as usize) - 0x2018] as usize,
        0x20AC..=0x20AC => CURRENCY_SYMBOLS[(c as usize) - 0x20AC] as usize,{{if .PrivateUseArea}}
        0xE700..=0xE70C => PRIVATE_USE_AREA[(c as usize) - 0xE700] as usize,{{end}}
        _ => SPECIALS[(0xFFFD as usize) - 0xFFFD] as usize,
    }
}

// Index to Unicode Basic Latin block glyph patterns
const BASIC_LATIN: [u16; ${basicLatin.length}] = [
    ${basicLatin.map(v => v.start + ", // '" + v.chr + "'").join("\n    ")}
];

// Index to Unicode Latin 1 block glyph patterns
const LATIN_1: [u16; ${latin1.length}] = [
    ${latin1.map(v => v.start + ", // '" + v.chr + "'").join("\n    ")}
];

// Index to Unicode Latin Extended A block glyph patterns
const LATIN_EXTENDED_A: [u16; ${latinExtendedA.length}] = [
    ${latinExtendedA.map(v => v.start + ", // '" + v.chr + "'").join("\n    ")}
];

// Index to General Punctuation block glyph patterns
const GENERAL_PUNCTUATION: [u16; ${generalPunctuation.length}] = [
    ${generalPunctuation.map(v => v.start + ", // '" + v.chr + "'").join("\n    ")}
];

// Index to Unicode Currency Symbols block glyph patterns
const CURRENCY_SYMBOLS: [u16; ${currencySymbols.length}] = [
    ${currencySymbols.map(v => v.start + ", // '" + v.chr +"'").join("\n    ")}
];{{if .PrivateUseArea}}

// Index to Unicode Private Use Area block glyph patterns (UI sprites)
const PRIVATE_USE_AREA: [u16; ${privateUseArea.length}] = [
    ${privateUseArea.map(v => v.start + ", // " + v.name).join("\n    ")}
];{{end}}

// Index to Unicode Specials block glyph patterns
const SPECIALS: [u16; ${specials.length}] = [
    ${specials.map(v => v.start + ", // '" + v.chr + "'").join("\n    ")}
];
`)
	t := template.Must(template.New("charMapIndex").Parse(templateString))
	context := struct {
		PrivateUseArea bool
	}{privateUseArea}
	var buff bytes.Buffer
	err := t.Execute(&buff, context)
	if err != nil {
		panic(err)
	}
	return buff.String()
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
		Fonts   []FontSpec
	}{confirm, outPath, fonts()}
	err := t.Execute(os.Stdout, context)
	if err != nil {
		panic(err)
	}
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
//   > Code licensed under the MIT License: http://opensource.org/licenses/MIT
//   > Graphics licensed under CC-BY 4.0: https://creativecommons.org/licenses/by/4.0/
//
// - Twemoji Source Code Link:
//   https://github.com/twitter/twemoji`

const chicago = `// This code incorporates glyphs from the Chicago bitmap typeface which was
// designed by Susan Kare and released by Apple in 1984. Chicago is a registered
// trademark of Apple Inc.`

const geneva = `// This code incorporates glyphs from the Geneva bitmap typeface which was
// designed by Susan Kare and released by Apple in 1984. Geneva is a registered
// trademark of Apple Inc.`

package main

import (
	"bytes"
	"fmt"
	"guilib/codegen/font"
	"image/png"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"text/template"
)

type CharSpec font.CharSpec

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

{{.Index}}

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
{{.Data}}
`, "\n")
	t := template.Must(template.New("usage").Parse(templateString))
	for _, f := range fonts() {
		var data string
		var index string
		switch f.Name {
		case "Emoji":
			index, data = emojiIndexData(f)
		case "Bold":
			index, data = sysLatinIndexData(f, true)
		case "Regular":
			index, data = sysLatinIndexData(f, true)
		case "Small":
			index, data = sysLatinIndexData(f, false)
		default:
			panic("unexpected FontSpec.Name")
		}
		context := struct {
			Font    font.FontSpec
			OutPath string
			Index   string
			Data    string
		}{f, outPath, index, data}
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

// Generate rust code for emoji glyphs packed as [u32] along with corresponding index
func emojiIndexData(font font.FontSpec) (index, data string) {
	index = "/* TODO: Emoji index */"
	data = "/* TODO: Emoji data */"
	return
}

// Generate rust code for sysLatin glyphs packed as [u32] along with corresponding index
func sysLatinIndexData(fs font.FontSpec, hasPUA bool) (index, data string) {
	indexTemplate := strings.TrimSpace(`
/// Return offset into DATA[] for start of pattern depicting glyph for character c
pub fn get_glyph_pattern_offset(c: char) -> usize {
    match c as u32 {
        0x20..=0x7E => BASIC_LATIN[(c as usize) - 0x20] as usize,
        0xA0..=0xFF => LATIN_1[(c as usize) - 0xA0] as usize,
        0x152..=0x153 => LATIN_EXTENDED_A[(c as usize) - 0x152] as usize,
        0x2018..=0x2022 => GENERAL_PUNCTUATION[(c as usize) - 0x2018] as usize,
        0x20AC..=0x20AC => CURRENCY_SYMBOLS[(c as usize) - 0x20AC] as usize,{{if gt (len .PrivateUseArea) 0}}
        0xE700..=0xE70C => PRIVATE_USE_AREA[(c as usize) - 0xE700] as usize,{{end}}
        _ => SPECIALS[(0xFFFD as usize) - 0xFFFD] as usize,
    }
}

// Index to Unicode Basic Latin block glyph patterns
const BASIC_LATIN: [u16; {{len .BasicLatin}}] = [
    {{.BasicLatin.FormatRustCode}}
];

// Index to Unicode Latin 1 block glyph patterns
const LATIN_1: [u16; {{len .Latin1}}] = [
    {{.Latin1.FormatRustCode}}
];

// Index to Unicode Latin Extended A block glyph patterns
const LATIN_EXTENDED_A: [u16; {{len .LatinExtendedA}}] = [
    {{.LatinExtendedA.FormatRustCode}}
];

// Index to General Punctuation block glyph patterns
const GENERAL_PUNCTUATION: [u16; {{len .GeneralPunctuation}}] = [
    {{.GeneralPunctuation.FormatRustCode}}
];

// Index to Unicode Currency Symbols block glyph patterns
const CURRENCY_SYMBOLS: [u16; {{len .CurrencySymbols}}] = [
    {{.CurrencySymbols.FormatRustCode}}
];{{if gt (len .PrivateUseArea) 0}}

// Index to Unicode Private Use Area block glyph patterns (UI sprites)
const PRIVATE_USE_AREA: [u16; {{len .PrivateUseArea}}] = [
    {{.PrivateUseArea.FormatRustCode}}
];{{end}}

// Index to Unicode Specials block glyph patterns
const SPECIALS: [u16; {{len .Specials}}] = [
    {{.Specials.FormatRustCode}}
];
`)
	dataBufTemplate := strings.TrimSpace(`
pub const DATA: [u32; {{.DataBufLen}}] = [
{{.DataBufRust}}];
`)
	// Arrays for compiling codepoint to blit pattern lookup tables
	var basicLatin CpToBlitList         // Block:     00..7E; Subset:     20..7E
	var latin1 CpToBlitList             // Block:     80..FF; Subset:     A0..FF
	var latinExtendedA CpToBlitList     // Block:   100..17F; Subset:   152..153
	var generalPunctuation CpToBlitList // Block: 2000..206F; Subset: 2018..2022
	var currencySymbols CpToBlitList    // Block: 20A0..20CF; Subset: 20AC..20AC
	var privateUseArea CpToBlitList     // Block: E000..F8FF; Subset: E700..E70C
	var specials CpToBlitList           // Block: FFF0..FFFF; Subset: FFFD..FFFD
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
	for _, cs := range font.SysLatinMap() {
		// ID unicode block for this character
		block, _ := font.BlockAndIndex(cs.Codepoint)
		if block == font.PRIVATE_USE_AREA && fs.Name != "Bold" && fs.Name != "Regular" {
			// Skip PUA block glyphs for fonts that don't include them
			continue
		}
		// Find the glyph and pack it into a [u32] blit pattern for the data buffer
		matrix, yOffset := font.ConvertGlyphBoxToMatrix(img, fs, cs.Row, cs.Col)
		//fmt.Println(font.ConvertMatrixToText(matrix, yOffset))
		blitPattern := font.ConvertMatrixToPattern(matrix, yOffset)
		headerIndex := uint32(len(dataBuf))
		dataBuf = append(dataBuf, blitPattern...)
		comment := fmt.Sprintf("[%d]: %X '%s'", headerIndex, cs.Codepoint, cs.Chr)
		if block == font.PRIVATE_USE_AREA {
			comment = fmt.Sprintf("[%d]: %X %s", headerIndex, cs.Codepoint, cs.Chr)
		}
		rustCode := font.ConvertPatternToRust(blitPattern, comment)
		dataBufRust += rustCode
		// Update the character index with the correct offset (DATA[n]) for pattern header
		if block == font.BASIC_LATIN {
			basicLatin = append(basicLatin, CpToBlit{headerIndex, cs})
		} else if block == font.LATIN_1 {
			latin1 = append(latin1, CpToBlit{headerIndex, cs})
		} else if block == font.LATIN_EXTENDED_A {
			latinExtendedA = append(latinExtendedA, CpToBlit{headerIndex, cs})
		} else if block == font.GENERAL_PUNCTUATION {
			generalPunctuation = append(generalPunctuation, CpToBlit{headerIndex, cs})
		} else if block == font.CURRENCY_SYMBOLS {
			currencySymbols = append(currencySymbols, CpToBlit{headerIndex, cs})
		} else if block == font.PRIVATE_USE_AREA {
			privateUseArea = append(privateUseArea, CpToBlit{headerIndex, cs})
		} else if block == font.SPECIALS {
			specials = append(specials, CpToBlit{headerIndex, cs})
		} else {
			panic("unexpected block")
		}
	}
	// Render index template
	indexT := template.Must(template.New("index").Parse(indexTemplate))
	indexContext := struct {
		BasicLatin         CpToBlitList
		Latin1             CpToBlitList
		LatinExtendedA     CpToBlitList
		GeneralPunctuation CpToBlitList
		CurrencySymbols    CpToBlitList
		PrivateUseArea     CpToBlitList
		Specials           CpToBlitList
	}{basicLatin, latin1, latinExtendedA, generalPunctuation, currencySymbols, privateUseArea, specials}
	var indexStrBuf bytes.Buffer
	err = indexT.Execute(&indexStrBuf, indexContext)
	if err != nil {
		panic(err)
	}
	// Render data template
	dataT := template.Must(template.New("dataBuf").Parse(dataBufTemplate))
	dataContext := struct {
		DataBufRust string
		DataBufLen  int
	}{dataBufRust, len(dataBuf)}
	var dataStrBuf bytes.Buffer
	err = dataT.Execute(&dataStrBuf, dataContext)
	if err != nil {
		panic(err)
	}
	// Return
	index = indexStrBuf.String()
	data = dataStrBuf.String()
	return
}

// An element of lookup tables to translate from codepoint block to blit pattern
type CpToBlit struct {
	DataIndex uint32
	CS        font.CharSpec
}

// Type for slice of CpToBlit (so they can be used with .FormatRustCode in template strings)
type CpToBlitList []CpToBlit

// Format the inner elements of a [u16; n] codepoint to blit pattern index translation table
func (ctbList CpToBlitList) FormatRustCode() string {
	var rustCode []string
	for _, ctb := range ctbList {
		block, _ := font.BlockAndIndex(ctb.CS.Codepoint)
		if block != font.PRIVATE_USE_AREA {
			rustCode = append(rustCode, fmt.Sprintf("%d, // '%s'", ctb.DataIndex, ctb.CS.Chr))
		} else {
			// PUA block chars have descriptions in ctb.CS.Chr, so omit quotes in comment
			rustCode = append(rustCode, fmt.Sprintf("%d, // %s", ctb.DataIndex, ctb.CS.Chr))
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

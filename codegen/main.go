package main

import (
	"fmt"
	"os"
	"strings"
	"text/template"
)

// Path for output files with generated font code
const outPath = "../src/fonts"

// Spec for how to generate font source code files from glyph grid sprite sheets
func fonts() []FontSpec {
	return []FontSpec{
		FontSpec{"emoji48x48_o3x0.png", 48, 16, 0, 0, "emoji_index.txt", twemoji, "emoji.rs"},
		FontSpec{"bold.png", 30, 16, 2, 2, "bold_index.txt", chicago, "bold.rs"},
		FontSpec{"regular.png", 30, 16, 2, 2, "regular_index.txt", geneva, "regular.rs"},
		FontSpec{"small.png", 24, 16, 2, 2, "small_index.txt", geneva, "small.rs"},
	}
}

// Holds description of sprite sheet and character map for generating a font
type FontSpec struct {
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
	fmt.Println("TODO: generating code...")
	templateString := strings.TrimLeft(`
{{.Font.Legal}}

// {{.Font.Sprites}}:
//   size:{{.Font.Size}}, cols:{{.Font.Cols}}, gutter:{{.Font.Gutter}}, border:{{.Font.Border}}
// {{.Font.Index}}
// {{.Font.RustOut}}
`, "\n")
	t := template.Must(template.New("usage").Parse(templateString))
	for _, f := range fonts() {
		fmt.Println("=======================================================")
		context := struct {
			Font    FontSpec
			OutPath string
		}{f, outPath}
		err := t.Execute(os.Stdout, context)
		if err != nil {
			panic(err)
		}
	}
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

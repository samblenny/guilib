package main

import (
	"fmt"
	"os"
)

const confirm = "--replace-font-files"
const path = "../src/fonts"

func usage() {
	fmt.Printf(`
This tool packs font glyphs from ./*.png into rust source code.

Are you sure you want to replace the font files in %s?
If so, use the %s switch.

Usage:
    go run main.go %s

`, path, confirm, confirm)
}

func codegen() {
	fmt.Println("TODO: generating code...")
}

func main() {
	if len(os.Args) == 2 && os.Args[1] == confirm {
		codegen()
	} else {
		usage()
	}
}

package font

import (
	"fmt"
	"image"
	"strings"
)

type Matrix [][]int
type MatrixRow []int
type BlitPattern []uint32

// Holds description of sprite sheet and character map for generating a font
type FontSpec struct {
	Name    string // Name of font
	Sprites string // Which file holds the sprite sheet image with the grid of glyphs?
	Size    int    // How many pixels on a side is each glyph (precondition: square glyphs)
	Cols    int    // How many glyphs wide is the grid?
	Gutter  int    // How many px between glyphs?
	Border  int    // How many px wide are top and left borders?
	Legal   string // What credits or license notices need to be included in font file comments?
	RustOut string // Where should the generated source code go?
}

// Extract matrix of pixels from an image containing grid of glyphs
// - img: image.Image from png file containing glyph grid
// - font: Glyph sheet specs (glyph size, border/gutter, etc)
// - row: source row in glyph grid
// - col: source column in glyph grid
func ConvertGlyphBoxToMatrix(img image.Image, font FontSpec, row int, col int) (Matrix, uint32) {
	imgRect := img.Bounds()
	rows := (imgRect.Max.Y - font.Border) / (font.Size + font.Gutter)
	if row < 0 || row >= rows || col < 0 || col >= font.Cols {
		panic("row or column out of range")
	}
	// Get pixels for grid cell, converting from RGBA to 1-bit
	gridSize := font.Size + font.Gutter
	border := font.Border
	pxMatrix := Matrix{}
	for y := border + (row * gridSize); y < (row+1)*gridSize; y++ {
		var row MatrixRow
		for x := border + (col * gridSize); x < (col+1)*gridSize; x++ {
			r, _, _, _ := img.At(x, y).RGBA()
			//fmt.Println(r, g, b, a)
			if r == 0 {
				row = append(row, 1)
			} else {
				row = append(row, 0)
			}
		}
		pxMatrix = append(pxMatrix, row)
	}
	// Trim left whitespace
	trblTrimLimit := trimLimits(font, row, col)
	pxMatrix = matrixTranspose(pxMatrix)
	pxMatrix = trimLeadingEmptyRows(pxMatrix, trblTrimLimit[3])
	// Trim right whitespace
	pxMatrix = reverseRows(pxMatrix)
	pxMatrix = trimLeadingEmptyRows(pxMatrix, trblTrimLimit[1])
	pxMatrix = reverseRows(pxMatrix)
	pxMatrix = matrixTranspose(pxMatrix)
	// Trim top whitespace and calculate y-offset
	preTrimH := len(pxMatrix)
	pxMatrix = trimLeadingEmptyRows(pxMatrix, trblTrimLimit[0])
	yOffset := preTrimH - len(pxMatrix)
	// Trim bottom whitespace
	pxMatrix = reverseRows(pxMatrix)
	pxMatrix = trimLeadingEmptyRows(pxMatrix, trblTrimLimit[2])
	pxMatrix = reverseRows(pxMatrix)
	// Return matrix and yOffset
	return pxMatrix, uint32(yOffset)
}

// Return glyph as text with one ASCII char per pixel
func ConvertMatrixToText(matrix Matrix, yOffset uint32) string {
	var ascii string
	for _, row := range matrix {
		for _, px := range row {
			if px == 1 {
				ascii += "#"
			} else {
				ascii += "."
			}
		}
		ascii += "\n"
	}
	return ascii
}

// Reverse the order of rows in a matrix
func reverseRows(src Matrix) Matrix {
	var dest Matrix
	for i := len(src) - 1; i >= 0; i-- {
		dest = append(dest, src[i])
	}
	return dest
}

// Trim whitespace rows from top of matrix
func trimLeadingEmptyRows(pxMatrix Matrix, limit int) Matrix {
	if len(pxMatrix) < 1 {
		return pxMatrix
	}
	for i := 0; i < limit; i++ {
		sum := 0
		for _, n := range pxMatrix[0] {
			sum += n
		}
		if len(pxMatrix) > 0 && sum == 0 {
			pxMatrix = pxMatrix[1:]
		} else {
			break
		}
	}
	return pxMatrix
}

// Look up trim limits based on row & column in glyph grid
func trimLimits(font FontSpec, row int, col int) [4]int {
	if font.Name == "Bold" || font.Name == "Regular" {
		// Radio strength bars get trimmed to match bounds of three bars
		if col == 0 && row >= 5 && row <= 9 {
			return [4]int{7, 5, 6, 4}
		}
	}
	if font.Name == "Bold" || font.Name == "Regular" || font.Name == "Small" {
		// Space gets 4px width and 2px height
		if col == 2 && row == 0 {
			lr := (font.Size / 2) - 2
			tb := (font.Size / 2) - 1
			return [4]int{tb, lr, tb, lr}
		}
	}
	// Everthing else gets max trim
	return [4]int{font.Size, font.Size, font.Size, font.Size}
}

// Return pixel matrix as pattern packed into a byte array
// pat[0]: pattern width in px
// pat[1]: pattern height in px
// pat[2]: y-offset from top of line (position properly relative to text baseline)
// pat[3:2+w*h/8]: pixels packed into bytes
func ConvertMatrixToPattern(pxMatrix Matrix, yOffset uint32) BlitPattern {
	// Pack trimmed pattern into a byte array
	patW := uint32(0)
	patH := uint32(0)
	if len(pxMatrix) > 0 && len(pxMatrix[0]) > 0 {
		patW = uint32(len(pxMatrix[0]))
		patH = uint32(len(pxMatrix))
	}
	pattern := BlitPattern{(patW << 16) | (patH << 8) | yOffset}
	bufWord := uint32(0)
	flushed := false
	for y := uint32(0); y < patH; y++ {
		for x := uint32(0); x < patW; x++ {
			if pxMatrix[y][patW-1-x] > 0 {
				bufWord = (bufWord << 1) | 1
			} else {
				bufWord = (bufWord << 1)
			}
			flushed = false
			if (y*patW+x)%32 == 31 {
				pattern = append(pattern, bufWord)
				bufWord = 0
				flushed = true
			}
		}
	}
	if !flushed {
		finalShift := 32 - ((patW * patH) % 32)
		pattern = append(pattern, bufWord<<finalShift)
	}
	return pattern
}

// Convert blit pattern to rust source code for part of an array of bytes
func ConvertPatternToRust(pattern BlitPattern, comment string) string {
	patternStr := fmt.Sprintf("    // %s\n    ", comment)
	wordsPerRow := uint32(8)
	ceilRow := uint32(len(pattern)) / wordsPerRow
	if uint32(len(pattern))%wordsPerRow > 0 {
		ceilRow += 1
	}
	for i := uint32(0); i < ceilRow; i++ {
		start := i * wordsPerRow
		end := min(uint32(len(pattern)), (i+1)*wordsPerRow)
		line := pattern[start:end]
		var s []string
		for _, word := range line {
			s = append(s, fmt.Sprintf("0x%08x", word))
		}
		patternStr += strings.Join(s, ", ") + ","
		if end < uint32(len(pattern)) {
			patternStr += "\n    "
		}
	}
	patternStr += "\n"
	return patternStr
}

// Transpose a matrix (flip around diagonal)
func matrixTranspose(matrix Matrix) Matrix {
	if len(matrix) < 1 {
		return matrix
	}
	w := len(matrix[0])
	h := len(matrix)
	var transposed Matrix
	for col := 0; col < w; col++ {
		var trRow []int
		for row := 0; row < h; row++ {
			trRow = append(trRow, matrix[row][col])
		}
		transposed = append(transposed, trRow)
	}
	return transposed
}

// Return lowest value among two integers
func min(a uint32, b uint32) uint32 {
	if b > a {
		return a
	}
	return b
}

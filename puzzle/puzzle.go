package puzzle

import (
	"encoding/binary"
	"fmt"
	"hash/crc64"
	"math/rand"
	"time"
)

var table *crc64.Table

func init() {
	table = crc64.MakeTable(crc64.ECMA)
	rand.Seed(time.Now().UnixNano())
}

// Convert a binary number to a slice consisting of uint8s of either 1 or 0
// ex: 10010 to [1, 0, 0, 1, 0]
func binaryToIntSlice(bin uint32, width uint8) []uint8 {
	ints := make([]uint8, width)

	for i := range ints {
		if (1 << (width - uint8(i) - 1)) & bin != 0 {
			ints[i] = 1
		} else {
			ints[i] = 0
		}
	}

	return ints
}

// Takes a randomly generated unsigned integert and uses its binary
// representation to generate an array of numbers to represent the
// segments in the row or column
func binaryToSegmentNums(bin uint32) []uint8 {
	row := make([]uint8, 0, 16)
	
	for i := 0; bin > 0; bin = bin>>1 {
		if bin % 2 == 1 {
			if i >= len(row) {
				row = append(row, 0)
			}

			row[i] = row[i] + 1
		} else if i < len(row) && row[i] != 0 {
			i = i + 1
		}
	}

	for i, j := 0, len(row) - 1; i < j; i, j = i + 1, j - 1 {
		row[i], row[j] = row[j], row[i]
	}

	return row
}

func binaryRowsToColumnNums(rows []uint32, width uint8) [][]uint8 {
	columns := make([][]uint8, width)

	for i := range columns {
		var col uint32

		for j := range rows {
			if 1 << (width - uint8(i) - 1) & rows[j] != 0 {
				col = col | (1 << (width - uint8(j) - 1))
			}
		}
		columns[i] = binaryToSegmentNums(col)
	}

	return columns
}

func hashPuzzle(grid [][]uint8, height, width uint8) []byte {
	bytes, n := make([]byte, uint16(height) * uint16(width)), 0

	for row := range grid {
		for col := range grid[row] {
			bytes[n] = grid[row][col]
			n = n + 1
		}
	}

	sum := crc64.Checksum(bytes, table)
	byteArray := make([]byte, 8)
	binary.LittleEndian.PutUint64(byteArray, sum)

	return byteArray
}

// Puzzle data structure including height, width, and the completed puzzle
// along with all information presented to the player to allow them to
// complete it.
type Puzzle struct {
	RowValues [][]uint8		`json:"row_values"`
	ColumnValues [][]uint8	`json:"column_values"`
	Height uint8			`json:"height"`
	Width uint8				`json:"width"`
	CompletedGrid [][]uint8 `json:"-"`
	Hash []byte				`json:"-"`
}

// Print out a formatted puzzle
func (p *Puzzle) Print() {
	for i := range p.CompletedGrid {
		for j := range p.CompletedGrid[i] {
			fmt.Printf("%v ", p.CompletedGrid[i][j])
		}
		fmt.Printf("\n")
	}

	fmt.Printf("height: %v\nwidth: %v", p.Height, p.Width)
}

// GeneratePuzzle randomly builds a new puzzle of the given dimensions
func GeneratePuzzle(height, width uint8) (*Puzzle, error) {
	if height <= 0 || width <= 0 || height > 30 || width > 30 {
		return nil, fmt.Errorf(
			"Height and width must be in the 1 to 30 range (h: %d, w: %d)",
			height,
			width,
		)
	}

	if height % 5 != 0 || width % 5 != 0 {
		return nil, fmt.Errorf(
			"Height and width must bother be divisible by 5 (h: %d, w: %d)",
			height,
			width,
		)
	}

	for {
		binaryGrid := make([]uint32, height)
		completedGrid := make([][]uint8, height)
		rowValues := make([][]uint8, height)

		for i := range binaryGrid {
			binaryGrid[i] = uint32(rand.Intn(1 << width))
			completedGrid[i] = binaryToIntSlice(binaryGrid[i], width)
			rowValues[i] = binaryToSegmentNums(binaryGrid[i])
		}

		puzzle := Puzzle{
			RowValues: rowValues,
			ColumnValues: binaryRowsToColumnNums(binaryGrid, width),
			CompletedGrid: completedGrid,
			Height: height,
			Width: width,
			Hash: hashPuzzle(completedGrid, height, width),
		}

		if puzzle.validate() {
			return &puzzle, nil
		}
	}
}

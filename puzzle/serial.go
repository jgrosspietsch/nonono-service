package puzzle

// SerialPuzzle is a struct specifically for serializing into JSON.
// This exists since uint8 slices are handled like byte slices and
// encoded into base64 automatically by "encoding/json"
type SerialPuzzle struct {
	RowValues [][]uint16	`json:"row_values"`
	ColumnValues [][]uint16	`json:"column_values"`
	Height uint8			`json:"height"`
	Width uint8				`json:"width"`
}

func eightBitSlicesToSixteenBitSlices(original [][]uint8) [][]uint16 {
	rows := make([][]uint16, len(original))

	for i := range rows {
		rows[i] = make([]uint16, len(original[i]))
		for j := range rows[i] {
			rows[i][j] = uint16(original[i][j])
		}
	}

	return rows
}

// FormatAsSerializable creates a JSON-encodable struct of puzzle info
func (p *Puzzle) FormatAsSerializable() *SerialPuzzle {
	return &SerialPuzzle{
		RowValues: eightBitSlicesToSixteenBitSlices(p.RowValues),
		ColumnValues: eightBitSlicesToSixteenBitSlices(p.ColumnValues),
		Height: p.Height,
		Width: p.Width,
	}
}

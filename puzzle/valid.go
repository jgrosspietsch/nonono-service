package puzzle

const (
	unknown = iota
	empty
	filled
)

type solution []uint8

func padRightSolution(progress solution, size uint8) solution {
	for uint8(len(progress)) < size {
		progress = append(progress, uint8(0))
	}

	return progress
}

func appendSegment(segment uint8, progress solution, addGap bool) solution {
	result := append(solution{}, progress...)
	
	for i := uint8(0); i < segment; i = i + 1 {
		result = append(result, uint8(1))
	}

	if addGap {
		result = append(result, uint8(0))
	}

	return result
}

func remainingLength(segments []uint8) (length uint8) {
	length = uint8(len(segments)) - 1

	for i := range segments {
		length = length + segments[i]
	}

	return
}

func solutionTree(progress solution, size uint8, segments []uint8) []solution {
	if len(progress) == int(size) {
		return []solution{progress}
	}

	if len(segments) == 0 {
		return []solution{padRightSolution(progress, size)}
	}

	// append the next segment
	segmentAppends := append(
		[]solution{},
		solutionTree(
			appendSegment(
				segments[0],
				progress,
				len(segments) > 1,
			),
			size,
			segments[1:],
		)...
	)

	zeroAppends := []solution{}
	// if we can, append a 0
	if remainingLength(segments) < (size - uint8(len(progress))) {
		zeroAppends = append(
			[]solution{},
			solutionTree(
				append(progress, 0),
				size,
				segments,
			)...
		)
	}

	return append(segmentAppends, zeroAppends...)
}

func enumeratePossible(segments []uint8, size uint8) []solution {
	return solutionTree(solution{}, size, segments)
}

func emptyBoard(height, width uint8) [][]uint8 {
	board := make([][]uint8, height)

	for i := range board {
		board[i] = make([]uint8, width)
		for j := range board[i] {
			board[i][j] = unknown
		}
	}

	return board
}

func dropBadSolutions(current []uint8, solutions []solution) ([]solution, uint8) {
	result := make([]solution, 0)
	actions := uint8(0)

	for i := range solutions {
		badSolution := false
		for j := range solutions[i] {
			if (current[j] == empty && solutions[i][j] != uint8(0)) || (current[j] == filled && solutions[i][j] != uint8(1)) {
				badSolution = true
				break
			}
		}

		if !badSolution {
			result = append(result, solutions[i])
		} else {
			actions = actions + 1
		}
	}

	return result, actions
}

func columnFromBoard(board [][]uint8, index uint8) []uint8 {
	col := make([]uint8, uint8(len(board)))

	for i := range board {
		col[i] = board[i][index]
	}

	return col
}

func commonVals(lists []solution, val, size uint8) []uint8 {
	vals := []uint8{}
	accumulator := make([]uint8, size)

	for i := range lists {
		for j := range lists[i] {
			if lists[i][j] == val {
				accumulator[j] = accumulator[j] + 1
			}
		}
	}

	for i := range accumulator {
		if accumulator[i] == uint8(len(lists)) {
			vals = append(vals, uint8(i))
		}
	}

	return vals
}

func isComplete(board [][]uint8) bool {
	for i := range board {
		for j := range board[i] {
			if board[i][j] == unknown {
				return false
			}
		}
	}

	return true
}

func (p *Puzzle) validate() bool {
	rows := make([][]solution, p.Height)
	columns := make([][]solution, p.Width)
	board := emptyBoard(p.Height, p.Width)

	for i := range rows {
		rows[i] = enumeratePossible(p.RowValues[i], p.Width)
	}

	for i := range columns {
		columns[i] = enumeratePossible(p.ColumnValues[i], p.Height)
	}

	emptyCycles := 0
	for {
		actions := uint8(0)
		for i := range rows {
			rowActions := uint8(0)
			commonFilled := commonVals(rows[i], 1, p.Width)
			for f := range commonFilled {
				board[i][commonFilled[f]] = filled
			}

			commonEmpty := commonVals(rows[i], 0, p.Width)
			for e := range commonEmpty {
				board[i][commonEmpty[e]] = empty
			}
			rows[i], rowActions = dropBadSolutions(board[i], rows[i])
			actions = actions + rowActions
		}

		for i := range columns {
			columnActions := uint8(0)
			commonFilled := commonVals(columns[i], 1, p.Height)
			for f := range commonFilled {
				board[commonFilled[f]][i] = filled
			}
			
			commonEmpty := commonVals(columns[i], 0, p.Height)
			for e := range commonEmpty {
				board[commonEmpty[e]][i] = empty
			}
			columns[i], columnActions = dropBadSolutions(columnFromBoard(board, uint8(i)), columns[i])
			actions = actions + columnActions
		}

		if actions == 0 {
			emptyCycles = emptyCycles + 1
			if isComplete(board) {
				return true
			} else if emptyCycles > 1 {
				return false
			}
		} else {
			emptyCycles = 0
		}
	}
}

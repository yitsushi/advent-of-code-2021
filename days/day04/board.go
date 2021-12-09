package day04

const (
	boardSize = 5
)

// Field is one field on the board.
type Field struct {
	Value  int64
	Marked bool
}

// Board is a Bingo board.
type Board []Field

// Add a value to the Bingo Board.
func (b *Board) Add(value int64) {
	*b = append(*b, Field{Value: value, Marked: false})
}

// Row returns with a specific row from the board.
func (b Board) Row(n int) []Field {
	return b[boardSize*n : boardSize*(n+1)]
}

// Column returns with a specific column from the board.
func (b Board) Column(n int) []Field {
	output := []Field{}

	for i := 0; i < boardSize; i++ {
		output = append(output, b[i*boardSize+n])
	}

	return output
}

// Check marks a field and returns true if it was not marked before.
// It shouldn't be marked twice, but just in case.
func (b *Board) Check(value int64) bool {
	for idx, field := range *b {
		if !field.Marked && field.Value == value {
			field.Marked = true
			(*b)[idx] = field

			return b.DidWin()
		}
	}

	return false
}

func (b Board) winCheck(values []Field) bool {
	for _, field := range values {
		if !field.Marked {
			return false
		}
	}

	return true
}

// DidWin returns true if the board is a winner board.
func (b Board) DidWin() bool {
	for i := 0; i < boardSize; i++ {
		if b.winCheck(b.Row(i)) || b.winCheck(b.Column(i)) {
			return true
		}
	}

	return false
}

// UnmarkedValues returns all the values without mark.
func (b Board) UnmarkedValues() []int64 {
	output := []int64{}

	for _, field := range b {
		if !field.Marked {
			output = append(output, field.Value)
		}
	}

	return output
}

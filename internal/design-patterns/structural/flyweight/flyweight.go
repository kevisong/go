package flyweight

type ChessPieceUnit struct {
	ID    uint
	Name  string
	Color string
}

var units = map[int]*ChessPieceUnit{
	1: {
		ID:    1,
		Name:  "車",
		Color: "red",
	},
	2: {
		ID:    2,
		Name:  "炮",
		Color: "red",
	},
}

type ChessPiece struct {
	Unit *ChessPieceUnit
	X    int
	Y    int
}

type ChessBoard struct {
	chessPieces map[int]*ChessPiece
}

func NewChessBoard() *ChessBoard {
	board := &ChessBoard{chessPieces: map[int]*ChessPiece{}}
	for id := range units {
		board.chessPieces[id] = &ChessPiece{
			Unit: units[id],
			X:    0,
			Y:    0,
		}
	}
	return board
}

func (c *ChessBoard) Move(id, x, y int) {
	c.chessPieces[id].X = x
	c.chessPieces[id].Y = y
}

func Run() {
	board1 := NewChessBoard()
	board1.Move(1, 1, 2)
}

package pkg

type Piece struct {
	colour     string
	identifier int
	row, col   int
}

type Square struct {
	piece    *Piece
	row, col int
}

type Board struct {
	squares [][]Square
}

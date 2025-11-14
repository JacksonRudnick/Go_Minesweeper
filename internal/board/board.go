package board

import (
	"math/rand"
)

type Board struct {
	width, height int
	full_board    [][]uint8
	revealed      [][]bool
	marked        [][]bool

	whitespace      float64
	revealed_count  int
	num_bombs       int
	remaining_bombs int
	game_over       bool
	win             bool
}

// NewBoard creates a new Minesweeper board with the given width, height, and whitespace ratio.
func NewBoard(width, height int, whitespace float64) *Board {
	n := int(float64(width*height) * (1.0 - whitespace))

	Board := &Board{
		width:           width,
		height:          height,
		whitespace:      whitespace,
		num_bombs:      n,
		remaining_bombs: n,
	}

	for i := 0; i < height; i++ {
		Board.full_board = append(Board.full_board, make([]uint8, width))
		Board.revealed = append(Board.revealed, make([]bool, width))
		Board.marked = append(Board.marked, make([]bool, width))
	}

	for i := 0; i < Board.num_bombs; i++ {
		x := rand.Intn(width)
		y := rand.Intn(height)

		for Board.full_board[y][x] == 9 {
			x = rand.Intn(width)
			y = rand.Intn(height)
		}

		Board.full_board[y][x] = 9
	}

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if Board.full_board[y][x] == 9 {
				continue
			}

			neighbors := Board.GetNeighbors(x, y)
			count := uint8(0)
			
			for _, n := range neighbors {
				nx, ny := n[0], n[1]
				if Board.full_board[ny][nx] == 9 {
					count++
				}
			}

			Board.full_board[y][x] = count
		}
	}

	return Board
}

func (b *Board) RevealCell(x, y int) {
	if b.marked[y][x] || b.game_over {
		//fmt.Println("Cell already revealed or marked, or game over.")
		return
	}

	if (b.revealed_count == 0) && (b.full_board[y][x] != 0) {
		//generete a new board until first revealed cell is whitespace
		for b.full_board[y][x] != 0 {
			*b = *NewBoard(b.width, b.height, b.whitespace)
		}
	}

	if !b.revealed[y][x] {
		b.revealed_count++
		b.revealed[y][x] = true	
	}

	if b.full_board[y][x] == 9 {
		b.game_over = true
		return
	}

	if b.full_board[y][x] == 0 {
		neighbors := b.GetNeighbors(x, y)
		for _, n := range neighbors {
			nx, ny := n[0], n[1]
			if !b.revealed[ny][nx] {
				b.RevealCell(nx, ny)
			}
		}
	} else {
		if b.IsFulfilled(x, y) {
			neighbors := b.GetNeighbors(x, y)
			for _, n := range neighbors {
				nx, ny := n[0], n[1]
				if !b.revealed[ny][nx] && !b.marked[ny][nx] {
					b.RevealCell(nx, ny)
				}
			}
		}
	}
}

func (b *Board) MarkCell(x, y int) {
	if b.revealed[y][x] || b.game_over {
		//fmt.Println("Cell already revealed or game over.")
		return
	}

	if b.marked[y][x] {
		b.marked[y][x] = false
		b.remaining_bombs++
	} else {
		b.marked[y][x] = true
		b.remaining_bombs--
	}
}

func (b *Board) CheckWin() {
	for y := 0; y < b.height; y++ {
		for x := 0; x < b.width; x++ {
			if b.full_board[y][x] != 9 && !b.revealed[y][x] {
				return
			}
		}
	}

	b.win = true
	b.game_over = true
}

func (b *Board) GetNeighbors(x, y int) [][2]int {
	neighbors := [][2]int{}

	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			nx, ny := x+dx, y+dy
			if nx >= 0 && nx < b.width && ny >= 0 && ny < b.height && (dx != 0 || dy != 0) {
				neighbors = append(neighbors, [2]int{nx, ny})
			}
		}
	}

	return neighbors
}

func (b *Board) GetFullBoard() [][]uint8 {
	return b.full_board
}

func (b *Board) GetRevealed() [][]bool {
	return b.revealed
}

func (b *Board) GetMarked() [][]bool {
	return b.marked
}

func (b *Board) GetRemainingBombs() int {
	return b.remaining_bombs
}

func (b *Board) GetWidth() int {
	return b.width
}

func (b *Board) GetHeight() int {
	return b.height
}

func (b *Board) IsGameOver() bool {
	return b.game_over
}

func (b *Board) IsWin() bool {
	return b.win
}

func (b *Board) IsRevealed(x, y int) bool {
	return b.revealed[y][x]
}

func (b *Board) IsMarked(x, y int) bool {
	return b.marked[y][x]
}

func (b *Board) GetCellValue(x, y int) uint8 {
	return b.full_board[y][x]
}

func (b *Board) GetWhitespace() float64 {
	return b.whitespace
}

func (b *Board) IsFulfilled(x, y int) bool {
	val := b.full_board[y][x]

	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			nx, ny := x+dx, y+dy
			if nx >= 0 && nx < b.width && ny >= 0 && ny < b.height && (dx != 0 || dy != 0) {
				if b.marked[ny][nx] {
					val--
				}
			}
		}
	}

	return val == 0 || val == 9
}

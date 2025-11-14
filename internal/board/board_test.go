package board_test

import (
	"jacksonrudnick/minesweeper/internal/board"
	"testing"
)

func TestNewBoard(t *testing.T) {
	b := board.NewBoard(10, 10, 0.5)
	if b == nil {
		t.Fatal("expected non-nil board")
	}

	if len(b.GetFullBoard()) != 10 || len(b.GetFullBoard()[0]) != 10 {
		t.Fatalf("expected board of size 10x10, got %dx%d", len(b.GetFullBoard()), len(b.GetFullBoard()[0]))
	}

	if b.GetRemainingBombs() != 50 {
		t.Fatalf("expected 50 bombs, got %d", b.GetRemainingBombs())
	}
}

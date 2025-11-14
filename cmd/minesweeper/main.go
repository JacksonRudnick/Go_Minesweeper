package main

import (
	"jacksonrudnick/minesweeper/internal/board"
	"jacksonrudnick/minesweeper/internal/ui"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	b := board.NewBoard(50, 50, 0.7)

	// with mouse cell motion gives mouse events relative to cell positions
	p := tea.NewProgram(ui.InitialModel(b), tea.WithMouseCellMotion())
	if _, err := p.Run(); err != nil {
		panic(err)
	}
}

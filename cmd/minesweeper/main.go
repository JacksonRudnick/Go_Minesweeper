package main

import (
	"jacksonrudnick/minesweeper/internal/ui"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	// with mouse cell motion gives mouse events relative to cell positions
	p := tea.NewProgram(ui.InitialModel(), tea.WithMouseCellMotion())
	if _, err := p.Run(); err != nil {
		panic(err)
	}
}

package ui

import (
	"jacksonrudnick/minesweeper/internal/board"
	"strconv"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	board *board.Board
}

func InitialModel(b *board.Board) model {
	return model{
		board: b,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "esc", "ctrl+c":
			return m, tea.Quit
		case "r":
			m.board = board.NewBoard(m.board.GetWidth(), m.board.GetHeight(), m.board.GetWhitespace())
		}
	case tea.MouseMsg:
		// Only read on mouse release actions
		if msg.Action != tea.MouseActionRelease {
			return m, nil
		}

		// each cell is 3 characters wide and there is a header of 3 lines
		const headerLines = 4
		const cellWidth = 3

		tea.Println("Mouse event at:", msg.X, msg.Y)

		gx := (msg.X) / cellWidth
		gy := msg.Y - headerLines

		// bounds check
		if gx < 0 || gx >= m.board.GetWidth() || gy < 0 || gy >= m.board.GetHeight() {
			return m, nil
		}

		switch msg.Button {
		case tea.MouseButtonLeft:
			m.board.RevealCell(gx, gy)
		case tea.MouseButtonRight:
			m.board.MarkCell(gx, gy)
		}
	}
	return m, nil
}

func (m model) View() string {
	s := ""

	s += "Minesweeper Game\n"
	s += "Use mouse to interact. Press q to quit. Press r to restart.\n"

	if m.board.IsGameOver() {
		s += "\nGame Over! You hit a bomb.\n"
	} else if m.board.IsWin() {
		s += "\nCongratulations! You win!\n"
	} else {
		s += "\nRemaining Bombs: " + strconv.Itoa(m.board.GetRemainingBombs()) + "\n"
	}

	const (
		colorReset     = "\x1b[0m"
		colorBlue      = "\x1b[34m"
		colorGreen     = "\x1b[32m"
		colorRed       = "\x1b[31m"
		colorMagenta   = "\x1b[35m"
		colorCyan      = "\x1b[36m"
		colorYellow    = "\x1b[33m"
		colorGray      = "\x1b[90m"
		colorBrightRed = "\x1b[91m"
	)

	// Render each cell as exactly 3 visible characters so mouse X maps cleanly.
	// Layout examples (visible):
	// unrevealed: [#]
	// marked:     [F]
	// bomb:       [*]
	// number:     [0]
	for y := 0; y < m.board.GetHeight(); y++ {
		for x := 0; x < m.board.GetWidth(); x++ {
			if m.board.IsRevealed(x, y) {
				val := m.board.GetCellValue(x, y)
				if val == 9 {
					s += colorBrightRed + "[*]" + colorReset
				} else {
					numColor := colorBlue
					switch val {
					case 1:
						numColor = colorBlue
					case 2:
						numColor = colorGreen
					case 3:
						numColor = colorRed
					case 4:
						numColor = colorMagenta
					case 5:
						numColor = colorCyan
					case 6:
						numColor = colorYellow
					case 7:
						numColor = colorGray
					case 8:
						numColor = colorMagenta
					}
					s += numColor + "[" + string('0'+val) + "]" + colorReset
				}
			} else if m.board.IsMarked(x, y) {
				s += colorYellow + "[F]" + colorReset
			} else {
				s += colorGray + "[#]" + colorReset
			}
		}
		s += "\n"
	}

	return s
}
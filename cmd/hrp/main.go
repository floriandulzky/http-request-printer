package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/floriandulzky/http-request-printer/internal/view"
	"log"
)

func main() {
	p := tea.NewProgram(view.NewMainScreen(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}

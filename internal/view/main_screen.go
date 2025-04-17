package view

import (
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
	"github.com/floriandulzky/http-request-printer/internal/model"
	"github.com/floriandulzky/http-request-printer/internal/view/commands"
	"sort"
	"strings"
)

type NewRequestMsg model.HttpRequest

type mainScreen struct {
	responseChan  chan model.HttpRequest
	responses     []model.HttpRequest
	responseIndex int
	serverRunning bool
	windowWidth   int
	bodyHeight    int
	version       string
}

func NewMainScreen(version string) *mainScreen {
	return &mainScreen{
		responseChan: make(chan model.HttpRequest),
		version:      version,
	}
}

func listenForResponses(ch <-chan model.HttpRequest) tea.Cmd {
	return func() tea.Msg {
		return NewRequestMsg(<-ch)
	}
}

func (m *mainScreen) Init() tea.Cmd {
	m.responses = make([]model.HttpRequest, 0)
	return listenForResponses(m.responseChan)
}

func (m *mainScreen) Update(message tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := message.(type) {
	case NewRequestMsg:
		req := model.HttpRequest(msg)
		m.responses = append(m.responses, req)
		m.responseIndex = len(m.responses) - 1
		return m, listenForResponses(m.responseChan)
	case commands.ServerState:
		m.serverRunning = msg == "running"
	case tea.WindowSizeMsg:
		m.bodyHeight = msg.Height
		m.windowWidth = msg.Width
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyRight:
			if m.responseIndex < len(m.responses)-1 {
				m.responseIndex++
			}
		case tea.KeyLeft:
			if m.responseIndex > 0 {
				m.responseIndex--
			}
		case tea.KeyCtrlC:
			return m, tea.Quit
		case tea.KeyEnter:
			return m, commands.StartServer(m.responseChan)
		}
	}
	return m, nil
}

func (m *mainScreen) View() string {
	var bodyStyle = lipgloss.NewStyle().
		BorderStyle(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("#fecb3f")).
		Width(m.windowWidth - 2). // -2 => left and right border
		Height(m.bodyHeight - 4). // -2 => upper and lower border + footer
		MaxHeight(m.bodyHeight).
		PaddingLeft(1).PaddingRight(1).
		AlignVertical(lipgloss.Top)
	var footer = table.New().
		Row([]string{
			fmt.Sprintf("< %d/%d >", m.responseIndex+1, len(m.responses)),
			m.version}...).
		Width(m.windowWidth - 4).
		StyleFunc(func(row, col int) lipgloss.Style {
			if col == 1 {
				return lipgloss.NewStyle().Align(lipgloss.Right)
			}
			return lipgloss.NewStyle().Align(lipgloss.Left)
		}).
		Border(lipgloss.HiddenBorder()).
		String()
	if m.serverRunning {
		if len(m.responses) == 0 {
			return bodyStyle.Render("Send any http request to port 8000") + footer
		}
		responseBuilder := strings.Builder{}
		responseBuilder.WriteString(m.responses[m.responseIndex].Method)
		responseBuilder.WriteString(" ")
		responseBuilder.WriteString(m.responses[m.responseIndex].Url)
		responseBuilder.WriteString("\n\nHEADERS:\n")
		// sort map by key
		keys := make([]string, 0, len(m.responses[m.responseIndex].Headers))
		for k := range m.responses[m.responseIndex].Headers {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		for _, key := range keys {
			responseBuilder.WriteString(key)
			responseBuilder.WriteString(": ")
			responseBuilder.WriteString(strings.Join(m.responses[m.responseIndex].Headers[key], ", "))
			responseBuilder.WriteString("\n")
		}
		responseBuilder.WriteString("\nBody:\n")
		responseBuilder.WriteString(string(m.responses[m.responseIndex].Body))

		return bodyStyle.Render(responseBuilder.String()) +
			footer

	} else {
		return bodyStyle.Render(`
Welcome to

 _   _ _____ ____  
| | | |_   _|  _ \ 
| |_| | | | | |_) |
|  _  | | | |  __/ 
|_| |_| |_| |_|

Press Enter to start HTTP Server on port 8000
`) + footer
	}
}

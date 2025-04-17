package commands

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/floriandulzky/http-request-printer/internal/model"
	"github.com/floriandulzky/http-request-printer/internal/service"
)

type ServerState string

func StartServer(responseChan chan model.HttpRequest) func() tea.Msg {
	return func() tea.Msg {
		go service.NewHTTPServer(responseChan).Start()
		return ServerState("running")
	}
}

package server

import (
	"context"
	"io"
	"net/http"

	mainModel "github.com/CelestialCrafter/games/mainModel"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/log"
	"github.com/coder/websocket"
	"github.com/labstack/echo/v4"
)

// this file is mostly a bad reimplementation of charmbracelet/bubbletea/tea.go's program.eventLoop and program.Run

type Session struct {
	model    tea.Model
	handlers chan struct{}
}

var sessions = map[string]Session{}
var mm = mainModel.MainModel{}

func NewSession(id string, game uint) {
	sessions[id] = Session{
		model:    mm.NewGame(game),
		handlers: make(chan struct{}),
	}
	sessions[id].model.Init()
}

func UpdateSession(id string, msg tea.Msg) {
	sessions[id].model.Update(msg)
}

func ServeHttp(w http.ResponseWriter, r *http.Request) {
	c, err := websocket.Accept(w, r, nil)
	if err != nil {
		log.Warn("could not accept websocket", "error", err)
		return
	}
	defer c.CloseNow()

	for {
		err := tick(r.Context(), c)
		if websocket.CloseStatus(err) == websocket.StatusNormalClosure {
			return
		}
		if err != nil {
			log.Warn("coult not run play tick", "ip", r.RemoteAddr, "error", err)
			return
		}
	}
}

func PlayPlanet(c echo.Context) error {
	ServeHttp(c.Response(), c.Request())
	return nil
}

func tick(ctx context.Context, c *websocket.Conn) error {
	typ, r, err := c.Reader(ctx)
	if err != nil {
		return err
	}

	w, err := c.Writer(ctx, typ)
	if err != nil {
		return err
	}

	_, err = io.Copy(w, r)
	if err != nil {
		return err
	}

	err = w.Close()
	return err
}

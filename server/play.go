package server

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"

	twenty48 "github.com/CelestialCrafter/games/apps/2048"
	"github.com/CelestialCrafter/games/common"
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
	game     uint
	id       string
}

var sessions = map[string]Session{}
var mm = mainModel.MainModel{}

func newSession(id string, game uint) *Session {
	sessions[id] = Session{
		model:    mm.NewGame(game),
		handlers: make(chan struct{}),
		game:     game,
		id:       id,
	}

	session := sessions[id]
	session.model.Init()

	return &session
}

func serveHttp(w http.ResponseWriter, r *http.Request, user *userClaims) {
	c, err := websocket.Accept(w, r, nil)
	if err != nil {
		log.Warn("could not accept websocket", "error", err)
		return
	}
	defer func() {
		err := c.CloseNow()
		if err != nil {
			log.Warn("could not close websocket")
		}
	}()

	// @TODO extract from planet features
	game := common.Twenty48.ID
	session := newSession(user.ID, game)
	initialized := false

	for {
		var err error
		if !initialized {
			err = initialize(r.Context(), c, session)
			initialized = true
		} else {
			err = tick(r.Context(), c, session)
		}
		if websocket.CloseStatus(err) == websocket.StatusNormalClosure {
			return
		}
		if err != nil {
			log.Warn("coult not process message", "ip", r.RemoteAddr, "error", err)
			return
		}
	}
}

func PlayPlanet(c echo.Context) error {
	tokenString := c.QueryParam("token")
	token, err := Verify(tokenString)
	if err != nil {
		return err
	}

	user := token.Claims.(*userClaims)
	serveHttp(c.Response(), c.Request(), user)
	return nil
}

func initialize(ctx context.Context, c *websocket.Conn, session *Session) error {
	w, err := c.Writer(ctx, websocket.MessageText)
	if err != nil {
		return err
	}

	_, err = w.Write([]byte(strconv.Itoa(int(session.game))))
	if err != nil {
		return err
	}

	return w.Close()
}

var codeToTea = map[string]tea.KeyType{
	"ArrowLeft":  tea.KeyLeft,
	"ArrowRight": tea.KeyRight,
	"ArrowUp":    tea.KeyUp,
	"ArrowDown":  tea.KeyDown,
}

func serializeSession(s *Session) ([]byte, error) {
	switch s.game {
	case common.Twenty48.ID:
		game, ok := s.model.(twenty48.Model)
		if !ok {
			return nil, errors.New("model did not match game id")
		}
		return json.Marshal(&struct {
			Board    [][]uint16 `json:"board"`
			Finished bool       `json:"finished"`
		}{
			Board:    game.Board,
			Finished: game.Finished,
		})
	default:
		return nil, errors.New("unsupported model")
	}
}

func tick(ctx context.Context, c *websocket.Conn, session *Session) error {
	typ, r, err := c.Reader(ctx)
	if err != nil {
		return err
	}

	keyCodeBytes, err := io.ReadAll(r)
	if err != nil {
		return err
	}

	keyType, ok := codeToTea[string(keyCodeBytes)]
	// silently fail
	if !ok {
		return nil
	}

	keyMsg := tea.KeyMsg(tea.Key{Type: keyType})
	session.model.Update(keyMsg)

	w, err := c.Writer(ctx, typ)
	if err != nil {
		return err
	}

	serialized, err := serializeSession(session)
	if err != nil {
		return err
	}

	_, err = w.Write(serialized)
	if err != nil {
		return err
	}

	return w.Close()
}

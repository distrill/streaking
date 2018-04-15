package main

import (
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"
)

type handler struct {
	db *sqlx.DB
}

func (h *handler) getUsers(c echo.Context) error {
	um := userModel{h.db}
	us := um.read(nil)
	return c.JSON(http.StatusOK, us)
}

func (h *handler) getUser(c echo.Context) error {
	um := userModel{h.db}
	gm := goalModel{h.db}
	sm := streakModel{h.db}

	id := c.Param("user_id")

	u := um.read(map[string]interface{}{"id": id})
	gs := gm.read(map[string]interface{}{"user_id": id})
	ss := sm.read(map[string]interface{}{"user_id": id})

	return c.JSON(http.StatusOK, &struct {
		User    user     `json:"user"`
		Goals   []goal   `json:"goals"`
		Streaks []streak `json:"streaks"`
	}{u[0], gs, ss})
}

func (h *handler) getGoals(c echo.Context) error {
	gm := goalModel{h.db}
	id := c.Param("user_id")
	gs := gm.read(map[string]interface{}{"user_id": id})
	return c.JSON(http.StatusOK, gs)
}

func (h *handler) getStreaks(c echo.Context) error {
	sm := streakModel{h.db}
	id := c.Param("user_id")
	ss := sm.read(map[string]interface{}{"user_id": id})
	return c.JSON(http.StatusOK, ss)
}

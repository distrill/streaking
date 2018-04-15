package main

import (
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"
)

type handler struct {
	db *sqlx.DB
}

// GET /users
func (h *handler) getUsers(c echo.Context) error {
	um := userModel{h.db}
	us := um.read(nil)
	return c.JSON(http.StatusOK, us)
}

// GET /users/:user_id
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

// GET /users/:user_id/goals
func (h *handler) getGoals(c echo.Context) error {
	gm := goalModel{h.db}
	id := c.Param("user_id")
	gs := gm.read(map[string]interface{}{"user_id": id})
	return c.JSON(http.StatusOK, gs)
}

// GET /users/:user_id/streaks
func (h *handler) getStreaks(c echo.Context) error {
	sm := streakModel{h.db}
	id := c.Param("user_id")
	ss := sm.read(map[string]interface{}{"user_id": id})
	return c.JSON(http.StatusOK, ss)
}

// POST /users/:user_id/goals
func (h *handler) createGoal(c echo.Context) error {

	return nil
}

// POST /users/:user_id/streaks
func (h *handler) createStreak(c echo.Context) error { return nil }

// PUT /users/:user_id/goals
func (h *handler) updateGoal(c echo.Context) error { return nil }

// PUT /users/:user_id/streaks
func (h *handler) updateStreak(c echo.Context) error { return nil }

// DELTE /users/:user_id
func (h *handler) deleteUser(c echo.Context) error { return nil }

// DELETE /users/:user_id/goals
func (h *handler) deleteGoal(c echo.Context) error { return nil }

// DELETE /users/:user_id/streaks
func (h *handler) deleteStreak(c echo.Context) error { return nil }

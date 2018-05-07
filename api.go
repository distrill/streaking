package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"
)

type handler struct {
	db *sqlx.DB
}

type successResponse struct {
	Success bool `json:"success"`
}

// GET /users
func (h *handler) getUsers(c echo.Context) error {
	um := userModel{h.db}
	us, err := um.read(nil)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, us)
}

// GET /users/:user_id
func (h *handler) getUser(c echo.Context) error {
	um := userModel{h.db}
	gm := goalModel{h.db}
	sm := streakModel{h.db}

	uid := c.Param("user_id")

	var us []user
	var gs []goal
	var ss []streak
	var err error

	if us, err = um.read(map[string]interface{}{"id": uid}); err != nil {
		fmt.Println(err)
		return err
	}
	if gs, err = gm.read(map[string]interface{}{"user_id": uid}); err != nil {
		fmt.Println(err)
		return err
	}
	if ss, err = sm.read(map[string]interface{}{"user_id": uid}); err != nil {
		fmt.Println(err)
		return err
	}

	return c.JSON(http.StatusOK, &struct {
		User    user     `json:"user"`
		Goals   []goal   `json:"goals"`
		Streaks []streak `json:"streaks"`
	}{us[0], gs, ss})
}

// GET /users/:user_id/goals
func (h *handler) getGoals(c echo.Context) error {
	gm := goalModel{h.db}
	uid := c.Param("user_id")
	var gs []goal
	var err error

	if gs, err = gm.read(map[string]interface{}{"user_id": uid}); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, gs)
}

// GET /users/:user_id/streaks
func (h *handler) getStreaks(c echo.Context) error {
	sm := streakModel{h.db}
	uid := c.Param("user_id")
	var ss []streak
	var err error

	if ss, err = sm.read(map[string]interface{}{"user_id": uid}); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, ss)
}

// POST /users/:user_id/goals
func (h *handler) createGoal(c echo.Context) error {
	g := goal{}
	if err := c.Bind(&g); err != nil {
		return err
	}

	uid, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		return err
	}
	g.UserID = uid

	gm := goalModel{h.db}
	if err := gm.create(g); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, successResponse{true})
}

// POST /users/:user_id/streaks
func (h *handler) createStreak(c echo.Context) error {
	s := streak{}
	if err := c.Bind(&s); err != nil {
		return err
	}

	sm := streakModel{h.db}

	if err := sm.create(s); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, successResponse{true})
}

// PUT /users/:user_id/goals/:goal_id
func (h *handler) updateGoal(c echo.Context) error {
	g := goal{}
	if err := c.Bind(&g); err != nil {
		return err
	}

	gid, err := strconv.Atoi(c.Param("goal_id"))
	uid, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		return err
	}
	g.UserID = uid

	gm := goalModel{h.db}
	if err := gm.update(gid, g); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, successResponse{true})
}

// PUT /users/:user_id/streaks/:streak_id
func (h *handler) updateStreak(c echo.Context) error {
	s := streak{}
	if err := c.Bind(&s); err != nil {
		return err
	}

	sid, err := strconv.Atoi(c.Param("streak_id"))
	if err != nil {
		return err
	}

	sm := streakModel{h.db}
	if err := sm.update(sid, s); err != nil {
		fmt.Println(err)
		return err
	}

	return c.JSON(http.StatusOK, successResponse{true})
}

// DELTE /users/:user_id
func (h *handler) deleteUser(c echo.Context) error {
	uid, err := strconv.Atoi(c.Param("user_id"))
	if err != nil {
		return err
	}

	um := userModel{h.db}
	um.delete(uid)

	return c.JSON(http.StatusOK, successResponse{true})
}

// DELETE /users/:user_id/goals/:goal_id
func (h *handler) deleteGoal(c echo.Context) error {
	gid, err := strconv.Atoi(c.Param("goal_id"))
	if err != nil {
		return err
	}

	gm := goalModel{h.db}
	gm.delete(gid)

	return c.JSON(http.StatusOK, successResponse{true})
}

// DELETE /users/:user_id/streaks/:streak_id
func (h *handler) deleteStreak(c echo.Context) error {
	sid, err := strconv.Atoi(c.Param("streak_id"))
	if err != nil {
		return err
	}

	sm := streakModel{h.db}
	sm.delete(sid)

	return c.JSON(http.StatusOK, successResponse{true})
}

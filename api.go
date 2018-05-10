package main

import (
	"bh/streaking/models"
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
	um := models.Users{h.db}
	us, err := um.Read(nil)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, us)
}

// GET /users/:user_id
func (h *handler) getUser(c echo.Context) error {
	um := models.Users{h.db}
	gm := models.Goals{h.db}
	sm := models.Streaks{h.db}

	uid := c.Param("user_id")

	var us []models.User
	var gs []models.Goal
	var ss []models.Streak
	var err error

	if us, err = um.Read(map[string]interface{}{"id": uid}); err != nil {
		fmt.Println(err)
		return err
	}
	if gs, err = gm.Read(map[string]interface{}{"user_id": uid}); err != nil {
		fmt.Println(err)
		return err
	}
	if ss, err = sm.Read(map[string]interface{}{"user_id": uid}); err != nil {
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
	gm := Goals{h.db}
	uid := c.Param("user_id")
	var gs []goal
	var err error

	if gs, err = gm.Read(map[string]interface{}{"user_id": uid}); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, gs)
}

// GET /users/:user_id/streaks
func (h *handler) getStreaks(c echo.Context) error {
	sm := Streaks{h.db}
	uid := c.Param("user_id")
	var ss []streak
	var err error

	if ss, err = sm.Read(map[string]interface{}{"user_id": uid}); err != nil {
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

	gm := Goals{h.db}
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

	sm := Streaks{h.db}

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

	gm := Goals{h.db}
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

	sm := Streaks{h.db}
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

	um := models.Users{h.db}
	um.delete(uid)

	return c.JSON(http.StatusOK, successResponse{true})
}

// DELETE /users/:user_id/goals/:goal_id
func (h *handler) deleteGoal(c echo.Context) error {
	gid, err := strconv.Atoi(c.Param("goal_id"))
	if err != nil {
		return err
	}

	gm := Goals{h.db}
	gm.delete(gid)

	return c.JSON(http.StatusOK, successResponse{true})
}

// DELETE /users/:user_id/streaks/:streak_id
func (h *handler) deleteStreak(c echo.Context) error {
	sid, err := strconv.Atoi(c.Param("streak_id"))
	if err != nil {
		return err
	}

	sm := Streaks{h.db}
	sm.delete(sid)

	return c.JSON(http.StatusOK, successResponse{true})
}

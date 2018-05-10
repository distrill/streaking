package main

import (
	"bh/streaking/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/sessions"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
)

type handler struct {
	db *sqlx.DB
}

type successResponse struct {
	Success bool `json:"success"`
}

// GET /me
func (h *handler) getUser(c echo.Context) error {
	um := models.Users{DB: h.db}
	gm := models.Goals{DB: h.db}
	sm := models.Streaks{DB: h.db}

	sess, _ := session.Get("session", c)
	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7,
		HttpOnly: true,
	}
	uid := sess.Values["user"]

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
		Users   models.User     `json:"user"`
		Goals   []models.Goal   `json:"goals"`
		Streaks []models.Streak `json:"streaks"`
	}{us[0], gs, ss})
}

// POST /me/goals
func (h *handler) createGoal(c echo.Context) error {
	g := models.Goal{}
	if err := c.Bind(&g); err != nil {
		return err
	}

	sess, _ := session.Get("session", c)
	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7,
		HttpOnly: true,
	}
	uid := sess.Values["user"].(int)
	g.UserID = uid

	gm := models.Goals{DB: h.db}
	if err := gm.Create(g); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, successResponse{true})
}

// POST /me/streaks
func (h *handler) createStreak(c echo.Context) error {
	s := models.Streak{}
	if err := c.Bind(&s); err != nil {
		return err
	}

	sm := models.Streaks{DB: h.db}

	if err := sm.Create(s); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, successResponse{true})
}

// PUT /me/goals/:goal_id
func (h *handler) updateGoal(c echo.Context) error {
	g := models.Goal{}
	if err := c.Bind(&g); err != nil {
		return err
	}

	sess, _ := session.Get("session", c)
	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7,
		HttpOnly: true,
	}
	uid := sess.Values["user"].(int)
	gid, err := strconv.Atoi(c.Param("goal_id"))

	if err != nil {
		return err
	}
	g.UserID = uid

	gm := models.Goals{DB: h.db}
	if err := gm.Update(gid, g); err != nil {
		return err
	}

	return c.JSON(http.StatusOK, successResponse{true})
}

// PUT /me/streaks/:streak_id
func (h *handler) updateStreak(c echo.Context) error {
	s := models.Streak{}
	if err := c.Bind(&s); err != nil {
		return err
	}

	sid, err := strconv.Atoi(c.Param("streak_id"))
	if err != nil {
		return err
	}

	sm := models.Streaks{DB: h.db}
	if err := sm.Update(sid, s); err != nil {
		fmt.Println(err)
		return err
	}

	return c.JSON(http.StatusOK, successResponse{true})
}

// DELTE /users/:user_id
func (h *handler) deleteUser(c echo.Context) error {
	sess, _ := session.Get("session", c)
	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7,
		HttpOnly: true,
	}
	uid := sess.Values["user"].(int)

	um := models.Users{DB: h.db}
	um.Delete(uid)

	return c.JSON(http.StatusOK, successResponse{true})
}

// DELETE /me/goals/:goal_id
func (h *handler) deleteGoal(c echo.Context) error {
	gid, err := strconv.Atoi(c.Param("goal_id"))
	if err != nil {
		return err
	}

	gm := models.Goals{DB: h.db}
	gm.Delete(gid)

	return c.JSON(http.StatusOK, successResponse{true})
}

// DELETE /me/streaks/:streak_id
func (h *handler) deleteStreak(c echo.Context) error {
	sid, err := strconv.Atoi(c.Param("streak_id"))
	if err != nil {
		return err
	}

	sm := models.Streaks{DB: h.db}
	sm.Delete(sid)

	return c.JSON(http.StatusOK, successResponse{true})
}

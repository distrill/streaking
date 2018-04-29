package main

import (
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// streaking init
	db, err := sqlx.Connect("mysql", "streaking:streaking@/streaking")
	if err != nil {
		log.Panic(err)
	}

	h := handler{db}
	e := echo.New()

	// middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// routes
	e.GET("/users", h.getUsers)
	e.GET("/users/:user_id", h.getUser)
	e.GET("/users/:user_id/goals", h.getGoals)
	e.GET("/users/:user_id/streaks", h.getStreaks)

	e.POST("/users/:user_id/goals", h.createGoal)
	e.POST("/users/:user_id/streaks", h.createStreak)

	e.PUT("/users/:user_id/goals", h.updateGoal)
	e.PUT("/users/:user_id/streaks", h.updateStreak)

	e.DELETE("/users/:user_id/goals", h.deleteGoal)
	e.DELETE("/users/:user_id/streaks", h.deleteStreak)

	// listen and serve
	e.Logger.Fatal(e.Start(":8080"))
}

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

	// echo instance
	e := echo.New()

	// middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// routes
	e.GET("/users", h.getUsers)
	e.GET("/users/:user_id/goals", h.getGoals)
	e.GET("/users/:user_id/streaks", h.getStreaks)

	// listen and serve
	e.Logger.Fatal(e.Start(":8080"))
}

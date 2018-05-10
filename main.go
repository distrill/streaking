package main

import (
	"log"
	"net/http"

	"bh/streaking/auth"
	"bh/streaking/auth/facebook"
	"bh/streaking/auth/github"
	"bh/streaking/auth/google"

	"github.com/gorilla/sessions"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/middleware"
)

const htmlIndex = `
<html>
	<body>
		Log in with <a href="/login/facebook">facebook</a>
		<br />
		Log in with <a href="/login/github">github</a>
		<br />
		Log in with <a href="/login/google">google</a>
	</body>
</html>
`

func handleMain(c echo.Context) error {
	sess, _ := session.Get("session", c)
	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7,
		HttpOnly: true,
	}
	user := sess.Values["user"]

	if user != nil {
		return c.Redirect(http.StatusFound, "/api/users/2")
	}
	return c.HTML(http.StatusOK, htmlIndex)
}

func main() {
	db, err := sqlx.Connect("mysql", "streaking:streaking@/streaking")
	if err != nil {
		log.Panic(err)
	}

	e := echo.New()
	a := e.Group("/api")
	api := handler{db}

	// 	global middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(session.Middleware(sessions.NewCookieStore([]byte("big giant dick session secret"))))
	// api middleware
	a.Use(auth.IsLoggedIn)

	// login/auth routes
	e.GET("/", handleMain)

	e.GET("/login/facebook", facebook.HandleLogin())
	e.GET("/callback/facebook", facebook.HandleCallback(db))

	e.GET("/login/github", github.HandleLogin())
	e.GET("/callback/github", github.HandleCallback(db))

	e.GET("/login/google", google.HandleLogin())
	e.GET("/callback/google", google.HandleCallback(db))

	e.GET("/logout", auth.Logout)

	// api routes
	a.GET("/me", api.getUser)

	a.POST("/goals", api.createGoal)
	a.POST("/streaks", api.createStreak)

	a.PUT("/goals/:goal_id", api.updateGoal)
	a.PUT("/streaks/:streak_id", api.updateStreak)

	a.DELETE("/goals/:goal_id", api.deleteGoal)
	a.DELETE("/streaks/:streak_id", api.deleteStreak)

	e.Logger.Fatal(e.Start(":8080"))
}

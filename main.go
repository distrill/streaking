package main

import (
	"log"
	"net/http"

	"bh/streaking/auth/facebook"
	"bh/streaking/auth/github"
	"bh/streaking/auth/google"

	"github.com/gorilla/sessions"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
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
	return c.HTML(http.StatusOK, htmlIndex)
}

func main() {
	db, err := sqlx.Connect("mysql", "streaking:streaking@/streaking")
	if err != nil {
		log.Panic(err)
	}

	e := echo.New()
	e.Use(session.Middleware(sessions.NewCookieStore([]byte("big giant dick session secret"))))

	e.GET("/", handleMain)

	e.GET("/login/facebook", facebook.HandleLogin())
	e.GET("/callback/facebook", facebook.HandleCallback(db))

	e.GET("/login/github", github.HandleLogin())
	e.GET("/callback/github", github.HandleCallback(db))

	e.GET("/login/google", google.HandleLogin())
	e.GET("/callback/google", google.HandleCallback(db))

	e.Logger.Fatal(e.Start(":8080"))
}

// package main

// import (
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"os"
// 	"strconv"

// 	"github.com/dghubble/gologin"
// 	"github.com/dghubble/gologin/github"
// 	"github.com/dghubble/sessions"
// 	_ "github.com/go-sql-driver/mysql"
// 	"github.com/jmoiron/sqlx"
// 	"github.com/labstack/echo"
// 	"github.com/labstack/echo/middleware"
// 	"golang.org/x/oauth2"
// 	githubOAuth2 "golang.org/x/oauth2/github"
// )

// type errorResponse struct {
// 	Error string `json:"error"`
// }

// // TODO pls no global
// const (
// 	sessionName    = "streaking-app"
// 	sessionSecret  = "streaking app cookie signing secret"
// 	sessionUserKey = "streaking"
// )

// // TODO pls no global
// // sessionStore encodes and decodes session data stored in signed cookies
// var sessionStore = sessions.NewCookieStore([]byte(sessionSecret), nil)

// // TODO pls no global
// // Config configures the main ServeMux.
// type Config struct {
// 	GithubClientID     string
// 	GithubClientSecret string
// }

// // TODO api pls?
// // welcomeHandler shows a welcome message and login button.
// // func welcomeHandler(w http.ResponseWriter, req *http.Request) {
// func welcomeHandler(c echo.Context) error {
// 	fmt.Println(c.Request().URL.Path)
// 	fmt.Println("/")
// 	fmt.Println(c.Request().URL.Path == "/")
// 	if c.Request().URL.Path != "/" {
// 		return c.JSON(http.StatusNotFound, errorResponse{"not found"})
// 	}
// 	fmt.Println("whatever")
// 	if isAuthenticated(c.Request()) {
// 		return c.Redirect(http.StatusFound, "/profile")
// 	}
// 	return c.File("home.html")
// }

// // TODO this will be api routes i think
// // profileHandler shows protected user content.
// // func profileHandler(w http.ResponseWriter, req *http.Request) {
// func profileHandler(c echo.Context) error {
// 	return c.HTML(http.StatusOK, `<p>You are logged in!</p><form action="/logout" method="post"><input type="submit" value="Logout"></form>`)
// }

// // TODO api pls?
// // logoutHandler destroys the session on POSTs and redirects to home.
// func logoutHandler(c echo.Context) error {
// 	if c.Request().Method == "POST" {
// 		sessionStore.Destroy(c.Response(), sessionName)
// 	}
// 	return c.Redirect(http.StatusFound, "/")
// }

// // TODO middleware file pls?
// // requireLogin redirects unauthenticated users to the login route.
// func requireLogin(next echo.HandlerFunc) echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		if isAuthenticated(c.Request()) {
// 			return next(c)
// 		}
// 		return c.JSON(http.StatusUnauthorized, errorResponse{"unauthorized"})
// 	}
// }

// // TODO middleware file pls
// // isAuthenticated returns true if the user has a signed session cookie.
// func isAuthenticated(req *http.Request) bool {
// 	if _, err := sessionStore.Get(req, sessionName); err == nil {
// 		return true
// 	}
// 	return false
// }

// func issueSession() http.Handler {
// 	fn := func(w http.ResponseWriter, req *http.Request) {
// 		fmt.Println("one")
// 		ctx := req.Context()
// 		fmt.Println("two")
// 		githubUser, err := github.UserFromContext(ctx)
// 		fmt.Println("three")
// 		fmt.Println(*githubUser.Name)
// 		fmt.Println(strconv.Itoa(int(*githubUser.ID)))
// 		fmt.Println("GITHUB")
// 		if err != nil {
// 			http.Error(w, err.Error(), http.StatusInternalServerError)
// 			return
// 		}
// 		// 2. Implement a success handler to issue some form of session
// 		session := sessionStore.New(sessionName)
// 		session.Values[sessionUserKey] = *githubUser.ID
// 		session.Save(w)
// 		http.Redirect(w, req, "/profile", http.StatusFound)
// 	}
// 	return http.HandlerFunc(fn)
// }

// func main() {
// 	// streaking init
// 	db, err := sqlx.Connect("mysql", "streaking:streaking@/streaking")
// 	if err != nil {
// 		log.Panic(err)
// 	}

// 	// TODO some github place?
// 	config := &Config{
// 		GithubClientID:     os.Getenv("GITHUB_CLIENT_ID"),
// 		GithubClientSecret: os.Getenv("GITHUB_CLIENT_SECRET"),
// 	}
// 	oauth2Config := &oauth2.Config{
// 		ClientID:     config.GithubClientID,
// 		ClientSecret: config.GithubClientSecret,
// 		RedirectURL:  "http://localhost:8080/github/callback",
// 		Endpoint:     githubOAuth2.Endpoint,
// 	}
// 	stateConfig := gologin.DebugOnlyCookieConfig

// 	h := handler{db}
// 	e := echo.New()

// 	// middleware
// 	e.Use(middleware.Logger())
// 	e.Use(middleware.Recover())

// 	// routes
// 	e.GET("/", welcomeHandler)
// 	e.GET("/profile", requireLogin(profileHandler))
// 	e.GET("/logout", logoutHandler)

// 	e.GET("/github/login", func(c echo.Context) error {
// 		github.StateHandler(stateConfig, github.LoginHandler(oauth2Config, nil))
// 		return nil
// 	})
// 	e.GET("/github/callback", func(c echo.Context) error {
// 		github.StateHandler(stateConfig, github.CallbackHandler(oauth2Config, issueSession(), nil))
// 		return nil
// 	})

// 	e.GET("/users", h.getUsers)
// 	e.GET("/users/:user_id", h.getUser)
// 	e.GET("/users/:user_id/goals", h.getGoals)
// 	e.GET("/users/:user_id/streaks", h.getStreaks)

// 	e.POST("/users/:user_id/goals", h.createGoal)
// 	e.POST("/users/:user_id/streaks", h.createStreak)

// 	e.PUT("/users/:user_id/goals/:goal_id", h.updateGoal)
// 	e.PUT("/users/:user_id/streaks/:streak_id", h.updateStreak)

// 	e.DELETE("/users/:user_id/goals/:goal_id", h.deleteGoal)
// 	e.DELETE("/users/:user_id/streaks/:streak_id", h.deleteStreak)

// 	// listen and serve
// 	e.Logger.Fatal(e.Start(":8080"))
// }

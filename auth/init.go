package auth

import (
	"bh/streaking/models"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/jmoiron/sqlx"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
	"golang.org/x/oauth2"
)

type errorResponse struct {
	Message bool `json:"success"`
}

var skipRoutes = map[string]bool{
	"/login":             true,
	"/logout":            true,
	"/login/facebook":    true,
	"/login/github":      true,
	"/login/google":      true,
	"/callback/facebook": true,
	"/callback/github":   true,
	"/callback/google":   true,
}

// Settings - settings for various login schemes
type Settings struct {
	OauthConf        *oauth2.Config
	OauthStateString string
	BaseURL          string
	DB               *sqlx.DB
	GetUser          func(string) models.User
}

// BuildLoginHandler build login handler given oauth conf and oauth state string
func BuildLoginHandler(settings Settings) echo.HandlerFunc {
	return func(c echo.Context) error {
		URL, err := url.Parse(settings.OauthConf.Endpoint.AuthURL)
		if err != nil {
			log.Fatal("Parse: ", err)
		}
		parameters := url.Values{}
		parameters.Add("client_id", settings.OauthConf.ClientID)
		parameters.Add("redirect_uri", settings.OauthConf.RedirectURL)
		parameters.Add("scope", strings.Join(settings.OauthConf.Scopes, " "))
		parameters.Add("response_type", "code")
		parameters.Add("invalid", "offline")
		parameters.Add("state", settings.OauthStateString)
		URL.RawQuery = parameters.Encode()
		url := URL.String()
		return c.Redirect(http.StatusTemporaryRedirect, url)
	}
}

// BuildCallbackHandler build callback handler given oauth conf and oauth state string
func BuildCallbackHandler(settings Settings) echo.HandlerFunc {
	return func(c echo.Context) error {
		// pull state and code params out of request
		query := new(struct {
			State string `query:"state"`
			Code  string `query:"code"`
		})
		if err := c.Bind(query); err != nil {
			return err
		}

		// ensure state matches what we set, to prevent phishing attacks
		if query.State != settings.OauthStateString {
			fmt.Printf("invalid oauth state, expected '%s', got '%s'\n", settings.OauthStateString, query.State)
			return c.Redirect(http.StatusTemporaryRedirect, "/")
		}

		// exchange code for access token (oauth step)
		token, err := settings.OauthConf.Exchange(oauth2.NoContext, query.Code)
		if err != nil {
			fmt.Printf("oauthConf.Exchange() failed with '%s'\n", err)
			return c.Redirect(http.StatusTemporaryRedirect, "/")
		}

		// grab user info with said access token
		resp, err := http.Get(settings.BaseURL + url.QueryEscape(token.AccessToken))
		if err != nil {
			fmt.Printf("Get: %s\n", err)
			return c.Redirect(http.StatusTemporaryRedirect, "/")
		}
		defer resp.Body.Close()
		response, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("ReadAll: %s\n", err)
			return c.Redirect(http.StatusTemporaryRedirect, "/")
		}

		// parse user object out of user info response and insert/get from db
		u := settings.GetUser(string(response))
		um := models.Users{DB: settings.DB}
		if err := um.Create(u); err != nil {
			log.Fatal(err)
		}
		users, err := um.Read(map[string]interface{}{
			"name":        u.Name,
			"email":       u.Email,
			"source":      u.Source,
			"external_id": u.ExternalID,
		})
		if err != nil {
			return err
		}
		if len(users) != 1 {
			return fmt.Errorf("invalid user records: %v", users)
		}
		user := users[0]
		fmt.Println(user)

		// set inserted user id in session
		sess, _ := session.Get("session", c)
		sess.Options = &sessions.Options{
			Path:     "/",
			MaxAge:   86400 * 7,
			HttpOnly: true,
		}
		sess.Values["user"] = user.ID
		sess.Save(c.Request(), c.Response())

		// should redirect to app url
		return c.Redirect(http.StatusFound, "/")
	}
}

// CheckLogIn - middleware to ensure user is logged in
func CheckLogIn(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// skip check for some routes
		if skipRoutes[c.Request().URL.Path] == true {
			return next(c)
		}

		sess, _ := session.Get("session", c)
		sess.Options = &sessions.Options{
			Path:     "/",
			MaxAge:   86400 * 7,
			HttpOnly: true,
		}
		user := sess.Values["user"]

		if user == nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "Please log in")
		}
		fmt.Println("two")

		return next(c)
	}
}

// Logout - log user out
func Logout(c echo.Context) error {
	sess, _ := session.Get("session", c)
	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7,
		HttpOnly: true,
	}
	sess.Values["user"] = nil
	sess.Save(c.Request(), c.Response())

	return c.Redirect(http.StatusFound, "/")
}

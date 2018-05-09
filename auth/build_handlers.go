package auth

import (
	"bh/streaking/models"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/labstack/echo"
	"golang.org/x/oauth2"
)

// Settings - settings for various login schemes
type Settings struct {
	OauthConf        *oauth2.Config
	OauthStateString string
	BaseURL          string
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
		query := new(struct {
			State string `query:"state"`
			Code  string `query:"code"`
		})

		if err := c.Bind(query); err != nil {
			return err
		}

		if query.State != settings.OauthStateString {
			fmt.Printf("invalid oauth state, expected '%s', got '%s'\n", settings.OauthStateString, query.State)
			return c.Redirect(http.StatusTemporaryRedirect, "/")
		}

		token, err := settings.OauthConf.Exchange(oauth2.NoContext, query.Code)
		if err != nil {
			fmt.Printf("oauthConf.Exchange() failed with '%s'\n", err)
			return c.Redirect(http.StatusTemporaryRedirect, "/")
		}

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

		// TODO upsert user in db, session
		u := settings.GetUser(string(response))
		fmt.Println(u)
		return c.Redirect(http.StatusTemporaryRedirect, "/")
	}
}

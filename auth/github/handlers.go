package github

import (
	"bh/streaking/auth"
	"bh/streaking/models"
	"encoding/json"
	"log"
	"strconv"

	"github.com/jmoiron/sqlx"

	"github.com/labstack/echo"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

var settings = auth.Settings{
	OauthConf: &oauth2.Config{
		ClientID:     "27664cbca31fbcd886db",
		ClientSecret: "9535df4affb9bd25ec44f6d00a32480a4fd9a078",
		RedirectURL:  "http://localhost:8080/callback/github",
		Scopes:       []string{"public_profile"},
		Endpoint:     github.Endpoint,
	},
	OauthStateString: "thisshouldberandom",
	BaseURL:          "https://api.github.com/user?access_token=",
	GetUser:          getUser,
}

func getUser(res string) models.User {
	temp := new(struct {
		ID    int    `json:"id"`
		Name  string `json:"name"`
		Email string `json:"email"`
	})
	if err := json.Unmarshal([]byte(res), &temp); err != nil {
		log.Fatal(err)
	}

	if temp.Name == "" {
		temp.Name = "NO_NAME_GIVEN"
	}
	if temp.Email == "" {
		temp.Email = "NO_EMAIL_GIVEN"
	}

	return models.User{
		Name:       temp.Name,
		Email:      temp.Email,
		Source:     "GITHUB",
		ExternalID: strconv.Itoa(temp.ID),
	}
}

// HandleLogin - handle facebook login
func HandleLogin() echo.HandlerFunc {
	return auth.BuildLoginHandler(settings)
}

// HandleCallback - handle facebook callback
func HandleCallback(db *sqlx.DB) echo.HandlerFunc {
	settings.DB = db
	return auth.BuildCallbackHandler(settings)
}

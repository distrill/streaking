package google

import (
	"bh/streaking/auth"
	"bh/streaking/models"
	"encoding/json"
	"log"

	"github.com/jmoiron/sqlx"

	"github.com/labstack/echo"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var settings = auth.Settings{
	OauthConf: &oauth2.Config{
		ClientID:     "443546063879-ocoa94kseo25apobl1dol3kqi2vkaqq1.apps.googleusercontent.com",
		ClientSecret: "WtR0ABtcDhWwfVcV3FR14SUI",
		RedirectURL:  "http://localhost:8080/callback/google",
		Scopes:       []string{"profile", "email"},
		Endpoint:     google.Endpoint,
	},
	OauthStateString: "thisshouldberandom",
	BaseURL:          "https://www.googleapis.com/oauth2/v3/tokeninfo?access_token=",
	GetUser:          getUser,
}

func getUser(res string) models.User {
	temp := new(struct {
		ID    string `json:"sub"`
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
		Source:     "GOOGLE",
		ExternalID: temp.ID,
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

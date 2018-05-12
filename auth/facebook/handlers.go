package facebook

import (
	"bh/streaking/auth"
	"bh/streaking/models"
	"encoding/json"
	"log"
	"os"

	"github.com/jmoiron/sqlx"

	"github.com/labstack/echo"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/facebook"
)

var settings = auth.Settings{
	OauthConf: &oauth2.Config{
		ClientID:     "226042608152816",
		ClientSecret: "617e257795853d28e562ebecd14e400f",
		RedirectURL:  os.Getenv("BASE_URL") + "/callback/facebook",
		Scopes:       []string{"public_profile", "email"},
		Endpoint:     facebook.Endpoint,
	},
	OauthStateString: "thisshouldberandom",
	BaseURL:          "https://graph.facebook.com/me?access_token=",
	GetUser:          getUser,
}

func getUser(res string) models.User {
	temp := new(struct {
		ID    string `json:"id"`
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
		Source:     "FACEBOOK",
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

package service

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"os"
	"split-that.com/split-that/v2/src/constants"
)

type AuthService struct {
	Cfg *oauth2.Config
}

func GetAuthService() *AuthService {
	cfg := &oauth2.Config{
		ClientID:     os.Getenv(constants.ClientId),
		ClientSecret: os.Getenv(constants.ClientSecret),
		Endpoint:     google.Endpoint,
		RedirectURL:  os.Getenv(constants.OAuthRedirectURL),
		Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email", "https://www.googleapis.com/auth/userinfo.profile"},
	}

	return &AuthService{Cfg: cfg}
}

func (as *AuthService) getCreds() {
}

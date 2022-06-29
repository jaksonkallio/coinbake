package config

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var OauthProviderConfigs = make(map[string]*oauth2.Config, 0)

func InitOauthProviderConfigs() {
	// Create the "google" oauth config.
	OauthProviderConfigs["google"] = &oauth2.Config{
		ClientID:     CurrentConfig.Oauth.Google.ClientId,
		ClientSecret: CurrentConfig.Oauth.Google.ClientSecret,
		RedirectURL:  "http://localhost:5010/api/v1/oauth_callback",
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
		},
		Endpoint: google.Endpoint,
	}
}

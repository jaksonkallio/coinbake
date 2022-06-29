package service

import (
	"log"
	"net/http"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/jaksonkallio/coinbake/config"
	"github.com/jaksonkallio/coinbake/random"
)

func Serve() {
	router := gin.Default()

	// All static files are accessible at `/static`
	router.Static("/static", "../frontend/dist/")

	// The index file is accessible at the root
	router.StaticFile("/", "../frontend/dist/index.html")

	// Set up the Gin session store.
	token, err := random.RandToken(64)
	if err != nil {
		log.Fatal("unable to generate random token: ", err)
	}

	store := sessions.NewCookieStore([]byte(token))
	store.Options(sessions.Options{
		Path:   "/",
		MaxAge: 86400 * 7,

		// Only set to unsecure when in dev.
		Secure: !config.IsDev(),

		// Cookie is not available to the frontend Javascript, and is only attached to API requests.
		HttpOnly: true,
	})

	router.Use(sessions.Sessions("coinbakesession", store))

	apiV1 := router.Group("/api/v1")
	apiV1.Use(Authenticate())

	apiV1.GET("/strategy", GetStrategy())
	apiV1.POST("/strategy", PostStrategy())
	apiV1.GET("/portfolio", GetPortfolio())
	apiV1.GET("/portfolios", GetPortfolios())
	apiV1.GET("/assets", GetAssets())
	apiV1.POST("/portfolio", PostPortfolio())
	apiV1.GET("/user", GetUser())
	apiV1.GET("/exchange_connection_valid", GetExchangeConnectionValid())
	apiV1.GET("/exchange_supported_assets", GetExchangeSupportedAssets())

	apiV1NonAuth := router.Group("/api/v1")
	apiV1NonAuth.GET("/oauth_login_url", GetOauthLoginUrl())
	apiV1NonAuth.GET("/oauth_callback", GetOauthCallback())

	http.ListenAndServe(":5010", router)
}

package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/jaksonkallio/coinbake/config"
	"github.com/jaksonkallio/coinbake/random"
	"golang.org/x/oauth2"
)

func GetStrategy() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// TODO: implement
		ctx.Status(http.StatusNotImplemented)
	}
}

func PostStrategy() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// TODO: implement
		ctx.Status(http.StatusNotImplemented)
	}
}

func GetPortfolio() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authUserId := DistillAuthUserId(ctx)

		portfolioId, err := strconv.Atoi(ctx.Query("id"))
		if err != nil {
			// TODO: error
		}

		portfolio, err := FindPortfolioById(uint(portfolioId))
		if err != nil {
			// TODO: error
		}

		if authUserId != uint(portfolio.UserID) {
			// TODO: error
		}

		buildStandardResponse(
			ctx,
			gin.H{
				"Portfolio": portfolio,
			},
		)
	}
}

func GetPortfolios() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authUserId := DistillAuthUserId(ctx)
		portfolios := FindPortfoliosByUserId(authUserId)

		buildStandardResponse(
			ctx,
			gin.H{
				"Portfolios": portfolios,
			},
		)
	}
}

func GetAssets() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var page int

		if len(ctx.Query("page")) > 0 {
			page, _ = strconv.Atoi(ctx.Query("page"))
		}

		// TODO: support other sorts

		// Get the assets ordered by marketcap
		assets := FindAssetsByMarketCap(20, page)

		buildStandardResponse(
			ctx,
			gin.H{
				"Assets": assets,
			},
		)
	}
}

func PostPortfolio() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// TODO: implement
		ctx.Status(http.StatusNotImplemented)
	}
}

func GetUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// TODO: implement
		ctx.Status(http.StatusNotImplemented)
	}
}

func GetExchangeSupportedAssets() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		portfolioId, err := strconv.Atoi(ctx.Query("portfolio_id"))
		if err != nil {
			// TODO: error
		}

		portfolio, err := FindPortfolioById(uint(portfolioId))
		if err != nil {
			// TODO: error
		}

		exchange, err := portfolio.Exchange()
		if err != nil {
			// TODO: error
		}

		supportedAssets, err := exchange.SupportedAssets(portfolio)
		if err != nil {
			// TODO: error
		}

		ctx.JSON(http.StatusOK, gin.H{
			"ExchangeSupportedAssets": supportedAssets,
		})
	}
}

func GetExchangeConnectionValid() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authUserId := DistillAuthUserId(ctx)

		portfolioId, err := strconv.Atoi(ctx.Query("portfolio_id"))
		if err != nil {
			// TODO: error
		}

		portfolio, err := FindPortfolioById(uint(portfolioId))
		if err != nil {
			// TODO: error
		}

		if authUserId != uint(portfolio.UserID) {
			// TODO: error
		}

		exchange, err := portfolio.Exchange()
		if err != nil {
			// TODO: error
		}

		testExchangeConnectionResult := exchange.ValidateConnection(portfolio)

		buildStandardResponse(
			ctx,
			gin.H{
				"TestExchangeConnectionResult": testExchangeConnectionResult,
			},
		)
	}
}

func GetOauthCallback() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Rehydrate the session.
		session := sessions.Default(ctx)

		// Get and compare the session state with the request query state.
		retrievedState := session.Get("state")
		queryState := ctx.Request.URL.Query().Get("state")
		if retrievedState != queryState {
			buildFailureResponse(ctx, http.StatusUnauthorized, "Invalid session state")
			return
		}

		// TODO: Dynamically select oauth provider (based on `provider` in session state) instead of hardcoding to "google". This will allow a variety of providers.
		// Get the oauth provider.
		oauthProvider := config.OauthProviderConfigs["google"]

		code := ctx.Request.URL.Query().Get("code")
		tok, err := oauthProvider.Exchange(oauth2.NoContext, code)
		if err != nil {
			buildFailureResponse(ctx, http.StatusBadRequest, "Login failed")
			return
		}

		// Create the oauth client.
		client := oauthProvider.Client(oauth2.NoContext, tok)
		// TODO: make google non-specific
		userInfoRes, err := client.Get("https://www.googleapis.com/oauth2/v3/userinfo")
		if err != nil {
			buildFailureResponse(ctx, http.StatusBadRequest, "Could not get user info: %s", err)
			return
		}
		defer userInfoRes.Body.Close()
		userInfoResBody, err := ioutil.ReadAll(userInfoRes.Body)
		if err != nil {
			buildFailureResponse(ctx, http.StatusInternalServerError, "Could not fetch user info response body: %s", err)
			return
		}

		// Parse user info.
		oauthUserInfo := &OauthUserInfo{}
		json.Unmarshal(userInfoResBody, &oauthUserInfo)
		if err != nil {
			buildFailureResponse(ctx, http.StatusInternalServerError, "Could not parse user info response body: %s", err)
			return
		}

		// Find user based on the email address from the OAuth user info.
		user := FindUserByEmailAddress(oauthUserInfo.Email)
		if user == nil {
			buildFailureResponse(ctx, http.StatusBadRequest, "User with that email address does not exist.")
			return
		}

		session.Set(UserIdSessionKey, user.ID)
		err = session.Save()
		if err != nil {
			buildFailureResponse(ctx, http.StatusInternalServerError, "Could not save session: %s", err)
			return
		}

		// Mark the user's last authed value.
		user.MarkLastAuthed()

		// Redirect the user to a "normal" page.
		ctx.Redirect(http.StatusTemporaryRedirect, "http://localhost:5010")
	}
}

func GetOauthLoginUrl() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Generate a session state token.
		state, err := random.RandToken(32)
		if err != nil {
			buildFailureResponse(ctx, http.StatusInternalServerError, "Could not generate session state token: %s", err)
			return
		}

		// Get the session and update with the random state token.
		session := sessions.Default(ctx)
		session.Set("state", state)
		err = session.Save()
		if err != nil {
			buildFailureResponse(ctx, http.StatusInternalServerError, "Could not save session: %s", err)
			return
		}

		// TODO: Dynamically select oauth provider (based on `provider` in session state) instead of hardcoding to "google". This will allow a variety of providers.
		// Get the oauth provider.
		oauthProvider := config.OauthProviderConfigs["google"]

		buildStandardResponse(ctx, gin.H{
			"OauthLoginUrl": oauthProvider.AuthCodeURL(state),
		})
	}
}

func buildStandardResponse(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, data)
}

func buildFailureResponse(ctx *gin.Context, statusCode int, message string, values ...interface{}) {
	e := fmt.Sprintf(message, values...)
	log.Printf("Request failed: %s", e)
	ctx.String(statusCode, e)
}

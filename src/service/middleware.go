package service

import (
	"net/http"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

const AuthUserIdKey string = "auth-user-id"

func Authenticate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		session := sessions.Default(ctx)
		userId := session.Get(UserIdSessionKey).(uint)
		if userId == 0 {
			buildFailureResponse(ctx, http.StatusUnauthorized, "Not logged in")
			return
		}

		ctx.Set(AuthUserIdKey, int(userId))

		ctx.Next()
	}
}

func DistillAuthUserId(ctx *gin.Context) uint {
	return uint(ctx.GetInt(AuthUserIdKey))
}

package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/jordanmarcelino/terradiscover-backend/internal/apperror"
	"github.com/jordanmarcelino/terradiscover-backend/internal/constant"
	"github.com/jordanmarcelino/terradiscover-backend/internal/utils/jwtutils"
)

func Authorization(jwtUtil jwtutils.JwtUtil) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		accessToken, err := parseAccessToken(ctx)
		if err != nil {
			ctx.Error(err)
			ctx.Abort()
			return
		}

		claims, err := jwtUtil.Parse(accessToken)
		if err != nil {
			ctx.Error(apperror.NewUnauthorizedError())
			ctx.Abort()
			return
		}

		ctx.Set(constant.JWT_USER_ID, claims.UserID)
		ctx.Next()
	}
}

func parseAccessToken(ctx *gin.Context) (string, error) {
	accessToken, err := ctx.Cookie(constant.COOKIE_ACCESS_TOKEN)
	if err != nil {
		return "", apperror.NewUnauthorizedError()
	}
	if accessToken == "" || len(accessToken) == 0 {
		return "", apperror.NewUnauthorizedError()
	}

	return accessToken, nil
}

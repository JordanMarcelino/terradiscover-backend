package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jordanmarcelino/terradiscover-backend/internal/apperror"
	"github.com/jordanmarcelino/terradiscover-backend/internal/constant"
	"github.com/jordanmarcelino/terradiscover-backend/internal/dto"
	"github.com/jordanmarcelino/terradiscover-backend/internal/usecase"
	"github.com/jordanmarcelino/terradiscover-backend/internal/utils/ginutils"
)

type AuthController struct {
	authUseCase usecase.AuthUseCase
}

func NewAuthController(authUseCase usecase.AuthUseCase) *AuthController {
	return &AuthController{
		authUseCase: authUseCase,
	}
}

func (c *AuthController) Register(ctx *gin.Context) {
	req := new(dto.UserRegisterRequest)
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctx.Error(err)
		return
	}

	res, err := c.authUseCase.Register(ctx, req)
	if err != nil {
		ctx.Error(err)
		return
	}

	ginutils.ResponseCreated(ctx, res)
}

func (c *AuthController) Login(ctx *gin.Context) {
	req := new(dto.UserLoginRequest)
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctx.Error(err)
		return
	}

	token, err := c.authUseCase.Login(ctx, req)
	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.SetCookie(constant.COOKIE_ACCESS_TOKEN, token, 86400, "/", "", false, false)

	ginutils.ResponseOKPlain(ctx)
}

func (c *AuthController) Logout(ctx *gin.Context) {
	token, err := ctx.Cookie(constant.COOKIE_ACCESS_TOKEN)
	if err != nil {
		ctx.Error(apperror.NewUnauthorizedError())
		return
	}

	ctx.SetCookie(constant.COOKIE_ACCESS_TOKEN, token, -1, "/", "", false, false)
	ginutils.ResponseOKPlain(ctx)
}

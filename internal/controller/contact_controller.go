package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jordanmarcelino/terradiscover-backend/internal/dto"
	"github.com/jordanmarcelino/terradiscover-backend/internal/usecase"
	"github.com/jordanmarcelino/terradiscover-backend/internal/utils/ginutils"
)

type ContactController struct {
	contactUseCase usecase.ContactUseCase
}

func NewContactController(contactUseCase usecase.ContactUseCase) *ContactController {
	return &ContactController{
		contactUseCase: contactUseCase,
	}
}

func (c *ContactController) Create(ctx *gin.Context) {
	req := new(dto.CreateContactRequest)
	if err := ctx.ShouldBindJSON(req); err != nil {
		ctx.Error(err)
		return
	}

	res, err := c.contactUseCase.Save(ctx, req)
	if err != nil {
		ctx.Error(err)
		return
	}

	ginutils.ResponseCreated(ctx, res)
}

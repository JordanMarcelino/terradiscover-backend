package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/jordanmarcelino/terradiscover-backend/internal/constant"
	"github.com/jordanmarcelino/terradiscover-backend/internal/dto"
	"github.com/jordanmarcelino/terradiscover-backend/internal/usecase"
	"github.com/jordanmarcelino/terradiscover-backend/internal/utils/ginutils"
	"github.com/jordanmarcelino/terradiscover-backend/internal/utils/pageutils"
)

type ContactController struct {
	contactUseCase usecase.ContactUseCase
}

func NewContactController(contactUseCase usecase.ContactUseCase) *ContactController {
	return &ContactController{
		contactUseCase: contactUseCase,
	}
}

func (c *ContactController) Search(ctx *gin.Context) {
	req := &dto.SearchContactRequest{UserID: ctx.GetInt64(constant.JWT_USER_ID)}
	if err := ctx.ShouldBindQuery(req); err != nil {
		ctx.Error(err)
		return
	}

	res, paging, err := c.contactUseCase.Search(ctx, req)
	if err != nil {
		ctx.Error(err)
		return
	}

	paging.Links = pageutils.CreateLinks(ctx.Request, int(paging.Page), int(paging.Size), int(paging.TotalItem), int(paging.TotalPage))
	ginutils.ResponseOKPagination(ctx, res, paging)
}

func (c *ContactController) Create(ctx *gin.Context) {
	req := &dto.CreateContactRequest{UserID: ctx.GetInt64(constant.JWT_USER_ID)}
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

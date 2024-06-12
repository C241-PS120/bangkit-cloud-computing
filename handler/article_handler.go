package handler

import (
	"context"
	"github.com/C241-PS120/bangkit-cloud-computing/helper"
	"time"

	"github.com/C241-PS120/bangkit-cloud-computing/dto"
	"github.com/C241-PS120/bangkit-cloud-computing/dto/converter"
	"github.com/C241-PS120/bangkit-cloud-computing/model"
	"github.com/C241-PS120/bangkit-cloud-computing/repository"
	"github.com/gofiber/fiber/v2"
)

type ArticleHandler struct {
	Repository repository.ArticleRepository
}

func NewArticleHandler(repo repository.ArticleRepository) *ArticleHandler {
	return &ArticleHandler{
		Repository: repo,
	}
}

func (h *ArticleHandler) GetArticleDetail(ctx *fiber.Ctx) error {
	timeoutCtx, cancel := context.WithTimeout(ctx.UserContext(), 5*time.Second)
	defer cancel()

	id, err := ctx.ParamsInt("id", 0)
	if id < 1 {
		return fiber.NewError(fiber.StatusBadRequest, "id is not valid, must be greater than 0")
	}
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "id is not valid, must be an integer")
	}

	var article model.Article
	err = h.Repository.GetArticleDetail(timeoutCtx, id, &article)
	if err != nil {
		return helper.HandleRequestError(err)
	}

	return ctx.JSON(dto.Envelope{
		Success: true,
		Data:    converter.ArticleToResponse(&article),
	})
}

func (h *ArticleHandler) GetArticleList(ctx *fiber.Ctx) error {
	timeoutCtx, cancel := context.WithTimeout(ctx.UserContext(), 5*time.Second)
	defer cancel()

	var articles []model.Article
	err := h.Repository.GetArticleList(timeoutCtx, &articles)
	if err != nil {
		return helper.HandleRequestError(err)
	}

	var response []interface{}
	for _, article := range articles {
		response = append(response, converter.ArticleToResponse(&article))
	}

	return ctx.JSON(dto.Envelope{
		Success: true,
		Data:    response,
	})
}

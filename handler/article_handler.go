package handler

import (
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
	id, err := ctx.ParamsInt("id", 0)
	if id < 1 {
		return fiber.NewError(fiber.StatusBadRequest, "id is not valid, must be greater than 0")
	}
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "id is not valid, must be an integer")
	}

	var article model.Article
	if err := h.Repository.GetArticleDetail(id, &article); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(dto.Envelope{
		Success: true,
		Data:    converter.ArticleToResponse(&article),
	})
}

func (h *ArticleHandler) GetArticleList(ctx *fiber.Ctx) error {
	var articles []model.Article
	if err := h.Repository.GetArticleList(&articles); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
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

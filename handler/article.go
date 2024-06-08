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
	if id == 0 {
		return err // bad request, id not found
	}
	if err != nil {
		return err // bad request, not a string
	}

	var article model.Article
	if err := h.Repository.GetArticleDetail(id, &article); err != nil {
		return err // dunno
	}

	return ctx.JSON(dto.Envelope{
		Success: true,
		Data:    converter.ArticleToResponse(&article),
	})
}

func (h *ArticleHandler) GetArticleList(ctx *fiber.Ctx) error {
	var articles []model.Article
	if err := h.Repository.GetArticleList(&articles); err != nil {
		return err
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

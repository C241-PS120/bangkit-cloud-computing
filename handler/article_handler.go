package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"strings"
	"time"

	"github.com/C241-PS120/bangkit-cloud-computing/helper"

	"github.com/C241-PS120/bangkit-cloud-computing/dto"
	"github.com/C241-PS120/bangkit-cloud-computing/dto/converter"
	"github.com/C241-PS120/bangkit-cloud-computing/model"
	"github.com/C241-PS120/bangkit-cloud-computing/repository"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
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

	id, err := helper.GetIdFromRequest(ctx)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var article model.Article
	err = h.Repository.GetArticleById(timeoutCtx, id, &article)
	if err != nil {
		log.Error(err)
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
		log.Error(err)
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

func (h *ArticleHandler) CreateArticle(ctx *fiber.Ctx) error {
	timeoutCtx, cancel := context.WithTimeout(ctx.UserContext(), 10*time.Second)
	defer cancel()
	// upload an image to Google Bucket
	form, err := ctx.MultipartForm()
	if err != nil {
		log.Error(err)
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	jsonStr := form.Value["json"]
	if len(jsonStr) == 0 {
		return fiber.NewError(fiber.StatusBadRequest, "Article data JSON is required")
	}

	var request dto.ArticleRequest
	err = json.Unmarshal([]byte(jsonStr[0]), &request)
	if err != nil {
		log.Error(err)
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	validate := validator.New()
	err = validate.Struct(request)
	if err != nil {
		log.Error(err)
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	article := converter.RequestToArticle(&request)
	// upload an image to Google Bucket
	images := form.File["image"]
	if len(images) == 0 {
		return fiber.NewError(fiber.StatusBadRequest, "Image is required")
	}

	image, err := images[0].Open()
	if err != nil {
		log.Error(err)
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}
	defer image.Close()

	uploader, err := helper.NewClientUploader(timeoutCtx)
	if err != nil {
		log.Error(err)
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	fileName := strings.ToLower(strings.ReplaceAll(article.Title, " ", "_"))
	imageURL, err := uploader.UploadFile(timeoutCtx, image, fileName)
	if err != nil {
		log.Error(err)
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	article.ImageURL = imageURL
	err = h.Repository.CreateArticle(timeoutCtx, article)
	if err != nil {
		log.Error(err)
		return helper.HandleRequestError(err)
	}

	return ctx.Status(fiber.StatusCreated).JSON(dto.Envelope{
		Success: true,
		Message: fmt.Sprintf("Article created successfully with ID %d", article.ArticleID),
	})
}

func (h *ArticleHandler) UpdateArticle(ctx *fiber.Ctx) error {
	timeoutCtx, cancel := context.WithTimeout(ctx.UserContext(), 10*time.Second)
	defer cancel()

	id, err := helper.GetIdFromRequest(ctx)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	var request dto.ArticleRequest
	err = ctx.BodyParser(&request)
	if err != nil {
		log.Error(err)
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	validate := validator.New()
	err = validate.Struct(request)
	if err != nil {
		log.Error(err)
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	article := converter.RequestToArticle(&request)
	article.ArticleID = id

	err = h.Repository.UpdateArticle(timeoutCtx, article)
	if err != nil {
		log.Error(err)
		return helper.HandleRequestError(err)
	}

	return ctx.Status(fiber.StatusNoContent).JSON(dto.Envelope{Message: "Article updated successfully"})
}

func (h *ArticleHandler) DeleteArticle(ctx *fiber.Ctx) error {
	timeoutCtx, cancel := context.WithTimeout(ctx.UserContext(), 5*time.Second)
	defer cancel()

	id, err := helper.GetIdFromRequest(ctx)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	err = h.Repository.DeleteArticle(timeoutCtx, id)
	if err != nil {
		log.Error(err)
		return helper.HandleRequestError(err)
	}

	return ctx.Status(fiber.StatusNoContent).JSON(dto.Envelope{Message: "Article deleted successfully"})
}

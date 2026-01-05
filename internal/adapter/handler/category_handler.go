package handler

import (
	"manews/internal/adapter/handler/response"
	"manews/internal/core/domain/entity"
	"manews/internal/core/service"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

var defaultSuccessResponse response.DefaultSuccessResponse

type CateforyHandler interface {
	GetCategories(c *fiber.Ctx) error
	GetCategoryByID(c *fiber.Ctx) error
	CreateCategory(c *fiber.Ctx) error
	EditCategory(c *fiber.Ctx) error
	DeleteCategory(c *fiber.Ctx) error
}

type categoryHandler struct {
	categoryService service.CategoryService
}

// CreateCategory implements CateforyHandler.
func (ch *categoryHandler) CreateCategory(c *fiber.Ctx) error {
	panic("unimplemented")
}

// DeleteCategory implements CateforyHandler.
func (ch *categoryHandler) DeleteCategory(c *fiber.Ctx) error {
	panic("unimplemented")
}

// EditCategory implements CateforyHandler.
func (ch *categoryHandler) EditCategory(c *fiber.Ctx) error {
	panic("unimplemented")
}

// GetCategories implements CateforyHandler.
func (ch *categoryHandler) GetCategories(c *fiber.Ctx) error {
	claims := c.Locals("user").(*entity.JwtData)
	userID := claims.UserID

	if userID == 0 {
		code = "[HANDLER] GetCategories - 1"
		log.Errorw(code, err)
		errorResp.Meta.Status = false
		errorResp.Meta.Message = "Unauthorized access"
		return c.Status(fiber.StatusUnauthorized).JSON(errorResp)
	}

	results, err := ch.categoryService.GetCategories(c.Context())
	if err != nil {
		code = "[HANDLER] GetCategories - 2"
		log.Errorw(code, err)
		errorResp.Meta.Status = false
		errorResp.Meta.Message = err.Error()

		return c.Status(fiber.StatusInternalServerError).JSON(errorResp)
	}

	categoryResponses := []response.SuccessCategoryResponse{}
	for _, result := range results {
		categoryResponse := response.SuccessCategoryResponse{
			ID:            result.ID,
			Title:         result.Title,
			Slug:          result.Slug,
			CreatedByName: result.User.Name,
		}
		categoryResponses = append(categoryResponses, categoryResponse)
	}

	defaultSuccessResponse.Meta.Status = true
	defaultSuccessResponse.Meta.Message = "Categories fetched successfully"
	defaultSuccessResponse.Data = categoryResponses

	return c.JSON(defaultSuccessResponse)
}

// GetCategoryByID implements CateforyHandler.
func (ch *categoryHandler) GetCategoryByID(c *fiber.Ctx) error {
	panic("unimplemented")
}

func NewCategoryHandler(categoryService service.CategoryService) CateforyHandler {
	return &categoryHandler{
		categoryService: categoryService,
	}
}

package handler

import (
	"manews/internal/core/service"

	"github.com/gofiber/fiber/v2"
)

type ContentHandler interface {
	GetContents(c *fiber.Ctx) error
	GetContentByID(c *fiber.Ctx) error
	CreateContent(c *fiber.Ctx) error
	EditContent(c *fiber.Ctx) error
	DeleteContent(c *fiber.Ctx) error
	upladImageR2(c *fiber.Ctx) error
}

type contentHandler struct {
	contentService service.ContentService
}

// CreateContent implements ContentHandler.
func (*contentHandler) CreateContent(c *fiber.Ctx) error {
	panic("unimplemented")
}

// DeleteContent implements ContentHandler.
func (*contentHandler) DeleteContent(c *fiber.Ctx) error {
	panic("unimplemented")
}

// EditContent implements ContentHandler.
func (*contentHandler) EditContent(c *fiber.Ctx) error {
	panic("unimplemented")
}

// GetContentByID implements ContentHandler.
func (*contentHandler) GetContentByID(c *fiber.Ctx) error {
	panic("unimplemented")
}

// GetContents implements ContentHandler.
func (*contentHandler) GetContents(c *fiber.Ctx) error {
	panic("unimplemented")
}

// upladImageR2 implements ContentHandler.
func (*contentHandler) upladImageR2(c *fiber.Ctx) error {
	panic("unimplemented")
}

func NewContentHandler(contentService service.ContentService) ContentHandler {
	return &contentHandler{
		contentService: contentService,
	}
}

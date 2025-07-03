package handler

import (
	"ai_marketplace/internal/ai"
	"ai_marketplace/internal/model"
	"ai_marketplace/internal/search"
	"github.com/gofiber/fiber/v2"
	"log"
)

type handler struct {
	AI     *ai.Serviсe
	Search *search.Serviсe
}

func NewAIHandler(aiService *ai.Serviсe, searchService *search.Serviсe) *handler {
	return &handler{
		AI:     aiService,
		Search: searchService,
	}
}

func (h *handler) HandleSuggest(c *fiber.Ctx) error {
	log.Println("handler triggered")
	var req model.SuggestRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request"})
	}

	log.Println("user query:", req.Query)

	suggestion := h.AI.GenerateSuggestions(req.Query)

	log.Println("suggestions:", suggestion)

	var allResults []search.ProductSuggestion
	for _, name := range suggestion {
		products := h.Search.FindGoogleProducts(name)
		allResults = append(allResults, products...)
	}

	return c.JSON(fiber.Map{"results": allResults})
}

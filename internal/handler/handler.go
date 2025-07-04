package handler

import (
	"ai_marketplace/internal/ai"
	"ai_marketplace/internal/model"
	"ai_marketplace/internal/search"
	"bytes"
	"github.com/gofiber/fiber/v2"
	"html/template"
)

type handler struct {
	AI     *ai.Serviсe
	Search *search.Serviсe
}

type BubbleData struct {
	Title       string
	Link        string
	ImageURL    string
	Price       string
	RandPercent string
	RandDelay   string
}

func NewAIHandler(aiService *ai.Serviсe, searchService *search.Serviсe) *handler {
	return &handler{
		AI:     aiService,
		Search: searchService,
	}
}

func (h *handler) HandleSuggest(c *fiber.Ctx) error {
	var req model.SuggestRequest

	if req.Query == "" {
		req.Query = c.FormValue("query")
	}

	if req.Query == "" {
		return c.Status(fiber.StatusBadRequest).SendString("Запит порожній")
	}

	suggestion := h.AI.GenerateSuggestions(req.Query)

	var allResults []search.ProductSuggestion
	for _, name := range suggestion {
		products := h.Search.FindGoogleProducts(name)
		allResults = append(allResults, products...)
	}

	tpl, err := template.ParseFiles("template/results.html")
	if err != nil {
		return c.Status(500).SendString("template error: " + err.Error())
	}

	var buf bytes.Buffer
	if err := tpl.Execute(&buf, allResults); err != nil {
		return c.Status(500).SendString("render error: " + err.Error())
	}

	return c.Type("html").Send(buf.Bytes())
}

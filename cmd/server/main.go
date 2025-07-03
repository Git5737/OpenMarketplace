package main

import (
	"ai_marketplace/internal/ai"
	"ai_marketplace/internal/config"
	"ai_marketplace/internal/handler"
	"ai_marketplace/internal/search"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		log.Fatal("failed to load config", err)
	}

	app := fiber.New()

	aiServise := ai.New(cfg)
	searchServise := search.New(cfg)
	handler := handler.NewAIHandler(aiServise, searchServise)

	app.Post("/ai/suggest", handler.HandleSuggest)

	app.Listen(":" + cfg.Port)
}

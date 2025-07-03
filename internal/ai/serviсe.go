package ai

import (
	"ai_marketplace/internal/config"
	"context"
	"fmt"
	"google.golang.org/genai"
	"log"
	"strings"
)

type Serviсe struct {
	apiKey string
}

func New(cfg *config.Config) *Serviсe {
	return &Serviсe{
		apiKey: cfg.GeminiAPIKey,
	}
}

func (s *Serviсe) GenerateSuggestions(query string) []string {
	apiKey := s.apiKey
	if apiKey == "" {
		return []string{"[API key not set]"}
	}

	ctx := context.Background()
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		Backend: genai.BackendGeminiAPI,
		APIKey:  apiKey,
	})
	if err != nil {
		return []string{"[API client error]"}
	}

	prompt := fmt.Sprintf(`Ти — AI помічник у маркетплейсі. Користувач написав: "%s". На основі опису запропонуй 3-5 назв товарів (товар це річ яка продається на маркетплейсах чи інших подібних сервісах ти розрахований на ринок України). Відповідь: список через кому. Без лапок, без крапки в кінці. Без вступу.`, query)

	result, err := client.Models.GenerateContent(
		ctx,
		"gemini-1.5-flash",
		genai.Text(prompt),
		nil,
	)
	if err != nil {
		log.Println("generation error:", err)
		return []string{"[generation error]"}
	}

	raw := result.Text()
	parts := strings.Split(raw, ",")
	for i := range parts {
		parts[i] = strings.TrimSpace(parts[i])
	}

	fmt.Println(parts)
	return parts
}

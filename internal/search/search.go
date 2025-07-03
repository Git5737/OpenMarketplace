package search

import (
	"ai_marketplace/internal/config"
	"encoding/json"
	"log"
	"net/http"
	"net/url"
)

type ProductSuggestion struct {
	Title    string `json:"title"`
	Link     string `json:"link"`
	ImageURL string `json:"image_url"`
	//Description string `json:"description"`
}

type Serviсe struct {
	apiKey string
	cx     string
}

func New(cfg *config.Config) *Serviсe {
	return &Serviсe{
		apiKey: cfg.GoogleSearchAPIKey,
		cx:     cfg.GoogleSearchCX,
	}
}

func (s *Serviсe) FindGoogleProducts(query string) []ProductSuggestion {
	apiKey := s.apiKey
	cx := s.cx

	baseURL := "https://www.googleapis.com/customsearch/v1"
	params := url.Values{}
	params.Set("key", apiKey)
	params.Set("cx", cx)
	params.Set("q", query)

	finalURL := baseURL + "?" + params.Encode()

	log.Println("URL:", finalURL)

	resp, err := http.Get(finalURL)
	if err != nil || resp.StatusCode != 200 {
		log.Printf("Google API error: %v, status: %d", err, resp.StatusCode)
		return nil
	}
	defer resp.Body.Close()

	var data struct {
		Items []struct {
			Title   string `json:"title"`
			Link    string `json:"link"`
			Snippet string `json:"snippet"`
			Pagemap struct {
				CseImage []struct {
					Src string `json:"src"`
				} `json:"cse_image"`
			} `json:"pagemap"`
		} `json:"items"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		log.Println("JSON decode error:", err)
		return nil
	}

	var suggestions []ProductSuggestion
	for _, item := range data.Items {
		imageURL := ""
		if len(item.Pagemap.CseImage) > 0 {
			imageURL = item.Pagemap.CseImage[0].Src
		}

		suggestions = append(suggestions, ProductSuggestion{
			Title: item.Title,
			Link:  item.Link,
			//Description: item.Snippet,
			ImageURL: imageURL,
		})
	}

	log.Printf("Found %d results\n", len(suggestions))
	return suggestions
}

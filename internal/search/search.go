package search

import (
	"ai_marketplace/internal/config"
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"time"
)

type ProductSuggestion struct {
	Title    string `json:"title"`
	Link     string `json:"link"`
	ImageURL string `json:"image_url"`
	Price    string `json:"price"`
	Top      int    `json:"top"`
	Left     int    `json:"left"`
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
	baseURL := "https://www.googleapis.com/customsearch/v1"
	params := url.Values{}
	params.Set("key", s.apiKey)
	params.Set("cx", s.cx)
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
			Pagemap struct {
				CseImage []struct {
					Src string `json:"src"`
				} `json:"cse_image"`
				Offer []struct {
					Price string `json:"price"`
				} `json:"offer"`
				MetaTags []map[string]string `json:"metatags"`
			} `json:"pagemap"`
		} `json:"items"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		log.Println("JSON decode error:", err)
		return nil
	}

	var suggestions []ProductSuggestion
	rand.Seed(time.Now().UnixNano())

	for _, item := range data.Items {
		img := ""
		price := ""

		// Отримати зображення
		if len(item.Pagemap.CseImage) > 0 {
			img = item.Pagemap.CseImage[0].Src
		}

		// Отримати ціну
		if len(item.Pagemap.Offer) > 0 {
			price = item.Pagemap.Offer[0].Price
		} else if len(item.Pagemap.MetaTags) > 0 {
			price = item.Pagemap.MetaTags[0]["product:price:amount"]
		}

		suggestions = append(suggestions, ProductSuggestion{
			Title:    item.Title,
			Link:     item.Link,
			ImageURL: img,
			Price:    price,
			Top:      rand.Intn(80), // 0–80% зверху
			Left:     rand.Intn(90),
		})
	}

	log.Printf("Found %d results\n", len(suggestions))
	return suggestions
}

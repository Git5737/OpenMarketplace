# OpenMarketplace💈

Open Marketplace — this is a simple marketplace that allows you to search for products using artificial intelligence (Gemini AI) and the Google Search API. Built on Go + Fiber, HTMX, TailwindCSS.

http://13.61.178.171:8080

<img width="1439" alt="Знімок екрана 2025-07-04 о 15 40 50" src="https://github.com/user-attachments/assets/6ae45e3c-4c59-4230-9308-cd4ec9c6c2c8" />

<img width="1395" alt="Знімок екрана 2025-07-04 о 15 40 10" src="https://github.com/user-attachments/assets/0120927c-0d87-47f6-9cee-5f30b59e9996" />



## Feature ⚙️

- ✨ Search products using AI (Google Gemini)
- 🔎 Search for relevant links through the Google Custom Search API
- 🧠 Generating product names based on user request
- 📦 Displaying the results in the form of beautiful cards


## Launch locally 🧪

### 1. Сloning
```bash
git clone https://github.com/your-username/ai-marketplace.git
cd ai-marketplace
```
### 2. Settings .env
```bash
PORT=8080
GOOGLE_SEARCH_API_KEY=your_google_key
GOOGLE_SEARCH_CX=your_cx
GEMINI_API_KEY=your_gemini_key
```

### 3. Run
```bash
  go run cmd/server/main.go
```

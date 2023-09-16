package open_ai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

const (
	OpenAIKeyEnv = "OPEN_AI_API_KEY"
)

type OpenAI struct {
	APIKey string
}

func NewOpenAI() *OpenAI {
	apiKey := os.Getenv(OpenAIKeyEnv)

	return &OpenAI{APIKey: apiKey}
}

func (o *OpenAI) GenerateSentences(englishWord string) ([]string, error) {
	endpoint := "https://api.openai.com/v1/chat/completions"

	payload := fmt.Sprintf(`{
		"model": "gpt-3.5-turbo",
		"messages": [{"role": "user", "content": "Generate 3 so short sentences & usually use in the life for this word? '%s'"}],
		"temperature": 0.7
	  }`, englishWord)

	// Create a new HTTP request
	req, err := http.NewRequest(http.MethodPost, endpoint, bytes.NewBuffer([]byte(payload)))
	if err != nil {
		return nil, err
	}

	// Set the request headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+o.APIKey)

	// Create an HTTP client
	client := &http.Client{}

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Decode the response JSON
	var responseData map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&responseData); err != nil {
		return nil, err
	}

	// Extract and parse the content into a list
	content := responseData["choices"].([]interface{})[0].(map[string]interface{})["message"].(map[string]interface{})["content"].(string)

	// remove 1. 2. 3. (numbering)
	for i := 1; i <= 3; i++ {
		content = strings.Replace(content, fmt.Sprintf("%d. ", i), "", 1)
	}

	contentList := strings.Split(content, "\n")

	return contentList, nil
}

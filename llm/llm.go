package llm

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/google/generative-ai-go/genai"
	"github.com/joho/godotenv"
	"google.golang.org/api/option"
)

type response_t struct {
	Answer tJudgement
}

const Judgement_t (
	Yes = iota
	No
  )



var (
	geminiKey = getKey("GEMINI_API_KEY")
)

func getKey(key_name string) string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	apiKey := os.Getenv(key_name)
	if apiKey == "" {
		log.Fatal("API_KEY not set in .env file")
	}
	return apiKey
}

// Function based upon examples at
// https://github.com/google/generative-ai-go
func CallGemini_struct(prompt string) {
	// Initialize model
	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(geminiKey))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	llm_model := client.GenerativeModel("gemini-2.0-flash")

	// Ask the model to respond with JSON.
	llm_model.ResponseMIMEType = "application/json"
	// Specify the schema.
	llm_model.ResponseSchema = &genai.Schema{
		Type: genai.TypeObject,
		Properties: response_t,
	}
	resp, err := llm_model.GenerateContent(ctx, genai.Text(prompt+" Using this JSON schema."))
	if err != nil {
		log.Fatal(err)
	}
	printResponse(resp)

}

// Print Gemini response helper function
func printResponse(resp *genai.GenerateContentResponse) string {
	var ret string
	for _, cand := range resp.Candidates {
		for _, part := range cand.Content.Parts {
			ret = ret + fmt.Sprintf("%v", part)
			fmt.Println(part)
		}
	}
	return ret
}

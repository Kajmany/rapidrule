package gemini_llm

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/google/generative-ai-go/genai"
	"github.com/joho/godotenv"
	"google.golang.org/api/option"
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

func CallGemini(prompt string) {
	geminiKey := getKey("GEMINI_API_KEY")

	// This example shows how to get a JSON response that conforms to a schema.
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
		Type:  genai.TypeArray,
		Items: &genai.Schema{Type: genai.TypeString},
	}
	resp, err := llm_model.GenerateContent(ctx, genai.Text("List a few popular cookie recipes using this JSON schema."))
	if err != nil {
		log.Fatal(err)
	}
	for _, part := range resp.Candidates[0].Content.Parts {
		if txt, ok := part.(genai.Text); ok {
			var recipes []string
			if err := json.Unmarshal([]byte(txt), &recipes); err != nil {
				log.Fatal(err)
			}
			fmt.Println(recipes)
		}
	}

}

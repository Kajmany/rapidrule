package llm

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/Kajmany/rapidrule/structs"
	"github.com/google/generative-ai-go/genai"
	"github.com/joho/godotenv"
	"google.golang.org/api/option"

	tea "github.com/charmbracelet/bubbletea"
)

// Eval message struct for evals
type PortEvalMsg struct {
	Evals []structs.Eval
}

type PortEvalError struct {
	Err error
}

// Eval message struct for overall evals
type TotalEvalMsg struct {
	TotalEval structs.TotalEval
}

type TotalEvalError struct {
	Err error
}

// Get API key from .env file
var (
	geminiKey = getKey("GEMINI_API_KEY")
)

// Helper function to get API key from .env file
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

// String method for PortEvalError
func (e PortEvalError) Error() string {
	return fmt.Sprintf("port eval error: %v", e.Err)
}

// Async get port Gemini evaluations command
func GetPortEvals(port_strings string) tea.Cmd {
	return func() tea.Msg {
		evals, err := GeminiPortsEval(port_strings)
		if err != nil {
			log.Println("problem evaluating ports with Gemini")
			return PortEvalError{Err: err}
		}
		return PortEvalMsg{Evals: evals}
	}
}

// Async get port Gemini evaluations command
func GetTotalEvals(ports_strings []string) tea.Cmd {
	return func() tea.Msg {
		total_eval, err := GeminiTotalEval(ports_strings)
		if err != nil {
			log.Println("problem evaluating overall posture with Gemini")
			return TotalEvalError{Err: err}
		}
		return TotalEvalMsg{TotalEval: total_eval}
	}
}

// Function based upon examples at
// https://github.com/google/generative-ai-go
func GeminiPortsEval(port_strings string) ([]structs.Eval, error) {
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
	response_schema := &genai.Schema{
		Type: genai.TypeObject,
		Properties: map[string]*genai.Schema{
			"concerns": {
				Type:        genai.TypeString,
				Description: "Describe possible concerns of this service and why they're concerning.",
			},
			"investigate": {
				Type:        genai.TypeString,
				Enum:        []string{"Yes", "Alert", "No"},
				Description: "Should this be investigated further. Yes, No, or just alert the user.",
			},
			"confidence": {
				Type:        genai.TypeString,
				Enum:        []string{"High", "Medium", "Low"},
				Description: "Confidence in choice of whether it should be investigated.",
			},
			"port": {
				Type:        genai.TypeInteger,
				Description: "What port is this service using?",
			},
		},
		Required: []string{"investigate", "confidence", "concerns", "port"},
	}

	llm_model.ResponseSchema = &genai.Schema{
		Type:  genai.TypeArray,
		Items: response_schema,
	}

	llm_request := "List out whether these services are suspicous: " + port_strings + " using this JSON schema:"

	resp, err := llm_model.GenerateContent(ctx, genai.Text(llm_request))
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Request: %s,\nResponse: %s", llm_request, stringifyResponse(resp))

	var evals []structs.Eval
	err = json.Unmarshal([]byte(stringifyResponse(resp)), &evals)
	return evals, err

}

func GeminiTotalEval(ports_strings []string) (structs.TotalEval, error) {
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
	response_schema := &genai.Schema{
		Type: genai.TypeObject,
		Properties: map[string]*genai.Schema{
			"overall": {
				Type:        genai.TypeString,
				Description: "Describe the overall security and possible concerns of a system with these open TCP ports.",
			},
			"alert": {
				Type:        genai.TypeString,
				Enum:        []string{"Red", "Yellow", "No"},
				Description: "Should the user be alerted of security of misconfiguration concerns. Respond with an urgent Red, less urgent Yellow, or no alert.",
			},
			"alert_message": {
				Type:        genai.TypeString,
				Description: "Describe in one sentence what concerns the user should be alerted of.",
			},
		},
		Required: []string{"overall", "alert", "alert_message"},
	}

	llm_model.ResponseSchema = &genai.Schema{
		Type:  genai.TypeArray,
		Items: response_schema,
	}

	llm_request := "Describe the overall security posture of a system with these open TCP ports: " +
		portsToprompt(ports_strings) +
		" using this JSON schema:"

	resp, err := llm_model.GenerateContent(ctx, genai.Text(llm_request))
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Request: %s,\nResponse: %s", llm_request, stringifyResponse(resp))

	var total_eval []structs.TotalEval
	err = json.Unmarshal([]byte(stringifyResponse(resp)), &total_eval)
	return total_eval[0], err

}

func portsToprompt(ports_strings []string) string {
	var prompt string
	for _, port_string := range ports_strings {
		prompt += port_string
	}
	return prompt
}

func stringifyResponse(resp *genai.GenerateContentResponse) string {
	var ret string
	for _, cand := range resp.Candidates {
		for _, part := range cand.Content.Parts {
			ret = ret + fmt.Sprintf("%v", part)
		}
	}
	return ret
}

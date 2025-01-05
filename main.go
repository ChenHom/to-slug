package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"strings"

	openai "github.com/sashabaranov/go-openai"
)

// 定義常數，表示可選的模型
const (
	ModelGpt4o     = "gpt-4o"
	ModelGpt4oMini = "gpt-4o-mini"
)

// 定義請求結構
type ChatCompletionRequest struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
}

// 定義訊息結構
type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// 定義回應結構
type ChatCompletionResponse struct {
	Choices []Choice `json:"choices"`
}

// 定義選擇結構
type Choice struct {
	Message Message `json:"message"`
}

var apiKey string
var client *openai.Client

func init() {
	apiKey = os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		fmt.Println("OPENAI_API_KEY is not set")
		os.Exit(1)
	}
	client = openai.NewClient(apiKey)
}

func main() {
	help := flag.Bool("h", false, "顯示使用說明")
	flag.Parse()

	if *help || len(os.Args) < 2 {
		fmt.Println("使用方法: to-slug \"需要翻譯的文字\"")
		fmt.Println("選項:")
		fmt.Println("  -h    顯示使用說明")
		os.Exit(0)
	}

	input := os.Args[1]

	// 翻譯輸入文字為英文
	translatedText := translateToEnglish(input)
	// 將翻譯後的英文轉換為 slug
	slug := translateToSlug(translatedText)

	fmt.Println("Slug: ", slug)
}

// 將文字翻譯為簡單的口語英文
func translateToEnglish(text string) string {
	req := openai.ChatCompletionRequest{
		Model: ModelGpt4o,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    "developer",
				Content: "Convert this text into simple spoken English, just give the result directly, no additional information is needed: " + text,
			},
		},
	}

	ctx := context.Background()
	resp, err := client.CreateChatCompletion(ctx, req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return ""
	}

	if len(resp.Choices) == 0 {
		fmt.Println("No choices returned from API")
		return ""
	}

	return strings.TrimSpace(resp.Choices[0].Message.Content)
}

// 將文字轉換為 URL slug
func translateToSlug(text string) string {
	req := openai.ChatCompletionRequest{
		Model: ModelGpt4o,
		Messages: []openai.ChatCompletionMessage{
			{
				Role: "developer",
				Content: "Make this English Slugify and adjust the number of words to simplify it." +
					" When responding, just give the result directly, no additional information is needed: " +
					text,
			},
		},
	}

	ctx := context.Background()
	resp, err := client.CreateChatCompletion(ctx, req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return ""
	}

	if len(resp.Choices) == 0 {
		fmt.Println("No choices returned from API")
		return ""
	}

	return strings.TrimSpace(resp.Choices[0].Message.Content)
}

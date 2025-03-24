package main

import (
	"context"
	"fmt"
	"os"

	openai "github.com/sashabaranov/go-openai"
)

func main() {

	apiKey := os.Getenv("OPENAI_API_KEY")
	baseURL := os.Getenv("OPENAI_BASE_URL")

	if apiKey == "" || baseURL == "" {
		fmt.Println("环境变量 OPENAI_API_KEY 或 OPENAI_BASE_URL 没有设置")
		return
	}

	config := openai.DefaultConfig(apiKey)
	//need"/v1"
	config.BaseURL = baseURL

	client := openai.NewClientWithConfig(config)
	resp, err := client.CreateTranscription(
		context.Background(),
		openai.AudioRequest{
			Model:    openai.Whisper1,
			FilePath: "test.mp3",
		},
	)
	if err != nil {
		fmt.Printf("Transcription error: %v\n", err)
		return
	}
	fmt.Println(resp.Text)
}

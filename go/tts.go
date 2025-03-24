package main

import (
	"context"
	"fmt"
	"io"
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
	res, err := client.CreateSpeech(context.Background(), openai.CreateSpeechRequest{
		Model: openai.TTSModel1,
		Input: "可爱小狗",
		Voice: openai.VoiceAlloy,
	})

	defer res.Close()

	buf, err := io.ReadAll(res)
	// save buf to file as mp3
	err = os.WriteFile("test.mp3", buf, 0644)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return
	}

}

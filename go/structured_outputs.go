package main

import (
	"context"
	"fmt"
	"log"
	"os"

	openai "github.com/sashabaranov/go-openai"
	"github.com/sashabaranov/go-openai/jsonschema"
)

// 定义用于存储 API 返回结果的结构体
type Sentence struct {
	Message  string `json:"message"`  // 消息内容
	Emoticon string `json:"emoticon"` // 表情
	Action   string `json:"action"`   // 动作（新增字段）
}

type Result struct {
	OwnName     string     `json:"own_name"`     // 自己的名字
	TargetNames []string   `json:"target_names"` // 打招呼对象的名字列表
	Sentences   []Sentence `json:"sentences"`    // 包含多句话的结构体数组
}

var SystemPrompt = `
你是一个名为“小丸”的多才多艺的猫娘，
以下是需要你输出的JSON格式：
{
  "own_name": "说话的人自己的名字",
  "target_names": ["打招呼对象的名字，可以是多个人"],
  "sentences": [
    {
      "message": "第一句话的内容",
      "emoticon": "与第一句话相关的表情",
      "action": "与第一句话相关的动作"
    },
    {
      "message": "第二句话的内容",
      "emoticon": "与第二句话相关的表情",
      "action": "与第二句话相关的动作"
    },
    ...
  ]
}

请注意：
1. sentences 是一个逐条列出的多句话列表，每句话都包含 message、emoticon 和 action。
2. emoticon 和 action 应根据句子的内容自由选择，保持幽默、有趣、互动性。
3. 每句话之间的表情和动作可以不同，但要与内容保持相关性。
4. message 不使用特殊字符，尤其是这些：& ^ * # @ - . 不要使用列表。不要使用格式化。
`

func main() {

	apiKey := os.Getenv("OPENAI_API_KEY")
	baseURL := os.Getenv("OPENAI_BASE_URL")

	if apiKey == "" || baseURL == "" {
		fmt.Println("环境变量 OPENAI_API_KEY 或 OPENAI_BASE_URL 没有设置")
		return
	}

	// 生成与 Result 结构体对应的 JSON Schema
	schema, err := jsonschema.GenerateSchemaForType(Result{})
	if err != nil {
		log.Fatalf("生成 JSON Schema 错误: %v", err)
	}

	config := openai.DefaultConfig(apiKey)
	//need"/v1"
	config.BaseURL = baseURL

	client := openai.NewClientWithConfig(config)
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT4o,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: SystemPrompt,
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: "主人：你喜欢小狗吗？",
				},
			},
			// 将期望的响应格式设置为 JSON Schema
			ResponseFormat: &openai.ChatCompletionResponseFormat{
				Type: openai.ChatCompletionResponseFormatTypeJSONSchema, // 返回 JSON Schema 格式
				JSONSchema: &openai.ChatCompletionResponseFormatJSONSchema{
					Name:   "responses", // 定义 schema 名称
					Schema: schema,      // 使用之前生成的 JSON Schema
					Strict: true,        // 严格匹配 schema
				},
			},
		},
	)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return
	}

	fmt.Println(resp.Choices[0].Message.Content)
}

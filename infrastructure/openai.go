package infrastructure

import "github.com/sashabaranov/go-openai"

func NewOpenAI(token string) *openai.Client {
	return openai.NewClient(token)
}

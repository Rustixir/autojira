package service

import (
	"context"
	"encoding/json"
	"github.com/sashabaranov/go-openai"
	"log"
	"smart/domain/model"
)

type AiService interface {
	ProcessIssue(content string) (*model.Issue, error)
}

type openAIService struct {
	client *openai.Client
}

func NewAiService(client *openai.Client) AiService {
	return openAIService{client}
}

func (ai openAIService) ProcessIssue(content string) (*model.Issue, error) {
	resp, err := ai.client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleAssistant,
					Content: content,
				},
			},
		},
	)

	if err != nil {
		log.Printf("ChatCompletion error: %v\n", err)
		return nil, err
	}

	issue := new(model.Issue)
	err = json.Unmarshal([]byte(resp.Choices[0].Message.Content), issue)
	if err != nil {
		return nil, err
	}
	return issue, nil
}

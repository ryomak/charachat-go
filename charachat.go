package charachat

import (
	"context"
	"errors"
	tokenizer "github.com/samber/go-gpt-3-encoder"
	openai "github.com/sashabaranov/go-openai"
)

type Charachat struct {
	personality     *Personality
	memory          Memory
	openAIClient    *openai.Client
	openAIModel     string
	openAIMaxTokens int
	encoder         *tokenizer.Encoder
}

func NewCharachat(
	openAIKey string,
	personality *Personality,
) (*Charachat, error) {
	encoder, err := tokenizer.NewEncoder()
	if err != nil {
		return nil, err
	}

	return &Charachat{
		personality:     personality,
		openAIClient:    openai.NewClient(openAIKey),
		memory:          NewLocalMemory(),
		openAIModel:     openai.GPT40613,
		openAIMaxTokens: 8000,
		encoder:         encoder,
	}, nil
}

func (c *Charachat) Talk(ctx context.Context, name, inputMessage string) (string, error) {
	initialToken := c.GetToken(inputMessage) + 1024
	messages := append(c.PastOpenAIMessages(c.personality.SystemPrompt(name), initialToken), openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleUser,
		Content: inputMessage,
	})

	res, err := c.openAIClient.CreateChatCompletion(ctx, openai.ChatCompletionRequest{
		Model:     openai.GPT40613,
		MaxTokens: 1024,
		Messages:  messages,
	})
	if err != nil {
		return "", err
	}
	if len(res.Choices) == 0 {
		return "", errors.New("no response")
	}
	c.memory.AddMessage(Message{
		Kind: MessageUserKindBot,
		Text: res.Choices[0].Message.Content,
	})
	return res.Choices[0].Message.Content, nil
}

func (c *Charachat) GetToken(content string) int {
	// TODO ChatGPTに対しては正確じゃないので代替案を考える
	encoded, err := c.encoder.Encode(content)

	if err != nil {
		return 0
	}

	return len(encoded)
}

func (c *Charachat) PastOpenAIMessages(systemPrompt string, token int) []openai.ChatCompletionMessage {
	maxToken := c.openAIMaxTokens - token - c.GetToken(systemPrompt)
	messages := c.memory.GetAllMessage()

	openAIMessages := make([]openai.ChatCompletionMessage, 0, len(messages))
	openAIMessages = append(openAIMessages, openai.ChatCompletionMessage{
		Role:    openai.ChatMessageRoleSystem,
		Content: systemPrompt,
	})

	for i := len(messages) - 1; i > 0; i-- {
		message := messages[i]
		maxToken = maxToken - c.GetToken(message.Text)
		if maxToken < 0 {
			break
		}
		role := openai.ChatMessageRoleUser
		if message.Kind == MessageUserKindBot {
			role = openai.ChatMessageRoleAssistant
		}
		openAIMessages = append(openAIMessages, openai.ChatCompletionMessage{
			Role:    role,
			Content: message.Text,
		})
	}
	return openAIMessages
}

func (c *Charachat) ResetConversation() error {
	c.memory.DeleteAll()
	return nil
}

type Option func(*Charachat) error

func WithPersonality(personality *Personality) Option {
	return func(c *Charachat) error {
		c.personality = personality
		return nil
	}
}

func WithMemory(memory Memory) Option {
	return func(c *Charachat) error {
		c.memory = memory
		return nil
	}
}
func WithOpenAIModel(openAIModel string) Option {
	return func(c *Charachat) error {
		c.openAIModel = openAIModel
		return nil
	}
}

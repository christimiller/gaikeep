package cmd

import (
	"context"
	"fmt"
	"log"
	"os"

	openai "github.com/sashabaranov/go-openai"
	"github.com/spf13/cobra"
)

var testgptCmd = &cobra.Command{
	Use:   "testgpt",
	Short: "Test ChatGPT interaction",
	Long:  `This command sends a test prompt to ChatGPT and displays the response.`,
	Run: func(cmd *cobra.Command, args []string) {
		apiKey := os.Getenv("OPENAI_API_KEY")
		if apiKey == "" {
			log.Fatal("OPENAI_API_KEY must be set in the environment variables")
		}

		client := openai.NewClient(apiKey)
		resp, err := client.CreateChatCompletion(
			context.Background(),
			openai.ChatCompletionRequest{
				Model: openai.GPT3Dot5Turbo, // Ensure to use the correct model for your needs
				Messages: []openai.ChatCompletionMessage{
					{
						Role:    openai.ChatMessageRoleUser,
						Content: "Hello, how can I help you today?",
					},
				},
			},
		)

		if err != nil {
			log.Fatalf("Failed to get response from OpenAI: %v", err)
			return
		}

		fmt.Printf("Response from ChatGPT: %s\n", resp.Choices[0].Message.Content)
	},
}

func init() {
	rootCmd.AddCommand(testgptCmd)
}

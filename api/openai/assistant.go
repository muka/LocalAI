package openai

import (
	config "github.com/go-skynet/LocalAI/api/config"
	"github.com/go-skynet/LocalAI/api/options"
	"github.com/gofiber/fiber/v2"
)

// https://platform.openai.com/docs/api-reference/completions
func CreateAssistantEndpoint(cm *config.ConfigLoader, o *options.Option) func(c *fiber.Ctx) error {
	return nil
}

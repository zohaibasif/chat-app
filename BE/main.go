package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/pusher/pusher-http-go"
)

func main() {
	app := fiber.New()

	app.Use(cors.New())

	app.Post("/api/messages", func(c *fiber.Ctx) error {
		data := map[string]string{}

		if err := c.BodyParser(&data); err != nil {
			return err
		}

		pusherClient := pusher.Client{
			AppID:   "1552487",
			Key:     "bcc34fc5bfaef1df32bc",
			Secret:  "bd7fa11891ad089ef0e1",
			Cluster: "ap2",
			Secure:  true,
		}

		pusherClient.Trigger("simple-chat-app", "message", data)

		return c.JSON([]string{})
	})

	app.Listen(":8000")
}

package utils

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
)

func SendJSON(c *fiber.Ctx, res interface{}) error {
	return json.NewEncoder(c.Type("json", "utf-8").Response().BodyWriter()).Encode(res)
}

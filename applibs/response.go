package applibs

import (
	"github.com/IshaqNiloy/go-rest-api/model"
	"github.com/gofiber/fiber/v2"
)

func Response(c *fiber.Ctx, codeObject CodeObject, status int, data *model.Products) error {
	err := c.Status(status).JSON(&fiber.Map{
		"lang":    codeObject.Lang,
		"message": codeObject.Message,
		"data":    data,
	})

	if err != nil {
		return err
	}
	return nil
}

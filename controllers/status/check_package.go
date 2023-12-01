package status

import (
	"github.com/ErikMora/quest2/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func (s *StatusController) CheckPackage(c *fiber.Ctx) error {
	ctx := c.Context()
	var result models.Packages
	order_id := c.Params("order_id")
	filter := bson.M{"identifier": order_id}

	if err := s.collection.FindOne(ctx, filter).Decode(&result); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Error": err.Error()})
	}

	return c.JSON(result)
}

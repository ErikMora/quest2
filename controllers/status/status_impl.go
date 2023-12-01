package status

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

type StatusController struct {
	collection  *mongo.Collection
	size_list   map[int]string
	status_list map[int]string
}

type StatusControllerInterface interface {
	ChangeStatus(c *fiber.Ctx) error
	CheckPackage(c *fiber.Ctx) error
}

func NewStatusController(collection *mongo.Collection, size_list map[int]string, status_list map[int]string) StatusControllerInterface {
	return &StatusController{
		collection:  collection,
		size_list:   size_list,
		status_list: status_list,
	}
}

package order

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

type OrderController struct {
	collection  *mongo.Collection
	size_list   map[int]string
	status_list map[int]string
}
type OrderControllerInterface interface {
	CreateOrder(c *fiber.Ctx) error
}

func NewOrderController(collection *mongo.Collection, size_list map[int]string, status_list map[int]string) OrderControllerInterface {
	return &OrderController{
		collection:  collection,
		size_list:   size_list,
		status_list: status_list,
	}
}

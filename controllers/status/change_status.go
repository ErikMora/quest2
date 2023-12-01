package status

import (
	"strconv"
	"time"

	"github.com/ErikMora/quest2/helper"
	"github.com/ErikMora/quest2/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

func (s *StatusController) ChangeStatus(c *fiber.Ctx) error {
	ctx := c.Context()
	var result models.Packages
	new_status := c.Params("new_status")
	order_id := c.Params("order_id")
	reembolso, err := strconv.ParseBool(c.Params("reembolso"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Error reembolso": err.Error()})
	}

	if !helper.ValidateSizeAndStatus(new_status, s.status_list) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Error": "Estatus no valido"})
	}

	filter := bson.M{"identifier": order_id}

	if err := s.collection.FindOne(ctx, filter).Decode(&result); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Error": err.Error()})
	}

	old_status := result.Status
	fecha_creacion := result.CreatedAt

	if (old_status == "en_ruta" || old_status == "entregado") && new_status == "cancelado" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Error": "El paquete no puede ser cancelado"})
	}

	if reembolso && new_status == "cancelado" {
		actual_date := time.Now()
		diff := actual_date.Sub(fecha_creacion)
		minutos := int(diff.Minutes())
		if minutos > 2 {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Error": "El reembolso ya no puede ser aplicado"})
		}
	}

	data := bson.M{
		"$set": bson.M{
			"status": new_status,
		},
	}

	_, err = s.collection.UpdateOne(ctx, filter, data)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Error": err.Error()})
	}

	if reembolso && new_status == "cancelado" {
		return c.JSON(fiber.Map{"success": "El paquete cambio a estatus " + new_status + " con reembolso"})
	} else if !reembolso && new_status == "cancelado" {
		return c.JSON(fiber.Map{"success": "El paquete cambio a estatus " + new_status + " sin reembolso"})
	} else {
		return c.JSON(fiber.Map{"success": "El paquete cambio a estatus " + new_status})
	}
}

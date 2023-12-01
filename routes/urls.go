package routes

import (
	"github.com/ErikMora/quest2/controllers/order"
	"github.com/ErikMora/quest2/controllers/status"
	"github.com/ErikMora/quest2/database"
	"github.com/gofiber/fiber/v2"
)

var collection = database.GetCollection("prueba2")
var size_list = map[int]string{0: "S", 1: "M", 2: "L"}
var status_list = map[int]string{0: "creado", 1: "recolectado", 2: "en_estacion", 3: "en_ruta", 4: "entregado", 5: "cancelado"}

func UrlRoutes(router fiber.Router) {
	orderController := order.NewOrderController(collection, size_list, status_list)
	statusController := status.NewStatusController(collection, size_list, status_list)
	router.Post("/api/new_order/:org_longitude/:org_latitude/:org_zipcode/:org_ext_num/:org_int_num/:dst_longitude/:dst_latitude/:dst_zipcode/:dst_ext_num/:dst_int_num/:size/:weight/:amount", orderController.CreateOrder)
	router.Post("/api/chage_status/:order_id/:new_status/:reembolso", statusController.ChangeStatus)
	router.Get("/api/CheckPackage/:order_id", statusController.CheckPackage)
}

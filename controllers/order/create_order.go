package order

import (
	"strconv"
	"strings"
	"time"

	"github.com/ErikMora/quest2/helper"
	"github.com/ErikMora/quest2/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (t *OrderController) CreateOrder(c *fiber.Ctx) error {
	m := c.Queries()
	org_longitude := c.Params("org_longitude")
	org_latitude := c.Params("org_latitude")
	org_zipcode := c.Params("org_zipcode")
	org_ext_num := c.Params("org_ext_num")
	org_int_num := c.Params("org_int_num")
	org_address := m["org_address"]

	dst_longitude := c.Params("dst_longitude")
	dst_latitude := c.Params("dst_latitude")
	dst_zipcode := c.Params("dst_zipcode")
	dst_ext_num := c.Params("dst_ext_num")
	dst_int_num := c.Params("dst_int_num")
	dst_address := m["dst_address"]

	size := strings.ToUpper(c.Params("size"))
	amount, _ := strconv.ParseInt(c.Params("amount"), 10, 0)
	weight, _ := strconv.ParseFloat(c.Params("weight"), 64)

	if !helper.ValidateSizeAndStatus(size, t.size_list) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Error": "Wrong Size"})
	}

	switch size {
	case "S":
		if weight <= 0 {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Error": "El peso no es valido"})
		} else if weight > 5 {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Error": "El peso sobrepasa el limite del tamaño S"})
		}
	case "M":
		if weight <= 5 {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Error": "El peso es menor al requerido"})
		} else if weight > 15 {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Error": "El peso sobrepasa el limite del tamaño M"})
		}
	case "L":
		if weight <= 15 {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Error": "El peso es menor al requerido"})
		} else if weight > 25 {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Error": "No cuenta con el servicio estándar para el peso elegido, favor de comunicarse con la empresa para realizar un convenio especial"})
		}
	}

	/////////////////// LONGITUDE VALIDATION
	org_lgt := strings.Split(org_longitude, ".")
	dst_lgt := strings.Split(dst_longitude, ".")

	org_lgt_int, _ := strconv.ParseInt(org_lgt[0], 10, 0)
	dst_lgt_int, _ := strconv.ParseInt(dst_lgt[0], 10, 0)

	if org_lgt_int < -180 || org_lgt_int > 180 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Error": "Longitud origen invalido"})
	}

	if dst_lgt_int < -180 || dst_lgt_int > 180 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Error": "Longitud destino invalido"})
	}
	////////////////// LATITUDE VALIDATION
	org_ltd := strings.Split(org_latitude, ".")
	dts_ltd := strings.Split(dst_latitude, ".")

	org_ltd_int, _ := strconv.ParseInt(org_ltd[0], 10, 0)
	dts_ltd_int, _ := strconv.ParseInt(dts_ltd[0], 10, 0)

	if org_ltd_int < -90 || org_ltd_int > 90 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Error": "Latitud origen invalido"})
	}

	if dts_ltd_int < -90 || dts_ltd_int > 90 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Error": "Latitud destino invalido"})
	}

	//////////// 	ORIGIN
	if len(org_address) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Error": "Ingrese Address origen"})
	}

	org_zipcode_, err := strconv.ParseInt(org_zipcode, 10, 0)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Error": "Ingrese zip code origen"})
	}

	org_int_num_, err := strconv.ParseInt(org_int_num, 10, 0)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Error": "Ingrese No Interno origen"})
	}

	org_ext_num_, err := strconv.ParseInt(org_ext_num, 10, 0)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Error": "Ingrese No Exterior origen"})
	}

	//////////////// DESTINY
	if len(dst_address) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Error": "Ingrese Address destino"})
	}

	dst_zipcode_, err := strconv.ParseInt(dst_zipcode, 10, 0)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Error": "Ingrese zip code destino"})
	}

	dst_int_num_, err := strconv.ParseInt(dst_int_num, 10, 0)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Error": "Ingrese No Interno destino"})
	}

	dst_ext_num_, err := strconv.ParseInt(dst_ext_num, 10, 0)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Error": "Ingrese No Exterior destino"})
	}

	/////////////////// CREATING PACKAGE
	var point_org = make([]float64, 2)
	var point_dst = make([]float64, 2)

	new_org_latitude, _ := strconv.ParseFloat(org_latitude, 64)
	new_org_longitude, _ := strconv.ParseFloat(org_longitude, 64)

	point_org[0] = new_org_longitude
	point_org[1] = new_org_latitude

	new_dst_latitude, _ := strconv.ParseFloat(dst_latitude, 64)
	new_dst_longitude, _ := strconv.ParseFloat(dst_longitude, 64)

	point_dst[0] = new_dst_longitude
	point_dst[1] = new_dst_latitude

	location_org := models.Location{
		Address:     org_address,
		Zipcode:     org_zipcode_,
		Extnum:      int(org_ext_num_),
		IntNum:      int(org_int_num_),
		Coordinates: point_org,
	}

	location_dst := models.Location{
		Address:     dst_address,
		Zipcode:     dst_zipcode_,
		Extnum:      int(dst_ext_num_),
		IntNum:      int(dst_int_num_),
		Coordinates: point_dst,
	}

	new_pkg := models.Packages{
		Identifier: uuid.New().String(),
		Size:       size,
		Amount:     int(amount),
		Weight:     weight,
		Status:     "creado",
		CreatedAt:  time.Now(),
		Origin:     &location_org,
		Destiny:    &location_dst,
	}

	_, err = t.collection.InsertOne(c.Context(), new_pkg)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"Error": err.Error()})
	}
	return c.JSON(fiber.Map{"success": "Packaged crated"})
}

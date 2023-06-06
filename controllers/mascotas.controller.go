package controllers

import (
	"fmt"
	"strconv"

	"github.com/HouseCham/MartonVet/models"
	"github.com/gofiber/fiber/v2"
)

func MascotasVer(c *fiber.Ctx) error {
	clientId := c.Params("id")
	if clientId == "" {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"mensaje": "No se ha especificado un id",
		})
	}
	var client models.Client
	clientIdInt, err := strconv.Atoi(clientId)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"mensaje": "El id no es valido",
		})
	}

	client.Id = clientIdInt
	result := db.First(&client, clientId)
	if result.Error != nil {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"mensaje": "No se encontro el cliente",
		})
	}

	var pets []models.Pet
	result = db.Find(&pets, "cliente_id = ?", clientIdInt)
	if result.Error != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"mensaje": "Hubo un error en el servidor",
			"errores": result.Error.Error(),
		})
	}

	return c.Render("clienteMascotas", fiber.Map{
		"nombre": fmt.Sprintf("%s %s %s", client.Nombre, client.ApellidoP, client.ApellidoM),
		"clientId": clientId,
		"pets": pets,
	})
}
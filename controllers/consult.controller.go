package controllers

import (
	"github.com/HouseCham/MartonVet/models"
	"github.com/gofiber/fiber/v2"
)

func ConsultasVer(c *fiber.Ctx) error {
	var clients []models.Client
	result := db.Find(&clients)
	if result.Error != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"mensaje": "Hubo un error en el servidor",
			"errores": result.Error.Error(),
		})
	}

	return c.Render("verConsultas", fiber.Map{
		"consults": clients,
	})
}

func ConsultasRegistrar(c *fiber.Ctx) error {
	return c.Render("nuevaConsulta", fiber.Map{
		"Title": "Registrar cliente",
	})
}
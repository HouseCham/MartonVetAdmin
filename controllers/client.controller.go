package controllers

import (
	"strconv"

	"github.com/HouseCham/MartonVet/models"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var db *gorm.DB

// InsertNewClient is a function that inserts a new client into the database
func InsertNewClient(c *fiber.Ctx) error {
	var request models.Client

	// Parse request body from JSON to struct
	c.BodyParser(&request)

	validate := validator.New()
	err := validate.Struct(request)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"mensaje": "Al parecer hubo un error en tu solicitud",
			"errores": err.Error(),
		})
	}

	// Insert new client
	result := db.Create(&request)
	if result.Error != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"mensaje": "Hubo un error en el servidor",
			"errores": result.Error.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"mensaje": "Se ha insertado un nuevo cliente",
	})
}

// UpdateClient is a function that updates a client in the database
func UpdateClient(c *fiber.Ctx) error {
	var request models.Client

	// Parse request body from JSON to struct
	c.BodyParser(&request)

	// Validate request body
	validate := validator.New()
	err := validate.Struct(request)
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"mensaje": "Al parecer hubo un error en tu solicitud",
			"errores": err.Error(),
		})
	}

	// look for client in database
	var client models.Client
	db.First(&client, request.Id)
	if client.Id == 0 {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"mensaje": "No se encontro el cliente",
		})
	}

	client.ParseClientParams(request)
	
	// update client
	db.Save(&client)

	return c.JSON(fiber.Map{
		"mensaje": "Se ha insertado un nuevo cliente",
	})
}

// ClientesRegistrar is a function that renders the new client page
func ClientesRegistrar(c *fiber.Ctx) error {
	return c.Render("nuevoCliente", fiber.Map{
		"Title": "Registrar cliente",
	})
}

// ClientesVer is a function that renders the view clients page
func ClientesVer(c *fiber.Ctx) error {
	var clients []models.Client
	result := db.Find(&clients)
	if result.Error != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"mensaje": "Hubo un error en el servidor",
			"errores": result.Error.Error(),
		})
	}

	return c.Render("verClientes", fiber.Map{
		"clients": clients,
	})
}

// VerClientesEditar is a function that renders the edit client page
func VerClientesEditar(c *fiber.Ctx) error {
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

	return c.Render("editarCliente", fiber.Map{
		"id": client.Id,
		"nombre": client.Nombre,
		"apellidoP": client.ApellidoP,
		"apellidoM": client.ApellidoM,
		"email": client.Email,
		"telefono": client.Telefono,
		"calle": client.Calle,
		"numExt": client.NumeroExt,
		"numInt": client.NumeroInt,
		"colonia": client.Colonia,
		"cp": client.CodigoPostal,
		"ciudad": client.Ciudad,
		"estado": client.Estado,
	})
}

// VerClientesEliminar is a function that renders the delete client page
func VerClienteEliminar(c *fiber.Ctx) error {
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

	return c.Render("eliminarCliente", fiber.Map{
		"id": client.Id,
		"nombre": client.Nombre,
		"apellidoP": client.ApellidoP,
		"apellidoM": client.ApellidoM,
	})
}

// DeleteClient is a function that deletes a client from the database
func DeleteClient(c *fiber.Ctx) error {
	clientId := c.Params("id")
	if clientId == "" {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"mensaje": "No se ha especificado un id",
		})
	}

	result := db.Delete(&models.Client{}, clientId)
	if result.Error != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"mensaje": "Hubo un error en el servidor",
			"errores": result.Error.Error(),
		})
	}
	return c.JSON(fiber.Map{
		"mensaje": "Se ha eliminado el cliente",
	})
}

func ShareDB(mainDB *gorm.DB) {
	db = mainDB
}
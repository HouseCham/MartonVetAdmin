package routes

import (
	"github.com/HouseCham/MartonVet/controllers"
	"github.com/gofiber/fiber/v2"
)

// SetAllRoutes is a function that sets up all routes for the application
func SetAllRoutes(app *fiber.App) {
	app.Get("/", controllers.Index)
	app.Get("/registrar_cliente", controllers.ClientesRegistrar)
	app.Get("/ver_clientes", controllers.ClientesVer)
	app.Get("/editar_cliente/:id", controllers.VerClientesEditar)
	app.Get("/ver_consultas", controllers.ConsultasVer)
	app.Get("/registrar_consulta", controllers.ConsultasRegistrar)
	app.Get("/eliminar_cliente/:id", controllers.VerClienteEliminar)
	app.Get("/ver_mascotas/:id", controllers.MascotasVer)
	
	app.Post("/insertar_cliente", controllers.InsertNewClient)

	app.Put("/actualizar_cliente", controllers.UpdateClient)

	app.Delete("/eliminar_cliente/:id", controllers.DeleteClient)
}


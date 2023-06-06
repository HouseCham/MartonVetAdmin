package main

import (
	"github.com/HouseCham/MartonVet/controllers"
	"github.com/HouseCham/MartonVet/models"
	"github.com/HouseCham/MartonVet/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	engine := html.New("./views", ".html")

	// Pass engine to Fiber's Views Engine
	app := fiber.New(fiber.Config{
		Views: engine,
	})
	app.Static("/static", "./static")

	// Creating gorm object
	dsn := "root:secret@tcp(127.0.0.1:3306)/MartonVet?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database: " + err.Error())
	}
	
	db.AutoMigrate(&models.Client{})
	db.AutoMigrate(&models.Pet{})
	
	controllers.ShareDB(db)
	
	routes.SetAllRoutes(app)

	app.Listen(":3000")
}

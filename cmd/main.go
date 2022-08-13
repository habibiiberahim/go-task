package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/habibiiiberahim/go-task/entities"
	"github.com/habibiiiberahim/go-task/routes"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main()  {
	app := fiber.New()

	db := SetupDatabase()

	app.Use(cors.New())

	//migrate table om database
	db.AutoMigrate(&entities.Task{})

	//add routes to app
	routes.InitTaskRoutes(db, app)

	app.Listen(":3000")
}

func SetupDatabase() *gorm.DB {
	//hard code database
	dsn := "root:12345678@tcp(localhost:3306)/go-task?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	
	if err != nil {
		fmt.Println(err)
	}

	return db
}
package main

import (
	"fiber/database"
	"fiber/model"
	"fiber/mw"
	"fiber/router"
	"fiber/tool"
	"log"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/cors"
	"github.com/gofiber/csrf"
	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/middleware"
	"github.com/gofiber/helmet"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func initDatabase() {
	var err error
	database.DB, err = gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("Failed to connect database")
	} else {
		log.Println("Connect database Success")
	}

	if database.DB.HasTable(&model.Users{}) {
		database.DB.AutoMigrate(&model.Users{})
	} else {
		database.DB.CreateTable(&model.Users{})
	}

	if database.DB.HasTable(&model.Infos{}) {
		database.DB.AutoMigrate(&model.Infos{})
	} else {
		database.DB.CreateTable(&model.Infos{})
	}
}

func main() {
	app := fiber.New()

	app.Use(middleware.Logger())
	app.Use(cors.New())
	app.Use(csrf.New())
	app.Use(helmet.New())

	initDatabase()

	app.Static("/", "./build", fiber.Static{
		Compress:  false,
		ByteRange: false,
		Browse:    true,
		Index:     "index.html",
	})

	// app.Get("/", func(c *fiber.Ctx) {
	// 	c.Send(c.Locals("csrf"))
	// })

	app.Get("/csrf", func(c *fiber.Ctx) {
		c.JSON(fiber.Map{
			"csrf": c.Locals("csrf"),
		})
	})

	app.Get("/restricted", mw.Protected(), func(c *fiber.Ctx) {
		user := c.Locals("user").(*jwt.Token)
		name := tool.ValidToken(user)
		c.Send("Welcome " + name)
	})

	router.Route(app)

	_ = app.Listen(9000)

	defer database.DB.Close()
}

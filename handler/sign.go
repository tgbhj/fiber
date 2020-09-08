package handler

import (
	"fiber/database"
	"fiber/model"
	"fiber/tool"
	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber"
	"time"
)

// SignUp 用户注册
func SignUp(c *fiber.Ctx) {
	db := database.DB
	user := new(model.Users)
	if err := c.BodyParser(user)
		err != nil {
		_ = c.Status(500).JSON(fiber.Map{"code": "50003", "message": "Review your input", "cb": err})
		return
	}

	user.Password = tool.PassMd5(user.Password)
	if err := db.Create(&user).Error
		err != nil {
		_ = c.Status(500).JSON(fiber.Map{"code": "50004", "message": "Couldn't create user", "cb": err})
		return
	}

	_ = c.JSON(fiber.Map{"code": "20000", "message": "Created user Success", "cb": "SignUp Success"})
}

// SignIn 用户登陆
func SignIn(c *fiber.Ctx) {
	username := c.FormValue("Username")
	password := c.FormValue("Password")
	password = tool.PassMd5(password)
	db := database.DB
	var user model.Users
	db.Where("username = ? and password = ?", username, password).First(&user)
	if user.Username == "" || user.Password == "" {
		_ = c.Status(500).JSON(fiber.Map{"code": "50005", "message": "User is non-existent", "cb": nil})
		return
	}

	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = user.Username
	claims["admin"] = true
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		c.SendStatus(fiber.StatusInternalServerError)
		return
	}

	if err := c.JSON(fiber.Map{"code": "20000", "message": "User SignIn Success", "cb": t, "csrf": c.Locals("csrf")})
		err != nil {
		c.Status(500).Send(err)
		return
	}
}

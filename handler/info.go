package handler

import (
	"fiber/database"
	"fiber/model"
	"github.com/gofiber/fiber"
)

// GetInfo 获取信息
func GetInfo(c *fiber.Ctx) {
	db := database.DB
	var infos []model.Infos
	db.Order("ID desc").Find(&infos)
	c.JSON(fiber.Map{
		"code": 20000,
		"msg":  "Success",
		"cb": infos,
	})
}

// PostInfo 提交信息
func PostInfo(c *fiber.Ctx) {
	db := database.DB
	info := new(model.Infos)
	if err := c.BodyParser(info)
		err != nil {
		_ = c.Status(500).JSON(fiber.Map{"code": "50006", "message": "Review your input", "cb": err})
		return
	}

	if err := db.Create(&info).Error
		err != nil {
		_ = c.Status(500).JSON(fiber.Map{"code": "50007", "message": "Couldn't create info", "cb": err})
		return
	}

	_ = c.JSON(fiber.Map{"code": "20000", "message": "Created info Success", "cb": "Success"})
}
package handlers

import (
	"auth-service/models"
	"auth-service/utils"

	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type AuthHandler struct {
    DB  *gorm.DB
    Rdb *redis.Client
    
}

func (h *AuthHandler) Register(c *fiber.Ctx) error {
	var u models.User
	if err := c.BodyParser(&u); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	pw, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}
	u.Password = pw

	if err := h.DB.Create(&u).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "could not register user"})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"id": u.ID})
}
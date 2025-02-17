package controller

import "github.com/gofiber/fiber/v2"

type ProductControllerInterface interface {
	Create(c *fiber.Ctx) error
	DeleteById(c *fiber.Ctx) error
	UpdateById(c *fiber.Ctx) error
	GetAll(c *fiber.Ctx) error
}

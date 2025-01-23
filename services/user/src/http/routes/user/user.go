package userroutes

import (
	userController "github.com/TIM-DEBUG-ProjectSprintBatch3/TutupLapak/user/src/http/controllers/user"
	"github.com/TIM-DEBUG-ProjectSprintBatch3/TutupLapak/user/src/http/middlewares"
	"github.com/gofiber/fiber/v2"
)

func SetRouteUsers(router fiber.Router, uc userController.UserControllerInterface) {
	router.Post("/register/email", uc.RegisterByEmail)
	router.Post("/register/phone", uc.RegisterByPhone)
	router.Post("/login/email", uc.LoginByEmail)
	router.Post("/login/phone", uc.LoginByPhone)

	router.Post("/user/link/email", middlewares.AuthMiddleware, uc.LinkEmail)
}

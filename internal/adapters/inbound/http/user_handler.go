package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/yduman/go-hexarch/internal/application/service"
	"github.com/yduman/go-hexarch/internal/domain/entity"
)

type UserHandler struct {
	UserService service.UserService
}

func NewUserHandler(userService service.UserService) *UserHandler {
	return &UserHandler{UserService: userService}
}

func (handler *UserHandler) CreateUser(ctx *fiber.Ctx) error {
	var user entity.User

	if err := ctx.BodyParser(&user); err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	if err := handler.UserService.CreateUser(&user); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return ctx.SendStatus(fiber.StatusCreated)
}

func (handler *UserHandler) GetUserByID(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	user, err := handler.UserService.GetUserByID(id)

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	if user == nil {
		return ctx.Status(fiber.StatusNotFound).SendString("User not found")
	}

	return ctx.JSON(user)
}

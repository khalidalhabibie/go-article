package http

import (
	"backend/app/user"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	userUsecase user.Usecase
}

func New(userUC user.Usecase) *Handler {
	return &Handler{
		userUsecase: userUC,
	}
}

func (h *Handler) Register(app *fiber.App) {

	app.Post("/login", h.Login)
	app.Post("/register", h.Registration)



}

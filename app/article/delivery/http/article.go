package http

import (
	"backend/app/article"

	"backend/pkg/middleware"

	fiber "github.com/gofiber/fiber/v2"
)

type Handler struct {
	articleUsecase article.Usecase
}

func New(articleUC article.Usecase) *Handler {
	return &Handler{
		articleUsecase: articleUC,
	}
}

func (h *Handler) Register(app *fiber.App) {
	route := app.Group("/articles")

	//public route
	public := route.Group("")
	public.Post("", middleware.JWTProtected(), h.Create)
	public.Get("", middleware.JWTProtected(), h.Index)
	public.Delete("", middleware.JWTProtected(), h.Delete)
	public.Patch("", middleware.JWTProtected(), h.Update)

}

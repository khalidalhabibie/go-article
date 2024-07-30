package utils

import (
	"backend/app/models"

	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
)

type GeneralResponse struct {
	Message string `json:"message"`
}

func ReturnResponse(c *fiber.Ctx, err error, data interface{}) error {

	var code int

	if err == nil {
		return c.Status(fiber.StatusOK).JSON(data)

	}

	switch err {
	case fiber.ErrNotFound:
		code = fiber.ErrNotFound.Code

	case fiber.ErrUnprocessableEntity:
		code = fiber.ErrUnprocessableEntity.Code

	case fiber.ErrBadRequest:
		code = fiber.ErrBadRequest.Code
	default:
		log.WithFields(LogFormat(models.LogLayerDelivery, models.ServiceAuth, err.Error())).Error("failed get error code")

		code = fiber.ErrUnprocessableEntity.Code
		// return err
	}

	return c.Status(code).JSON(GeneralResponse{
		Message: err.Error(),
	})

}

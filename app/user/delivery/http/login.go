package http

import (
	"backend/app/models"
	"backend/app/user/delivery/http/request"
	"backend/pkg/utils"

	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
)

// Login
// @Tags User
// @Summary Login user
// @Description Login new user
// @Accept json
// @Produce json
// @Param request body request.Login true "body"
// @Success 200 {object} models.Tokens
// @Failure 400 {object} utils.ValidatorResponse
// @Failure 422 {object} utils.GeneralResponse
// @Router /login [post]
func (h *Handler) Login(c *fiber.Ctx) error {
	// logId := uuid.NewUUID()
	log.WithFields(utils.LogFormat(models.LogLayerDelivery, models.UserService, "start")).Info("login")

	// Create a create in struct.
	request := &request.Login{}

	// Checking received data from JSON body.
	if err := c.BodyParser(request); err != nil {
		err := fiber.ErrBadRequest

		log.WithFields(utils.LogFormat(models.LogLayerDelivery, models.UserService, err.Error())).Error("failed to parse request")

		// Return status 400 and error message
		return utils.ReturnResponse(c, err, nil)

	}

	// Create a new validator for a model.
	validate := utils.NewValidator()

	// Validate sign up fields.
	if err := validate.Struct(request); err != nil {
		log.WithFields(utils.LogFormat(models.LogLayerDelivery, models.ServiceArticle, err.Error())).Error("error validate request body")

		errM := fiber.ErrBadRequest
		errM.Message = err.Error()

		return utils.ValidatorErrorsBind(err, *c)

	}

	userM, err := h.userUsecase.Login(request.Email, request.Password)
	if err != nil {
		return utils.ReturnResponse(c, err, nil)

	}

	dataM, err := utils.MarshalUsers(userM, models.AuthRoleNamePublic)
	if err != nil {
		log.WithFields(utils.LogFormat(models.LogLayerDelivery, models.UserService, err.Error())).Error("error marshal to user")

		return utils.ReturnResponse(c, err, nil)

	}

	log.WithFields(utils.LogFormat(models.LogLayerDelivery, models.UserService, "end")).Info("login")

	return utils.ReturnResponse(c, nil, dataM)

}

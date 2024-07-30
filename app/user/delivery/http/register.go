package http

import (
	"backend/app/models"
	"backend/app/user/delivery/http/request"
	"backend/pkg/utils"

	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
)

// Register
// @Tags User
// @Summary Create user
// @Description Create new user
// @Accept json
// @Produce json
// @Param request body request.SignUp true "body"
// @Success 200 {object} models.User
// @Failure 400 {object} utils.ValidatorResponse
// @Failure 422 {object} utils.GeneralResponse
// @Router /register [post]
func (h *Handler) Registration(c *fiber.Ctx) error {
	log.WithFields(utils.LogFormat(models.LogLayerDelivery, models.UserService, "start")).Info("register")

	// Create a create in struct.
	request := &request.SignUp{}

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

	userM, err := h.userUsecase.Registration(*request)
	if err != nil {

		return utils.ReturnResponse(c, err, nil)

	}

	dataM, err := utils.MarshalUsers(userM, models.AuthRoleNamePublic)
	if err != nil {
		log.WithFields(utils.LogFormat(models.LogLayerDelivery, models.UserService, err.Error())).Error("error marshal to user")

		return utils.ReturnResponse(c, err, nil)

	}

	log.WithFields(utils.LogFormat(models.LogLayerDelivery, models.UserService, "end")).Info("create")

	return utils.ReturnResponse(c, nil, dataM)

}

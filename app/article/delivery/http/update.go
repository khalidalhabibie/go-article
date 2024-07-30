package http

import (
	"backend/app/article/delivery/http/request"
	"backend/app/models"
	"backend/pkg/utils"

	"github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
)

// UpdateArticle
// @Tags Articles
// @Summary update article
// @Description update  article,
// @Accept json
// @Produce json
// @Param id path string true "UUID of the Article"
// @Param request body request.Create true "body"
// @Success 200 {object} models.Article
// @Failure 400 {object} utils.ValidatorResponse
// @Failure 422 {object} utils.GeneralResponse
// @Router /articles [pathc]
func (h *Handler) Update(c *fiber.Ctx) error {
	log.WithFields(utils.LogFormat(models.LogLayerDelivery, models.ServiceArticle, "start")).Info("create")

	// Create a create in struct.
	request := &request.Create{}

	// Checking received data from JSON body.
	if err := c.BodyParser(request); err != nil {
		err := fiber.ErrBadRequest

		log.WithFields(utils.LogFormat(models.LogLayerDelivery, models.ServiceArticle, err.Error())).Error("failed to parse request")

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

	token, err := utils.ExtractTokenMetadata(c)
	if err != nil {
		log.WithFields(utils.LogFormat(models.LogLayerUsecase, models.ServiceArticle, err)).Error("failed to extract token")

		err := fiber.ErrUnprocessableEntity

		return utils.ReturnResponse(c, err, nil)
	}

	articleM, err := h.articleUsecase.Create(*request, *token)
	if err != nil {

		return utils.ReturnResponse(c, err, nil)

	}

	dataM, err := utils.MarshalUsers(articleM, models.AuthRoleNamePublic)
	if err != nil {
		log.WithFields(utils.LogFormat(models.LogLayerDelivery, models.ServiceArticle, err.Error())).Error("error marshal to user")

		return utils.ReturnResponse(c, err, nil)

	}

	log.WithFields(utils.LogFormat(models.LogLayerDelivery, models.ServiceArticle, "end")).Info("create")

	return utils.ReturnResponse(c, nil, dataM)

}

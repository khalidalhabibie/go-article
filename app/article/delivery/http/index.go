package http

import (
	"backend/app/article/delivery/http/request"
	"backend/app/models"
	"backend/pkg/utils"
	"net/url"

	fiber "github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
)

// IndexArticls
// @Tags Articles
// @Summary Search articles
// @Description Get articles by body and title with pagination
// @Accept json
// @Produce json
// @Param search query string false "search"
// @Param limit query string false "limit"
// @Param offset query string false "offset"
// @Success 200 {object} response.Index
// @Failure 422 {object} utils.GeneralResponse
// @Router /articles [get]
func (h *Handler) Index(c *fiber.Ctx) error {

	log.WithFields(utils.LogFormat(models.LogLayerDelivery, models.ServiceArticle, "start")).Info("start index")

	// parsing query string
	values, err := url.ParseQuery(string(c.Request().URI().QueryString()))
	if err != nil {
		log.WithFields(utils.LogFormat(models.LogLayerDelivery, models.ServiceArticle, err.Error())).Error("failed get query string")

		err := fiber.ErrBadRequest

		return utils.ReturnResponse(c, err, nil)
	}

	// get data form usecase
	articles, err := h.articleUsecase.Index(request.PaginationConfig(values))
	if err != nil {
		log.WithFields(utils.LogFormat(models.LogLayerDelivery, models.ServiceArticle, err.Error())).Error("failed to extract data user")

		return utils.ReturnResponse(c, err, nil)

	}

	// serialization data based on role
	dataM, err := utils.MarshalUsers(articles, models.AuthRoleNamePublic)
	if err != nil {
		log.WithFields(utils.LogFormat(models.LogLayerDelivery, models.ServiceArticle, err.Error())).Error("error marshal to user")

		return utils.ReturnResponse(c, err, nil)

	}

	log.WithFields(utils.LogFormat(models.LogLayerDelivery, models.ServiceArticle, "start")).Info("end index")

	return utils.ReturnResponse(c, err, dataM)

}

package http

import (
	"backend/app/models"
	"backend/pkg/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

// DeleteArticle
// @Tags Articles
// @Summary Delete article
// @Description Delete an article by UUID
// @Accept json
// @Produce json
// @Param id path string true "UUID of the Article"
// @Success 204 "Article deleted successfully"
// @Failure 400 {object} utils.ValidatorResponse
// @Failure 404 {object} utils.GeneralResponse
// @Router /articles/{uuid} [delete]
func (h *Handler) Delete(c *fiber.Ctx) error {
	idStr := c.Params("id")
	log.WithFields(utils.LogFormat(models.LogLayerDelivery, models.ServiceArticle, "start delete")).Info("delete article")

	if idStr == "" {
		err := fiber.ErrBadRequest
		err.Message = "UUID must be provided"
		log.WithFields(utils.LogFormat(models.LogLayerDelivery, models.ServiceArticle, err.Message)).Error("invalid UUID")
		return utils.ReturnResponse(c, err, nil)
	}

	// Validate and parse UUID
	id, err := uuid.Parse(idStr)
	if err != nil {
		log.WithFields(utils.LogFormat(models.LogLayerUsecase, models.ServiceArticle, err)).Error("invalid id")

		err := fiber.ErrBadRequest
		err.Message = "invalid id"

		return utils.ReturnResponse(c, err, nil)

	}

	token, err := utils.ExtractTokenMetadata(c)
	if err != nil {
		log.WithFields(utils.LogFormat(models.LogLayerUsecase, models.ServiceArticle, err)).Error("failed to extract token")

		err := fiber.ErrUnprocessableEntity

		return utils.ReturnResponse(c, err, nil)
	}

	dataM, err := h.articleUsecase.Delete(id, *token)
	if err != nil {
		return utils.ReturnResponse(c, err, nil)
	}

	log.WithFields(utils.LogFormat(models.LogLayerDelivery, models.ServiceArticle, "end delete")).Info("article deleted")
	return utils.ReturnResponse(c, nil, dataM)
}

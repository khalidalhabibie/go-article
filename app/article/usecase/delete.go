package usecase

import (
	"backend/app/models"
	"backend/pkg/utils"

	fiber "github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

func (u *Usecase) Delete(ID uuid.UUID, token utils.TokenMetadata) (*models.Article, error) {
	articleM, err := u.articleRepo.FindById(ID)
	if err != nil {
		log.WithFields(utils.LogFormat(models.LogLayerUsecase, models.ServiceArticle, err)).Error("failed copier message")

		err := fiber.ErrNotFound
		err.Message = "article not found"

		return nil, err
	}

	if articleM.CreatedBy != token.UserID {
		log.WithFields(utils.LogFormat(models.LogLayerUsecase, models.ServiceArticle, err)).Error("forbidden")
		err := fiber.ErrForbidden
		err.Message = "you can't update this article"

		return nil, err

	}

	err = u.articleRepo.Delete(*articleM, nil)
	if err != nil {
		log.WithFields(utils.LogFormat(models.LogLayerUsecase, models.ServiceArticle, err)).Error("failed to delete article")

		err := fiber.ErrUnprocessableEntity
		err.Message = "Failed to delete article"
		return nil, err

	}

	return articleM, err

}

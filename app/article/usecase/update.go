package usecase

import (
	"backend/app/article/delivery/http/request"
	"backend/app/models"
	"backend/pkg/utils"

	fiber "github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	log "github.com/sirupsen/logrus"
)

func (u *Usecase) Update(request request.Update, ID uuid.UUID, token utils.TokenMetadata) (*models.Article, error) {
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

	err = copier.Copy(articleM, request)
	if err != nil {
		log.WithFields(utils.LogFormat(models.LogLayerUsecase, models.ServiceArticle, err)).Error("failed copier message")
		err := fiber.ErrUnprocessableEntity
		return nil, err

	}

	err = u.articleRepo.Update(*articleM, nil)
	if err != nil {
		log.WithFields(utils.LogFormat(models.LogLayerUsecase, models.ServiceArticle, err)).Error("failed to update article")

		err := fiber.ErrUnprocessableEntity
		err.Message = "Failed to update article"
		return nil, err

	}

	// flash product
	go u.articleCache.FlushAll()

	return articleM, err

}

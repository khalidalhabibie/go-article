package usecase

import (
	"backend/app/article/delivery/http/request"
	"backend/app/models"
	"backend/pkg/utils"

	fiber "github.com/gofiber/fiber/v2"
	"github.com/jinzhu/copier"
	log "github.com/sirupsen/logrus"
)

func (u *Usecase) Create(request request.Create, token utils.TokenMetadata) (*models.Article, error) {
	articleM := &models.Article{}

	err := copier.Copy(articleM, request)
	if err != nil {
		log.WithFields(utils.LogFormat(models.LogLayerUsecase, models.ServiceArticle, err)).Error("failed copier message")

		err := fiber.ErrUnprocessableEntity

		return nil, err

	}

	articleM.CreatedBy = token.UserID

	err = u.articleRepo.Insert(articleM, nil)
	if err != nil {
		log.WithFields(utils.LogFormat(models.LogLayerUsecase, models.ServiceArticle, err)).Error("failed to insert article")

		err := fiber.ErrUnprocessableEntity
		err.Message = "Failed  o insert article"
		return nil, err

	}

	// flash product
	go u.articleCache.FlushAll()

	return articleM, err

}

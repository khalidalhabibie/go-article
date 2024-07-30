package usecase

import (
	"backend/app/article/delivery/http/response"
	"backend/app/models"
	"backend/pkg/utils"

	fiber "github.com/gofiber/fiber/v2"
	log "github.com/sirupsen/logrus"
)

// func (u *Usecase) Index(request utils.PaginationConfig) ([]models.Article, utils.PaginationMeta, error) {// Index(request utils.PaginationConfig) ([]models.Article, utils.PaginationMeta, error)
func (u *Usecase) Index(request utils.PaginationConfig) (*response.Index, error) {

	// check request in cache
	dataCache, _ := u.articleCache.Get(request)
	// if err != nil {
	// 	return nil, err
	// }

	// if data cache is nul
	if dataCache != nil {
		return dataCache, nil
	}

	// if data  is not cached
	meta := utils.PaginationMeta{
		Offset: request.Offset(),
		Limit:  request.Limit(),
		Total:  0,
	}

	articles, err := u.articleRepo.FindAll(request)
	if err != nil {
		log.WithFields(utils.LogFormat(models.LogLayerUsecase, models.ServiceArticle, err.Error())).Error("error find articles")
		err := fiber.ErrUnprocessableEntity

		// return nil, meta, err
		return nil, err
	}

	total, err := u.articleRepo.Count(request)
	if err != nil {
		log.WithFields(utils.LogFormat(models.LogLayerUsecase, models.ServiceArticle, err.Error())).Error("error count articles")

		err := fiber.ErrUnprocessableEntity

		// return nil, meta, err
		return nil, err
	}

	meta.Total = total

	response := response.Index{
		Data: articles,
		Meta: meta,
	}

	// cache data
	go u.articleCache.Set(request, response)

	// return articles, meta, nil
	return &response, nil
}

package redis

import (
	"backend/app/article/delivery/http/response"
	"backend/app/models"
	"backend/pkg/utils"
	"encoding/json"
	"fmt"
	"log"
	"time"
)

func (c *Cache) Set(request utils.PaginationConfig, response response.Index) error {

	marshalResponse, err := json.Marshal(response)
	if err != nil {
		log.Panicln("salah, err : ", err)

		return err
	}

	// data := c.cache.Set(key, response, time.Duration(time.Hour*models.ArticleCacheTimeHour))
	// err := c.cache.Set(fmt.Sprintf("%v-%v", models.ServiceArticle, request), fmt.Sprint(response), time.Duration(time.Hour*models.ArticleCacheTimeHour)).Err()
	err = c.cache.Set(fmt.Sprintf("%v-%v", models.ServiceArticle, request), string(marshalResponse), time.Duration(time.Hour*models.ArticleCacheTimeHour)).Err()
	// err := c.cache.Set("1", fmt.Sprint(response), time.Duration(time.Hour*models.ArticleCacheTimeHour)).Err()
	if err != nil {
		log.Println("Failed to set cache, err ", err)
		return err
	}

	return nil

}

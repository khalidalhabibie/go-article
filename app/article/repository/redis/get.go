package redis

import (
	"backend/app/article/delivery/http/response"
	"backend/app/models"
	"backend/pkg/utils"
	"encoding/json"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

func (c *Cache) Get(request utils.PaginationConfig) (*response.Index, error) {
	// func (c *Cache) Get(request utils.PaginationConfig) (interface{}, error) {

	dataFromRedis := c.cache.Get(fmt.Sprintf("%v-%v", models.ServiceArticle, request))
	// dataFromRedis := c.cache.Get("1")
	// if dataFromRedis == nil {
	// 	log.Println("data not found, data : ", dataFromRedis)

	// 	err := fiber.ErrNotFound
	// 	return nil, err
	// }

	if err := dataFromRedis.Err(); err != nil {
		log.Println("data not found, data : ", dataFromRedis)

		err := fiber.ErrNotFound
		return nil, err
	}

	dataString, err := dataFromRedis.Result()
	if err != nil {
		log.Printf("failed to get data , data : %v , err : %v  ", dataFromRedis, err)

		err := fiber.ErrUnprocessableEntity
		return nil, err
	}

	// return dataString, err

	response := &response.Index{}

	// log.Println("xxxx , ", dataString)

	err = json.Unmarshal([]byte(dataString), &response)
	if err != nil {
		log.Printf("failed to marshal data , data : %v , err : %v  ", dataFromRedis, err)

		err := fiber.ErrUnprocessableEntity
		return nil, err

	}

	return response, nil

}

package redis

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func (c *Cache) FlushAll() error {
	err := c.cache.FlushAll().Err()
	if err != nil {
		log.Println("error flush all , err ", err)

		err := fiber.ErrUnprocessableEntity
		// err.Message = "Failed to get flush"

		return err
	}

	return nil

}

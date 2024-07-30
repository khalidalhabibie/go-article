package redis

import (
	"backend/app/models"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

func (c *Cache) DelVerificationCode(email string) error {

	codeFromRedis := c.cache.Del(fmt.Sprintf("%v-%v", models.UserRegistrationKeyVerificationCode, email))

	if err := codeFromRedis.Err(); err != nil {
		log.Printf("verification code %v  and email %v not found", codeFromRedis, email)

		err := fiber.ErrNotFound
		return err
	}

	return nil
}

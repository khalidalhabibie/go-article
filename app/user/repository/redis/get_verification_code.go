package redis

import (
	"backend/app/models"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

func (c *Cache) GetVerificationCode(email, verificationCode string) (*string, error) {

	codeFromRedis := c.cache.Get(fmt.Sprintf("%v-%v", models.UserRegistrationKeyVerificationCode, email))

	if err := codeFromRedis.Err(); err != nil {
		log.Printf("verification code %v  and email %v not found", codeFromRedis, email)

		err := fiber.ErrNotFound
		return nil, err
	}

	codeString, err := codeFromRedis.Result()
	if err != nil {
		log.Printf("failed to verification code %v  and email %v, err : %v  ", codeFromRedis, email, err)

		err := fiber.ErrUnprocessableEntity
		return nil, err
	}

	return &codeString, nil
}

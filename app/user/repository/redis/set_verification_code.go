package redis

import (
	"backend/app/models"
	"fmt"
	"log"
	"time"
)

func (c *Cache) SetVerificationCode(email, verificationCode string) error {

	err := c.cache.Set(fmt.Sprintf("%v-%v", models.UserRegistrationKeyVerificationCode, email), verificationCode, time.Duration(time.Hour*models.ArticleCacheTimeHour)).Err()
	if err != nil {
		log.Println("Failed SetVerificationCode to set cache, err ", err)
		return err
	}

	return nil
}

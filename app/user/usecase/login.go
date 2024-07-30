package usecase

import (
	"backend/app/models"
	"backend/pkg/utils"

	"github.com/gofiber/fiber/v2"
)

func (u *Usecase) Login(email, password string) (*models.Tokens, error) {

	// check, is email exists?
	user, err := u.userRepo.FindByEmail(email)
	if err != nil {
		err := fiber.ErrNotFound
		err.Message = "wrong user email or password"
		return nil, err
	}

	// check, is email  verified?
	if user.VerifiedAt == nil {
		err := fiber.ErrUnprocessableEntity
		err.Message = "confirmation you email, please"
		return nil, err

	}

	// Compare given user password with stored in found user.
	isPasswordCorrect := utils.ComparePasswords(user.Password, password)
	if !isPasswordCorrect {
		// Return, if password is not compare to stored in database.
		err := fiber.ErrBadRequest
		err.Message = "wrong user email or password"
		return nil, err
	}

	// generated the jwt
	// Generate a new pair of access and refresh tokens.
	tokens, err := utils.GenerateNewTokens(user.ID, email)
	if err != nil {
		err := fiber.ErrUnprocessableEntity
		err.Message = "Error generate token"
		return nil, err

	}

	return tokens, nil
}

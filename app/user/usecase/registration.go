package usecase

import (
	"backend/app/models"
	"backend/app/user/delivery/http/request"
	"backend/pkg/utils"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (u *Usecase) Registration(request request.SignUp) (*models.Tokens, error) {

	now := time.Now()

	// mapping user
	userM := models.User{
		ID:         uuid.New(),
		Name:       request.Name,
		Email:      request.Email,
		Address:    request.Address,
		PhoneNo:    request.PhoneNo,
		Password:   utils.GeneratePassword(request.Password),
		VerifiedAt: &now,
		Channel:    models.AuthInternalApps,
		CreatedAt:  now,
		UpdatedAt:  now,
	}

	// check is email used?
	user, _ := u.userRepo.FindByEmail(request.Email)
	if user != nil {
		err := fiber.ErrUnprocessableEntity
		err.Message = "email already exists"
		return nil, err

	}

	// TODO  : consider with my transaction
	err := u.userRepo.Insert(userM, nil)
	if err != nil {
		// tx.Rollback()
		return nil, err
	}

	// generated the jwt
	// Generate a new pair of access and refresh tokens.
	tokens, err := utils.GenerateNewTokens(user.ID, request.Email)
	if err != nil {
		err := fiber.ErrUnprocessableEntity
		err.Message = "Error generate token"
		return nil, err

	}

	return tokens, nil
}

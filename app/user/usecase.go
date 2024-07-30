package user

import (
	"backend/app/models"
	"backend/app/user/delivery/http/request"
)

type Usecase interface {
	Registration(request request.SignUp) (*models.Tokens, error)
	Login(email, password string) (*models.Tokens, error)
}

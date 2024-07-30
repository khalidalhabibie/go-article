package usecase

import (
	"backend/app/user"
)

type Usecase struct {
	userRepo  user.Repository
	userCache user.Cache
}

func New(
	userRepo user.Repository,
) user.Usecase {
	return &Usecase{
		userRepo: userRepo,
	}
}

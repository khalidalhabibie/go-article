package usecase

import "backend/app/article"

type Usecase struct {
	articleRepo  article.Repository
	articleCache article.Cache
}

func New(
	articleRepo article.Repository,
	articleCache article.Cache,

) article.Usecase {
	return &Usecase{
		articleRepo:  articleRepo,
		articleCache: articleCache,
	}
}

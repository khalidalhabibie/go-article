package article

import (
	"backend/app/article/delivery/http/request"
	"backend/app/article/delivery/http/response"
	"backend/app/models"
	"backend/pkg/utils"

	"github.com/google/uuid"
)

type Usecase interface {
	Create(requsest request.Create, token utils.TokenMetadata) (*models.Article, error)
	Index(request utils.PaginationConfig) (*response.Index, error)
	Update(request request.Update, ID uuid.UUID, token utils.TokenMetadata) (*models.Article, error)
	Delete(ID uuid.UUID, token utils.TokenMetadata) (*models.Article, error)
}

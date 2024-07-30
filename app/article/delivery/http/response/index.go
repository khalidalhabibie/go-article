package response

import (
	"backend/app/models"
	"backend/pkg/utils"
)

type Index struct {
	Data []models.Article     `json:"data" groups:"public"`
	Meta utils.PaginationMeta `json:"meta" groups:"public"`
}

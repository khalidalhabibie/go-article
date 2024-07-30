package request

import (
	"backend/pkg/utils"
)

func PaginationConfig(conditions map[string][]string) utils.PaginationConfig {

	utils.CopyKey(conditions, "search", "title")
	utils.CopyKey(conditions, "search", "body")
	utils.DeleteKey(conditions, "search")

	filterable := map[string]string{
		"title": utils.StringType,
		"body":  utils.StringType,
	}

	conditions["sort"] = []string{"created_at DESC"}

	// x := utils.NewRequestPaginationConfig(conditions, filterable)

	return utils.NewRequestPaginationConfig(conditions, filterable)
}

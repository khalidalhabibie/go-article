package request

type Update struct {
	Body string `json:"body" form:"body" validate:"required,gte=20"`
}

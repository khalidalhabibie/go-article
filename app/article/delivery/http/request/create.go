package request

type Create struct {
	Author string `json:"author" form:"author" validate:"required,gte=4,lte=50"`
	Title  string `json:"title" form:"title" validate:"required,gte=8"`
	Body   string `json:"body" form:"body" validate:"required,gte=20"`
}

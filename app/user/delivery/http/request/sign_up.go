package request

type SignUp struct {
	Name     string `json:"name" form:"name" validate:"required,gte=4,lte=50"`
	Email    string `json:"email" form:"email" validate:"required,email"`
	Address  string `json:"address" form:"address" validate:"required,gte=20"`
	PhoneNo  string `json:"phone_no" form:"phone_no" validate:"required,phone_number"`
	Password string `json:"password" form:"password" validate:"required,lte=20,gte=8"`
}

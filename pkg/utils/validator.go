package utils

import (
	"errors"
	"fmt"
	"regexp"

	"github.com/go-playground/validator/v10"
	strcase "github.com/stoewer/go-strcase"

	fiber "github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/liip/sheriff"
)

// NewValidator func for create a new validator for model fields.
func NewValidator() *validator.Validate {
	validate := validator.New()

	// Custom validation for uuid.UUID fields.
	_ = validate.RegisterValidation("uuid", func(fl validator.FieldLevel) bool {
		field := fl.Field().String()
		if _, err := uuid.Parse(field); err != nil {
			return true
		}
		return false
	})

	// validate email
	_ = validate.RegisterValidation("email", func(fl validator.FieldLevel) bool {
		field := fl.Field().String()
		emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)

		return emailRegex.MatchString(field)

	})

	return validate

}

// struct validator response
type ValidatorResponse struct {
	Message map[string]string `json:"message"`
}

// ValidatorErrors func for show validation errors for each invalid fields.
func ValidatorErrorsBind(err error, c fiber.Ctx) error {
	var ve validator.ValidationErrors
	errors.As(err, &ve)
	list := make(map[string]string)

	for _, err := range ve {
		list[strcase.SnakeCase(err.Field())] = validationErrorToText(err)
	}

	// c.Status(fiber.ErrBadRequest.Code).JSON(fiber.Map{"message": list})
	return c.Status(fiber.ErrBadRequest.Code).JSON(ValidatorResponse{Message: list})

}

// error handling middleware binding translator
func validationErrorToText(e validator.FieldError) string {
	errorField := strcase.SnakeCase(e.Field())
	switch e.Tag() {
	case "required":
		return fmt.Sprintf("%s is required", errorField)
	case "gte":
		return fmt.Sprintf("%s must equal to or greater than %s", errorField, e.Param())
	case "lte":
		return fmt.Sprintf("%s must equal to or less than %s", errorField, e.Param())
	}
	return fmt.Sprintf("%s is not valid", errorField)
}

func MarshalUsers(data interface{}, groups ...string) (interface{}, error) {
	o := &sheriff.Options{
		Groups: groups,
	}

	data, err := sheriff.Marshal(o, data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

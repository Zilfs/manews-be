package validatorLib

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

func ValidateStruct(s interface{}) error {
	var errrorMessage []string
	err := validate.Struct(s)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			switch err.Tag() {
			case "email":
				errrorMessage = append(errrorMessage, "Invalid email format")
			case "required":
				errrorMessage = append(errrorMessage, err.Field()+" is required")
			case "min":
				if err.Field() == "Password" {
					errrorMessage = append(errrorMessage, err.Field()+" must be at least "+err.Param()+" characters long")
				}
			case "eqfield":
				errrorMessage = append(errrorMessage, err.Field()+" must match "+err.Param())
			default:
				errrorMessage = append(errrorMessage, "Field "+err.Field()+" is not valid")
			}
		}
		return errors.New("Validation Failed: " + joinMessages(errrorMessage))
	}

	return nil
}

func joinMessages(messages []string) string {
	result := ""
	for i, msg := range messages {
		if i > 0 {
			result += ", "
		}
		result += msg
	}
	return result
}

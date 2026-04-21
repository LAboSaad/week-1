package validation

//request validation
import "github.com/go-playground/validator/v10"

var Validate = validator.New()

type CreateUserRequest struct {
	Name  string `json:"name"  validate:"required,min=3"`
	Email string `json:"email" validate:"required,email"`
}

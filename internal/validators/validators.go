package validators

import (
	"github.com/go-playground/validator/v10"
	"github.com/steveiliop56/puck/internal/config"
)

var validate *validator.Validate

func ValidateServer(server config.Server) (error) {
	validate = validator.New(validator.WithRequiredStructEnabled())

	var validateErr = validate.Struct(server)
	if validateErr != nil {
		return validateErr
	}

	return nil
}
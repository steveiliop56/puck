package validators

import (
	"github.com/go-playground/validator/v10"
	"github.com/steveiliop56/puck/internal/config"
)

var validate *validator.Validate

// Given a server struct it validates that all required fields exist, if they don't it returns an error
func ValidateServer(server config.Server) (error) {
	validate = validator.New(validator.WithRequiredStructEnabled())

	var validateErr = validate.Struct(server)
	if validateErr != nil {
		return validateErr
	}

	return nil
}

// Given the root config struct it validates it, if there is an error it returns an error
func ValidateConfig(config config.Config) (error) {
	validate = validator.New(validator.WithRequiredStructEnabled())

	var validateErr = validate.Struct(config)
	if validateErr != nil {
		return validateErr
	}

	return nil
}
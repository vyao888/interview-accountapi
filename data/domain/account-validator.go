package domain

import (
	"github.com/go-playground/validator"
)

// use a single instance of Validate, it caches struct info
var validate *validator.Validate

func validateAccount(acc *Account, logger *hclog.logger) (bool, error) {

	err := validate.Struct(acc)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			logger.Error("Account failed in validation", err)
			return false, err
		}
	}
	return true, nil
}

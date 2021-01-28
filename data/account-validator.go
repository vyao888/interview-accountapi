package data

import (
	"github.com/go-playground/validator/v10"
	"github.com/hashicorp/go-hclog"
)

func ValidateAccount(acc Account, l hclog.Logger) (bool, error) {
	v := NewValidator()
	err := v.Struct(acc)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			l.Error("Account failed in validation", err)
			return false, err
		}
	}
	return true, nil
}


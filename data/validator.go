package data

import (
	"fmt"
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
	for _, err := range err.(validator.ValidationErrors) {

		fmt.Println(err.Namespace())
		fmt.Println(err.Field())
		fmt.Println(err.StructNamespace())
		fmt.Println(err.StructField())
		fmt.Println(err.Tag())
		fmt.Println(err.ActualTag())
		fmt.Println(err.Kind())
		fmt.Println(err.Type())
		fmt.Println(err.Value())
		fmt.Println(err.Param())
		fmt.Println()
	}

	// from here you can create your own error messages in whatever language you wish
	return true, nil
}


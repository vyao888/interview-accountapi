package data

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/hashicorp/go-hclog"
	"regexp"
)

func ValidateAccount(acc Account, l hclog.Logger) (bool, error) {
	v := NewValidator()

	err := v.Struct(acc)

	if err != nil {

		// this check is only needed when your code could produce
		// an invalid value for validation such as interface with nil
		// value most including myself do not usually have code like this.
		if _, ok := err.(*validator.InvalidValidationError); ok {
			fmt.Println(err)
			return false, err
		}

		for _, err := range err.(validator.ValidationErrors) {

			fmt.Println(err.Namespace()) // can differ when a custom TagNameFunc is registered or
			fmt.Println(err.Field())     // by passing alt name to ReportError like below
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
		return false, err
	}
	return true, nil
}

func ValidateDate(fl validator.FieldLevel) bool {
	re := regexp.MustCompile(`(^\\d{4})-((?<=/)\\d{2}(?=/))-(\\d{2}$)`)
	date := re.FindAllString(fl.Field().String(), -1)

	if len(date) == 1 {
		return true
	}

	return false
}


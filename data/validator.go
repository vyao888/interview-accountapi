package data

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"regexp"
)

func ValidateAccount(acc Account) (bool, error) {
	v := NewValidator()
	v.RegisterValidation("date", ValidateDate)
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
			fmt.Printf(ValidationError(err))
		}

		return false, err
	}
	return true, nil
}

func ValidateDate(fl validator.FieldLevel) bool {
	re := regexp.MustCompile(`^\d{4}-\d{2}-\d{2}$`)
	date := re.FindAllString(fl.Field().String(), -1)
	return len(date) == 1
}

func ValidationError(err validator.FieldError) string {
	return fmt.Sprintf("%s:%s:%s", err.Namespace(), err.Tag(), err.Value())
}


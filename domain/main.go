package main

import (
	"fmt"
	"regexp"

	"github.com/go-playground/validator"
	//"github.com/go-playground/validator/v10"
	//	"github.com/hashicorp/go-hclog"
)

// use a single instance of Validate, it caches struct info
var validate *validator.Validate

func main() {
	validate := validator.New()
	validate.RegisterValidation("datetime", validateDate)
	//	logger := hclog.Default()

	a := &AccountHolder{
		Identification: "ad27e265-9605-4b4b-a0e5-3003ea9cc4dc",
		BirthDate:      "1988-03-08",
		BirthCountry:   "CN",
		Address:        []string{"line1", "line2", "line3"},
		City:           "London", Country: "UK"}

	err := validate.Struct(a)
	fmt.Println("before error check")
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			fmt.Println(err)
			return
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
		return
		//logger.Error("Failed to validate", err)
	}
	fmt.Printf("AccountHolder %f ", a)
}

type AccountHolder struct {
	Identification string   `json:"identification" validate:"unique"`
	BirthDate      string   `json:"birth_date" validate:"datetime"`
	BirthCountry   string   `json:"birth_country" validate:"uppercase,len=2"`
	Address        []string `json:"address" validate:"alphanum"`
	City           string   `json:"city" validate:"alpha"`
	Country        string   `json:"country" validate:"uppercase,len=2"`
}

func validateDate(fl validator.FieldLevel) bool {
	re := regexp.MustCompile(`(^\\d{4})-((?<=/)\\d{2}(?=/))-(\\d{2}$)`)
	date := re.FindAllString(fl.Field().String(), -1)

	if len(date) == 1 {
		return true
	}

	return false
}

package data

import (
	"fmt"
	"regexp"
	"github.com/go-playground/validator/v10"
)

//// User contains user information
//type User struct {
//	FirstName      string     `validate:"required"`
//	LastName       string     `validate:"required"`
//	Age            uint8      `validate:"gte=0,lte=130"`
//	Email          string     `validate:"required,email"`
//	FavouriteColor string     `validate:"iscolor"`                // alias for 'hexcolor|rgb|rgba|hsl|hsla'
//	Addresses      []*Address `validate:"required,dive,required"` // a person can have a home and cottage...
//	AddressLines   []string   `json:"address-lines"`
//}

//// Address houses a users address information
//type Address struct {
//	Street string `validate:"required"`
//	City   string `validate:"required"`
//	Planet string `validate:"required"`
//	Phone  string `validate:"required"`
//}



func validateStruct() {

	address := &Address{
		Street: "Eavesdown Docks",
		Planet: "Persphone",
		Phone:  "none",
	}

	user := &User{
		FirstName:      "Badger",
		LastName:       "Smith",
		Age:            135,
		Email:          "Badger.Smith@gmail.com",
		FavouriteColor: "#000-",
		Addresses:      []*Address{address},
		AddressLines:   []string{"line 1", "line 2", "line 3"},
	}

	// returns nil or ValidationErrors ( []FieldError )
	validate := NewValidator()
	validate.RegisterValidation("date", validateDate)
	err := validate.Struct(user)
	if err != nil {
		fmt.Println("Validation failed.")

		// this check is only needed when your code could produce
		// an invalid value for validation such as interface with nil
		// value most including myself do not usually have code like this.
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
	}

	// save user to database
}

func NewValidator() *validator.Validate {
	return validator.New()
}

func validateVariable() {

	myEmail := "joeybloggs@gmail.com"
	validate := NewValidator()
	errs := validate.Var(myEmail, "required,email")

	if errs != nil {
		fmt.Println(errs) // output: Key: "" Error:Field validation for "" failed on the "email" tag
		return
	}

	// email ok, move on
}

func validateDate(fl validator.FieldLevel) bool {
	re := regexp.MustCompile(`(^\\d{4})-((?<=/)\\d{2}(?=/))-(\\d{2}$)`)
	date := re.FindAllString(fl.Field().String(), -1)

	if len(date) == 1 {
		return true
	}

	return false
}

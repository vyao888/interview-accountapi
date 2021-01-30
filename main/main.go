package main

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/hashicorp/go-hclog"

	//	"github.com/hashicorp/go-hclog"
	"github.com/vyao888/interview-accountapi/data"
)

type Address struct {
	Street string `validate:"required"`
	City   string `validate:"required"`
	Planet string `validate:"required"`
	Phone  string `validate:"required"`
}

type User struct {
	FirstName      string     `validate:"required"`
	LastName       string     `validate:"required"`
	Age            uint8      `validate:"gte=0,lte=130"`
	Email          string     `validate:"required,email"`
	FavouriteColor string     `validate:"iscolor"`                // alias for 'hexcolor|rgb|rgba|hsl|hsla'
	Addresses      []*Address `validate:"required,dive,required"` // a person can have a home and cottage...
	AddressLines   []string   `json:"address-lines"`
}

func main() {

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

	l := hclog.Default()
	validate := data.NewValidator()
	validate.RegisterValidation("datetime", data.ValidateDate)
	err := validate.Struct(user)
	if err != nil {

		// this check is only needed when your code could produce
		// an invalid value for validation such as interface with nil
		// value most including myself do not usually have code like this.
		if _, ok := err.(*validator.InvalidValidationError); ok {
			fmt.Println(err)
			l.Error("failed validation", err)
			return
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

	}


	a := data.GetAccounts()
	fmt.Println(len(a))
	e := *a[0]
	fmt.Printf("%s", e)
	fmt.Printf("%s", data.Account.Json(e))
	b, err := data.ValidateAccount(e)
	if err != nil {
		l.Error("Account validation failed.", err)
	}
	fmt.Printf("%s", b)

}



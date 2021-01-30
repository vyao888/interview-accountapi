package main

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/hashicorp/go-hclog"

	//	"github.com/hashicorp/go-hclog"
	"github.com/vyao888/interview-accountapi/data"
)

//type Address struct {
//	Street string `validate:"required"`
//	City   string `validate:"required"`
//	Planet string `validate:"required"`
//	Phone  string `validate:"required"`
//}
//
//type User struct {
//	FirstName      string     `validate:"required"`
//	LastName       string     `validate:"required"`
//	Age            uint8      `validate:"gte=0,lte=130"`
//	Email          string     `validate:"required,email"`
//	FavouriteColor string     `validate:"iscolor"`                // alias for 'hexcolor|rgb|rgba|hsl|hsla'
//	Addresses      []*Address `validate:"required,dive,required"` // a person can have a home and cottage...
//	AddressLines   []string   `json:"address-lines"`
//}

type Actor struct {
	Name      [4]string `json:"name" validate:"alpha"`
	BirthDate string `json:"birth_date" validate:"date"`
	Residency string `json:"residency" validate:"uppercase,len=2"`
}

func main() {

	address := &data.Address{
		Street: "Eavesdown Docks",
		Planet: "Persphone",
		Phone:  "none",
	}

	user := &data.User{
		FirstName:      "Badger",
		LastName:       "Smith",
		Age:            135,
		Email:          "Badger.Smith@gmail.com",
		FavouriteColor: "#000-",
		Addresses:      []*data.Address{address},
		AddressLines:   []string{"line 1", "line 2", "line 3"},
		Name:      [4]string{"Jeff Page"},
		BirthDate: "2006-01-02T15:04:05",
	}

	l := hclog.Default()
	validate := data.NewValidator()
	validate.RegisterValidation("date", data.ValidateDate)
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
			fmt.Printf("%s:%s:%s\n", err.Namespace(), err.Tag(), err.Value())
		}

	}

	a := data.GetAccounts()
	fmt.Println(len(a))
	e := *a[0]
	fmt.Printf("%s", e)
	fmt.Printf("%s", data.Account.Json(e))

	ac := &Actor{
			Name:      [4]string{"Jeff Page"},
			BirthDate: "2006-01-02",
			Residency: "GB",
	}
	err = validate.Struct(ac)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			l.Error("Validation failed", err.Field() + ":" + err.Tag())
		}
	}
	//fmt.Printf("%s", b)

	//p := &data.AccountHolder{
	//	Identification: "13YH458762",
	//	BirthDate:      "2017-07-23",
	//	BirthCountry:   "GB",
	//	Address:        []string{"10 Avenue des Champs"},
	//	City:           "London",
	//	Country:        "GB",
	//}
	//
	//err = validate.Struct(p)
	//if err != nil {
	//	l.Error("AccountHolder validation failed.", err)
	//}
	//fmt.Printf("%s", b)

}



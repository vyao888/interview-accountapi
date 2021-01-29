package main

import (
	"fmt"
	"github.com/hashicorp/go-hclog"
	"github.com/vyao888/interview-accountapi/data"
)

func main() {
	l := hclog.Default()
	a := data.GetAccounts()
	fmt.Println(len(a))
	e := *a[0]
	fmt.Printf("%s", e)
	fmt.Printf("%s", data.Account.Json(e))

	b, err := data.ValidateAccount(e, l)
	if err != nil {
		l.Error("Account validation failed.", err)
	}
	fmt.Printf("%s", b)
}

package main

import (
	"fmt"
	"github.com/vyao888/interview-accountapi/data"
)

func main() {
//	l := hclog.Default()
	a := data.GetAccounts()
	fmt.Println(len(a))
	e := a[0]
	fmt.Printf("%s", *e)
	fmt.Printf("%s", data.Account.Json(e))
}

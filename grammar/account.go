package main

import (
	"fmt"

	"../grammar/accounts"
)

func main() {
	// account := accounts.Account{Owner: "nico"}
	// account.Owner = "pepe"
	// fmt.Println(account)

	account := accounts.NewAccount("nico")
	fmt.Println(account)
}

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
	account.Deposit(10)
	fmt.Println(account)
	fmt.Println(account.Balance())
}

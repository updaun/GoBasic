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
	account.Deposit(10)
	err := account.Withdraw(20)
	if err != nil {
		fmt.Println(err)
		// log.Fatalln(err) // log.Fatal will stop the program
	}
	fmt.Println(account)
}

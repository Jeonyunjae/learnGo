package main

import (
	"fmt"
	"name/accounts"
)

func main() {
	account := accounts.NewAccount("yunjae")
	account.Deposit(10)
	err := account.Withdraw(20)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(account.String())
}

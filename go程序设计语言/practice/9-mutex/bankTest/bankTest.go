package main

import (
	"fmt"

	"github.com/my/repo/go程序设计语言/practice/9-mutex/bank"
)

func main() {
	fmt.Println(bank.Balance())
	bank.Deposit(20)
	fmt.Println(bank.Balance())
	fmt.Println(bank.Withdraw(10))
	fmt.Println(bank.Balance())
}

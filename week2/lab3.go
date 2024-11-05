package main

import (
	"fmt"
)

type BankAccount struct {
	balance int
}

func (account *BankAccount) Deposit(a int) error {
	if a < 0 {
		return fmt.Errorf("입금 금액은 0보다 작을 수 없습니다.")
	}
	account.balance += a
	fmt.Println("입금 성공! 현재 잔액: ", account.balance)
	return nil
}

func (account *BankAccount) Withdraw(a int) error {
	if a < 0 {
		return fmt.Errorf("출금 금액은 0보다 작을 수 없습니다.")
	} else if account.balance < a {
		return fmt.Errorf("잔액이 부족합니다. 현재 잔액: %d원, 요청한 금액: %d원", account.balance, a)
	}
	account.balance -= a
	fmt.Println("출금 성공! 현재 잔액: ", account.balance)
	return nil
}

func main() {
	var mode int
	var input int
	myAccount := BankAccount{100}
	fmt.Print("입금 (1), 출금 (2), 종료 (Others) : ")
	fmt.Scan(&mode)

	if mode == 1 {
		fmt.Print("입금할 금액을 입력하세요: ")
		fmt.Scan(&input)
		myAccount.Deposit(input)

		if err := myAccount.Deposit(input); err != nil {
			fmt.Println(err)
		}

	} else if mode == 2 {
		fmt.Print("출금할 금액을 입력하세요: ")
		fmt.Scan(&input)
		myAccount.Withdraw(input)
		if err := myAccount.Withdraw(input); err != nil {
			fmt.Println(err)
		}
	} else {
		return
	}
}

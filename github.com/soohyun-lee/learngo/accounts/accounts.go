package accounts

import "errors"

// Account struct private
type Account struct {
	owner   string
	balance int
}

// NewAccount creates Account
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

//Deposit X amount on your account
func (a *Account) Deposit(amount int) {
	a.balance += amount
}

// Balance of your account
func (a Account) Balance() int {
	return a.balance
}

// go 는 복사본을 만든다.

// Withdraw x amount
func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		return errors.New("can't withdraw")
	}
	a.balance -= amount
	return nil
}

// changeowner of the account
func (a *Account) ChangeOwner(newOwner string) {
	a.owner = newOwner
}

// owner of the account
func (a Account) Owner() string {
	return a.owner
}

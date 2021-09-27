package main

import "fmt"

type payment interface {
	Pay()
}

type CashPayment struct {
}

func (CashPayment) Pay() {
	fmt.Println("Payment using cash")
}

func ProcessPayment(p payment) {
	p.Pay()
}

type BankPayment struct {
}

func (BankPayment) Pay(bankAccount int) {
	fmt.Printf("Paying using BankAccount %d\n", bankAccount)
}

type BankPaymentAdapter struct {
	BankPayment *BankPayment
	bankAccount int
}

func (bpa *BankPaymentAdapter) Pay() {
	bpa.BankPayment.Pay(bpa.bankAccount)
}

func main() {
	cash := &CashPayment{}
	ProcessPayment(cash)
	bpa := &BankPaymentAdapter{
		bankAccount: 5,
		BankPayment: &BankPayment{},
	}

	ProcessPayment(bpa)
}

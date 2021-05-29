package hello

import (
	cust "github.com/Anmol-M-0/goSimpleCLIapp/business/customer"
)

func Hello() string {
	return "hello world!!"
}

var customer *cust.Customer

func NewCustomer(name, email string, age int, creditLimit float64) (c cust.Customer, err error) {
	c, err = customer.New(name, email, age, creditLimit)
	return
}

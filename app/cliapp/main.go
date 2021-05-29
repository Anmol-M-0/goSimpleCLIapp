package main

import (
	"fmt"

	"github.com/Anmol-M-0/goSimpleCLIapp/app/cliapp/hello"
	cust "github.com/Anmol-M-0/goSimpleCLIapp/business/customer"
)

func main() {
	fmt.Println("hello world")
	fmt.Println(hello.Hello())
	var customer cust.Customer
	var err error
	for i := 0; i < 1; i++ {
		customer, err = createCustomer()
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println(customer)
	}

}

func createCustomer() (c cust.Customer, err error) {
	var name, email string
	var age int
	var creditLimit float64
	fmt.Scanf("%s %s %d %f", name, email, age, creditLimit)
	c, err = hello.NewCustomer(name, email, age, creditLimit)
	return
}

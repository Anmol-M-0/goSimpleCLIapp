package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Anmol-M-0/goSimpleCLIapp/app/cliapp/hello"
	cust "github.com/Anmol-M-0/goSimpleCLIapp/business/customer"
)

var sc *bufio.Scanner = bufio.NewScanner(os.Stdin)

func main() {
	fmt.Println("hello world")
	fmt.Println(hello.Hello())
	fmt.Println("enter number of input")
	var size int = scanInt()
	// fmt.Scanf("%d", &size)
	// sc.Scan()

	var customerList []cust.Customer = make([]cust.Customer, size)
	var tmpCustomer cust.Customer
	var err error

	for i := range customerList {
		tmpCustomer, err = createCustomer()
		if err != nil {
			fmt.Println(err)
			continue
		}
		customerList[i] = tmpCustomer
	}
	for _, v := range customerList {
		fmt.Println(v)
	}
	var command string
loop:
	for {
		fmt.Println("please enter case to execute, q to quit")
		command = scanString()
		switch command {
		case "stop", "1000", "exit", "q":
			break loop
		case "1", "new", "new customer":
			tmpCustomer, err = createCustomer()
			if err != nil {
				continue
			}
			customerList = append(customerList, tmpCustomer)
		case "2", "show", "show all":
			for _, v := range customerList {
				fmt.Println(v)
			}
		}

	}
}

func createCustomer() (c cust.Customer, err error) {
	var name, email string
	var age int
	var creditLimit float64
	fmt.Println("<name> <email> <age> <creditLimit>")
	fmt.Scanf("%s %s %d %f", &name, &email, &age, &creditLimit)
	sc.Scan()
	c, err = hello.NewCustomer(name, email, age, creditLimit)
	return
}
func scanInt() (i int) {
	fmt.Scanf("%d", &i)
	sc.Scan()
	return
}

func scanString() (str string) {
	sc.Scan()
	str = sc.Text()
	str = strings.TrimSpace(str)
	return
}

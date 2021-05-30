package customerCreditApplication

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/Anmol-M-0/goSimpleCLIapp/app/cliapp/customerCredit/hello"
	cust "github.com/Anmol-M-0/goSimpleCLIapp/business/customer"
)

var sc *bufio.Scanner = bufio.NewScanner(os.Stdin)

func RunCustomerCreditApp() {
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
		if err = validateByEmail(customerList, tmpCustomer.Email); err != nil {
			fmt.Println(err)
			continue
		}
		customerList[i] = tmpCustomer
	}
	for _, v := range customerList {
		fmt.Println(v)
	}

	var command string
	var tmpString string
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
			if err = validateByEmail(customerList, tmpCustomer.Email); err != nil {
				fmt.Println(err)
				continue
			}
			customerList = append(customerList, tmpCustomer)
		case "2", "show", "show all":
			for _, v := range customerList {
				fmt.Println(v)
			}
		case "3", "find":
			fmt.Println("enter 'email' to  find by email, 'name' to find by name")
			tmpString = scanString()
			switch tmpString {
			case "email":
				fmt.Println("enter email")
				tmpString = scanString()

				tmpCustomer, _, err = findByEmail(customerList, tmpString)
				if err != nil {
					fmt.Println(err)
					continue
				}
				fmt.Println(tmpCustomer)
			case "name":
				fmt.Println("enter name")
				tmpString = scanString()

				tmpCustomer, _, err = findByName(customerList, tmpString)
				if err != nil {
					fmt.Println(err)
					continue
				}
				fmt.Println(tmpCustomer)
			default:
				fmt.Println("invalid option")
			}
		case "change credit limit", "4":
			var index int
			fmt.Println("enter email")
			tmpString = scanString()
			tmpCustomer, index, err = findByEmail(customerList, tmpString)
			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Println("changing the credit limit of customer", tmpCustomer.Name)
			fmt.Println("enter 'add' to add to credit limit, 'rem' to subtract from the same")
			tmpString = scanString()
			switch tmpString {
			case "add":
				fmt.Println("enter the amount to add to the credit limit of", tmpCustomer.Name)
				customerList[index].CreditLimit = customerList[index].CreditLimit + scanFloat()
			case "rem":
				fmt.Println("enter the amount to remove from the credit limit of", tmpCustomer.Name)
				customerList[index].CreditLimit = customerList[index].CreditLimit - scanFloat()
			default:
				fmt.Println("invalid option")
			}
		}

	}
}

//create a customer instance by reading values from std input (console),
//return empty customer, error if invalid input
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

//scan for a float64 from std input (console)
func scanFloat() (f float64) {
	fmt.Scanf("%f", &f)
	sc.Scan()
	return
}

//scan for a int from std input (console)
func scanInt() (i int) {
	fmt.Scanf("%d", &i)
	sc.Scan()
	return
}

//scan for a string from std input (console)
func scanString() (str string) {
	sc.Scan()
	str = sc.Text()
	str = strings.TrimSpace(str)
	return
}

//from <customerList> find the customer and it's index, return (empty customer, -1, error) otherwise
func findByName(customerList []cust.Customer, name string) (c cust.Customer, index int, err error) {
	for i, v := range customerList {
		if v.Name == name {
			return v, i, nil
		}
	}
	return cust.Customer{}, -1, errors.New("customer not found")
}

func validateByEmail(customerList []cust.Customer, email string) error {
	for _, v := range customerList {
		if v.Email == email {
			return errors.New("this email is already present")
		}
	}
	return nil
}

func findByEmail(customerList []cust.Customer, email string) (c cust.Customer, index int, err error) {
	for i, v := range customerList {
		if v.Email == email {
			return v, i, nil
		}
	}
	return cust.Customer{}, -1, errors.New("customer not found")
}

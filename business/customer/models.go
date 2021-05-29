package customer

import "fmt"

type Error string

func (e Error) Error() string {
	return string(e)
}

const (
	invalidName        Error = "please provide valid name"
	invalidEmail       Error = "please provide valid email"
	invalidAge         Error = "please provide valid age, humans dont live that long"
	underAge           Error = "you have to be over 16 years of age to use our services"
	invalidCreditLimit Error = "credit limit cannot be negative"
)

type Customer struct {
	Name, Email string
	Age         int
	CreditLimit float64
}

func (c *Customer) New(name, email string, age int, creditLimit float64) (cs Customer, err error) {
	if len(name) < 1 {
		return cs, invalidName
	}
	if len(email) < 4 {
		return cs, invalidEmail
	}
	if age < 16 {
		return cs, underAge
	}
	if age > 135 {
		return cs, invalidAge
	}
	if creditLimit < 0 {
		return cs, invalidCreditLimit
	}
	cs = Customer{
		Name:        name,
		Email:       email,
		Age:         age,
		CreditLimit: creditLimit,
	}
	return cs, nil
}
func (c Customer) String() string {
	return fmt.Sprintf("Name : %s\nEmail: %s\nAge: %d\nCredit Limit: %f", c.Name, c.Email, c.Age, c.CreditLimit)
}

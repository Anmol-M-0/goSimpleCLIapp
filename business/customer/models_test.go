package customer

import (
	"testing"
)

var customer *Customer

func TestError(t *testing.T) {
	_, err := customer.New("", "", 12, 123.55)
	var invalidName Error = Error("please provide valid name")
	testError(invalidName, err, t)

	_, err = customer.New("Ken", "", 12, 123.55)
	var invalidEmail Error = Error("please provide valid email")
	testError(invalidEmail, err, t)

	_, err = customer.New("Ken", "ken@example.com", 12, 123.55)
	var underAge Error = Error("you have to be over 16 years of age to use our services")
	testError(underAge, err, t)

	_, err = customer.New("Ken", "ken@example.com", 1200, 123.55)
	var invalidAge Error = Error("please provide valid age, humans dont live that long")
	testError(invalidAge, err, t)

	_, err = customer.New("Ken", "ken@example.com", 25, -123.55)
	var invalidCreditLimit Error = Error("credit limit cannot be negative")
	testError(invalidCreditLimit, err, t)

}

func testError(want, got error, t *testing.T) {
	if want != got {
		t.Errorf("test fail, \nwant: '%v'\ngot : '%v' ", want, got)
	}
}

package customer

import (
	"testing"
)

var customer *Customer

func TestError(t *testing.T) {
	_, err := customer.New("", "", 12, 123.55)
	var invalidName Error = Error("please provide valid name")
	if err != invalidName {
		t.Errorf("test fail, \nwant: '%v'\ngot : '%v' ", invalidName, err)
	}
	_, err = customer.New("Ken", "", 12, 123.55)
	var invalidEmail Error = Error("please provide valid email")
	if err != invalidEmail {
		t.Errorf("test fail, \nwant: '%v'\ngot : '%v' ", invalidEmail, err)
	}
	_, err = customer.New("Ken", "ken@example.com", 12, 123.55)
	var underAge Error = Error("you have to be over 16 years of age to use our services")
	if err != underAge {
		t.Errorf("test fail, \nwant: '%v'\ngot : '%v' ", underAge, err)
	}
	_, err = customer.New("Ken", "", 1200, 123.55)
	var invalidAge Error = Error("please provide valid age, humans dont live that long")
	if err != invalidEmail {
		t.Errorf("test fail, \nwant: '%v'\ngot : '%v' ", invalidAge, err)
	}
	_, err = customer.New("Ken", "", 25, -123.55)
	var invalidCreditLimit Error = Error("credit limit cannot be negative")
	if err != invalidEmail {
		t.Errorf("test fail, \nwant: '%v'\ngot : '%v' ", invalidCreditLimit, err)
	}

}

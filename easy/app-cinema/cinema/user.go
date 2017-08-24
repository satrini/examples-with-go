package cinema

import "errors"

type User struct {
	Name   string
	Age    int
	Ticket string
	Price  float64
}

func SetUser(name string, age int) *User {
	var ticket string
	var price float64

	if age < 18 {
		ticket = "meia"
		price = 5.0
	} else {
		ticket = "inteira"
		price = 10.0
	}

	u := &User{Name: name, Age: age, Ticket: ticket, Price: price}
	return u
}

func (u *User) BuyTicket(value float64) (float64, error) {
	if value < u.Price {
		err := errors.New("value not allowed")
		return value, err
	}

	change := (value - u.Price)
	return change, nil
}

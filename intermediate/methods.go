package main

import "fmt"

func main() {

	// Creating an instance of Rectangle
	rect := Rectangle{Width: 10, Height: 5}

	// Calling the method
	println("Area of Rectangle:", rect.Area()) // Output: Area of Rectangle: 50

	// Creating an instance of Account
	acc := Account{Owner: "Radosław", Balance: 100.0}

	// Depositing money
	acc.Deposit(50.0)
	println("Account after deposit:", acc.String()) // Output: Account after deposit: Radosław: 150.00

}

// Defining a method for a struct
type Rectangle struct {
	Width  float64
	Height float64
}

// Method with value receiver
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Account struct {
	Owner   string
	Balance float64
}

func (a *Account) Deposit(amount float64) {
	a.Balance += amount
}

func (a Account) String() string {
	return fmt.Sprintf("%s: %.2f", a.Owner, a.Balance)
}

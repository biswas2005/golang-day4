package functions

import "fmt"

type Details struct {
	Name    string
	ID      int
	Balance float64
}

func (d *Details) deposit(num ...float64) float64 {

	for _, n := range num {
		d.Balance += n
	}
	return d.Balance

}

func (d *Details) withdraw(num ...float64) {
	for _, n := range num {
		if n > d.Balance {
			fmt.Println("Insuficient Balance!")
			return
		}
		d.Balance -= n

	}

}

func (d Details) display() {
	fmt.Println("Name:", d.Name)
	fmt.Println("ID:", d.ID)
	fmt.Println("Balance:", d.Balance)
}

func BankAccount() {
	a := Details{Name: "Abhi", ID: 2244, Balance: 2055.72}
	a.display()
	a.deposit(1000)
	fmt.Println("Balance:", a.Balance)
	a.withdraw(100)
	fmt.Println("Balance:", a.Balance)
}

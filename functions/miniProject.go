package functions

import "fmt"

type Details struct {
	Name    string
	ID      int
	Balance float64
}

//deposit() to deposit certain amount to account
func (d *Details) deposit(num ...float64) float64 {

	for _, n := range num {
		d.Balance += n
	}
	return d.Balance

}

//withdraw() to take amount from account
func (d *Details) withdraw(num ...float64) {
	for _, n := range num {
		
		//checks if the Balance is suficient to withdraw money
		if n > d.Balance {
			fmt.Println("Insuficient Balance!")
			return
		}
		d.Balance -= n

	}

}


//display() gives the account holder details
func (d Details) display() {
	fmt.Println("Name:", d.Name)
	fmt.Println("ID:", d.ID)
	fmt.Println("Balance:", d.Balance)
}

func BankAccount() {

	for {
		fmt.Println("<How can we help you>")
		fmt.Println("1.Account Details.")
		fmt.Println("2.Cash Deposit.")
		fmt.Println("3.Account Withdraw.")

		var userhelp int
		fmt.Println("Enter your choice: ")
		fmt.Scanln(&userhelp)

		a := Details{Name: "Abhi", ID: 2244, Balance: 2022.22}

		if userhelp == 1 {
			a.display()
			return
		}

		var amount float64
		fmt.Println("Enter any amount:")
		fmt.Scanln(&amount)

		switch userhelp {

		case 2:
			fmt.Println("Enter amount to deposit:", amount)

			a.deposit(amount)
			fmt.Println("Updated Balance:", a.deposit())
			return
		case 3:
			fmt.Println("Enter amount to withdraw:", amount)
			a.withdraw(amount)
			fmt.Println("Current Balance:", a.Balance)
			return
		}
	}
}

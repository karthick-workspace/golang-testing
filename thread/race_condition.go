package thread

import "time"

var balance int = 100

func Spend(amt int) {
	b := balance
	time.Sleep(time.Second)

	b -= amt
	balance = b
}

func GetBalance() int {
	return balance
}

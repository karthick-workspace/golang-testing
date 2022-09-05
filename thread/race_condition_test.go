package thread

import (
	"fmt"
	"testing"
	"time"
)

func TestSpend(t *testing.T) {
	go Spend(40)
	go Spend(20)

	time.Sleep(5 * time.Second)

	fmt.Println("GET BALANCE :", GetBalance())
}

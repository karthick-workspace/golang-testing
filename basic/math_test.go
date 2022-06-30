package basic

import (
	"fmt"
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("Test Sum Function", func(t *testing.T) {
		fmt.Println("Running Tests")
		nums := []int{1, 2, 3, 4}
		sum_result := sum(nums)

		if sum_result != 10 {
			t.Errorf("Wanted 10 but received %d", sum_result)
		}
	})

}

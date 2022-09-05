package parallel

import (
	"fmt"
	"testing"
	"time"
)

func TestSomething(t *testing.T) {
	t.Parallel()
}

func TestA(t *testing.T) {
	t.Parallel()
}

func TestB(t *testing.T) {
	fmt.Println("Setup")

	defer fmt.Println("defered teardown")
	t.Run("sub1", func(t *testing.T) {
		t.Parallel()

		time.Sleep(time.Second)
		fmt.Println("Sub1 Done")
	})

	t.Run("sub2", func(t *testing.T) {
		t.Parallel()

		time.Sleep(time.Second)
		fmt.Println("Sub2 Done")
	})

	fmt.Println("teardown")
}

func TestGroup(t *testing.T) {
	fmt.Println("Setup")

	defer fmt.Println("defered teardown")

	t.Run("group", func(t *testing.T) {
		t.Run("sub1", func(t *testing.T) {
			t.Parallel()

			time.Sleep(time.Second)
			fmt.Println("Sub1 Done")
		})

		t.Run("sub2", func(t *testing.T) {
			t.Parallel()

			time.Sleep(time.Second)
			fmt.Println("Sub2 Done")
		})
	})

	fmt.Println("teardown")
}

func TestGotcha(t *testing.T) {
	for i := 0; i < 10; i++ {
		t.Run(fmt.Sprintf("i=%d", i), func(t *testing.T) {
			t.Parallel()
			t.Logf("Testing with i=%d", i)
		})
	}
}

func Test_should_fail_but_not(t *testing.T) {
	testCases := []struct {
		arg  int
		want int
	}{
		{2, 5},
		{3, 9},
		{4, 16},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("arg=%d", tc.arg), func(t *testing.T) {
			t.Parallel()
			t.Logf("Testing with: arg:%d, want=%d", tc.arg, tc.want)

			if tc.arg*tc.arg != tc.want {
				t.Errorf("%d^2 != %d", tc.arg, tc.want)
			}
		})
	}
}

func Test_should_fail(t *testing.T) {
	testCases := []struct {
		arg  int
		want int
	}{
		{2, 5},
		{3, 9},
		{4, 16},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(fmt.Sprintf("arg=%d", tc.arg), func(t *testing.T) {
			t.Parallel()
			t.Logf("Testing with: arg:%d, want=%d", tc.arg, tc.want)

			if tc.arg*tc.arg != tc.want {
				t.Errorf("%d^2 != %d", tc.arg, tc.want)
			}
		})
	}
}

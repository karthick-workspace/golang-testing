package example

import "fmt"

// go test -v
func ExampleHello() {
	greeting := Hello("Karthick")
	fmt.Println(greeting)

	// Output:
	// Hello, Karthick
}

func ExamplePage() {
	checkIns := map[string]bool{
		"P1":  true,
		"P2":  false,
		"P3":  false,
		"P4":  true,
		"P5":  true,
		"AP6": false,
	}
	Page(checkIns)

	// Unordered Output:
	// Paging P2; please see the front desk to check in.
	// Paging P3; please see the front desk to check in.
	// Paging AP6; please see the front desk to check in.

}

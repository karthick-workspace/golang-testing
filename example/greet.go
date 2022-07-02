// Package example is a package for demonstrating examples in source code
package example

import "fmt"

func Hello(name string) string {
	return fmt.Sprintf("Hello, %s", name)
}

//Page will print out a message asking each person who hasn;t checked in to do so
func Page(checkIns map[string]bool) {
	for name, checkIn := range checkIns {
		if !checkIn {
			fmt.Printf("Paging %s; please see the front desk to check in.\n", name)
		}
	}
}

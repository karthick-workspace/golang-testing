package main

import "fmt"

type Department struct {
	DeptId   string
	DeptName string
}
type Employee struct {
	EmpName string
	EmpNum  string
	Dept    *Department
}

func main() {
	emp1 := Employee{}
	emp1.Dept = &Department{} //  emp1.Dept = new(Department)
	updateEmpDetails(&emp1)

	// new allocates zeroed storage in memory and returns a pointer to it
	emp2 := new(Employee)
	emp2.Dept = new(Department)
	updateEmpDetails(emp2)

	fmt.Println(emp1)
	fmt.Println(emp2)
	fmt.Println(*emp2)
}

func updateEmpDetails(emp *Employee) {
	emp.EmpNum = "1"
	emp.EmpName = "KKKKVVV"

	emp.Dept.DeptId = "D1"
	emp.Dept.DeptName = "DDD1"
}

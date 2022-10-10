package main

import "fmt"

type MyDepartment struct {
	DeptId   string
	DeptName string
}

type HrDepartment struct {
	CompanyName   string
	EmployeeCount int64
	MyDepartment
}

func (dept MyDepartment) Show() {
	fmt.Println(dept)
}

func (hr HrDepartment) Show() {
	fmt.Println(hr)
}

func main() {
	hr := HrDepartment{
		MyDepartment: MyDepartment{
			DeptId:   "D1",
			DeptName: "HR",
		},
		CompanyName:   "ABC",
		EmployeeCount: 1000,
	}

	hr.MyDepartment.Show()
	hr.Show()
}

package main

import (
	"errors"
	"fmt"
)

const (
	HR = "HR"
	IT = "IT"
	MG = "Management"
)

type Employee struct {
	ID         int
	Name       string
	Age        int
	Department string
}

var employees []Employee

func addEmp(id int, name string, age int, department string) error {
	if age <= 18 {
		return errors.New("Age must be greater than 18")
	}
	for _, emp := range employees {
		if emp.ID == id {
			return errors.New("Employee ID must be unique")
		}
	}
	employees = append(employees, Employee{ID: id, Name: name, Age: age, Department: department})
	return nil
}
func searchEmp(searchTerm string) (*Employee, error) {
	for _, emp := range employees {
		if fmt.Sprint(emp.ID) == searchTerm || emp.Name == searchTerm {
			return &emp, nil
		}
	}
	return nil, errors.New("Employee not found")
}
func listEmpByDepartment(department string) []Employee {
	var departmentEmployees []Employee
	for _, emp := range employees {
		if emp.Department == department {
			departmentEmployees = append(departmentEmployees, emp)
		}
	}
	return departmentEmployees
}

func countEmp(department string) int {
	count := 0
	for _, emp := range employees {
		if emp.Department == department {
			count++
		}
	}
	return count
}
func displayEmp(department string) {
	employeesInDept := listEmpByDepartment(department)
	if len(employeesInDept) == 0 {
		fmt.Printf("No employees found in the %s department.\n", department)
	} else {
		fmt.Printf("Employees in %s Department:\n", department)
		for _, emp := range employeesInDept {
			fmt.Printf("ID: %d, Name: %s, Age: %d, Department: %s\n", emp.ID, emp.Name, emp.Age, emp.Department)
		}
	}
}

func main() {
	// Adding some employees
	err := addEmp(1, "abc", 19, MG)
	if err != nil {
		fmt.Println("Error:", err)
	}
	err = addEmp(2, "def", 25, IT)
	if err != nil {
		fmt.Println("Error:", err)
	}
	err = addEmp(3, "ghi", 17, HR)
	if err != nil {
		fmt.Println("Error:", err)
	}
	err = addEmp(4, "jkl", 29, HR)
	if err != nil {
		fmt.Println("Error:", err)
	}
	searchTerm := "def"
	employee, err := searchEmp(searchTerm)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Found employee: ID: %d, Name: %s, Age: %d, Department: %s\n", employee.ID, employee.Name, employee.Age, employee.Department)
	}
	displayEmp(IT)
	hrCount := countEmp(HR)
	fmt.Printf("Total employees in HR department: %d\n", hrCount)
	_, err = searchEmp("xyz")
	if err != nil {
		fmt.Println("Error:", err)
	}
}

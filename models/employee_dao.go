package models

import (
	"echo-mysql-api/db"
	"errors"
	"fmt"
	"log"
)

type employeeDaoInterface interface {
	GetEmployee(userId int64) (*Employee, error)
}

func init() {
	EmployeeDao = &employeeDao{}
}

type employeeDao struct{}

var EmployeeDao employeeDaoInterface

func GetAllEmployees() Employees {
	con := db.CreateConnection()
	sqlStatement := "SELECT id, employee_name, employee_salary, employee_age FROM Employees ORDER BY id"

	rows, err := con.Query(sqlStatement)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	result := Employees{}

	for rows.Next() {
		employee := Employee{}
		err = rows.Scan(&employee.Id, &employee.Name, &employee.Salary, &employee.Age)
		if err != nil {
			fmt.Println(err)
		}
		result.Employees = append(result.Employees, employee)
	}
	return result
}

func (e *employeeDao) GetEmployee(userId int64) (*Employee, error) {
	log.Println("We're connecting to the external database!")
	con := db.CreateConnection()
	sqlStatement := fmt.Sprintf("SELECT id, employee_name, employee_salary, employee_age FROM Employees WHERE id=%d", userId)

	var result Employee
	rows, err := con.Query(sqlStatement)
	if err != nil {
		return nil, err
	}

	if rows.Next() {
		err = rows.Scan(&result.Id, &result.Name, &result.Salary, &result.Age)
		if err != nil {
			return nil, err
		}
		return &result, nil
	} else {
		return nil, errors.New("User is not found!")
	}
}

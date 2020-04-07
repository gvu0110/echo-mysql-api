package models

import (
	"echo-mysql-api/db"
	"fmt"
)

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

func GetEmployee(userId int64) Employee {
	con := db.CreateConnection()
	sqlStatement := fmt.Sprintf("SELECT id, employee_name, employee_salary, employee_age FROM Employees WHERE id=%d", userId)

	rows, err := con.Query(sqlStatement)
	if err != nil {
		fmt.Println(err)
	}
	defer rows.Close()
	// TODO: If rows is nill

	// TODO: If rows has 2 more employees

	result := Employee{}

	for rows.Next() {
		err = rows.Scan(&result.Id, &result.Name, &result.Salary, &result.Age)
		if err != nil {
			fmt.Println(err)
		}
	}
	return result
}

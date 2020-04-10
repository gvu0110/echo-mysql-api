package models

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

var EmployeeDaoMock employeesDaoMock
var getEmployeeFunction func(int64) (*Employee, error)

type employeesDaoMock struct{}

func init() {
	EmployeeDao = &EmployeeDaoMock
}

func (m *employeesDaoMock) GetEmployee(userId int64) (*Employee, error) {
	return getEmployeeFunction(userId)
}

func TestGetEmployeeNotFound(t *testing.T) {
	getEmployeeFunction = func(int64) (*Employee, error) {
		return nil, errors.New("User 0 is not found!")
	}

	employee, err := EmployeeDao.GetEmployee(0)
	if employee != nil {
		t.Error("We're expecting an error user 0 is not found!")
	}

	if err == nil {
		t.Error("We're expecting an not found error when employee ID is 0!")
	}
}

func TestGetUserNoError(t *testing.T) {
	getEmployeeFunction = func(int64) (*Employee, error) {
		return &Employee{
			Id:     1,
			Name:   "Adam Vu",
			Salary: 1234.56,
			Age:    25,
		}, nil
	}
	employee, err := EmployeeDao.GetEmployee(1)
	assert.Nil(t, err)
	assert.NotNil(t, employee)
	assert.EqualValues(t, 1, employee.Id)
	assert.EqualValues(t, "Adam Vu", employee.Name)
	assert.EqualValues(t, 1234.56, employee.Salary)
	assert.EqualValues(t, 25, employee.Age)
}

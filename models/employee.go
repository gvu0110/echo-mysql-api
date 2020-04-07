package models

type Employee struct {
	Id     uint64  `json:"id"`
	Name   string  `json:"name"`
	Salary float64 `json:"salary"`
	Age    int     `json:"age"`
}

type Employees struct {
	Employees []Employee `json:"employees"`
}

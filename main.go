package main

import (
	"encoding/json"
	"fmt"
)

type role string

const (
	Admin role = "admin"
	Hr    role = "hr"
	User  role = "user"
)

type Employee struct {
	ID       string
	Name     string
	Salary   int
	Password string
	Rights   map[string]bool
	Boss     *Employee
	Role     role
}

func (e Employee) MarshalJSON() ([]byte, error) {
	serialize := fromRole(e)

	return json.Marshal(serialize)
}

func fromRole(e Employee) interface{} {
	result := struct {
		ID       string          `json:"id,omitempty"`
		Name     string          `json:"name,omitempty"`
		Salary   *int            `json:"salary,omitempty"`
		Password string          `json:"password,omitempty"`
		Rights   map[string]bool `json:"rights,omitempty"`
		Boss     *Employee       `json:"boss,omitempty"`
	}{
		ID:       e.ID,
		Name:     e.Name,
		Salary:   &e.Salary,
		Password: e.Password,
		Rights:   e.Rights,
		Boss:     e.Boss,
	}

	if e.Role != Admin {
		result.Password = ""
		result.Rights = nil
	}

	if e.Role != Admin && e.Role != Hr {
		result.Salary = nil
		result.Boss = nil
	}
	return result
}

func main() {
	fuser := Employee{
		"1",
		"Jack",
		100000,
		"password321",
		map[string]bool{"create": false, "update": false},
		&Employee{
			"2",
			"John",
			120000,
			"pwd",
			map[string]bool{"create": true, "update": true},
			nil,
			Admin,
		},
		User,
	}

	fhr := fuser
	fhr.Role = Hr
	fadmin := fuser
	fadmin.Role = Admin

	buser, err := json.MarshalIndent(fuser, "", "  ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Filtering with role user: ")
	fmt.Println(string(buser))

	bhr, err := json.MarshalIndent(fhr, "", "  ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("\nFiltering with role hr: ")
	fmt.Println(string(bhr))

	badmin, err := json.MarshalIndent(fadmin, "", "  ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("\nFiltering with role admin: ")
	fmt.Println(string(badmin))
}

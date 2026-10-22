package main

import (
	"fmt"
)

const (
	DefaultMinTemp = 15
	DefaultMaxTemp = 30
	MinCount       = 1
	MaxCount       = 1000
)

type TemperatureManager struct {
	minTemp int
	maxTemp int
}

func NewTemperatureManager() *TemperatureManager {
	return &TemperatureManager{
		minTemp: DefaultMinTemp,
		maxTemp: DefaultMaxTemp,
	}
}

func (tm *TemperatureManager) AddConstraint(operation string, value int) error {
	if operation == ">=" {
		if value > tm.minTemp {
			tm.minTemp = value
		}
	} else if operation == "<=" {
		if value < tm.maxTemp {
			tm.maxTemp = value
		}
	} else {
		return fmt.Errorf("unknown operation: %s", operation)
	}
	return nil
}

func (tm *TemperatureManager) GetOptimalTemperature() int {
	if tm.minTemp <= tm.maxTemp {
		return tm.minTemp
	}
	return -1
}

func processEmployees(employeesCount int) error {
	tm := NewTemperatureManager()

	for range employeesCount {
		var (
			operation string
			val       int
		)

		if _, err := fmt.Scan(&operation, &val); err != nil {
			return fmt.Errorf("error while reading employees data: %w", err)
		}

		if err := tm.AddConstraint(operation, val); err != nil {
			return fmt.Errorf("error processing constraint - %w", err)
		}
		fmt.Println(tm.GetOptimalTemperature())
	}
	return nil
}

func main() {
	var departmentNumber int
	if _, err := fmt.Scan(&departmentNumber); err != nil {
		fmt.Printf("error while reading number of departments: %v\n", err)
		return
	}

	if departmentNumber < MinCount || departmentNumber > MaxCount {
		fmt.Println("incorrect number of departments")
		return
	}

	for range departmentNumber {
		var employeesCount int
		if _, err := fmt.Scan(&employeesCount); err != nil {
			fmt.Printf("error while reading number of employees: %v\n", err)
			return
		}

		if employeesCount < MinCount || employeesCount > MaxCount {
			fmt.Println("incorrect employees number")
			return
		}

		if err := processEmployees(employeesCount); err != nil {
			fmt.Println(err)
			return
		}
	}
}

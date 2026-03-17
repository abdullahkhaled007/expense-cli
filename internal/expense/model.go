package expense

import (
	"errors"
	"time"
)

type Expense struct {
	ID          int
	Name        string
	Date        time.Time
	Price       float64
	Description string
}

func NewExpense(name string, price float64, description string) (*Expense, error) {
	if name == "" {
		return nil, errors.New("name cannot be empty")
	}

	if price <= 0 {
		return nil, errors.New("price must be greater than zero")
	}

	return &Expense{
		Name:        name,
		Date:        time.Now(),
		Price:       price,
		Description: description,
	}, nil
}

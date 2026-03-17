package expense

import (
	"encoding/json"
	"errors"
	"os"
	"slices"
)

type Store struct {
	expenses []Expense
	nextID   int
	filePath string
}

func NewStore() *Store {
	return &Store{
		expenses: []Expense{},
		nextID:   1,
	}
}

// NewPersistentStore creates a store that saves to and loads from a JSON file.
// If the file exists it will be loaded; the store will write the full list
// after Add/Delete operations.
func NewPersistentStore(path string) *Store {
	s := &Store{
		expenses: []Expense{},
		nextID:   1,
		filePath: path,
	}

	// Attempt to load existing file
	if path != "" {
		if data, err := os.ReadFile(path); err == nil {
			var exs []Expense
			if err := json.Unmarshal(data, &exs); err == nil {
				s.expenses = exs
				// compute nextID
				max := 0
				for _, e := range exs {
					if e.ID > max {
						max = e.ID
					}
				}
				s.nextID = max + 1
			}
		}
	}

	return s
}

func (s *Store) Add(expense *Expense) (*Expense, error) {
	if expense == nil {
		return nil, errors.New("expense must not be nil")
	}
	expense.ID = s.nextID
	s.nextID++
	s.expenses = append(s.expenses, *expense)
	// Persist if configured
	if s.filePath != "" {
		if err := s.persist(); err != nil {
			return expense, err
		}
	}
	return expense, nil
}

func (s *Store) List() ([]Expense, error) {
	return s.expenses, nil
}

func (s *Store) Delete(id int) (Expense, error) {
	var index int = -1
	for i, current := range s.expenses {
		if current.ID == id {
			index = i
			break
		}
	}
	if index == -1 {
		return Expense{}, errors.New("expense not found")
	}
	deletedExpense := s.expenses[index]
	s.expenses = slices.Delete(s.expenses, index, index+1)
	// Persist if configured
	if s.filePath != "" {
		if err := s.persist(); err != nil {
			return deletedExpense, err
		}
	}
	return deletedExpense, nil
}

func (s *Store) persist() error {
	if s.filePath == "" {
		return nil
	}
	data, err := json.MarshalIndent(s.expenses, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(s.filePath, data, 0644)
}

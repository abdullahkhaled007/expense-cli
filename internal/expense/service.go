package expense

type Service struct {
	store *Store
}

func NewService() *Service {
	return &Service{
		store: NewStore(),
	}
}

// NewServiceWithFile returns a service backed by a persistent JSON store at path.
func NewServiceWithFile(path string) *Service {
	return &Service{
		store: NewPersistentStore(path),
	}
}

func (s *Service) Add(name string, price float64, description string) (*Expense, error) {
	e, err := NewExpense(name, price, description)
	if err != nil {
		return nil, err
	}
	addedValue, err := s.store.Add(e)
	if err != nil {
		return nil, err
	}
	return addedValue, nil
}

func (s *Service) List() ([]Expense, error) {
	return s.store.List()
}

func (s *Service) Delete(id int) (Expense, error) {
	return s.store.Delete(id)
}

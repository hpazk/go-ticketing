package transaction

type Services interface {
	SaveTransaction(req *request) (Transaction, error)
	FetchTransactions() ([]Transaction, error)
	FetchTransaction(id uint) (Transaction, error)
	EditTransaction(id uint) (Transaction, error)
	RemoveTransaction(id uint) error
	FetchTransactionsByEvent(id uint) ([]Transaction, error)
}

type services struct {
	repo repository
}

func transactionService(repo repository) *services {
	return &services{repo}
}

func (s *services) SaveTransaction(req *request) (Transaction, error) {
	var tsx Transaction
	return tsx, nil
}

func (s *services) FetchTransactions() ([]Transaction, error) {
	tsxs, _ := s.repo.Fetch()
	return tsxs, nil
}

func (s *services) FetchTransaction(id uint) (Transaction, error) {
	var tsx Transaction
	return tsx, nil
}

func (s *services) EditTransaction(id uint) (Transaction, error) {
	var tsx Transaction
	return tsx, nil
}

func (s *services) RemoveTransaction(id uint) error {
	return nil
}

// Transaction - Event
func (s *services) FetchTransactionsByEvent(id uint) ([]Transaction, error) {
	tsxs, _ := s.repo.FindByEventID(id)

	return tsxs, nil
}

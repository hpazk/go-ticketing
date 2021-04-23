package transaction

import "github.com/hpazk/go-booklib/database/model"

type Services interface {
	SaveTransaction(req *request) (model.Transaction, error)
	FetchTransactions() ([]model.Transaction, error)
	FetchTransaction(id uint) (model.Transaction, error)
	EditTransaction(id uint) (model.Transaction, error)
	RemoveTransaction(id uint) error
	FetchTransactionsByEvent(id uint) ([]model.Transaction, error)
}

type services struct {
	repo repository
}

func transactionService(repo repository) *services {
	return &services{repo}
}

func (s *services) SaveTransaction(req *request) (model.Transaction, error) {
	var tsx model.Transaction
	return tsx, nil
}

func (s *services) FetchTransactions() ([]model.Transaction, error) {
	tsxs, _ := s.repo.Fetch()
	return tsxs, nil
}

func (s *services) FetchTransaction(id uint) (model.Transaction, error) {
	var tsx model.Transaction
	return tsx, nil
}

func (s *services) EditTransaction(id uint) (model.Transaction, error) {
	var tsx model.Transaction
	return tsx, nil
}

func (s *services) RemoveTransaction(id uint) error {
	return nil
}

// model.Transaction - Event
func (s *services) FetchTransactionsByEvent(id uint) ([]model.Transaction, error) {
	tsxs, _ := s.repo.FindByEventID(id)

	return tsxs, nil
}

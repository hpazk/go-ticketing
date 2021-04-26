package transaction

import (
	"github.com/hpazk/go-ticketing/apps/event"
	"github.com/hpazk/go-ticketing/database/model"
	"github.com/hpazk/go-ticketing/helper"
)

type Services interface {
	SaveTransaction(req *request, participant model.User) (model.Transaction, error)
	FetchTransactions() ([]model.Transaction, error)
	FetchTransaction(id uint) (model.Transaction, error)
	EditTransaction(id uint, req *updateRequest) (model.Transaction, error)
	UploadPaymentOrder(participanID uint, imagePath string) (model.Transaction, error)
	RemoveTransaction(id uint) error
	FetchTransactionsByEvent(id uint) ([]model.Transaction, error)
}

type services struct {
	repo repository
}

func TransactionService() *services {
	repo := TransactionRepository()
	return &services{repo}
}

func (s *services) SaveTransaction(req *request, participant model.User) (model.Transaction, error) {
	// TODO jwt: id, email

	var transaction model.Transaction
	transaction.EventID = req.EventID
	transaction.ParticipantID = participant.ID
	transaction.Amount = req.Amount

	savedTransaction, err := s.repo.Store(transaction)
	if err != nil {
		return savedTransaction, nil
	}

	eventService := event.EventService()
	orderedEvent, err := eventService.FetchEvent(savedTransaction.EventID)
	if err != nil {
		return savedTransaction, nil
	}

	emailBody := helper.PaymentOrderTemplate(orderedEvent)
	helper.SendEmail(participant.Email, "Webinar Payment Order", emailBody)

	return savedTransaction, nil
}

func (s *services) FetchTransactions() ([]model.Transaction, error) {
	tsxs, _ := s.repo.Fetch()
	return tsxs, nil
}

func (s *services) EditTransaction(id uint, req *updateRequest) (model.Transaction, error) {
	transaction, err := s.repo.FindById(id)
	if err != nil {
		return transaction, nil
	}

	transaction.StatusPayment = req.StatusPayment
	editedTransaction, err := s.repo.Update(transaction)
	if err != nil {
		return editedTransaction, nil
	}

	return transaction, nil
}

func (s *services) FetchTransaction(id uint) (model.Transaction, error) {
	var tsx model.Transaction
	return tsx, nil
}

func (s *services) UploadPaymentOrder(participanID uint, imagePath string) (model.Transaction, error) {
	transaction, err := s.repo.FindByParticipant(participanID)
	if err != nil {
		return transaction, nil
	}

	transaction.ImagePath = imagePath
	editedTransaction, err := s.repo.Update(transaction)
	if err != nil {
		return editedTransaction, nil
	}

	return transaction, nil
}

func (s *services) RemoveTransaction(id uint) error {
	return nil
}

// model.Transaction - Event
func (s *services) FetchTransactionsByEvent(id uint) ([]model.Transaction, error) {
	tsxs, _ := s.repo.FindByEventID(id)

	return tsxs, nil
}

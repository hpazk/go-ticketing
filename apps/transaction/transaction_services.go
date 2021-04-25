package transaction

import (
	"github.com/hpazk/go-booklib/apps/event"
	"github.com/hpazk/go-booklib/database/model"
	"github.com/hpazk/go-booklib/helper"
)

type Services interface {
	SaveTransaction(req *request) (model.Transaction, error)
	FetchTransactions() ([]model.Transaction, error)
	FetchTransaction(id uint) (model.Transaction, error)
	UploadPaymentOrder(id uint, imagePath string) (model.Transaction, error)
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

func (s *services) SaveTransaction(req *request) (model.Transaction, error) {
	// TODO jwt: id, email
	var id uint = 1
	var participantEmail string = "email@email.com"

	var transaction model.Transaction
	transaction.EventID = req.EventID
	transaction.ParticipantID = id

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
	helper.SendEmail(participantEmail, "Webinar Payment Order", emailBody)

	return savedTransaction, nil
}

func (s *services) FetchTransactions() ([]model.Transaction, error) {
	tsxs, _ := s.repo.Fetch()
	return tsxs, nil
}

func (s *services) FetchTransaction(id uint) (model.Transaction, error) {
	var tsx model.Transaction
	return tsx, nil
}

func (s *services) UploadPaymentOrder(id uint, imagePath string) (model.Transaction, error) {
	var transaction model.Transaction
	transaction.ID = id
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

package transaction

import (
	"github.com/hpazk/go-ticketing/apps/event"
	"github.com/hpazk/go-ticketing/database/model"
	"github.com/hpazk/go-ticketing/helper"
	"github.com/hpazk/go-ticketing/template"
)

type Services interface {
	SaveTransaction(req *request, participant model.User) error
	FetchTransactions() ([]model.Transaction, error)
	FetchTransaction(id uint) (model.Transaction, error)
	EditTransaction(id uint, req *updateRequest) error
	UploadPaymentOrder(participanID uint, imagePath string) error
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

func (s *services) SaveTransaction(req *request, participant model.User) error {
	var transaction model.Transaction
	eventService := event.EventService()

	orderedEvent, err := eventService.FetchEvent(req.EventID)
	if err != nil {
		return err
	}

	transaction.EventID = req.EventID
	transaction.ParticipantID = participant.ID

	err = s.repo.Store(transaction)
	if err != nil {
		return err
	}

	emailBody := template.InvoiceTemplate(orderedEvent)
	go helper.SendEmail(participant.Email, "Webinar Payment Order", emailBody)

	return nil
}

func (s *services) FetchTransactions() ([]model.Transaction, error) {
	tsxs, _ := s.repo.Fetch()
	return tsxs, nil
}

func (s *services) EditTransaction(id uint, req *updateRequest) error {
	transaction, err := s.repo.FindById(id)
	if err != nil {
		return err
	}

	transaction.StatusPayment = req.StatusPayment
	transaction.Amount = req.Amount
	err = s.repo.Update(transaction)
	if err != nil {
		return err
	}

	participant, err := s.repo.FindDetil(id)
	if err != nil {
		return nil
	}

	emailBody := template.PaymentSuccessLayout(participant)
	helper.SendEmail(participant.Email, "Webinar Detil", emailBody)

	return nil
}

func (s *services) FetchTransaction(id uint) (model.Transaction, error) {
	var tsx model.Transaction
	return tsx, nil
}

func (s *services) UploadPaymentOrder(participanID uint, imagePath string) error {
	transaction, err := s.repo.FindByParticipant(participanID)
	if err != nil {
		return err
	}

	transaction.ImagePath = imagePath
	err = s.repo.Update(transaction)
	if err != nil {
		return err
	}

	return nil
}

func (s *services) RemoveTransaction(id uint) error {
	return nil
}

// model.Transaction - Event
func (s *services) FetchTransactionsByEvent(id uint) ([]model.Transaction, error) {
	tsxs, _ := s.repo.FindByEventID(id)

	return tsxs, nil
}

package service

import (
	"github.com/fabioqmarsiaj/payment-gateway/go-gateway/internal/domain"
	"github.com/fabioqmarsiaj/payment-gateway/go-gateway/internal/dto"
)

type InvoiceService struct {
	invoiceRepository domain.InvoiceRepository
	accountservice    AccountService
}

func NewInvoiceService(
	invoiceRepository domain.InvoiceRepository,
	accountService AccountService) *InvoiceService {
	return &InvoiceService{
		invoiceRepository: invoiceRepository,
		accountservice:    accountService,
	}
}

func (s *InvoiceService) Create(input *dto.CreateInvoiceInput) (*dto.InvoiceOutput, error) {
	accountOutput, err := s.accountservice.FindByAPIKey(input.APIKey)
	if err != nil {
		return nil, err
	}

	invoice, err := dto.ToInvoice(input, accountOutput.ID)
	if err != nil {
		return nil, err
	}

	if err := invoice.Process(); err != nil {
		return nil, err
	}

	if invoice.Status == domain.StatusApproved {
		_, err = s.accountservice.UpdateBalance(input.APIKey, invoice.Amount)
		if err != nil {
			return nil, err
		}
	}

	if err := s.invoiceRepository.Save(invoice); err != nil {
		return nil, err
	}

	return dto.FromInvoice(invoice), nil
}

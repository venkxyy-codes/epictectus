package communication

import (
	"context"
	"epictectus/contract"
	"epictectus/domain"
)

// MockCommunicationService is a mock implementation of the CommunicationService interface.
type MockCommunicationService struct {
	SendPaymentLinkToCustomerOnWhatsappFunc func(ctx context.Context, req contract.SendPaymentLinkToCustomer, provider domain.WhatsappProvider) error
}

func (m *MockCommunicationService) SendPaymentLinkToCustomerOnWhatsapp(ctx context.Context, req contract.SendPaymentLinkToCustomer, provider domain.WhatsappProvider) error {
	if m.SendPaymentLinkToCustomerOnWhatsappFunc != nil {
		return m.SendPaymentLinkToCustomerOnWhatsappFunc(ctx, req, provider)
	}
	return nil
}

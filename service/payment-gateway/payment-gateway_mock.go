package payment_gateway

import (
	"context"
	"epictectus/contract"
	"epictectus/domain"
)

type MockPaymentGatewayService struct {
	CreateStandardPaymentLinkRazorpayCalled bool
	CreateStandardPaymentLinkRazorpayErr    error
}

func (m *MockPaymentGatewayService) CreateStandardPaymentLinkRazorpay(ctx context.Context, req contract.CreateStandardPaymentLink, notifyCrm bool, crmProvider domain.CrmProvider, notifyCustomer bool, whatsappProvider domain.WhatsappProvider) error {
	m.CreateStandardPaymentLinkRazorpayCalled = true
	return m.CreateStandardPaymentLinkRazorpayErr
}

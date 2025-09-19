package webhook_processor

import (
	"context"
	"epictectus/contract"
	"epictectus/service/crm"
	payment_gateway "epictectus/service/payment-gateway"
	"github.com/stretchr/testify/assert"
	"testing"
)

// Mock dependencies and setup
// ...existing code...

func TestWebhookProcessorService_HandleLeadsquaredWebhook(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	mockCrm := &crm.MockCrmService{}
	mockPg := &payment_gateway.MockPaymentGatewayService{}
	service := NewWebhookProcessorService(mockCrm, mockPg)

	t.Run("happy flow", func(t *testing.T) {
		webhook := contract.LeadsquaredActivityWebhook{
			ActivityEvent:     "1",
			Data:              struct{ MxCustom1, MxCustom2 string }{MxCustom1: "100", MxCustom2: "INR"},
			Current:           struct{ FirstName, Phone, EmailAddress string }{"John", "9876543210", "john@example.com"},
			RelatedProspectId: "prospect123",
		}
		service.HandleLeadsquaredWebhook(ctx, webhook)
		assert.True(t, mockPg.CreateStandardPaymentLinkRazorpayCalled)
	})

	t.Run("invalid amount", func(t *testing.T) {
		webhook := contract.LeadsquaredActivityWebhook{
			ActivityEvent:     "1",
			Data:              struct{ MxCustom1, MxCustom2 string }{MxCustom1: "invalid", MxCustom2: "INR"},
			Current:           struct{ FirstName, Phone, EmailAddress string }{"John", "9876543210", "john@example.com"},
			RelatedProspectId: "prospect123",
		}
		service.HandleLeadsquaredWebhook(ctx, webhook)
		assert.False(t, mockPg.CreateStandardPaymentLinkRazorpayCalled)
	})
}

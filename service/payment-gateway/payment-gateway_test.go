package payment_gateway

import (
	"context"
	"epictectus/contract"
	"epictectus/domain"
	"epictectus/service/crm"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

// Mock dependencies and setup
// ...existing code...

func TestPaymentGatewayService_CreateStandardPaymentLinkRazorpay(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	mockCrm := &crm.MockCrmService{}
	service := NewPaymentGatewayService(mockCrm)

	t.Run("happy flow", func(t *testing.T) {
		err := service.CreateStandardPaymentLinkRazorpay(ctx, contract.CreateStandardPaymentLink{}, true, domain.Leadsquared)
		assert.NoError(t, err)
	})

	t.Run("provider not found", func(t *testing.T) {
		pg := service.(*paymentGatewayService)
		pg.paymentGatewayDetails = map[domain.PaymentProvider]interface{}{}
		err := pg.CreateStandardPaymentLinkRazorpay(ctx, contract.CreateStandardPaymentLink{}, true, domain.Leadsquared)
		assert.Error(t, err)
	})

	t.Run("invalid details struct", func(t *testing.T) {
		pg := service.(*paymentGatewayService)
		pg.paymentGatewayDetails = map[domain.PaymentProvider]interface{}{domain.PaymentProvider(strconv.Itoa(1)): "invalid"}
		err := pg.CreateStandardPaymentLinkRazorpay(ctx, contract.CreateStandardPaymentLink{}, true, domain.Leadsquared)
		assert.Error(t, err)
	})
}

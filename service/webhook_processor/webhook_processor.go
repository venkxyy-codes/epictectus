package webhook_processor

import (
	"context"
	"epictectus/blog"
	"epictectus/contract"
	"epictectus/domain"
	"epictectus/service/crm"
	payment_gateway "epictectus/service/payment-gateway"
	"strconv"
)

type webhookProcessorService struct {
	crmService            crm.CrmService
	paymentGatewayService payment_gateway.PaymentGatewayService
}

type WebhookProcessorService interface {
	HandleLeadsquaredWebhook(ctx context.Context, webhook contract.LeadsquaredActivityWebhook)
}

func NewWebhookProcessorService(crmService crm.CrmService, paymentGatewayService payment_gateway.PaymentGatewayService) WebhookProcessorService {
	return &webhookProcessorService{
		crmService:            crmService,
		paymentGatewayService: paymentGatewayService,
	}
}

func (w *webhookProcessorService) HandleLeadsquaredWebhook(ctx context.Context, webhook contract.LeadsquaredActivityWebhook) {
	amount, err := strconv.ParseInt(webhook.Data.MxCustom1, 10, 64)
	if err != nil {
		blog.ErrorCtx(ctx, err, "err-failed-to-parse-amount")
		return
	}
	switch webhook.ActivityEvent {
	case strconv.FormatInt(int64(domain.GetPaymentRazorpayLink), 10):
		_ = w.paymentGatewayService.CreateStandardPaymentLinkRazorpay(ctx,
			contract.CreateStandardPaymentLink{
				Amount:          amount * 100, // converting to paise
				ProspectId:      webhook.RelatedProspectId,
				Currency:        webhook.Data.MxCustom2,
				CustomerName:    webhook.Current.FirstName,
				CustomerContact: webhook.Current.Phone,
				CustomerEmail:   webhook.Current.EmailAddress,
				NotifySms:       "true",
				NotifyEmail:     "false",
			}, true, domain.Leadsquared)
	}
}

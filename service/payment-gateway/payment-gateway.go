package payment_gateway

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"epictectus/blog"
	"epictectus/clients"
	"epictectus/config"
	"epictectus/contract"
	"epictectus/domain"
	"epictectus/service/communication"
	"epictectus/service/crm"
	"epictectus/utils"
	"epictectus/view"
	"fmt"
	"github.com/google/uuid"
	"net/http"
	"slices"
	"time"
)

type paymentGatewayService struct {
	crmService            crm.CrmService
	communicationService  communication.CommunicationService
	baseClient            clients.BaseClient
	paymentGatewayDetails map[domain.PaymentProvider]interface{}
}

type PaymentGatewayService interface {
	CreateStandardPaymentLinkRazorpay(ctx context.Context, createStandardPaymentLinkRequest contract.CreateStandardPaymentLink, notifyCrm bool, crmProvider domain.CrmProvider, notifyCustomer bool, whatsappProvider domain.WhatsappProvider) error
}

func populatePaymentGatewayConfig() map[domain.PaymentProvider]interface{} {
	return map[domain.PaymentProvider]interface{}{
		domain.Razorpay: struct {
			Credentials config.RazorpayCredentials
			HttpConfig  config.RazorpayHttpConfig
		}{
			Credentials: config.GetRazorpayCredentials(),
			HttpConfig:  config.GetRazorpayHttpConfig(),
		},
	}
}

func NewPaymentGatewayService(crmService crm.CrmService, commService communication.CommunicationService, baseClient clients.BaseClient) PaymentGatewayService {
	return &paymentGatewayService{
		crmService:            crmService,
		communicationService:  commService,
		baseClient:            baseClient,
		paymentGatewayDetails: populatePaymentGatewayConfig(),
	}
}

func (r *paymentGatewayService) CreateStandardPaymentLinkRazorpay(ctx context.Context, req contract.CreateStandardPaymentLink, notifyCrm bool, crmProvider domain.CrmProvider, notifyCustomer bool, whatsappProvider domain.WhatsappProvider) error {
	paymentGatewayDetails, ok := r.paymentGatewayDetails[domain.Razorpay]
	if !ok {
		return fmt.Errorf("err-payment-gateway-credentials-not-identified")
	}
	paymentGatewayDetailsStruct, ok := paymentGatewayDetails.(struct {
		Credentials config.RazorpayCredentials
		HttpConfig  config.RazorpayHttpConfig
	})
	if !ok {
		return fmt.Errorf("err-payment-gateway-config-and-credentials-malformed")
	}
	paymentGatewayPayload := contract.CreateStandardPaymentLinkRequestRazorpay{
		Amount:      req.Amount,
		Currency:    req.Currency,
		ReferenceID: uuid.New().String(),
		Description: "",
		ExpireBy:    time.Now().Add(time.Hour * 12).UnixMilli(),
		Notes:       nil,
		CustomerDetails: struct {
			Name    string `json:"name"`
			Contact string `json:"contact"`
			Email   string `json:"email"`
		}(struct {
			Name    string
			Contact string
			Email   string
		}{
			Name:    req.CustomerName,
			Contact: req.CustomerContact,
			Email:   req.CustomerEmail,
		}),
		Notify: struct {
			Sms   string `json:"sms,omitempty"`
			Email string `json:"email,omitempty"`
		}(struct {
			Sms   string
			Email string
		}{
			Sms:   req.NotifySms,
			Email: req.NotifyEmail,
		}),
		AcceptPartial:  false,
		CallbackUrl:    "https://ulc-api.leadsquaredapps.com/v1/UniversalLeadSync.svc/RealtimeWebhook/Connector/78089/3893d0d6-54f6-4b09-a4d9-ff540c2b03c6/0a1bcb55f3ff4c58be5b54dd4686e7aa0db8aff2c2c6c784d3e9f2b80ecbb552917ebc4455fb1831e92db6456ca69e56",
		CallbackMethod: "get",
	}
	blog.InfoCtx(ctx, "info-credentials", "credentials", paymentGatewayDetailsStruct.Credentials)
	blog.InfoCtx(ctx, "info-http-config", "httpConfig", paymentGatewayDetailsStruct.HttpConfig)
	headers := map[string]string{
		"Authorization": "Basic " + base64.StdEncoding.EncodeToString([]byte(paymentGatewayDetailsStruct.Credentials.KeyId+":"+paymentGatewayDetailsStruct.Credentials.KeySecret)),
		"Content-Type":  "application/json",
	}
	blog.InfoCtx(ctx, "info-calling-create-payment-link-api", "payload", paymentGatewayPayload, "headers", headers)
	url := fmt.Sprintf("%s%s", paymentGatewayDetailsStruct.HttpConfig.Host, paymentGatewayDetailsStruct.HttpConfig.CreatePaymentLinkEndpoint)
	responseBody, err := r.baseClient.Do(http.MethodPost, paymentGatewayPayload, url, nil, headers)
	if err != nil {
		blog.ErrorCtx(ctx, err, "err-creating-payment-link-from-gateway", "response", responseBody)
		return err
	}
	var razorpayResponse view.RazorpayCreatePaymentLinkResponse
	err = json.Unmarshal([]byte(responseBody), &razorpayResponse)
	if err != nil {
		blog.ErrorCtx(ctx, err, "err-unmarshalling-response", "response", responseBody)
		return err
	}
	blog.InfoCtx(ctx, "info-created-payment-link")

	if notifyCrm {
		// Create Activity on CRM platform
		if crmProvider == "" || !slices.Contains(domain.GetCrmProviders(), string(crmProvider)) {
			return fmt.Errorf("err-crm-provider-not-identified")
		}
		switch crmProvider {
		case domain.Leadsquared:
			if req.ProspectId == "" {
				leadDetails, fetchErr := r.crmService.FetchLeadByPhoneNumberLeadsquared(ctx, req.CustomerContact)
				if fetchErr != nil {
					blog.ErrorCtx(ctx, fetchErr, "err-fetch-lead-by-phone-number", "response", responseBody)
					return fetchErr
				}
				blog.InfoCtx(ctx, "info-fetched-lead", "lead_details", *leadDetails)
				req.ProspectId = leadDetails.ProspectId
			}
			err := r.crmService.PostLeadActivityLeadsquared(ctx, contract.PostActivityLeadsquared{
				RelatedProspectId: req.ProspectId,
				ActivityEvent:     int64(domain.PostRazorpayPaymentLink),
				ActivityNote:      fmt.Sprintf("Payment link created for %s, for %d - %s, payment link is %s", req.CustomerContact, req.Amount/100, req.Currency, razorpayResponse.ShortUrl),
				ProcessFilesAsync: true,
				ActivityDateTime:  utils.GetCurrentUtcTimeInIso8086(),
				Fields: []struct {
					SchemaName string      `json:"SchemaName"`
					Value      interface{} `json:"Value"`
				}{
					{
						SchemaName: "mx_Custom_1",
						Value:      req.Amount / 100,
					},
				},
			})
			if err != nil {
				blog.ErrorCtx(ctx, err, "err-posting-payment-link-to-lead-activity", "response", responseBody)
			}
			blog.InfoCtx(ctx, "info-posted-activity-on-leadsquared")
		}
	}
	if notifyCustomer {
		// Send payment link to customer on whatsapp
		if whatsappProvider == "" || !slices.Contains(domain.GetWhatsappProviders(), string(whatsappProvider)) {
			return fmt.Errorf("err-whatsapp-provider-not-identified")
		}
		switch whatsappProvider {
		case domain.Angoor:
			err = r.communicationService.SendPaymentLinkToCustomerOnWhatsapp(ctx, contract.SendPaymentLinkToCustomer{
				CustomerPhoneNumber:   req.CustomerContact,
				PaymentLink:           razorpayResponse.ShortUrl,
				PaymentAmount:         float64(req.Amount / 100),
				LeadsquaredProspectId: req.ProspectId,
			}, domain.Angoor)
			if err != nil {
				blog.ErrorCtx(ctx, err, "err-sending-payment-link-to-customer-on-whatsapp")
				return err
			}
			blog.InfoCtx(ctx, "info-sent-payment-link-to-customer-on-whatsapp")
		}
	}
	return nil
}

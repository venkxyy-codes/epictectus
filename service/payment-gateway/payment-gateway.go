package payment_gateway

import (
	"bitbucket.org/apps-for-bharat/gotools/blog"
	"context"
	"encoding/base64"
	"epictectus/clients"
	"epictectus/config"
	"epictectus/contract"
	"epictectus/domain"
	"epictectus/service/crm"
	"epictectus/utils"
	"fmt"
	"net/http"
)

type paymentGatewayService struct {
	leadsquaredService    crm.LeadsquaredService
	baseClient            clients.BaseClient
	paymentGatewayDetails map[domain.PaymentProvider]interface{}
}

type PaymentGatewayService interface {
	CreateStandardPaymentLink(ctx context.Context, createStandardPaymentLinkRequest contract.CreateStandardPaymentLinkRequestRazorpay, paymentProvider domain.PaymentProvider) error
}

func populatePaymentGatewayConfig() map[domain.PaymentProvider]interface{} {
	return map[domain.PaymentProvider]interface{}{
		domain.Razorpay: map[string]interface{}{
			"credentials": config.GetRazorpayCredentials(),
			"http_config": config.GetRazorpayHttpConfig(),
		},
	}
}

func NewPaymentGatewayService(leadsquaredService crm.LeadsquaredService) PaymentGatewayService {
	return &paymentGatewayService{
		leadsquaredService:    leadsquaredService,
		baseClient:            clients.NewBaseClient(),
		paymentGatewayDetails: populatePaymentGatewayConfig(),
	}
}

func (r paymentGatewayService) getHttpConfigAndCredentialsForPaymentGateway(pgDetails interface{}) (map[string]interface{}, map[string]interface{}, error) {
	pgDetailsMap, isMap := pgDetails.(map[string]interface{})
	if !isMap {
		return nil, nil, fmt.Errorf("err-pgDetails-is-not-a-map")
	}
	httpConfig, ok := pgDetailsMap["http_config"]
	if !ok {
		return nil, nil, fmt.Errorf("err-httpConfig-is-not-present")
	}
	httpConfigMap, isMap := httpConfig.(map[string]interface{})
	if !isMap {
		return nil, nil, fmt.Errorf("err-httpConfig-is-not-a-map")
	}
	credentials, ok := pgDetailsMap["credentials"]
	if !ok {
		return nil, nil, fmt.Errorf("err-credentials-is-not-present")
	}
	credentialsMap, isMap := credentials.(map[string]interface{})
	if !isMap {
		return nil, nil, fmt.Errorf("err-credentials-is-not-a-map")
	}
	return httpConfigMap, credentialsMap, nil
}

func (r paymentGatewayService) CreateStandardPaymentLink(ctx context.Context, req contract.CreateStandardPaymentLinkRequestRazorpay, paymentProvider domain.PaymentProvider) error {
	paymentGatewayDetails, ok := r.paymentGatewayDetails[paymentProvider]
	if !ok {
		return fmt.Errorf("err-payment-gateway-credentials-not-identified")
	}
	req.AcceptPartial = false
	req.CallbackUrl = "https://ulc-api.leadsquaredapps.com/v1/UniversalLeadSync.svc/RealtimeWebhook/Connector/78089/3893d0d6-54f6-4b09-a4d9-ff540c2b03c6/0a1bcb55f3ff4c58be5b54dd4686e7aa0db8aff2c2c6c784d3e9f2b80ecbb552917ebc4455fb1831e92db6456ca69e56"
	req.CallbackMethod = "get"
	httpConfig, credentials, err := r.getHttpConfigAndCredentialsForPaymentGateway(paymentGatewayDetails)
	blog.InfoCtx(ctx, "info-credentials", "credentials", credentials)
	blog.InfoCtx(ctx, "info-http-config", "httpConfig", httpConfig)
	keyId, err := utils.GetStringValueFromMap(credentials, "key_id")
	if err != nil {
		return err
	}
	keySecret, err := utils.GetStringValueFromMap(credentials, "key_secret")
	if err != nil {
		return err
	}
	host, err := utils.GetStringValueFromMap(httpConfig, "host")
	if err != nil {
		blog.ErrorCtx(ctx, err, "err-geting-host-from-http-config")
	}
	endpoint, err := utils.GetStringValueFromMap(httpConfig, "create_payment_link_endpoint")
	if err != nil {
		return err
	}
	headers := map[string]string{
		"Authorization": "Basic " + base64.StdEncoding.EncodeToString([]byte(keyId+":"+keySecret)),
		"Content-Type":  "application/json",
	}
	blog.InfoCtx(ctx, "info-calling-create-payment-link-api", "payload", req, "headers", headers, "host", host)
	responseBody, err := r.baseClient.Do(http.MethodPost, req, fmt.Sprintf("%s%s", host, endpoint), nil, headers)
	if err != nil {
		blog.ErrorCtx(ctx, err, "err-creating-payment-link-from-gateway", "response", responseBody)
		return err
	}
	return nil
}

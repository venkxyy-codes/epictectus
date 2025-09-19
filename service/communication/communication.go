package communication

import (
	"context"
	"epictectus/blog"
	"epictectus/clients"
	"epictectus/config"
	"epictectus/contract"
	"epictectus/domain"
	"fmt"
	"net/http"
)

type commService struct {
	CommDetails map[domain.CommunicationChannels]interface{}
	baseClient  clients.BaseClient
}

type CommunicationService interface {
	SendPaymentLinkToCustomerOnWhatsapp(ctx context.Context, req contract.SendPaymentLinkToCustomer, provider domain.WhatsappProvider) error
}

func populateCommunicationProviderDetails() map[domain.CommunicationChannels]interface{} {
	return map[domain.CommunicationChannels]interface{}{
		domain.Whatsapp: map[domain.WhatsappProvider]interface{}{
			domain.Angoor: struct {
				AngoorHttpConfig  config.AngoorHttpConfig
				AngoorCredentials config.AngoorCredentials
			}{
				AngoorHttpConfig:  config.GetAngoorHttpConfig(),
				AngoorCredentials: config.GetAngoorCredentials(),
			},
		},
	}
}

func NewCommService(baseClient clients.BaseClient) CommunicationService {
	return &commService{
		baseClient:  baseClient,
		CommDetails: populateCommunicationProviderDetails(),
	}
}

func (c commService) SendPaymentLinkToCustomerOnWhatsapp(ctx context.Context, request contract.SendPaymentLinkToCustomer, provider domain.WhatsappProvider) error {
	whatsappDetails, ok := c.CommDetails[domain.Whatsapp]
	if !ok {
		return fmt.Errorf("err-communication-channel-unidentified")
	}
	providerDetails, ok := whatsappDetails.(map[domain.WhatsappProvider]interface{})
	if !ok {
		return fmt.Errorf("err-communication-channel-details-malformed")
	}
	providerConfig, ok := providerDetails[provider]
	if !ok {
		return fmt.Errorf("err-communication-provider-unidentified")
	}

	angoorDetails, ok := providerConfig.(struct {
		AngoorHttpConfig  config.AngoorHttpConfig
		AngoorCredentials config.AngoorCredentials
	})
	if !ok {
		return fmt.Errorf("err-communication-provider-config-and-credentials-malformed")
	}
	url := fmt.Sprintf("%s%s", angoorDetails.AngoorHttpConfig.Host, angoorDetails.AngoorHttpConfig.TriggerWebhook)
	blog.InfoCtx(ctx, "info-posting-activity-on-leadsquared", "url", url, "payload", request)
	_, err := c.baseClient.Do(http.MethodPost, request, url, nil, map[string]string{
		"Content-Type": "application/json",
		"x-api-key":    angoorDetails.AngoorCredentials.AccessKey,
	})
	if err != nil {
		return err
	}
	return nil
}

package crm

import (
	"context"
	"encoding/json"
	"epictectus/blog"
	"epictectus/clients"
	"epictectus/config"
	"epictectus/contract"
	"epictectus/domain"
	"epictectus/view"
	"fmt"
	"net/http"
)

type crmService struct {
	CrmDetails map[domain.CrmProvider]interface{}
	baseClient clients.BaseClient
}

type CrmService interface {
	FetchLeadByPhoneNumberLeadsquared(ctx context.Context, phoneNumber string) (*view.LeadDetailsLeadsquared, error)
	PostLeadActivityLeadsquared(ctx context.Context, request contract.PostActivityLeadsquared) error
}

func populateCrmDetails() map[domain.CrmProvider]interface{} {
	return map[domain.CrmProvider]interface{}{
		domain.Leadsquared: struct {
			LeadsquaredHttpConfig  config.LeadsquaredHttpConfig
			LeadsquaredCredentials config.LeadsquaredCredentials
		}{
			LeadsquaredHttpConfig:  config.GetLeadsquaredHttpConfig(),
			LeadsquaredCredentials: config.GetLeadsquaredCredentials(),
		},
	}
}

func NewCrmService(baseClient clients.BaseClient) CrmService {
	return &crmService{
		baseClient: baseClient,
		CrmDetails: populateCrmDetails(),
	}
}

func (c *crmService) FetchLeadByPhoneNumberLeadsquared(ctx context.Context, phoneNumber string) (*view.LeadDetailsLeadsquared, error) {
	lsqDetails, ok := c.CrmDetails[domain.Leadsquared]
	if !ok {
		return nil, fmt.Errorf("err-crm-provider-unidentified")
	}
	lsqDetailsStruct, ok := lsqDetails.(struct {
		LeadsquaredHttpConfig  config.LeadsquaredHttpConfig
		LeadsquaredCredentials config.LeadsquaredCredentials
	})
	if !ok {
		return nil, fmt.Errorf("err-crm-provider-config-and-credentials-malformed")
	}
	queryParams := map[string]string{
		"phone":     phoneNumber,
		"accessKey": lsqDetailsStruct.LeadsquaredCredentials.AccessKey,
		"secretKey": lsqDetailsStruct.LeadsquaredCredentials.SecretAccess,
	}
	url := fmt.Sprintf("%s%s", lsqDetailsStruct.LeadsquaredHttpConfig.Host, lsqDetailsStruct.LeadsquaredHttpConfig.FetchLeadUsingPhoneNumberEndpoint)
	responseBody, err := c.baseClient.Do(http.MethodGet, nil, url, queryParams, map[string]string{
		"Content-Type": "application/json",
	})
	if err != nil {
		return nil, err
	}
	var leadsDetails []view.LeadDetailsLeadsquared
	err = json.Unmarshal([]byte(responseBody), &leadsDetails)
	if err != nil {
		return nil, err
	}
	blog.InfoCtx(ctx, "info-response", "lead_details", leadsDetails[0])
	return &leadsDetails[0], nil
}

func (c *crmService) PostLeadActivityLeadsquared(ctx context.Context, request contract.PostActivityLeadsquared) error {
	lsqDetails, ok := c.CrmDetails[domain.Leadsquared]
	if !ok {
		return fmt.Errorf("err-crm-provider-unidentified")
	}
	lsqDetailsStruct, ok := lsqDetails.(struct {
		LeadsquaredHttpConfig  config.LeadsquaredHttpConfig
		LeadsquaredCredentials config.LeadsquaredCredentials
	})
	if !ok {
		return fmt.Errorf("err-crm-provider-config-and-credentials-malformed")
	}
	queryParams := map[string]string{
		"accessKey": lsqDetailsStruct.LeadsquaredCredentials.AccessKey,
		"secretKey": lsqDetailsStruct.LeadsquaredCredentials.SecretAccess,
	}
	url := fmt.Sprintf("%s%s", lsqDetailsStruct.LeadsquaredHttpConfig.Host, lsqDetailsStruct.LeadsquaredHttpConfig.PostActivityToLead)
	blog.InfoCtx(ctx, "info-posting-activity-on-leadsquared", "url", url, "payload", request)
	_, err := c.baseClient.Do(http.MethodPost, request, url, queryParams, map[string]string{
		"Content-Type": "application/json",
	})
	if err != nil {
		return err
	}
	return nil
}

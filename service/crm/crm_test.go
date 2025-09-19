package crm

import (
	"context"
	"epictectus/clients"
	"epictectus/contract"
	"epictectus/domain"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestCrmService_FetchLeadByPhoneNumberLeadsquared(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	service := NewCrmService(clients.NewBaseClient())

	t.Run("happy flow", func(t *testing.T) {
		lead, err := service.FetchLeadByPhoneNumberLeadsquared(ctx, "9876543210")
		assert.NoError(t, err)
		assert.NotNil(t, lead)
	})

	t.Run("provider not found", func(t *testing.T) {
		// Simulate missing provider
		crm := service.(*crmService)
		crm.CrmDetails = map[domain.CrmProvider]interface{}{}
		lead, err := crm.FetchLeadByPhoneNumberLeadsquared(ctx, "9876543210")
		assert.Error(t, err)
		assert.Nil(t, lead)
	})

	t.Run("invalid details struct", func(t *testing.T) {
		crm := service.(*crmService)
		crm.CrmDetails = map[domain.CrmProvider]interface{}{domain.CrmProvider(strconv.Itoa(1)): "invalid"}
		lead, err := crm.FetchLeadByPhoneNumberLeadsquared(ctx, "9876543210")
		assert.Error(t, err)
		assert.Nil(t, lead)
	})
}

func TestCrmService_PostLeadActivityLeadsquared(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	service := NewCrmService(clients.NewBaseClient())

	t.Run("happy flow", func(t *testing.T) {
		err := service.PostLeadActivityLeadsquared(ctx, contract.PostActivityLeadsquared{})
		assert.NoError(t, err)
	})

	t.Run("provider not found", func(t *testing.T) {
		crm := service.(*crmService)
		crm.CrmDetails = map[domain.CrmProvider]interface{}{}
		err := crm.PostLeadActivityLeadsquared(ctx, contract.PostActivityLeadsquared{})
		assert.Error(t, err)
	})

	t.Run("invalid details struct", func(t *testing.T) {
		crm := service.(*crmService)
		crm.CrmDetails = map[domain.CrmProvider]interface{}{domain.CrmProvider(strconv.Itoa(1)): "invalid"}
		err := crm.PostLeadActivityLeadsquared(ctx, contract.PostActivityLeadsquared{})
		assert.Error(t, err)
	})
}

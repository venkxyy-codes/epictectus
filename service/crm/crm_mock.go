package crm

import (
	"context"
	"epictectus/contract"
	"epictectus/view"
)

type MockCrmService struct {
	FetchLeadByPhoneNumberLeadsquaredCalled bool
	PostLeadActivityLeadsquaredCalled       bool
	FetchLeadByPhoneNumberLeadsquaredErr    error
	PostLeadActivityLeadsquaredErr          error
	LeadDetails                             *view.LeadDetailsLeadsquared
}

func (m *MockCrmService) FetchLeadByPhoneNumberLeadsquared(ctx context.Context, phoneNumber string) (*view.LeadDetailsLeadsquared, error) {
	m.FetchLeadByPhoneNumberLeadsquaredCalled = true
	return m.LeadDetails, m.FetchLeadByPhoneNumberLeadsquaredErr
}

func (m *MockCrmService) PostLeadActivityLeadsquared(ctx context.Context, request contract.PostActivityLeadsquared) error {
	m.PostLeadActivityLeadsquaredCalled = true
	return m.PostLeadActivityLeadsquaredErr
}

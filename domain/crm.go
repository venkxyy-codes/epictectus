package domain

type CrmProvider string

const (
	Leadsquared CrmProvider = "leadsquared"
)

func GetCrmProviders() []string {
	return []string{string(Leadsquared)}
}

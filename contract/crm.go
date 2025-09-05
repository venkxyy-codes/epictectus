package contract

import (
	"epictectus/domain"
	"fmt"
	"slices"
)

type LeadsquaredActivityWebhook struct {
	ProspectActivityId string `json:"ProspectActivityId"`
	RelatedProspectId  string `json:"RelatedProspectId"`
	ActivityEvent      string `json:"ActivityEvent"`
	ActivityEventName  string `json:"ActivityEventName"`
	ActivityType       string `json:"ActivityType"`
	Score              int    `json:"Score"`
	CreatedBy          string `json:"CreatedBy"`
	CreatedOn          string `json:"CreatedOn"`
	Data               struct {
		Owner     string `json:"Owner"`
		MxCustom1 string `json:"mx_Custom_1"`
		MxCustom2 string `json:"mx_Custom_2"`
		MxCustom3 string `json:"mx_Custom_3"`
	} `json:"Data"`
	Attachments interface{} `json:"Attachments"`
	Current     struct {
		ProspectID        string      `json:"ProspectID"`
		ProspectAutoId    string      `json:"ProspectAutoId"`
		FirstName         string      `json:"FirstName"`
		EmailAddress      string      `json:"EmailAddress"`
		Phone             string      `json:"Phone"`
		Mobile            interface{} `json:"Mobile"`
		Source            string      `json:"Source"`
		Score             string      `json:"Score"`
		RelatedProspectId interface{} `json:"RelatedProspectId"`
	} `json:"Current"`
	RelatedOpportunityEvent interface{} `json:"RelatedOpportunityEvent"`
	RelatedOpportunityId    interface{} `json:"RelatedOpportunityId"`
	AdditionalData          struct {
		ProcessFormId     string `json:"ProcessFormId"`
		ProcessDesignerId string `json:"ProcessDesignerId"`
	} `json:"AdditionalData"`
}

type PostActivityLeadsquared struct {
	RelatedProspectId string `json:"RelatedProspectId"`
	ActivityEvent     int64  `json:"ActivityEvent"`
	ActivityNote      string `json:"ActivityNote"`
	ProcessFilesAsync bool   `json:"ProcessFilesAsync"`
	ActivityDateTime  string `json:"ActivityDateTime"`
	Fields            []struct {
		SchemaName string      `json:"SchemaName"`
		Value      interface{} `json:"Value"`
	} `json:"Fields"`
}

func (w *LeadsquaredActivityWebhook) Validate() error {
	// Validate ActivityEvent
	validEvents := domain.GetValidActivityEventCodesLeadsquared()
	if !slices.Contains(validEvents, w.ActivityEvent) {
		return fmt.Errorf("invalid ActivityEvent: %s", w.ActivityEvent)
	}
	if w.ProspectActivityId == "" {
		return fmt.Errorf("ProspectActivityId cannot be empty")
	}
	if w.RelatedProspectId == "" {
		return fmt.Errorf("RelatedProspectId cannot be empty")
	}
	return nil
}

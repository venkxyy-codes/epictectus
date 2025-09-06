package contract

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

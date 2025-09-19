package clients

// MockBaseClient is a mock implementation of the BaseClient interface.
type MockBaseClient struct {
	DoFunc func(httpMethod string, requestBody interface{}, baseURL string, queryParams map[string]string, headers map[string]string) (string, error)
}

func (m *MockBaseClient) Do(httpMethod string, requestBody interface{}, baseURL string, queryParams map[string]string, headers map[string]string) (string, error) {
	if m.DoFunc != nil {
		return m.DoFunc(httpMethod, requestBody, baseURL, queryParams, headers)
	}
	return "", nil
}

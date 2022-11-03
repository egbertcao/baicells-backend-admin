package model

type Context struct {
	ApiVersion string `json:"api_version"`
	Realm      string `json:"realm"`
	Nonce      string `json:"nonce"`
	Accept     string `json:"accept"`
	Sn         string `json:"sn"`
	DeviceId   string `json:"device_id"`
	DomainId   string `json:"domain_id"`
	Model      string `json:"model"`
	Time       string `json:"time"`
}

type Login struct {
	ID      string  `json:"id"`
	Command string  `json:"command"`
	Context Context `json:"context"`
	Data    string  `json:"data"`
}

type Msg struct {
	ResponseTopic   string `json:"response_topic"`
	CorrelationData string `json:"correlation_data"`
	Data            []byte `json:"data"`
}

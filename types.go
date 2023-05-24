package HotRemote

type Request struct {
	Action    string `json:"Action"`
	Press     any    `json:"Press"`
	TargetBox string `json:"TargetBox"`
	SourceBox string `json:"SourceBox"`
	Event     string `json:"Event"`
	Params    string `json:"Params"`
	Token     string `json:"Token"`
}

type ListDevicesResponse struct {
	RemoteResponseCode string          `json:"RemoteResponseCode"`
	Message            string          `json:"Message"`
	Data               [][]interface{} `json:"Data"`
	Params             struct {
		SourceBox string `json:"SourceBox"`
	} `json:"Params"`
	MethodName string `json:"MethodName"`
}

type Device struct {
	ID   string
	Name string
}

type KeepAliveResponse struct {
	Action        string        `json:"Action"`
	State         string        `json:"State"`
	MethodName    string        `json:"MethodName"`
	GUID          int64         `json:"GUID"`
	Token         string        `json:"Token"`
	SourceBox     string        `json:"SourceBox"`
	TargetBox     string        `json:"TargetBox"`
	Params        []interface{} `json:"Params"`
	CallerStbName string        `json:"CallerStbName"`
}

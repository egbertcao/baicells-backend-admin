package model

type Instance struct {
	ProductId string      `json:"productId" form:"productId"`
	DeviceId  string      `json:"deviceId" form:"deviceId"`
	Command   string      `json:"command" form:"command"`
	InputData interface{} `json:"inputData" form:"inputData"`
}

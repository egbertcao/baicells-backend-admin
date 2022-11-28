package model

type Action struct {
	ID     string `json:"id"`
	Method string `json:"method"`
}

type ActionData struct {
	Action  Action `json:"action"`
	Payload string `json:"payload"`
}

type DeviceAction struct {
	ID      string     `json:"id"`
	Command string     `json:"command"`
	Data    ActionData `json:"data"`
}

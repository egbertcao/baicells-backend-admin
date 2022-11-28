package service

import (
	"encoding/json"
	"fmt"
	"time"

	backendmodel "github.com/flipped-aurora/gin-vue-admin/server/plugin/deviceManagement/backend/model"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/deviceManagement/model"
)

type InstanceService struct{}

func (instance *InstanceService) CallMethod(e model.Instance) (err error) {
	action := backendmodel.Action{
		ID:     "action-" + time.Now().String(),
		Method: e.Command,
	}

	payload, _ := json.Marshal(e.InputData)

	actionData := backendmodel.ActionData{
		Action:  action,
		Payload: string(payload[:]),
	}

	deviceAction := backendmodel.DeviceAction{
		ID:      time.Now().String(),
		Command: "device_action_req",
		Data:    actionData,
	}

	fmt.Printf("%v\n", deviceAction)

	return
}

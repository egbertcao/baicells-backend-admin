package response

import "github.com/flipped-aurora/gin-vue-admin/server/plugin/deviceManagement/model"

type DeviceResponse struct {
	Device model.Device `json:"device"`
}

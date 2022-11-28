package service

import (
	"encoding/json"

	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/deviceManagement/config"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/deviceManagement/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/deviceManagement/model"
	deviceReq "github.com/flipped-aurora/gin-vue-admin/server/plugin/deviceManagement/model/request"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DeviceService struct{}

var authentication AuthenticationSerivice

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateDevice
//@description: 创建客户
//@param: e model.Device
//@return: err error
func (exa *DeviceService) Test(c *gin.Context) {
	println("123test")
	response.OkWithMessage("创建成功", c)
}

func (exa *DeviceService) CreateDevice(e model.Device) (err error) {
	objectid := primitive.NewObjectID()
	e.ID = objectid.Hex()
	e.CreatedAt = objectid.Timestamp()
	e.UpdatedAt = objectid.Timestamp()

	mongo := config.MongoParam{
		DataBase:   "mesh",
		Collection: "Device",
	}
	err = global.MongoSession.MongoInsert(mongo, "serialnumber", e)
	if err != nil {
		return
	}
	param := model.MqttAuthentication{
		UserName: e.SerialNumber,
		Password: e.SerialNumber,
		Super:    false,
	}
	authentication.CreateAuthentication(param)

	instanceRouter := global.Group.Group("TEST").Use(middleware.OperationRecord())
	instanceRouter.POST("TESTrouter", exa.Test)

	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteFileChunk
//@description: 删除客户
//@param: e model.Device
//@return: err error

func (exa *DeviceService) DeleteDevice(e model.Device) (err error) {
	return err
}

func (deviceService *DeviceService) DeleteDeviceByIds(ids request.IdsReq) (err error) {
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateDevice
//@description: 更新客户
//@param: e *model.Device
//@return: err error

func (exa *DeviceService) UpdateDevice(e model.Device) (err error) {
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetDevice
//@description: 获取客户信息
//@param: id uint
//@return: customer model.Device, err error

func (exa *DeviceService) GetDevice(id string) (customer model.Device, err error) {
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetCustomerInfoList
//@description: 分页获取客户列表
//@param: sysUserAuthorityID string, info request.PageInfo
//@return: list interface{}, total int64, err error

func (exa *DeviceService) GetDeviceInfoList(info deviceReq.DeviceSearch) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	filter := bson.D{}
	mongo := config.MongoParam{
		DataBase:   "mesh",
		Collection: "Device",
	}
	data, total, err := global.MongoSession.MongoGet(mongo, limit, offset, filter)
	if err != nil {
		return
	}
	var response []model.Device
	data_byte, err := json.Marshal(data)
	json.Unmarshal(data_byte, &response)

	return response, total, err
}

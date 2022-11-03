package service

import (
	"encoding/json"

	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/deviceManagement/config"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/deviceManagement/global"
	"github.com/flipped-aurora/gin-vue-admin/server/plugin/deviceManagement/model"
	deviceReq "github.com/flipped-aurora/gin-vue-admin/server/plugin/deviceManagement/model/request"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AuthorizationSerivice struct{}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateAuthorization
//@description: 创建客户
//@param: e model.Authorization
//@return: err error

func (exa *AuthorizationSerivice) CreateAuthorization(e model.MqttAuthorization) (err error) {
	objectid := primitive.NewObjectID()
	e.ID = objectid.Hex()
	e.CreatedAt = objectid.Timestamp()
	e.UpdatedAt = objectid.Timestamp()

	mongo := config.MongoParam{
		DataBase:   "mqtt",
		Collection: "mqtt_acl",
	}
	err = global.MongoSession.MongoInsert(mongo, "username", e)
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteFileChunk
//@description: 删除客户
//@param: e model.Authorization
//@return: err error

func (exa *AuthorizationSerivice) DeleteAuthorization(e model.MqttAuthorization) (err error) {
	return err
}

func (AuthorizationSerivice *AuthorizationSerivice) DeleteAuthorizationByIds(ids request.IdsReq) (err error) {
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateAuthorization
//@description: 更新客户
//@param: e *model.Authorization
//@return: err error

func (exa *AuthorizationSerivice) UpdateAuthorization(e model.MqttAuthorization) (err error) {
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetAuthorization
//@description: 获取客户信息
//@param: id uint
//@return: customer model.Authorization, err error

func (exa *AuthorizationSerivice) GetAuthorization(id string) (customer model.MqttAuthorization, err error) {
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetCustomerInfoList
//@description: 分页获取客户列表
//@param: sysUserAuthorityID string, info request.PageInfo
//@return: list interface{}, total int64, err error

func (exa *AuthorizationSerivice) GetAuthorizationInfoList(info deviceReq.MqttAuthorizationSearch) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	filter := bson.D{}
	mongo := config.MongoParam{
		DataBase:   "mqtt",
		Collection: "mqtt_acl",
	}
	data, total, err := global.MongoSession.MongoGet(mongo, limit, offset, filter)
	if err != nil {
		return
	}
	var response []model.MqttAuthorization
	data_byte, err := json.Marshal(data)
	json.Unmarshal(data_byte, &response)

	return response, total, err
}

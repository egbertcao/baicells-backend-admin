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

type AuthenticationSerivice struct{}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateAuthentication
//@description: 创建客户
//@param: e model.Authentication
//@return: err error

func (exa *AuthenticationSerivice) CreateAuthentication(e model.MqttAuthentication) (err error) {
	objectid := primitive.NewObjectID()
	e.ID = objectid.Hex()
	e.CreatedAt = objectid.Timestamp()
	e.UpdatedAt = objectid.Timestamp()
	mongo := config.MongoParam{
		DataBase:   "mqtt",
		Collection: "users",
	}
	err = global.MongoSession.MongoInsert(mongo, "username", e)
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteFileChunk
//@description: 删除客户
//@param: e model.Authentication
//@return: err error

func (exa *AuthenticationSerivice) DeleteAuthentication(e model.MqttAuthentication) (err error) {
	return err
}

func (AuthenticationSerivice *AuthenticationSerivice) DeleteAuthenticationByIds(ids request.IdsReq) (err error) {
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateAuthentication
//@description: 更新客户
//@param: e *model.Authentication
//@return: err error

func (exa *AuthenticationSerivice) UpdateAuthentication(e model.MqttAuthentication) (err error) {
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetAuthentication
//@description: 获取客户信息
//@param: id uint
//@return: customer model.Authentication, err error

func (exa *AuthenticationSerivice) GetAuthentication(id string) (customer model.MqttAuthentication, err error) {
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetCustomerInfoList
//@description: 分页获取客户列表
//@param: sysUserAuthorityID string, info request.PageInfo
//@return: list interface{}, total int64, err error

func (exa *AuthenticationSerivice) GetAuthenticationInfoList(info deviceReq.MqttAuthenticationSearch) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	filter := bson.D{}
	mongo := config.MongoParam{
		DataBase:   "mqtt",
		Collection: "users",
	}
	data, total, err := global.MongoSession.MongoGet(mongo, limit, offset, filter)
	if err != nil {
		return
	}
	var response []model.MqttAuthentication
	data_byte, err := json.Marshal(data)
	json.Unmarshal(data_byte, &response)
	return response, total, err
}

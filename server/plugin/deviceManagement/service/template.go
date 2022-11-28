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

type TemplateService struct{}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: CreateTemplate
//@description: 创建客户
//@param: e model.Template
//@return: err error

func (exa *TemplateService) CreateTemplate(e model.Template) (err error) {
	objectid := primitive.NewObjectID()
	e.ID = objectid.Hex()
	e.CreatedAt = objectid.Timestamp()
	e.UpdatedAt = objectid.Timestamp()

	mongo := config.MongoParam{
		DataBase:   "mesh",
		Collection: "Template",
	}
	err = global.MongoSession.MongoInsert(mongo, "model", e)
	if err != nil {
		return
	}
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteFileChunk
//@description: 删除客户
//@param: e model.Template
//@return: err error

func (exa *TemplateService) DeleteTemplate(e model.Template) (err error) {
	return err
}

func (deviceService *TemplateService) DeleteTemplateByIds(ids request.IdsReq) (err error) {
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateTemplate
//@description: 更新客户
//@param: e *model.Template
//@return: err error

func (exa *TemplateService) UpdateTemplate(e model.Template) (err error) {
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetTemplate
//@description: 获取客户信息
//@param: id uint
//@return: customer model.Template, err error

func (exa *TemplateService) GetTemplate(id string) (customer model.Template, err error) {
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetCustomerInfoList
//@description: 分页获取客户列表
//@param: sysUserAuthorityID string, info request.PageInfo
//@return: list interface{}, total int64, err error

func (exa *TemplateService) GetTemplateInfoList(info deviceReq.TemplateSearch) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	filter := bson.D{}
	mongo := config.MongoParam{
		DataBase:   "mesh",
		Collection: "Template",
	}
	data, total, err := global.MongoSession.MongoGet(mongo, limit, offset, filter)
	if err != nil {
		return
	}
	var response []model.Template
	data_byte, err := json.Marshal(data)
	json.Unmarshal(data_byte, &response)

	return response, total, err
}

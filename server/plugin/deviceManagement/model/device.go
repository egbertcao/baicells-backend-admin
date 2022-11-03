package model

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
)

const (
	// 拥有读取查看权限
	READ = "read"

	// 拥有操作权限
	Write = "Write"
)

type Authorization struct {
	Type string `json:"type" form:"type" gorm:"comment:权限类型"` // 权限类型
}

type ShareOwn struct {
	OwnerID        uint            `json:"ownerid" form:"ownerid" gorm:"comment:拥有人ID"` // 拥有人ID
	SysUser        system.SysUser  `json:"sysuser" form:"sysuser" gorm:"comment:拥有人详情"` // 拥有人详情
	Authorizations []Authorization `json:"authorizations" form:"authorizations" gorm:"comment:权限"`
}

type Device struct {
	ID           string         `json:"id" bson:"_id" form:"id" gorm:"column:id;comment:ID;"`
	CreatedAt    time.Time      // 创建时间
	UpdatedAt    time.Time      // 更新时间
	SerialNumber string         `json:"serialnumber" form:"serialnumber" gorm:"comment:序列号"` // 序列号
	MacAddress   string         `json:"macaddress" form:"macaddress" gorm:"comment:MAC地址"`   // MAC地址
	IpAddress    string         `json:"ipaddress" form:"ipaddress" gorm:"comment:IP地址"`      // IP地址
	Status       string         `json:"status" form:"status" gorm:"comment:设备状态"`            // 设备状态
	OwnerID      uint           `json:"ownerid" form:"ownerid" gorm:"comment:拥有人ID"`         // 拥有人ID
	SysUser      system.SysUser `json:"sysuser" form:"sysuser" gorm:"comment:拥有人详情"`         // 拥有人详情
	ShareOwns    []ShareOwn     `json:"shareowns" form:"shareowns" gorm:"comment:共享人详情"`
}

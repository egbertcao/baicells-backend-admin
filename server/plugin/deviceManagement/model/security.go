package model

import "time"

// 认证
type MqttAuthentication struct {
	ID        string    `json:"id" bson:"_id" form:"id" gorm:"column:id;comment:ID;"`
	CreatedAt time.Time // 创建时间
	UpdatedAt time.Time // 更新时间
	UserName  string    `json:"username" form:"username" gorm:"comment:用户名"`
	Salt      string    `json:"salt" form:"salt" gorm:"comment:加盐"`
	Super     bool      `json:"super" form:"super" gorm:"comment:是否为超级用户"`
	Password  string    `json:"password" form:"password" gorm:"comment:密码"`
}

// 授权
type MqttAuthorization struct {
	ID         string    `json:"id" bson:"_id" form:"id" gorm:"column:id;comment:ID;"`
	CreatedAt  time.Time // 创建时间
	UpdatedAt  time.Time // 更新时间
	UserName   string    `json:"username" form:"username" gorm:"comment:用户名"`
	Permission string    `json:"permission" form:"permission" gorm:"comment:授权结果deny/allow"`
	Action     string    `json:"action" form:"action" gorm:"comment:动作publish/subscribe/all"`
	Topics     []string  `json:"topics" form:"topics" gorm:"comment:主题"`
}

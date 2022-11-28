package model

import "time"

type Spec struct {
	Min  string `json:"min"`
	Max  string `json:"max"`
	Step string `json:"step"`
}

type DataType struct {
	Type string `json:"type"`
	Spec Spec   `json:"spec"`
}

type InputData struct {
	Name     string   `json:"name"`
	DataType DataType `json:"dataType"`
}

type OutputData struct {
	Name     string   `json:"name"`
	DataType DataType `json:"dataType"`
}

type Attribute struct {
	Identifier string   `json:"identifier"`
	Desc       string   `json:"desc"`
	Name       string   `json:"name"`
	AccessMode string   `json:"accessMode"`
	DataType   DataType `json:"dataType"`
}

type Event struct {
	Identifier string       `json:"identifier"`
	Desc       string       `json:"desc"`
	Name       string       `json:"name"`
	OutputData []OutputData `json:"outputData"`
}

type Service struct {
	Identifier string       `json:"identifier"`
	Desc       string       `json:"desc"`
	Name       string       `json:"name"`
	CallType   string       `json:"calltype"`
	InputData  []InputData  `json:"inputData"`
	OutputData []OutputData `json:"outputData"`
}

type Template struct {
	ID         string      `json:"_id" bson:"_id" form:"id" gorm:"column:id;comment:ID;"`
	CreatedAt  time.Time   // 创建时间
	UpdatedAt  time.Time   // 更新时间
	Name       string      `json:"name"`
	Model      string      `json:"model"`
	Desc       string      `json:"desc"`
	Attributes []Attribute `json:"attributes"`
	Services   []Service   `json:"services"`
	Events     []Event     `json:"events"`
}

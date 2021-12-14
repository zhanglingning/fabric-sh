package models

import "time"

type ChaincodeInfo struct {
	Uuid    string    `json:"uuid"`
	TxId    string    `json:"txId"`
	Time    time.Time `json:"time"`
	Payload string    `json:"payload"`
}

type Education struct {
	ObjectType string `json:"docType" validate:"required,gt=0"`
	Name       string `json:"Name" validate:"required,gt=0"`     // 姓名
	Gender     string `json:"Gender" validate:"required,gt=0"`   // 性别
	Nation     string `json:"Nation" validate:"required,gt=0"`   // 民族
	EntityID   string `json:"EntityID" validate:"required,gt=0"` // 身份证号
	Place      string `json:"Place" validate:"required,gt=0"`    // 籍贯
	BirthDay   string `json:"BirthDay" validate:"required,gt=0"` // 出生日期

	EnrollDate     string `json:"EnrollDate" validate:"required,gt=0"`     // 入学日期
	GraduationDate string `json:"GraduationDate" validate:"required,gt=0"` // 毕（结）业日期
	SchoolName     string `json:"SchoolName" validate:"required,gt=0"`     // 学校名称
	Major          string `json:"Major" validate:"required,gt=0"`          // 专业
	QuaType        string `json:"QuaType" validate:"required,gt=0"`        // 学历类别
	Length         string `json:"Length" validate:"required,gt=0"`         // 学制
	Mode           string `json:"Mode" validate:"required,gt=0"`           // 学习形式
	Level          string `json:"Level" validate:"required,gt=0"`          // 层次
	Graduation     string `json:"Graduation" validate:"required,gt=0"`     // 毕（结）业
	CertNo         string `json:"CertNo" validate:"required,gt=0"`         // 证书编号

	Photo string `json:"Photo"` // 照片

	Historys []HistoryItem // 当前edu的历史记录
}

type HistoryItem struct {
	TxId      string
	Education Education
}

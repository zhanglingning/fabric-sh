package services

import (
	"fmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/fab"
	log "github.com/sirupsen/logrus"
	"time"
)

type Education struct {
	ObjectType     string `json:"docType" validate:"required,gt=0"`
	Name           string `json:"Name" validate:"required,gt=0"`           // 姓名
	Gender         string `json:"Gender" validate:"required,gt=0"`         // 性别
	Nation         string `json:"Nation" validate:"required,gt=0"`         // 民族
	EntityID       string `json:"EntityID" validate:"required,gt=0"`       // 身份证号
	Place          string `json:"Place" validate:"required,gt=0"`          // 籍贯
	BirthDay       string `json:"BirthDay" validate:"required,gt=0"`       // 出生日期
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

type ServiceSetup struct {
	ChaincodeID string
	Client      *channel.Client
}

func regitserEvent(client *channel.Client, chaincodeID, eventID string) (fab.Registration, <-chan *fab.CCEvent) {

	reg, notifier, err := client.RegisterChaincodeEvent(chaincodeID, eventID)
	if err != nil {
		log.Errorf("注册链码事件失败: %s\n", err)
		return reg, notifier
	}
	return reg, notifier
}

func eventResult(notifier <-chan *fab.CCEvent, eventID string) error {
	select {
	case ccEvent := <-notifier:
		log.Printf("接收到链码事件: %v\n", ccEvent)
	case <-time.After(time.Second * 20):
		return fmt.Errorf("不能根据指定的事件ID接收到相应的链码事件(%s)", eventID)
	}
	return nil
}

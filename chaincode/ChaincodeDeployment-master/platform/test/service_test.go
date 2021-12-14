package test

import (
	"encoding/json"
	"fmt"
	"github.com/FuradWho/ChaincodeDeployment/platform/models"
	"github.com/FuradWho/ChaincodeDeployment/platform/web/services"
	"testing"
)

func TestService(t *testing.T) {
	fabricClient := new(models.FabricClient)

	err := fabricClient.Init()
	if err != nil {
		t.Errorf("%s", err)
		return
	}

	serviceSetup := services.ServiceSetup{
		ChaincodeID: "chsi_0",
		Client:      fabricClient.ChannelClient,
	}

	//// 根据身份证号码查询信息
	result, err := serviceSetup.FindEduInfoByEntityID("101")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		var edu services.Education
		json.Unmarshal(result, &edu)
		fmt.Println("根据身份证号码查询信息成功：")
		fmt.Printf("%+v \n", edu)
	}

	//edu := services.Education{
	//	Name:           "张三",
	//	Gender:         "男",
	//	Nation:         "汉",
	//	EntityID:       "101",
	//	Place:          "北京",
	//	BirthDay:       "1991年01月01日",
	//	EnrollDate:     "2009年9月",
	//	GraduationDate: "2013年7月",
	//	SchoolName:     "中国政法大学",
	//	Major:          "社会学",
	//	QuaType:        "普通",
	//	Length:         "四年",
	//	Mode:           "普通全日制",
	//	Level:          "本科",
	//	Graduation:     "毕业",
	//	CertNo:         "111",
	//	Photo:          "/static/phone/11.png",
	//}
	//
	//msg, err := serviceSetup.SaveEdu(edu)
	//if err != nil {
	//	fmt.Println(err.Error())
	//} else {
	//	fmt.Println("信息发布成功, 交易编号为: " + msg)
	//}

	// 修改/添加信息
	//info := services.Education{
	//	Name:           "张三",
	//	Gender:         "男",
	//	Nation:         "汉",
	//	EntityID:       "101",
	//	Place:          "北京",
	//	BirthDay:       "1991年01月01日",
	//	EnrollDate:     "2013年9月",
	//	GraduationDate: "2015年7月",
	//	SchoolName:     "中国政法大学",
	//	Major:          "社会学",
	//	QuaType:        "普通",
	//	Length:         "两年",
	//	Mode:           "普通全日制",
	//	Level:          "研究生",
	//	Graduation:     "毕业",
	//	CertNo:         "333",
	//	Photo:          "/static/phone/11.png",
	//}
	//msg, err := serviceSetup.ModifyEdu(info)
	//if err != nil {
	//	fmt.Println(err.Error())
	//} else {
	//	fmt.Println("信息操作成功, 交易编号为: " + msg)
	//}

}

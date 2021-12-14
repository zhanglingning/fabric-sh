package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/FuradWho/ChaincodeDeployment/platform/models"
	"github.com/FuradWho/ChaincodeDeployment/platform/web/services"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-playground/validator/v10"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	log "github.com/sirupsen/logrus"
)

type EduController struct {
	Ctx     iris.Context
	Service services.ServiceSetup
}

var validate = validator.New()

func (c *EduController) GetTest() models.ResponseBean {
	path := c.Ctx.Path()
	log.Infoln(path)
	return models.SuccessMsg("1111")
}

func (c *EduController) SaveEdu() models.ResponseBean {

	path := c.Ctx.Path()
	log.Infoln(path)

	var eduInfo services.Education

	err := c.Ctx.ReadJSON(&eduInfo)
	if err != nil {
		log.Errorf("Failed to json read to struct error : %s", err)
		return models.FailedMsg("Failed to json read to struct error")
	}
	err = validate.Struct(eduInfo)
	if err != nil {
		log.Errorf("Failed to struct format error : %s", err)
		return models.FailedMsg("Failed to struct format error")

	}
	txId, err := c.Service.SaveEdu(eduInfo)
	if err != nil {
		log.Errorf("Failed to service save edu info : %s", err)
		return models.FailedMsg("Failed to service save edu info")
	}
	return models.SuccessData(map[string]string{
		"txId": txId,
	})

}

func (c *EduController) FindEduInfoByEntityID(entityID string) models.ResponseBean {

	path := c.Ctx.Path()
	log.Infoln(path)

	if entityID == "" {
		log.Errorf("Failed to Get Info because entity id is empty")
		return models.FailedMsg("Failed to Get Info because entity id is empty")
	}

	fmt.Println(entityID)
	result, err := c.Service.FindEduInfoByEntityID(entityID)
	if err != nil {
		log.Errorf("Failed to FindEduInfoByEntityID : %s", err)
		return models.FailedMsg("Failed to FindEduInfoByEntityID")
	}
	var eduInfo services.Education
	if err != nil {
		fmt.Println(err.Error())
	} else {
		json.Unmarshal(result, &eduInfo)
	}
	return models.SuccessData(map[string]services.Education{
		"eduInfo": eduInfo,
	})
}

func (c *EduController) BeforeActivation(b mvc.BeforeActivation) {

	b.Handle(
		"POST",
		"/save_edu",
		"SaveEdu",
	)

	b.Handle(
		"GET",
		"/find_by_entityID/{entityID:string}",
		"FindEduInfoByEntityID",
	)

}

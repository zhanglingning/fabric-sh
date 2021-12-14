package controllers

import (
	"github.com/FuradWho/ChaincodeDeployment/platform/models"
	"github.com/FuradWho/ChaincodeDeployment/platform/web/services"
	"github.com/iris-contrib/swagger/v12"
	"github.com/iris-contrib/swagger/v12/swaggerFiles"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"

	log "github.com/sirupsen/logrus"

	_ "github.com/FuradWho/ChaincodeDeployment/platform/docs"
)

// @title ChaincodeDeployment Platform
// @version 1.0
// @description go sdk for ChaincodeDeployment

// @contact.name FuradWho
// @contact.email liu1337543811@gmail.com

// @license.name Fabric 2.3.3
// @license.url https://hyperledger-fabric.readthedocs.io/zh_CN/release-2.2/who_we_are.html

// StartIris
// @host localhost:9089
// @BasePath /
func StartIris() {
	app := iris.New()
	app.Use(Cors)

	config := &swagger.Config{
		URL:         "http://33p67e8007.qicp.vip/swagger/doc.json",
		DeepLinking: true,
	}

	app.Get("/swagger/{any:path}", swagger.CustomWrapHandler(config, swaggerFiles.Handler))
	mvc.Configure(app.Party("/edu"), edu)
	app.Listen(":9099")

}

// Cors Resolve the CORS
func Cors(ctx iris.Context) {

	ctx.Header("Access-Control-Allow-Origin", "*")
	if ctx.Request().Method == "OPTIONS" {
		ctx.Header("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,PATCH,OPTIONS")
		ctx.Header("Access-Control-Allow-Headers", "Content-Type, Accept, Authorization")
		ctx.StatusCode(204)
		return
	}
	ctx.Next()
}

func edu(app *mvc.Application) {
	fabricClient := new(models.FabricClient)

	err := fabricClient.Init()
	if err != nil {
		log.Errorln(err)
		return
	}

	serviceSetup := services.ServiceSetup{
		ChaincodeID: "chsi_0",
		Client:      fabricClient.ChannelClient,
	}

	app.Register(serviceSetup)
	app.Handle(new(EduController))

}

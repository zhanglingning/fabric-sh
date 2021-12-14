package main

import (
	_ "github.com/FuradWho/ChaincodeDeployment/platform/third_party/logger"
	"github.com/FuradWho/ChaincodeDeployment/platform/web/controllers"
)

func main() {

	controllers.StartIris()

}

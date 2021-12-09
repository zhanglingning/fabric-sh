package main

import (
	"fmt"

	"github.com/FuradWho/GoChaincode/sacc/chaincode"
	"github.com/hyperledger/fabric-chaincode-go/shim"
)

func main() {
	err := shim.Start(new(chaincode.SimpleAsset))
	if err != nil {
		fmt.Printf("启动SimpleAsset时发生错误: %s", err)
	}
}

package test

import (
	"fmt"
	"github.com/FuradWho/GoChaincode/sacc/chaincode"
	"github.com/hyperledger/fabric-chaincode-go/shimtest"
	"testing"
)

func TestInvoke(t *testing.T) {
	cc := new(chaincode.SimpleAsset)
	stub := shimtest.NewMockStub("sacc", cc) // 创建MockStub对象
	// 调用Init接口，将a的值设为90
	stub.MockInit("1", [][]byte{[]byte("a"), []byte("90")})
	// 调用get接口查询a的值
	res := stub.MockInvoke("1", [][]byte{[]byte("get"), []byte("a")})
	fmt.Println("The value of a is", string(res.Payload))
	// 调用set接口设置a为100
	stub.MockInvoke("1", [][]byte{[]byte("set"), []byte("a"), []byte("100")})
	// 再次查询a的值
	res = stub.MockInvoke("1", [][]byte{[]byte("get"), []byte("a")})
	fmt.Println("The new value of a is", string(res.Payload))

	// TODO
}

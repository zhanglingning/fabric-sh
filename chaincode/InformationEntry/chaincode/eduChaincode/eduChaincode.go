package eduChaincode

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/FuradWho/GoChaincode/chsi/model/education"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-protos-go/peer"
)

const (
	DOC_TYPE = "eduObj"
)

type EducationChaincode struct {
}

func (e *EducationChaincode) Init(stub shim.ChaincodeStubInterface) peer.Response {
	return shim.Success(nil)
}
func (e *EducationChaincode) Invoke(stub shim.ChaincodeStubInterface) peer.Response {

	fun, args := stub.GetFunctionAndParameters()

	if fun == "addEdu" {
		return e.addEdu(stub, args) // 添加信息
	} else if fun == "queryEduByCertNoAndName" {
		return e.queryEduByCertNoAndName(stub, args) // 根据证书编号及姓名查询信息
	} else if fun == "queryEduInfoByEntityID" {
		return e.queryEduInfoByEntityID(stub, args) // 根据身份证号码及姓名查询详情
	} else if fun == "updateEdu" {
		return e.updateEdu(stub, args) // 根据证书编号更新信息
	} else if fun == "delEdu" {
		return e.delEdu(stub, args) // 根据证书编号删除信息
	}

	return shim.Error("指定的函数名称错误")

}

func PutEdu(stub shim.ChaincodeStubInterface, edu education.Education) ([]byte, bool) {

	edu.ObjectType = DOC_TYPE

	marshal, err := json.Marshal(edu)
	if err != nil {
		return nil, false
	}

	err = stub.PutState(edu.EntityID, marshal)
	if err != nil {
		return nil, false
	}

	return marshal, true

}

func GetEduInfo(stub shim.ChaincodeStubInterface, entityID string) (education.Education, bool) {

	var edu education.Education

	b, err := stub.GetState(entityID)
	if err != nil {
		return edu, false
	}

	if b == nil {
		return edu, false
	}

	err = json.Unmarshal(b, &edu)
	if err != nil {
		return edu, false
	}

	return edu, true
}

// 根据指定的查询字符串实现富查询
func getEduByQueryString(stub shim.ChaincodeStubInterface, queryString string) ([]byte, error) {

	resultsIterator, err := stub.GetQueryResult(queryString)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing QueryRecords
	var buffer bytes.Buffer

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}

		// Record is a JSON object, so we write as-is
		buffer.WriteString(string(queryResponse.Value))
		bArrayMemberAlreadyWritten = true
	}

	fmt.Printf("- getQueryResultForQueryString queryResult:\n%s\n", buffer.String())

	return buffer.Bytes(), nil

}

func (e *EducationChaincode) addEdu(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 2 {
		return shim.Error("input parameter error!")
	}

	var edu education.Education

	err := json.Unmarshal([]byte(args[0]), &edu)
	if err != nil {
		return shim.Error("unmarshal has error!")
	}

	_, exist := GetEduInfo(stub, edu.EntityID)
	if exist {
		return shim.Error("entity has exist!")
	}

	_, saveFlag := PutEdu(stub, edu)
	if !saveFlag {
		return shim.Error("save edu has error!")
	}

	err = stub.SetEvent(args[1], []byte{})
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success([]byte("save edu success!"))
}

// 根据证书编号及姓名查询信息
// args: CertNo, name
func (e *EducationChaincode) queryEduByCertNoAndName(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) != 2 {
		return shim.Error("给定的参数个数不符合要求")
	}
	CertNo := args[0]
	name := args[1]

	// 拼装CouchDB所需要的查询字符串(是标准的一个JSON串)
	// queryString := fmt.Sprintf("{\"selector\":{\"docType\":\"eduObj\", \"CertNo\":\"%s\"}}", CertNo)
	queryString := fmt.Sprintf("{\"selector\":{\"docType\":\"%s\", \"CertNo\":\"%s\", \"Name\":\"%s\"}}", DOC_TYPE, CertNo, name)

	// 查询数据
	result, err := getEduByQueryString(stub, queryString)
	if err != nil {
		return shim.Error("根据证书编号及姓名查询信息时发生错误")
	}
	if result == nil {
		return shim.Error("根据指定的证书编号及姓名没有查询到相关的信息")
	}
	return shim.Success(result)
}

// 根据身份证号码查询详情（溯源）
// args: entityID
func (e *EducationChaincode) queryEduInfoByEntityID(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 1 {
		return shim.Error("给定的参数个数不符合要求")
	}

	// 根据身份证号码查询edu状态
	b, err := stub.GetState(args[0])
	if err != nil {
		return shim.Error("根据身份证号码查询信息失败")
	}

	if b == nil {
		return shim.Error("根据身份证号码没有查询到相关的信息")
	}

	// 对查询到的状态进行反序列化
	var edu education.Education
	err = json.Unmarshal(b, &edu)
	if err != nil {
		return shim.Error("反序列化edu信息失败")
	}

	// 获取历史变更数据
	iterator, err := stub.GetHistoryForKey(edu.EntityID)
	if err != nil {
		return shim.Error("根据指定的身份证号码查询对应的历史变更数据失败")
	}
	defer iterator.Close()

	// 迭代处理
	var historys []education.HistoryItem
	var hisEdu education.Education
	for iterator.HasNext() {
		hisData, err := iterator.Next()
		if err != nil {
			return shim.Error("获取edu的历史变更数据失败")
		}

		var historyItem education.HistoryItem
		historyItem.TxId = hisData.TxId
		err = json.Unmarshal(hisData.Value, &hisEdu)
		if err != nil {
			return shim.Error("反序列化edu信息失败")
		}

		if hisData.Value == nil {
			var empty education.Education
			historyItem.Education = empty
		} else {
			historyItem.Education = hisEdu
		}

		historys = append(historys, historyItem)

	}

	edu.Historys = historys

	// 返回
	result, err := json.Marshal(edu)
	if err != nil {
		return shim.Error("序列化edu信息时发生错误")
	}
	return shim.Success(result)
}

// 根据身份证号更新信息
// args: educationObject
func (e *EducationChaincode) updateEdu(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 2 {
		return shim.Error("给定的参数个数不符合要求")
	}

	var info education.Education
	err := json.Unmarshal([]byte(args[0]), &info)
	if err != nil {
		return shim.Error("反序列化edu信息失败")
	}

	// 根据身份证号码查询信息
	result, bl := GetEduInfo(stub, info.EntityID)
	if !bl {
		return shim.Error("根据身份证号码查询信息时发生错误")
	}

	result.EnrollDate = info.EnrollDate
	result.GraduationDate = info.GraduationDate
	result.SchoolName = info.SchoolName
	result.Major = info.Major
	result.QuaType = info.QuaType
	result.Length = info.Length
	result.Mode = info.Mode
	result.Level = info.Level
	result.Graduation = info.Graduation
	result.CertNo = info.CertNo

	_, bl = PutEdu(stub, result)
	if !bl {
		return shim.Error("保存信息信息时发生错误")
	}

	err = stub.SetEvent(args[1], []byte{})
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success([]byte("信息更新成功"))
}

// 根据身份证号删除信息（暂不对外提供）
// args: entityID
func (e *EducationChaincode) delEdu(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 2 {
		return shim.Error("给定的参数个数不符合要求")
	}

	/*var edu Education
	  result, bl := GetEduInfo(stub, info.EntityID)
	  err := json.Unmarshal(result, &edu)
	  if err != nil {
	      return shim.Error("反序列化信息时发生错误")
	  }*/

	err := stub.DelState(args[0])
	if err != nil {
		return shim.Error("删除信息时发生错误")
	}

	err = stub.SetEvent(args[1], []byte{})
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success([]byte("信息删除成功"))
}
package services

import (
	"encoding/json"
	"errors"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	log "github.com/sirupsen/logrus"
)

func (t *ServiceSetup) SaveEdu(edu Education) (string, error) {

	eventID := "eventAddEdu"
	reg, notifier := regitserEvent(t.Client, t.ChaincodeID, eventID)
	defer t.Client.UnregisterChaincodeEvent(reg)

	// 将edu对象序列化成为字节数组
	b, err := json.Marshal(edu)
	if err != nil {
		log.Errorf("指定的edu对象序列化时发生错误: %s \n", err)
		return "", errors.New("指定的edu对象序列化时发生错误")
	}

	req := channel.Request{
		ChaincodeID: t.ChaincodeID,
		Fcn:         "addEdu",
		Args:        [][]byte{b, []byte(eventID)},
	}

	response, err := t.Client.Execute(req)
	if err != nil {
		log.Errorln(err)
		return "", err
	}

	err = eventResult(notifier, eventID)
	if err != nil {
		log.Errorln(err)
		return "", err
	}

	return string(response.TransactionID), nil
}

func (t *ServiceSetup) FindEduInfoByEntityID(entityID string) ([]byte, error) {

	req := channel.Request{
		ChaincodeID: t.ChaincodeID,
		Fcn:         "queryEduInfoByEntityID",
		Args:        [][]byte{[]byte(entityID)},
	}
	response, err := t.Client.Query(req)
	if err != nil {
		log.Errorln(err)
		return []byte{0x00}, err
	}

	return response.Payload, nil
}

func (t *ServiceSetup) ModifyEdu(edu Education) (string, error) {

	eventID := "eventModifyEdu"
	reg, notifier := regitserEvent(t.Client, t.ChaincodeID, eventID)
	defer t.Client.UnregisterChaincodeEvent(reg)

	// 将edu对象序列化成为字节数组
	b, err := json.Marshal(edu)
	if err != nil {
		log.Errorf("指定的edu对象序列化时发生错误: %s \n", err)
		return "", errors.New("指定的edu对象序列化时发生错误")
	}

	req := channel.Request{
		ChaincodeID: t.ChaincodeID,
		Fcn:         "updateEdu",
		Args:        [][]byte{b, []byte(eventID)},
	}
	response, err := t.Client.Execute(req)
	if err != nil {
		log.Errorln(err)
		return "", err
	}

	err = eventResult(notifier, eventID)
	if err != nil {
		log.Errorln(err)
		return "", err
	}

	return string(response.TransactionID), nil
}

func (t *ServiceSetup) FindEduByCertNoAndName(certNo, name string) ([]byte, error) {

	req := channel.Request{
		ChaincodeID: t.ChaincodeID,
		Fcn:         "queryEduByCertNoAndName",
		Args:        [][]byte{[]byte(certNo), []byte(name)},
	}
	response, err := t.Client.Query(req)
	if err != nil {
		log.Errorln(err)
		return []byte{0x00}, err
	}

	return response.Payload, nil
}

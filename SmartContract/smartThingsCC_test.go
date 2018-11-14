package main

import (
	"fmt"
	"testing"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func Test_Init(t *testing.T) {
	simpleCC := new(SimpleAsset)
	mockStub := shim.NewMockStub("mockstub", simpleCC)
	txId := "mockTxID"

	mockStub.MockTransactionStart(txId)
	response := simpleCC.Init(mockStub)
	mockStub.MockTransactionEnd(txId)
	if s := response.GetStatus(); s != 200 {
		fmt.Println("Init test failed")
		t.FailNow()
	}
}

func Test_Invoke_noFunction(t *testing.T) {
	simpleCC := new(SimpleAsset)
	mockStub := shim.NewMockStub("mockstub", simpleCC)
	txId := "mockTxID"
	mockStub.MockTransactionStart(txId)
	args := []string{"noFunction", "key1", "value1"}
	argsAsBytes := make([][]byte, len(args))
	for i, v := range args {
		argsAsBytes[i] = []byte(v)
	}
	response := mockStub.MockInvoke(txId, argsAsBytes)
	mockStub.MockTransactionEnd(txId)
	if s := response.GetStatus(); s != 404 {
		fmt.Println("Invoke_noFunction test failed")
		t.FailNow()
	}
}
func Test_Invoke_saveNewEvent(t *testing.T) {
	simpleCC := new(SimpleAsset)
	mockStub := shim.NewMockStub("mockstub", simpleCC)
	txId := "mockTxID"
	mockStub.MockTransactionStart(txId)
	args := []string{"saveNewEvent", "sswitch", "sswitch", "true", "e9221d8090a911e8b7f70af90b9e3042", "", "sswitch switch is off", "null", "false", "false", "84242cad1b7a4541acb2f7af98899c8e", "locationCreateNew", "8e25e0b1c2a94be78db8a9f558de518f", "DEVICE", "null", "off", "switch", "20180727T08:00:07.000Z"}
	argsAsBytes := make([][]byte, len(args))
	for i, v := range args {
		argsAsBytes[i] = []byte(v)
	}
	response := mockStub.MockInvoke(txId, argsAsBytes)
	mockStub.MockTransactionEnd(txId)
	if s := response.GetStatus(); s != 200 {
		fmt.Println("Invoke_saveNewEvent test failed")
		t.FailNow()
	}
}

func Test_Main(t *testing.T) {
	main()
}

func Test_saveNewEvent(t *testing.T) {
	simpleCC := new(SimpleAsset)
	mockStub := shim.NewMockStub("mockstub", simpleCC)
	txId := "mockTxID"

	args := []string{"sswitch", "sswitch", "true", "e9221d8090a911e8b7f70af90b9e3042", "", "sswitch switch is off", "null", "false", "false", "84242cad1b7a4541acb2f7af98899c8e", "locationCreateNew", "8e25e0b1c2a94be78db8a9f558de518f", "DEVICE", "null", "off", "switch", "20180727T08:00:07.000Z"}
	mockStub.MockTransactionStart(txId)
	response := simpleCC.saveNewEvent(mockStub, args)
	mockStub.MockTransactionEnd(txId)
	fmt.Println("Status: " + fmt.Sprint(response.GetStatus()))
	fmt.Println("Payload: " + string(response.GetPayload()))
	fmt.Println("Message: " + response.GetMessage())

	if s := response.GetStatus(); s != 200 {
		fmt.Println("saveNewEvent test failed")
		t.FailNow()
	}
}

func Test_saveNewEvent_incorrectArgs(t *testing.T) {
	simpleCC := new(SimpleAsset)
	mockStub := shim.NewMockStub("mockstub", simpleCC)
	txId := "mockTxID"

	args := []string{"value1", "value2"}
	mockStub.MockTransactionStart(txId)
	response := simpleCC.saveNewEvent(mockStub, args)
	mockStub.MockTransactionEnd(txId)
	fmt.Println("Status: " + fmt.Sprint(response.GetStatus()))
	fmt.Println("Payload: " + string(response.GetPayload()))
	fmt.Println("Message: " + response.GetMessage())

	if s := response.GetStatus(); s != 400 {
		fmt.Println("saveNewEvent_incorrectArgs test failed")
		t.FailNow()
	}
}

func Test_saveNewEvent_emptyKey(t *testing.T) {
	simpleCC := new(SimpleAsset)
	mockStub := shim.NewMockStub("mockstub", simpleCC)
	txId := "mockTxID"

	args := []string{"sswitch", "sswitch", "true", "e9221d8090a911e8b7f70af90b9e3042", "", "sswitch switch is off", "null", "false", "false", "", "locationCreateNew", "8e25e0b1c2a94be78db8a9f558de518f", "DEVICE", "null", "off", "switch", ""}
	mockStub.MockTransactionStart(txId)
	response := simpleCC.saveNewEvent(mockStub, args)
	mockStub.MockTransactionEnd(txId)
	fmt.Println("Status: " + fmt.Sprint(response.GetStatus()))
	fmt.Println("Payload: " + string(response.GetPayload()))
	fmt.Println("Message: " + response.GetMessage())

	if s := response.GetStatus(); s != 400 {
		fmt.Println("emptyKey_incorrectArgs test failed")
		t.FailNow()
	}
}

func Test_queryByDate_incorrectArgs(t *testing.T) {
	simpleCC := new(SimpleAsset)
	mockStub := shim.NewMockStub("mockstub", simpleCC)
	txId := "mockTxID"

	args := []string{"key1"}
	mockStub.MockTransactionStart(txId)
	response := simpleCC.queryByDate(mockStub, args)
	mockStub.MockTransactionEnd(txId)
	fmt.Println("Status: " + fmt.Sprint(response.GetStatus()))
	fmt.Println("Payload: " + string(response.GetPayload()))
	fmt.Println("Message: " + response.GetMessage())

	if s := response.GetStatus(); s != 400 {
		fmt.Println("queryByDate_incorrectArgs test failed")
		t.FailNow()
	}
}

/*
 * Copyright IBM Corp All Rights Reserved
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package main

import (
	"bytes"
	"fmt"
	"strings"
	"time"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

// SimpleAsset implements a simple chaincode to manage an asset
type SimpleAsset struct {
}

// Init is called during chaincode instantiation to initialize any
// data. Note that chaincode upgrade also calls this function to reset
// or to migrate data.
func (t *SimpleAsset) Init(stub shim.ChaincodeStubInterface) peer.Response {
	return shim.Success(nil)
}

// Invoke is called per transaction on the chaincode. Each transaction is
// either a 'get' or a 'set' on the asset created by Init function. The Set
// method may create a new asset by specifying a new key-value pair.
func (t *SimpleAsset) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	// Extract the function and args from the transaction proposal
	function, args := stub.GetFunctionAndParameters()

	if function == "saveNewEvent" {
		return t.saveNewEvent(stub, args)
	} else if function == "getDeviceLastEvent" {
		return t.getDeviceLastEvent(stub)
	} else if function == "getHistoryForDevice" {
		return t.getHistoryForDevice(stub, args)
	} else if function == "getEvent" {
		return t.getEvent(stub, args)
	}

	return shim.Error("Invalid function name for 'invoke'")
}

// Set stores the asset (both key and value) on the ledger. If the key exists,
// it will override the value with the new one
func (t *SimpleAsset) saveNewEvent(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 17 {
		return shim.Error("incorrect arguments. Expecting full event details")
	}

	displayName := strings.ToLower(args[0])
	device := strings.ToLower(args[1])
	isStateChange := strings.ToLower(args[2])
	id := strings.ToLower(args[3])
	description := strings.ToLower(args[4])
	descriptionText := strings.ToLower(args[5])
	installedSmartAppID := strings.ToLower(args[6])
	isDigital := strings.ToLower(args[7])
	isPhysical := strings.ToLower(args[8])
	deviceID := strings.ToLower(args[9])
	location := strings.ToLower(args[10])
	locationID := strings.ToLower(args[11])
	source := strings.ToLower(args[12])
	unit := strings.ToLower(args[13])
	value := strings.ToLower(args[14])
	name := strings.ToLower(args[15])
	time := strings.ToLower(args[16])

	//Building the event json string manually without struct marshalling
	eventJSONasString := `{"docType":"Event",  "displayName": "` + displayName + `",
	 "device": "` + device + `", "isStateChange": "` + isStateChange + `",
	 "id": "` + id + `", "description": "` + description + `",
	 "descriptionText": "` + descriptionText + `", "installedSmartAppId": "` + installedSmartAppID + `",
	 "isDigital": "` + isDigital + `", "isPhysical": "` + isPhysical + `", "deviceId": "` + deviceID + `",
	 "location": "` + location + `", "locationId": "` + locationID + `", "source": "` + source + `",
	 "unit": "` + unit + `", "value": "` + value + `", "name": "` + name + `", "time": "` + time + `"}`
	eventJSONasBytes := []byte(eventJSONasString)

	err1 := stub.PutState(deviceID, eventJSONasBytes)
	if err1 != nil {
		return shim.Error("Failed to set asset")
	}
	return shim.Success([]byte(device))
}

// Get returns the value of the specified asset key
func (t *SimpleAsset) getEvent(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	var userID, jsonResp string
	var err error

	if len(args) != 1 {
		return shim.Error("Incorrect arguments. Expecting a key")
	}

	valueAsBytes, err := stub.GetState(args[0])
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + userID + "\"}"
		return shim.Error(jsonResp)
	}
	if valueAsBytes == nil {
		jsonResp = "{\"Error\":\"Transaction does not exist: " + userID + "\"}"
		return shim.Error(jsonResp)
	}
	return shim.Success(valueAsBytes)
}

// Gets the last transactions for all unique keys i.e., all devices
func (t *SimpleAsset) getDeviceLastEvent(stub shim.ChaincodeStubInterface) peer.Response {
	resultsIterator, err := stub.GetStateByRange("", "")
	if err != nil {
		return shim.Error("Failed to get data " + err.Error())
	}
	defer resultsIterator.Close()

	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false

	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}

		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"DeviceId\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.Key)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Record\":")
		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	return shim.Success(buffer.Bytes())
}

// For a given device, return its historical data
func (t *SimpleAsset) getHistoryForDevice(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 1 {
		return shim.Error("Expected argument lenght doesnt match")
	}

	resultsIterator, err := stub.GetHistoryForKey(args[0])
	if err != nil {
		return shim.Error("Failed to get data " + err.Error())
	}
	defer resultsIterator.Close()

	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false

	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}

		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"TxId\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.TxId)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Timestamp\":")
		buffer.WriteString("\"")
		buffer.WriteString(time.Unix(queryResponse.Timestamp.Seconds, int64(queryResponse.Timestamp.Nanos)).String())
		buffer.WriteString("\"")

		buffer.WriteString(", \"Record\":")
		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	return shim.Success(buffer.Bytes())
}

// main function starts up the chaincode in the container during instantiate
func main() {
	if err := shim.Start(new(SimpleAsset)); err != nil {
		fmt.Printf("Error starting SimpleAsset chaincode: %s", err)
	}
}

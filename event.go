/*
 * Copyright IBM Corp All Rights Reserved
 *
 * SPDX-License-Identifier: Apache-2.0
 */

package main

import (
	"fmt"
	"strings"

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
	fn, args := stub.GetFunctionAndParameters()

	var result string
	var err error
	if fn == "saveNewEvent" {
		result, err = saveNewEvent(stub, args)
	} else { // assume 'getEvent' even if fn is nil
		result, err = getEvent(stub, args)
	}
	if err != nil {
		return shim.Error(err.Error())
	}

	// Return the result as success payload
	return shim.Success([]byte(result))
}

// Set stores the asset (both key and value) on the ledger. If the key exists,
// it will override the value with the new one
func saveNewEvent(stub shim.ChaincodeStubInterface, args []string) (string, error) {
	if len(args) != 17 {
		return "", fmt.Errorf("incorrect arguments. Expecting full event details")
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

	err1 := stub.PutState(id, eventJSONasBytes)
	if err1 != nil {
		return "", fmt.Errorf("Failed to set asset: %s", args[0])
	}
	return id, nil
}

// Get returns the value of the specified asset key
func getEvent(stub shim.ChaincodeStubInterface, args []string) (string, error) {
	var userID, jsonResp string
	var err error

	if len(args) != 1 {
		return "", fmt.Errorf("Incorrect arguments. Expecting a key")
	}

	valueAsBytes, err := stub.GetState(args[0])
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + userID + "\"}"
		return "", fmt.Errorf(jsonResp)
	}
	if valueAsBytes == nil {
		jsonResp = "{\"Error\":\"Marble does not exist: " + userID + "\"}"
		return "", fmt.Errorf(jsonResp)
	}
	return string(valueAsBytes), nil
}

// main function starts up the chaincode in the container during instantiate
func main() {
	if err := shim.Start(new(SimpleAsset)); err != nil {
		fmt.Printf("Error starting SimpleAsset chaincode: %s", err)
	}
}

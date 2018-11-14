/**
 *  Xooa SmartThings Event Logger Chaincode
 *
 *  Copyright 2018 Xooa
 *
 *  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except
 *  in compliance with the License. You may obtain a copy of the License at:
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing, software distributed under the License is distributed
 *  on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License
 *  for the specific language governing permissions and limitations under the License.
 */
/*
 * Original source via IBM Corp:
 *  https://hyperledger-fabric.readthedocs.io/en/release-1.2/chaincode4ade.html#pulling-it-all-together
 *
 * Modifications from: Arisht Jain:
 *  https://github.com/xooa/samples
 */

package main

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

var logger = shim.NewLogger("smartThingsCC")

// SimpleAsset implements a simple smart contract to manage an asset
type SimpleAsset struct {
}

// Init is called during smart contract instantiation to initialize any
// data. Note that smart contract upgrade also calls this function to reset
// or to migrate data.
func (t *SimpleAsset) Init(stub shim.ChaincodeStubInterface) peer.Response {
	logger.Debug("Init() called.")
	return shim.Success(nil)
}

// Invoke is called per transaction on the smart contract. Each transaction is
// either updating the state or retreiving the state created by Init function.
func (t *SimpleAsset) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	logger.Debug("Invoke() called.")
	// Extract the function and args from the transaction proposal
	function, args := stub.GetFunctionAndParameters()

	if function == "saveNewEvent" {
		return t.saveNewEvent(stub, args)
	} else if function == "queryByDate" {
		return t.queryByDate(stub, args)
	} else if function == "queryLocation" {
		return t.queryLocation(stub)
	}
	logger.Error("Function declaration not found for ", function)
	resp := shim.Error("Invalid function name : " + function)
	resp.Status = 404
	return resp
}

// saveNewEvent stores the event on the ledger. For each device
// it will override the current state with the new one
func (t *SimpleAsset) saveNewEvent(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	logger.Debug("saveNewEvent() called.")
	// Essential check to verify number of arguments
	if len(args) != 17 {
		logger.Error("Incorrect number of arguments passed in saveNewEvent.")
		resp := shim.Error("Incorrect number of arguments. Expecting 17 arguments: " + strconv.Itoa(len(args)) + " given.")
		resp.Status = 400
		return resp
	}

	displayName := args[0]
	device := args[1]
	isStateChange := args[2]
	id := args[3]
	description := args[4]
	descriptionText := args[5]
	installedSmartAppID := args[6]
	isDigital := args[7]
	isPhysical := args[8]
	deviceID := args[9]
	location := args[10]
	locationID := args[11]
	source := args[12]
	unit := args[13]
	value := args[14]
	name := args[15]
	time := args[16]

	if deviceID == "" || time == "" {
		logger.Error("Empty key passed to saveNewResponse()")
		resp := shim.Error("Key must not be empty.")
		resp.Status = 400
		return resp
	} else {
		date := strings.Replace(time, "-", "", -1)
		date = strings.Split(date, "T")[0]

		//Building the event json string manually without struct marshalling
		eventJSONasString := `{"docType":"Event",  "displayName": "` + displayName + `",
			"device": "` + device + `", "isStateChange": "` + isStateChange + `",
			"id": "` + id + `", "description": "` + description + `",
			"descriptionText": "` + descriptionText + `", "installedSmartAppId": "` + installedSmartAppID + `",
			"isDigital": "` + isDigital + `", "isPhysical": "` + isPhysical + `", "deviceId": "` + deviceID + `",
			"location": "` + location + `", "locationId": "` + locationID + `", "source": "` + source + `",
			"unit": "` + unit + `", "value": "` + value + `", "name": "` + name + `", "time": "` + time + `", "date": "` + date + `"}`
		eventJSONasBytes := []byte(eventJSONasString)

		logger.Debug("eventJSONasString: ", eventJSONasString)

		eventLessArgsString := `{"docType":"EventLess",  "displayName": "` + displayName + `", "value": "` + value + `",
	 		"time": "` + time + `", "locationId": "` + locationID + `"}`
		eventLessArgs := []byte(eventLessArgsString)
		err := stub.PutState(deviceID, eventLessArgs)

		if err != nil {
			logger.Error("Error occured while calling PutState(): ", err)
			return shim.Error("Error in updating ledger.")
		}

		arr := []string{deviceID, time}
		myCompositeKey, err := stub.CreateCompositeKey("deviceId~time", arr)
		if err != nil {
			logger.Error("Error occured while calling CreateCompositeKey(): ", err)
			return shim.Error("Failed to set composite key")
		}
		err = stub.PutState(myCompositeKey, eventJSONasBytes)
		if err != nil {
			logger.Error("Error occured while calling PutState(): ", err)
			return shim.Error("Error in updating ledger.")
		}
	}
	return shim.Success([]byte(device))
}

// main function starts up the smart contract in the container during instantiate
func main() {
	logger.Debug("main() called.")
	if err := shim.Start(new(SimpleAsset)); err != nil {
		logger.Error("Error starting SimpleAsset smart contract: ", err)
		fmt.Printf("Error starting SimpleAsset smart contract: %s", err)
	}
}

// getQueryResultForQueryString retrieves the data from couchdb
// for rich queries passed as a string
func getQueryResultForQueryString(stub shim.ChaincodeStubInterface, queryString string) ([]byte, error) {
	logger.Debug("getQueryResultForQueryString called.")
	fmt.Printf("- getQueryResultForQueryString queryString:\n%s\n", queryString)

	resultsIterator, err := stub.GetQueryResult(queryString)
	if err != nil {
		logger.Error("Error occured while calling GetQueryResult(): ", err)
		return nil, err
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing QueryRecords
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			logger.Error("Error occured while calling resultsIterator.Next(): ", err)
			return nil, err
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.Key)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Record\":")
		// Record is a JSON object, so we write as-is
		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("- getQueryResultForQueryString queryResult:\n%s\n", buffer.String())

	return buffer.Bytes(), nil
}

// queryLocation creates a rich query to query the location using locationId.
// It retrieve all the devices and their last states for that location.
func (t *SimpleAsset) queryLocation(stub shim.ChaincodeStubInterface) peer.Response {
	logger.Debug("queryLocation called.")
	queryString := fmt.Sprintf("{\r\n    \"selector\": {\r\n        \"docType\": \"EventLess\"\r\n    },\r\n    \"fields\": [\"displayName\", \"value\",\"time\"]\r\n}")

	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		logger.Error("Error occured while calling getQueryResultForQueryString(): ", err.Error())
		return shim.Error("Error occured while calling getQueryResultForQueryString: " + err.Error())
	}

	return shim.Success(queryResults)
}

// queryByDate creates a rich query to query using locationId, deviceId and date.
// It retrieves all the history of the device for a particular date.
func (t *SimpleAsset) queryByDate(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	logger.Debug("queryByDate called.")
	// Essential check to verify number of arguments
	if len(args) != 2 {
		logger.Error("Incorrect number of arguments passed in queryByDate.")
		resp := shim.Error("Incorrect number of arguments. Expecting 2 arguments: " + strconv.Itoa(len(args)) + " given.")
		resp.Status = 400
		return resp
	}

	deviceId := args[0]
	date := args[1]
	queryString := fmt.Sprintf("{\r\n    \"selector\": {\r\n        \"docType\": \"Event\",\r\n        \"deviceId\": \"%s\",\r\n        \"date\": \"%s\"\r\n    },\r\n    \"fields\": [\"value\",\"time\"]\r\n}", deviceId, date)

	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		logger.Error("Error occured while calling getQueryResultForQueryString(): ", err.Error())
		return shim.Error("Error occured while calling getQueryResultForQueryString: " + err.Error())
	}
	queryResultsString := strings.Replace(string(queryResults), "\u0000", "||", -1)
	queryResults = []byte(queryResultsString)
	return shim.Success(queryResults)
}

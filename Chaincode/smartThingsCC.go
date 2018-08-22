/**
 *  Blockchain Event Logger
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
 *  https://github.com/xooa/smartThings-xooa
 *
 * Changes:
 *  Logs to Xooa blockchain platform from SmartThings instead from user
 */

package main

import (
	"bytes"
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
// either updating the state or retreiving the state created by Init function.
func (t *SimpleAsset) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	// Extract the function and args from the transaction proposal
	function, args := stub.GetFunctionAndParameters()

	if function == "saveNewEvent" {
		return t.saveNewEvent(stub, args)
	} else if function == "queryByDate" {
		return t.queryByDate(stub, args)
	} else if function == "queryLocation" {
		return t.queryLocation(stub, args)
	}

	return shim.Error("Invalid function name for 'invoke'")
}

// saveNewEvent stores the event on the ledger. For each device
// it will override the current state with the new one
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
	date := strings.Replace(time, "-", "", -1)
	date = strings.Split(date, "t")[0]

	//Building the event json string manually without struct marshalling
	eventJSONasString := `{"docType":"Event",  "displayName": "` + displayName + `",
	 "device": "` + device + `", "isStateChange": "` + isStateChange + `",
	 "id": "` + id + `", "description": "` + description + `",
	 "descriptionText": "` + descriptionText + `", "installedSmartAppId": "` + installedSmartAppID + `",
	 "isDigital": "` + isDigital + `", "isPhysical": "` + isPhysical + `", "deviceId": "` + deviceID + `",
	 "location": "` + location + `", "locationId": "` + locationID + `", "source": "` + source + `",
	 "unit": "` + unit + `", "value": "` + value + `", "name": "` + name + `", "time": "` + time + `", "date": "` + date + `"}`
	eventJSONasBytes := []byte(eventJSONasString)

	eventLessArgsString := `{"docType":"EventLess",  "displayName": "` + displayName + `", "value": "` + value + `",
	 "time": "` + time + `", "locationId": "` + locationID + `"}`
	eventLessArgs := []byte(eventLessArgsString)
	err := stub.PutState(deviceID, eventLessArgs)
	if err != nil {
		return shim.Error("Failed to set asset")
	}

	arr := []string{deviceID, time}
	myCompositeKey, err := stub.CreateCompositeKey("combined", arr)
	if err != nil {
		return shim.Error("Failed to set composite key")
	}
	err = stub.PutState(myCompositeKey, eventJSONasBytes)
	if err != nil {
		return shim.Error("Failed to set asset")
	}
	return shim.Success([]byte(device))
}

// main function starts up the chaincode in the container during instantiate
func main() {
	if err := shim.Start(new(SimpleAsset)); err != nil {
		fmt.Printf("Error starting SimpleAsset chaincode: %s", err)
	}
}

// getQueryResultForQueryString retrieves the data from couchdb
// for rich queries passed as a string
func getQueryResultForQueryString(stub shim.ChaincodeStubInterface, queryString string) ([]byte, error) {

	fmt.Printf("- getQueryResultForQueryString queryString:\n%s\n", queryString)

	resultsIterator, err := stub.GetQueryResult(queryString)
	if err != nil {
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
func (t *SimpleAsset) queryLocation(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) < 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	locationId := args[0]

	queryString := fmt.Sprintf("{\r\n    \"selector\": {\r\n        \"docType\": \"EventLess\",\r\n        \"locationId\": \"%s\"\r\n    },\r\n    \"fields\": [\"displayName\", \"value\",\"time\"]\r\n}", locationId)

	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}

	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(queryResults)
}

// queryByDate creates a rich query to query using locationId, deviceId and date.
// It retrieves all the history of the device for a particular date.
func (t *SimpleAsset) queryByDate(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) < 2 {
		return shim.Error("Incorrect number of arguments. Expecting 3")
	}

	locationId := args[0]
	deviceId := args[1]
	date := args[2]
	queryString := fmt.Sprintf("{\r\n    \"selector\": {\r\n        \"docType\": \"Event\",\r\n        \"locationId\": \"%s\",\r\n        \"deviceId\": \"%s\",\r\n        \"date\": \"%s\"\r\n    },\r\n    \"fields\": [\"value\",\"time\"]\r\n}", locationId, deviceId, date)

	queryResults, err := getQueryResultForQueryString(stub, queryString)
	queryResultsString := strings.Replace(string(queryResults), "\u0000", "||", -1)
	if err != nil {
		return shim.Error(err.Error())
	}
	queryResults = []byte(queryResultsString)
	return shim.Success(queryResults)
}

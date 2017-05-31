/*
Copyright IBM Corp 2016 All Rights Reserved.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
		 http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"errors"
	"fmt"

	"github.com/IBM-Bluemix/go-cloudant"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct {
}

func main() {
	err := shim.Start(new(SimpleChaincode))

	fmt.Println("****** Starting to send information to my ledger")

	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}

// Init resets all the things
func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting 1")
	}

	err := stub.PutState("hello_world", []byte(args[0]))

	if err != nil {
		return nil, err
	}

	return nil, nil
}

// Invoke isur entry point to invoke a chaincode function
func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("invoke is running " + function)

	// Handle different functions
	if function == "init" {
		return t.Init(stub, "init", args)
	} else if function == "lumos" {
		return t.myMethod(stub, args)
	}
	fmt.Println("invoke did not find func: " + function)

	return nil, errors.New("Received unknown function invocation: " + function)
}

// Query is our entry point for queries
func (t *SimpleChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("query is running " + function)

	// Handle different functions
	if function == "read" { //read a variable
		return t.read(stub, args)
	}

	fmt.Println("query did not find func: " + function)

	return nil, errors.New("Received unknown function query: " + function)
}

// write - invoke function to write key/value pair
func (t *SimpleChaincode) myMethod(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	var key, value string
	var err error

	fmt.Println("****** Starting to send information to my ledger")

	type data struct {
		ID   string `json:"tax_id"`
		Name string `json:"name"`
	}
	testData := &data{
		ID:   "1453",
		Name: "constantinople",
	}

	fmt.Println("****** Starting to create my client")

	client, error := cloudant.NewClient("attaidifieveredgedatingi", "85363815ee8b0396585f5cf7d3a12508aba2a3d2")

	fmt.Println("****** Completed creating my client, ensuring db exists")

	fmt.Println("**** Error while creating client: " + error.Error())

	client.CreateDB("https://ab5e5a7c-76de-4d8a-8516-64e21e8c4042-bluemix:9d794d83dc3913e7958a6ce546293269aab45ef097d5610b6eb43b9c6bf0bd7e@ab5e5a7c-76de-4d8a-8516-64e21e8c4042-bluemix.cloudant.com")

	fmt.Println("****** Completed ensuring db exists, creating DB object with url- https://ab5e5a7c-76de-4d8a-8516-64e21e8c4042-bluemix:9d794d83dc3913e7958a6ce546293269aab45ef097d5610b6eb43b9c6bf0bd7e@ab5e5a7c-76de-4d8a-8516-64e21e8c4042-bluemix.cloudant.com")

	dbObj := client.DB("https://ab5e5a7c-76de-4d8a-8516-64e21e8c4042-bluemix:9d794d83dc3913e7958a6ce546293269aab45ef097d5610b6eb43b9c6bf0bd7e@ab5e5a7c-76de-4d8a-8516-64e21e8c4042-bluemix.cloudant.com")

	fmt.Println(" ******** got DB object dbObj, creating document by calling dbObj.CreateDocument(testData)")

	dbObj.CreateDocument(testData)

	fmt.Println("running write()")

	if len(args) != 2 {
		return nil, errors.New("Incorrect number of arguments. Expecting 2. name of the key and value to set")
	}

	key = args[0] //rename for funsies
	value = args[1]
	err = stub.PutState(key, []byte(value)) //write the variable into the chaincode state

	if err != nil {
		return nil, err
	}
	return nil, nil
}

// read - query function to read key/value pair
func (t *SimpleChaincode) read(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var key, jsonResp string
	var err error

	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting name of the key to query")
	}

	key = args[0]
	valAsbytes, err := stub.GetState(key)
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + key + "\"}"
		return nil, errors.New(jsonResp)
	}

	return valAsbytes, nil
}

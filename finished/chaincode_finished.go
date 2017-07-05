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
	"encoding/json"
	"errors"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

type MarketerStruct struct {
	EId                   string `json:"EId"`
	TaxId                 string `json:"TaxId"`
	BeginDate             string `json:"BeginDate"`
	MarketerTypeFlag      string `json:"MarketerTypeFlag"`
	MarketerType          string `json:"MarketerType"`
	MarketerRole          string `json:"MarketerRole"`
	MarketerStatus        string `json:"MarketerStatus"`
	LegalName             string `json:"LegalName"`
	Gender                string `json:"Gender"`
	DoB                   string `json:"DoB"`
	RegStateName          string `json:"RegStateName"`
	MarketerEffectiveDate string `json:"MarketerEffectiveDate"`
	MarketerEndDate       string `json:"MarketerEndDate"`
	FirstName             string `json:"FirstName"`
	LastName              string `json:"LastName"`
	BusinessAddress       string `json:"BusinessAddress"`
	City                  string `json:"City"`
	State                 string `json:"State"`
	PostalCode            string `json:"PostalCode"`
	PhoneNumber           string `json:"PhoneNumber"`
	EMail                 string `json:"EMail"`
	MarketerEaRole        string `json:"“MarketerEaRole”"`
	OwnerRole             string `json:"OwnerRole"`
	OrgName               string `json:"OrgName"`
}

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
	} else if function == "write" {
		return t.write(stub, args)
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
func (t *SimpleChaincode) write(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	var key string
	var err error

	mktrStruct := MarketerStruct{
		EId:                   args[0],
		TaxId:                 args[1],
		BeginDate:             args[2],
		MarketerTypeFlag:      args[3],
		MarketerType:          args[4],
		MarketerRole:          args[5],
		MarketerStatus:        args[6],
		LegalName:             args[7],
		Gender:                args[8],
		DoB:                   args[9],
		RegStateName:          args[10],
		MarketerEffectiveDate: args[11],
		MarketerEndDate:       args[12],
		FirstName:             args[13],
		LastName:              args[14],
		BusinessAddress:       args[15],
		City:                  args[16],
		State:                 args[17],
		PostalCode:            args[18],
		PhoneNumber:           args[19],
		EMail:                 args[20],
		MarketerEaRole:        args[21],
		OwnerRole:             args[22],
		OrgName:               args[23],
	}

	mktrStructBytes, err := json.Marshal(mktrStruct)
	_ = err //ignore errors
	key = args[0]
	stub.PutState(key, mktrStructBytes)

	fmt.Println("*** successfully wrote marketer to state")

	if err != nil {
		return nil, err
	}
	return nil, nil
}

// read - query function to read key/value pair
func (t *SimpleChaincode) read(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var key, jsonResp string
	var err error

	var retrievedStruct MarketerStruct
	key = args[0]
	retrievedBytes, err := stub.GetState(key)
	json.Unmarshal(retrievedBytes, retrievedStruct)

	fmt.Println("Retrieved struct: ", retrievedStruct)

	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + key + "\"}"
		return nil, errors.New(jsonResp)
	}

	return retrievedBytes, nil
}

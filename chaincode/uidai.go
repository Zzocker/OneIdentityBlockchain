package main

import (
	"encoding/json"
	. "fmt"

	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-protos-go/peer"
)

func getState(stub shim.ChaincodeStubInterface, RKey string) ([]byte, error) {
	RByte, err := stub.GetState(RKey)
	if err != nil {
		return nil, Errorf(err.Error())
	}
	if len(RByte) == 0 {
		return nil, Errorf("Request doesn't exists")
	}
	return RByte, nil
}
func responRequest(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	// [0] = RKey , [1]= Resonse if response = accept [2]=GiveDate
	if len(args) != 3 {
		return shim.Error("Provide RKey and response")
	}
	RByte, err := getState(stub, args[0])
	if err != nil {
		return shim.Error(err.Error())
	}
	var request Request
	json.Unmarshal(RByte, &request)
	response := args[1]
	if request.Status != "0" {
		return shim.Error("Request has not been created")
	}
	if response == "-1" {
		stub.DelState(args[0])
		return shim.Error("Sorry your request has been rejected")
	}
	if len(args) != 3 {
		return shim.Error("Provide data for verification to be done")
	}
	if response == "1" {
		request.Status = "1"
		request.GivenData = args[2]
	}
	RByte, _ = json.Marshal(request)
	stub.PutState(request.ID, RByte)
	result := struct {
		Msg string
	}{
		Msg: Sprintf("Please visit UIDAI office on %s", args[2]),
	}
	output, _ := json.Marshal(result)
	return shim.Success(output)
}
func verifyPersonal(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	// args[0] = Rkey
	if len(args) != 1 {
		return shim.Error("Provide RKey")
	}
	RByte, err := getState(stub, args[0])
	if err != nil {
		return shim.Error(err.Error())
	}
	var request Request
	json.Unmarshal(RByte, &request)
	if request.Status != "1" {
		return shim.Error("Request hasn't been accepted")
	}
	IByte, _ := getState(stub, request.From)
	var identity Identity
	json.Unmarshal(IByte, &identity)
	PByte, _ := getState(stub, identity.PersonalDetails)
	var personal Personal
	json.Unmarshal(PByte, &personal)
	personal.Status = "1"
	PByte, _ = json.Marshal(personal)
	stub.PutState(personal.ID, PByte)
	stub.DelState(args[0])
	return shim.Success([]byte("Success!! Person verified"))
}

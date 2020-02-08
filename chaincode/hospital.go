package main

import (
	"encoding/json"

	"github.com/google/uuid"
	"github.com/hyperledger/fabric-chaincode-go/shim"
)

func getDOBCert(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	// args[0]= Doctor , args[1]=DOB
	if len(args) != 2 {
		return shim.Error("Please Provide Doctor name and DOB")
	}
	id := uuid.New().String()
	dob := DoB{
		DOBDetails: HealthReport{
			DocType: DOB,
			ID:      id,
			Doctor:  args[0],
			Type:    "DOB",
		},
		DOB:       args[1],
		IssueTime: time.Now().Unix(),
	}
	DByte ,_ := json.Marshal(dob)
	err:= stub.PutState(id,DByte)
	if err!=nil{
		return shim.Error(err.Error())
	}
	return shim.Success([]byte(id))
}

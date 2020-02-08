package main

import (
	"encoding/json"
	
	"time"

	"github.com/google/uuid"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-protos-go/peer"
)

func addQualification(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	// args [0] = IdeKey ,args[1] type of quali
	if len(args) != 2 {
		return shim.Error("Please provide IKey and type of quali")
	}
	DByte, err := getState(stub, args[0])
	if err != nil {
		return shim.Error(err.Error())
	}
	var identity Identity
	qKey := uuid.New().String()
	json.Unmarshal(DByte, &identity)
	qulai := Qualification{
		DocType:   QUALIFICATION,
		ID:        qKey,
		IssueTime: time.Now().Unix(),
		Status:    "1",
	}
	DByte, _ = getState(stub, identity.EducationDetails)
	var educationDet Education
	json.Unmarshal(DByte, &educationDet)
	educationDet.Qualification[args[1]] = qKey

	DByte, _ = json.Marshal(qulai)
	stub.PutState(qKey, DByte)

	DByte, _ = json.Marshal(educationDet)
	stub.PutState(educationDet.ID, DByte)
	return shim.Success(nil)
}

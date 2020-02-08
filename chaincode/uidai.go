package main

import (
	"encoding/json"
	. "fmt"
	"github.com/google/uuid"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-protos-go/peer"
)

func createIdentity(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	// [0]= name ,[1]= PAddress , [2]=CurrentAddress , [3]=Phone , [4]=Parent
	if len(args) != 5 {
		return shim.Error("[0]= name ,[1]= PAddress , [2]=CurrentAddress , [3]=Phone , [4]=Parent")
	}
	IKey := uuid.New().String()
	HKey := uuid.New().String()
	PKey := uuid.New().String()
	EKey := uuid.New().String()
	RKey := uuid.New().String()
	health := Health{
		DocType: HEALTH,
		ID:      HKey,
		Reports: make(map[in64]string),
	}
	personal := Personal{
		DocType:  PERSONAL,
		ID:       PKey,
		Name:     args[0],
		PAddress: args[1],
		CAddress: args[2],
		Phone:    args[3],
		Parent:   args[4],
		Status:   "0",
		Photo:    "null",
	}
	education := Education{
		DocType:       EDUCATION,
		ID:            EKey,
		Qualification: make(map[string]string),
	}
	Iden := Identity{
		DocType:          IDENTITY,
		ID:               IKey,
		PersonalDetails:  PKey,
		HealthDetails:    HKey,
		EducationDetails: EKey,
		Requests:         make(map[string]string),
	}
	request := Request{
		DocType:     PERSONALREQUEST,
		ID:          RKey,
		From:        IKey,
		Status:      "0",
		RequestTime: time.Now().Unix(),
	}
	Iden.Requests[PERSONALREQUEST] = RKey

	DByte, _ := json.Marshal(health)
	stub.PutState(HKey, DByte)

	DByte, _ = json.Marshal(personal)
	stub.PutState(PKey, DByte)

	DByte, _ = json.Marshal(education)
	stub.PutState(EKey, DByte)

	DByte, _ = json.Marshal(Iden)
	stub.PutState(IKey, DByte)

	DByte, _ = json.Marshal(request)
	stub.PutState(RKey, DByte)

	return shim.Success([]byte(IKey))
}
func getState(stub shim.ChaincodeStubInterface, RKey) ([]byte, error) {
	if len(args) != 1 {
		return nil, Errorf("Please provide Request key")
	}
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
	if len(args)!=2{
		return shim.Error()
	}
}

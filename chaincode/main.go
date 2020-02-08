package main

import (
	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-protos-go/peer"
)

type Chaincode struct {
}

// bcFunctions
const bcFunction = map[string]func(shim.ChaincodeStubInterface, []string) peer.Response{
	"get_dob_cert_fileName" :getDOBCert,
	"createIdentity" : createIdentity
}

func (c *Chaincode) Init(stub shim.ChaincodeStubInterface) peer.Response {
	return shim.Success(nil)
}
func (c *Chaincode) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	funcName, args := stub.GetFunctionAndParameters()
	bcFunc := bcFunction[funcName]
	if bcFunc == nil {
		return shim.Error("No function avaiable of given name")
	}
	return bcFunc(stub, args)
}
func main() {
	err := shim.Start(new(Chaincode))
	if err != nil {
		return shim.Error(err.Error())
	}
}

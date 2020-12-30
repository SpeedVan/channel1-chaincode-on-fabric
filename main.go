package main

import (
	"log"

	"github.com/SpeedVan/channel1-chaincode-on-fabric/chaincode"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func main() {
	assetChaincode, err := contractapi.NewChaincode(&chaincode.SmartContract{})
	if err != nil {
		log.Panicf("Error creating fabric-demo chaincode: %v", err)
	}

	if err := assetChaincode.Start(); err != nil {
		log.Panicf("Error starting fabric-demo chaincode: %v", err)
	}
}

package main

import (
	"log"
	"os"

	"github.com/SpeedVan/channel1-chaincode-on-fabric/chaincode"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

func main() {
	cc, err := contractapi.NewChaincode(&chaincode.SmartContract{})
	if err != nil {
		log.Panicf("Error creating fabric-demo chaincode: %v", err)
	}

	// if err := assetChaincode.Start(); err != nil {
	// 	log.Panicf("Error starting fabric-demo chaincode: %v", err)
	// }

	// if err != nil {
	// 	fmt.Println("Error starting a new ContractApi Chaincode:", err)
	// }
	
	server := &shim.ChaincodeServer{
		CCID:    os.Getenv("CHAINCODE_CCID"),
		Address: os.Getenv("CHAINCODE_ADDRESS"),
		CC:      cc,
		TLSProps: shim.TLSProperties{
				Disabled: true,
		},
	}

	// Start the chaincode external server
	err = server.Start()

	if err != nil {
		log.Panicf("Error starting FabCar chaincode server: %v", err)
	} else {
		log.Println("Succesfully started new Fabcar Chaincode server with the new ContractApi")
	}
}

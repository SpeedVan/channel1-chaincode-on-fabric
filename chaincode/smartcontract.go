package chaincode

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// VoucherAsset voucher assets
type VoucherAsset struct {
	ID           string    `json:"ID"`
	ParentID     string    `json:"parentID"`
	Issuer       string    `json:"issuer"`
	Owner        string    `json:"owner"`
	Amount       int64     `json:"amount"`
	CreateDate   time.Time `json:"createDate"`
	EndDate      time.Time `json:"endDate"`
	ContractHash string    `json:"contractHash"`
	InvoiceHash  string    `json:"invoiceHash"`
	Status       string    `json:"status"`
	Signature    string    `json:"signature"`
}

// SmartContract todo
type SmartContract struct {
	contractapi.Contract
}

// ExistsVA todo
func (s *SmartContract) ExistsVA(ctx contractapi.TransactionContextInterface, id string) (bool, error) {
	assetJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		return false, fmt.Errorf("failed to read from world state: %v", err)
	}

	return assetJSON != nil, nil
}

// CreateVA create voucher asset
func (s *SmartContract) CreateVA(ctx contractapi.TransactionContextInterface, id string, issuer string, owner string, amount int64, contractHash string, invoiceHash string) error {
	exists, err := s.ExistsVA(ctx, id)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("the asset %s already exists", id)
	}

	asset := VoucherAsset{
		ID:           id, // or invoiceHash
		Issuer:       issuer,
		Owner:        owner,
		Amount:       amount,
		CreateDate:   time.Now(),
		ContractHash: contractHash,
		InvoiceHash:  invoiceHash,
	}
	assetJSON, err := json.Marshal(asset)
	if err != nil {
		return err
	}

	return ctx.GetStub().PutState(id, assetJSON)
}

// DeleteVA todo delete voucher asset
func (s *SmartContract) DeleteVA(ctx contractapi.TransactionContextInterface, id string) error {
	exists, err := s.ExistsVA(ctx, id)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("the asset %s does not exist", id)
	}

	return ctx.GetStub().DelState(id)
}

// GetAllVA returns all voucher assets found in world state
func (s *SmartContract) GetAllVA(ctx contractapi.TransactionContextInterface) ([]*VoucherAsset, error) {
	// range query with empty string for startKey and endKey does an
	// open-ended query of all assets in the chaincode namespace.
	resultsIterator, err := ctx.GetStub().GetStateByRange("", "")
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	var assets []*VoucherAsset
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}

		asset := &VoucherAsset{}
		err = json.Unmarshal(queryResponse.Value, asset)
		if err != nil {
			return nil, err
		}
		assets = append(assets, asset)
	}

	return assets, nil
}

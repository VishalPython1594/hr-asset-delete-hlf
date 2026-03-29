package main

import (
	"fmt"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

type SmartContract struct {
	contractapi.Contract
}

// DeleteAsset deletes an asset from the ledger
func (s *SmartContract) DeleteAsset(ctx contractapi.TransactionContextInterface, id string) error {

	// TODO:
	// 1. Check if asset exists using AssetExists
	// 2. If asset does not exist → return error
	// 3. Delete asset using DelState

	return fmt.Errorf("not implemented")
}

// AssetExists checks if asset exists in world state
func (s *SmartContract) AssetExists(ctx contractapi.TransactionContextInterface, id string) (bool, error) {

	assetJSON, err := ctx.GetStub().GetState(id)

	if err != nil {
		return false, err
	}

	return assetJSON != nil, nil
}
package txdefs

import (
	"encoding/json"

	"github.com/goledgerdev/cc-tools/assets"
	"github.com/goledgerdev/cc-tools/errors"
	sw "github.com/goledgerdev/cc-tools/stubwrapper"
	tx "github.com/goledgerdev/cc-tools/transactions"
)

var GetLastRenters = tx.Transaction{
	Tag:         "getLastRenters",
	Label:       "Get last Renters of a Vehicle",
	Description: "Get last Renters of a Vehicle",
	Method:      "GET",

	Args: []tx.Argument{
		{
			Tag:         "vehicle",
			Label:       "Vehicle",
			Description: "Vehicle",
			DataType:    "->vehicle",
			Required:    true,
		},
	},
	Routine: func(stub *sw.StubWrapper, req map[string]interface{}) ([]byte, errors.ICCError) {
		vehicleKey, ok := req["vehicle"].(assets.Key)
		if !ok {
			return nil, errors.WrapError(nil, "Parameter vehicle must be an asset")
		}

		// Return history of vehicle asset
		historyAsset, err := stub.GetHistoryForKey(vehicleKey.Key())
		if err != nil {
			return nil, errors.WrapError(err, "Failed to get vehicle asset history")
		}

		var renters []interface{}
		var count int64 = 0

		defer historyAsset.Close()
		for historyAsset.HasNext() {

			historyResponse, err := historyAsset.Next()
			if err != nil {
				return nil, errors.WrapErrorWithStatus(err, "error iterating response", 500)
			}

			vehicleData := make(map[string]interface{})

			err = json.Unmarshal(historyResponse.GetValue(), &vehicleData)
			if err != nil {
				return nil, errors.WrapError(err, "failed to unmarshal historyResponse's values")
			}

			if vehicleData["renter"] != nil {
				renterKey, _ := vehicleData["renter"].(map[string]interface{})
				renter, _ := assets.NewKey(renterKey)

				renterAsset, err := renter.Get(stub)
				if err != nil {
					return nil, errors.WrapError(err, "failed to get renter assets from the ledger")
				}

				renterMap := (map[string]interface{})(*renterAsset)
				renterInfo := make(map[string]interface{})
				renterInfo["Renter Name"] = renterMap["renterName"]
				renterInfo["CPF"] = renterMap["id"]

				if !dataRenter(renters, renterInfo,
					func(data1, data2 interface{}) bool {
						return data1.(map[string]interface{})["cpf"] == data2.(map[string]interface{})["cpf"]
					}) {
					renters = append(renters, renterInfo)
					count++
				}
			}
		}

		responseJSON, nerr := json.Marshal(renters)
		if nerr != nil {
			return nil, errors.WrapError(err, "failed to marshal response")
		}
		return responseJSON, nil
	},
}

type returnData func(interface{}, interface{}) bool

// Confer values
func dataRenter(array []interface{}, value interface{}, confer returnData) bool {
	for _, key := range array {
		if confer(key, value) {
			return true
		}
	}
	return false
}

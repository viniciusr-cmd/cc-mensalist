package txdefs

import (
	"encoding/json"
	"time"

	"github.com/goledgerdev/cc-mensalist/chaincode/datatypes"

	"github.com/goledgerdev/cc-tools/assets"
	"github.com/goledgerdev/cc-tools/errors"
	sw "github.com/goledgerdev/cc-tools/stubwrapper"
	tx "github.com/goledgerdev/cc-tools/transactions"
)

// End a rent for a vehicle
var EndRent = tx.Transaction{
	Tag:         "endRent",
	Label:       "End Rent",
	Description: "End the rent of a vehicle",
	Method:      "POST",
	Callers:     []string{"$org3MSP"}, // Only org3 can call this transaction

	Args: []tx.Argument{
		{
			Tag:         "renter",
			Label:       "Renter",
			Description: "Renter",
			DataType:    "->renter",
			Required:    true,
		},
		{
			Tag:         "vehicle",
			Label:       "Vehicle",
			Description: "Vehicle",
			DataType:    "->vehicle",
			Required:    true,
		},
		{
			Tag:         "rent",
			Label:       "Rent",
			Description: "Rent",
			DataType:    "->rent",
			Required:    true,
		},
	},

	Routine: func(stub *sw.StubWrapper, req map[string]interface{}) ([]byte, errors.ICCError) {
		renterKey, ok := req["renter"].(assets.Key)
		if !ok {
			return nil, errors.WrapError(nil, "Parameter renter must be an asset")
		}

		vehicleKey, ok := req["vehicle"].(assets.Key)
		if !ok {
			return nil, errors.WrapError(nil, "Parameter vehicle must be an asset")
		}

		rentKey, ok := req["rent"].(assets.Key)
		if !ok {
			return nil, errors.WrapError(nil, "Parameter rent must be an asset")
		}

		// Get assets from ledger
		renterAsset, err := renterKey.Get(stub)
		if err != nil {
			return nil, errors.WrapError(err, "failed to get renter asset from the ledger")
		}
		renterMap := (map[string]interface{})(*renterAsset)

		vehicleAsset, err := vehicleKey.Get(stub)
		if err != nil {
			return nil, errors.WrapError(err, "failed to get vehicle asset from the ledger")
		}
		vehicleMap := (map[string]interface{})(*vehicleAsset)

		rentAsset, err := rentKey.Get(stub)
		if err != nil {
			return nil, errors.WrapError(err, "failed to get rent asset from the ledger")
		}
		rentMap := (map[string]interface{})(*rentAsset)

		// Vehicle MUST be rented
		rentStatus := rentMap["status"].(datatypes.RentStatus)
		if rentStatus != datatypes.Rented {
			return nil, errors.WrapError(nil, "Vehicle IS NOT rented")
		}

		// Get make renter, vehicle and rent key
		renterRentKey := make((map[string]interface{}))
		renterRentKey["@assetType"] = "renter"
		renterRentKey["@key"] = renterMap["@key"]

		vehicleRentKey := make((map[string]interface{}))
		vehicleRentKey["@assetType"] = "vehicle"
		vehicleRentKey["@key"] = vehicleMap["@key"]

		rentedKey := make((map[string]interface{}))
		rentedKey["@assetType"] = "rent"
		rentedKey["@key"] = rentMap["@key"]

		// Update renter
		renterMap, err = renterAsset.Update(stub, renterMap)
		if err != nil {
			return nil, errors.WrapError(err, "failed to update renter asset")
		}
		createdAt, _ := stub.Stub.GetTxTimestamp()

		// Update vehicle rent
		vehicleMap["available"] = true
		vehicleMap["renter"] = renterRentKey
		vehicleMap["rent"] = rentedKey

		// Update status rent
		rentMap["status"] = datatypes.Ended
		rentMap["rentDate"] = nil
		rentMap["endDate"] = createdAt.AsTime().Format(time.RFC3339)

		rentMap, err = rentAsset.Update(stub, rentMap)
		if err != nil {
			return nil, errors.WrapError(err, "failed to update rent asset")
		}

		vehicleMap, err = vehicleAsset.Update(stub, vehicleMap)
		if err != nil {
			return nil, errors.WrapError(err, "failed to update vehicle asset")
		}

		// Marshal to JSON
		responseMap := map[string]interface{}{
			"renter":  renterMap,
			"vehicle": vehicleMap,
			"rent":    rentMap,
		}

		responseJSON, nerr := json.Marshal(responseMap)
		if nerr != nil {
			return nil, errors.WrapError(err, "failed to marshal response")
		}

		return responseJSON, nil
	},
}

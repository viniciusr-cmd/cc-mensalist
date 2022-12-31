package txdefs

import (
	"encoding/json"

	"github.com/goledgerdev/cc-mensalist/chaincode/datatypes"
	"github.com/goledgerdev/cc-tools/assets"
	"github.com/goledgerdev/cc-tools/errors"
	sw "github.com/goledgerdev/cc-tools/stubwrapper"
	tx "github.com/goledgerdev/cc-tools/transactions"
)

// Create a new Agency on channel
// POST Method
var CreateNewVehicle = tx.Transaction{
	Tag:         "createNewVehicle",
	Label:       "Create New Vehicle",
	Description: "Create a New Vehicle",
	Method:      "POST",
	Callers:     []string{"$org2MSP", "$orgMSP"}, // Only org2 can call this transaction

	Args: []tx.Argument{
		{
			Tag:         "id",
			Label:       "Plate",
			Description: "Mercosur Plate of Vehicle",
			DataType:    "plate",
			Required:    true,
		},
		{
			Tag:         "type",
			Label:       "Vehicle type",
			Description: "Vehicle Type",
			DataType:    "vehicleType",
			Required:    true,
		},
		{
			Tag:         "model",
			Label:       "Vehicle Model Name",
			Description: "Vehicle model name",
			DataType:    "string",
			Required:    true,
		},
		{
			Tag:         "renter",
			Label:       "Renter",
			Description: "renter",
			DataType:    "->renter",
			Required:    false,
		},
		{
			Tag:         "rent",
			Label:       "Rent ID",
			Description: "rent",
			DataType:    "->rent",
			Required:    false,
		},
	},
	Routine: func(stub *sw.StubWrapper, req map[string]interface{}) ([]byte, errors.ICCError) {
		plateID, _ := req["id"].(string)
		vehicleType, _ := req["type"].(datatypes.VehicleType)
		model, _ := req["model"].(string)
		renter, _ := req["renter"].(assets.Key)
		rent, _ := req["rent"].(assets.Key)

		vehicleMap := make(map[string]interface{})

		vehicleMap["@assetType"] = "vehicle"
		vehicleMap["id"] = plateID
		vehicleMap["type"] = vehicleType
		vehicleMap["model"] = model

		if renter != nil {
			vehicleMap["renter"] = renter
		}

		if rent != nil {
			vehicleMap["rent"] = rent
		}

		vehicleAsset, err := assets.NewAsset(vehicleMap) //Essa linha: 75
		if err != nil {
			return nil, errors.WrapError(err, "Failed to create a new asset")
		}

		// Save the new agency on channel
		_, err = vehicleAsset.PutNew(stub)
		if err != nil {
			return nil, errors.WrapError(err, "Error saving asset on blockchain")
		}

		// Marshal asset back to JSON format
		vehicleJSON, nerr := json.Marshal(vehicleAsset)
		if nerr != nil {
			return nil, errors.WrapError(nil, "failed to encode asset to JSON format")
		}

		return vehicleJSON, nil
	},
}

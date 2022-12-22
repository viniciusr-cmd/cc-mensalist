package datatypes

import (
	"fmt"
	"strconv"

	"github.com/goledgerdev/cc-tools/assets"
	"github.com/goledgerdev/cc-tools/errors"
)

type VehicleType int

const (
	VehicleTypeCar VehicleType = iota
	VehicleTypeMotor
)

func (status VehicleType) CheckType() errors.ICCError {
	switch status {
	case VehicleTypeCar, VehicleTypeMotor:
		return nil
	default:
		return errors.NewCCError("invalid type", 400)
	}
}

var vehicleType = assets.DataType{
	DropDownValues: map[string]interface{}{
		"Car":        VehicleTypeCar,
		"Motorcycle": VehicleTypeMotor,
	},
	Parse: func(data interface{}) (string, interface{}, errors.ICCError) {
		var dataValue int

		switch vehicle := data.(type) {
		case VehicleType:
			dataValue = (int)(vehicle)
		case int:
			dataValue = vehicle
		case float64:
			dataValue = (int)(vehicle)
		case string:
			var err error
			dataValue, err = strconv.Atoi(vehicle)
			if err != nil {
				return "", nil, errors.WrapErrorWithStatus(err, "asset property must be an integer, is %t", 400)
			}
		default:
			return "", nil, errors.NewCCError("asset property must be an integer, is %t", 400)
		}

		retVal := (VehicleType)(dataValue)
		err := retVal.CheckType()
		return fmt.Sprint(retVal), retVal, err
	},
}

package datatypes

import (
	"fmt"
	"strconv"

	"github.com/goledgerdev/cc-tools/assets"
	"github.com/goledgerdev/cc-tools/errors"
)

type RentStatus int

const (
	Rented RentStatus = iota
	NotRented
	Ended
)

func (status RentStatus) CheckType() errors.ICCError {
	switch status {
	case Rented:
		return nil
	case NotRented:
		return nil
	case Ended:
		return nil
	default:
		return errors.NewCCError("invalid type", 400)
	}
}

var rentStatus = assets.DataType{
	DropDownValues: map[string]interface{}{
		"Rented":     Rented,
		"Not Rented": NotRented,
		"Ended":      Ended,
	},
	Parse: func(data interface{}) (string, interface{}, errors.ICCError) {
		var dataValue int

		switch v := data.(type) {
		case RentStatus:
			dataValue = (int)(v)
		case int:
			dataValue = v
		case float64:
			dataValue = (int)(v)
		case string:
			var err error
			dataValue, err = strconv.Atoi(v)
			if err != nil {
				return "", nil, errors.WrapErrorWithStatus(err, "asset property must be an integer", 400)
			}
		default:
			return "", nil, errors.NewCCError("asset property must be an integer", 400)
		}

		retVal := (RentStatus)(dataValue)
		err := retVal.CheckType()

		return fmt.Sprint(retVal), retVal, err
	},
}

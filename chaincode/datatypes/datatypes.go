package datatypes

import (
	"github.com/goledgerdev/cc-tools/assets"
)

// CustomDataTypes contain the user-defined primary data types
var CustomDataTypes = map[string]assets.DataType{
	"cpf":      cpf,
	"cnpj":	 	cnpj,
	"vehicleType": vehicleType,
	"plate":    plate,
}

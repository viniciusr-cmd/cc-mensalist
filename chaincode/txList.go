package main

import (
	txdefs "github.com/goledgerdev/cc-mensalist/chaincode/txdefs"

	tx "github.com/goledgerdev/cc-tools/transactions"
)

var txList = []tx.Transaction{
	tx.CreateAsset,
	tx.UpdateAsset,
	tx.DeleteAsset,
	txdefs.RentVehicle,
	txdefs.EndRent,
	txdefs.CreateNewAgency,
	txdefs.GetNumberOfVehiclesFromAgency,
	// txdefs.GetBooksByAuthor, TODO: GET VEHICLES BY RENTER
}

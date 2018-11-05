/*
   Medium Chaincode
   Enrique Macias
   macienrique@hotmail.com
*/

package main

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

type product struct {
	Id     string `json:"Id"`
	Name   string `json:"Name"`
	Color  string `json:"Color"`
	Length string `json:"Length"`
	Width  string `json:"Width"`
}

type productPrice struct {
	Id        string  `json:"Id"`
	BuyPrice  float64 `json:"BuyPrice"`
	SellPrice float64 `json:"SellPrice"`
}

type MediumChaincode struct { // define to implement CC interface
}

func main() {

	err := shim.Start(new(MediumChaincode))

	if err != nil {
		fmt.Printf("Error starting the Medium Contract: %s", err)
	}

}

// Implement Barebone Init
func (t *MediumChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {

	fmt.Println("Successfully init chaincode")

	return shim.Success(nil)

}

// Implement Invoke
func (t *MediumChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {

	fmt.Println("Start Invoke")
	defer fmt.Println("Stop Invoke")

	// Get function name and args
	function, args := stub.GetFunctionAndParameters()

	switch function {

	case "createProduct":
		return t.createProduct(stub, args)
	case "getProduct":
		return t.getProduct(stub, args)
	case "getProductPrice":
		return t.getProductPrice(stub, args)
	default:
		return shim.Error("Invalid invoke function name.")
	}

}

func (t *MediumChaincode) createProduct(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	id := args[0]
	name := args[1]
	color := args[2]
	length := args[3]
	width := args[4]
	buyPrice, err1 := strconv.ParseFloat(args[5], 32)
	sellPrice, err2 := strconv.ParseFloat(args[6], 32)

	if err1 != nil || err2 != nil {
		return shim.Error("Error parsing the values")
	}

	product := &product{id, name, color, length, width}
	productBytes, err3 := json.Marshal(product)

	if err3 != nil {
		return shim.Error(err1.Error())
	}

	productPrice := &productPrice{id, buyPrice, sellPrice}
	productPriceBytes, err4 := json.Marshal(productPrice)

	if err4 != nil {
		return shim.Error(err2.Error())
	}

	err5 := stub.PutPrivateData("collectionMedium", id, productBytes)

	if err5 != nil {
		return shim.Error(err5.Error())
	}

	err6 := stub.PutPrivateData("collectionPrivate", id, productPriceBytes)

	if err6 != nil {
		return shim.Error(err6.Error())
	}

	jsonProduct, err7 := json.Marshal(product)
	if err7 != nil {
		return shim.Error(err7.Error())
	}

	return shim.Success(jsonProduct)

}

func (t *MediumChaincode) getProduct(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	id := args[0]
	product := product{}

	productBytes, err1 := stub.GetPrivateData("collectionMedium", id)
	if err1 != nil {
		return shim.Error(err1.Error())
	}

	err2 := json.Unmarshal(productBytes, &product)

	if err2 != nil {
		fmt.Println("Error unmarshalling object with id: " + id)
		return shim.Error(err2.Error())
	}

	jsonProduct, err3 := json.Marshal(product)
	if err3 != nil {
		return shim.Error(err3.Error())
	}

	return shim.Success(jsonProduct)

}

func (t *MediumChaincode) getProductPrice(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	id := args[0]
	productPrice := productPrice{}

	productPriceBytes, err1 := stub.GetPrivateData("collectionPrivate", id)
	if err1 != nil {
		return shim.Error(err1.Error())
	}

	err2 := json.Unmarshal(productPriceBytes, &productPrice)

	if err2 != nil {
		fmt.Println("Error unmarshalling object with id: " + id)
		return shim.Error(err2.Error())
	}

	jsonProductPrice, err3 := json.Marshal(productPrice)
	if err3 != nil {
		return shim.Error(err3.Error())
	}

	return shim.Success(jsonProductPrice)

}

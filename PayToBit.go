
// 

package main

import (
	"fmt"
	"strconv"
	"encoding/json"

	"github.com/hyperledger/fabric/common/util"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

type PayToBitChaincode struct {
}


var (
	sellerList map[string]bool
)

// Called when first deploy the code
func (t *PayToBitChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("PayToBit Init")
	_, args := stub.GetFunctionAndParameters()
	var cashAddr, bitAddr string    // neutral account: bitcoin address and another payment address 
	var err error
	sellerList = make(map[string]bool)

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	cashAddr = args[0]

	bitAddr = args[1]

	publicInfo := "{\"cashAddr\":\"" + cashAddr + "\", \"totalCash\":0 , \"bitAddr\":\"" + bitAddr + "\", \"bitAddr\": 0}"
	fmt.Println(publicInfo)

	// Write the state to the ledger
	err = stub.PutState("publicInfo", []byte(publicInfo))
	if err != nil {
		return shim.Error("Put State Error.")
	}

	return shim.Success(nil)
}

func (t *PayToBitChaincode) applyForSell(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var dat map[string]string

	publicInfoBytes, err := stub.GetState("publicInfo")
	if err := json.Unmarshal(publicInfoBytes, &dat); err != nil {
        return shim.Error("Parsing json error.")
    }
    //fmt.Println(dat)

	jsonResp := "{\"addr\":\"" + dat["bitAddr"] + "\"}"
	fmt.Printf("ApplyForSell Response:%s\n", jsonResp)
	return shim.Success([]byte(jsonResp))
}

func (t *PayToBitChaincode) bundingCoin(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var recvAddr, totalCoin, txID, transHash string
	var err error
	recvAddr = args[0]
	totalCoin = args[1]
	// transHash = args[2]  // the transaction ID which seller paid to the public address, chaincode will check the transaction
	txID = util.GenerateUUID() // generate a serial number randomly, call from fabric util

	// check if the transaction exist
	// res = lib.CheckTx(transHash)

	seller := "{\"recvAddr\":\"" + recvAddr + "\", \"totalCoin\":\"" + totalCoin + "\", \"txID\":\"" + txID + "\"}"
	fmt.Println(seller)

	sellerList[txID] = true
	err = stub.PutState(txID, []byte(seller))
	if err != nil {
		return shim.Error("Put State Error.")
	}
	return shim.Success([]byte(seller))
}

// Deletes an entity from state
func (t *PayToBitChaincode) revokeTx(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	txID := args[0]
	sellerList[txID] = false

	// Delete the key from the state in ledger
	err := stub.DelState(txID)
	if err != nil {
		jsonResp := "{\"status\":\"Failed to delete transition " + txID + ", \"error " + err.Error() + ",\"}"
		return shim.Error(jsonResp)
	}

	jsonResp := "{\"status\":\"transition " + txID + " has been canceled.\"}"
	return shim.Success([]byte(jsonResp))
}

func (t *PayToBitChaincode) getSellingList(stub shim.ChaincodeStubInterface,  args []string) pb.Response {
	var res []string
	threshold, _ := strconv.Atoi(args[0])
	
	i := 0
	for seller, status := range sellerList{
		
		if status {
			res = append(res, seller)
		}
		i++
		if i >= threshold {  // return the first #threshold seller information
			break
		}
	}

	jsonResp, err := json.Marshal(res)
	if err != nil {
		return shim.Error("Parsing json error.")
	}
	return shim.Success([]byte(jsonResp))
}

func (t *PayToBitChaincode) getTxByID(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	txID := args[0]
	txInfoBytes, err := stub.GetState("txID")
	txInfo := string(txInfoBytes)
	if err != nil {
		return shim.Error("Get State Error.")
	}
	return shim.Success([]byte(txInfo))
}

/**** wait for implementation 

func (t *PayToBitChaincode) submitPaymentProof(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	
}

*/

func (t *PayToBitChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	function, args := stub.GetFunctionAndParameters()
	if function == "applyForSell" {
		return t.applyForSell(stub, args)
	} else if function == "bundingCoin" {
		return t.bundingCoin(stub, args)
	}else if function == "revokeTx" {
		return t.revokeTx(stub, args)
	}else if function == "bundingCoin" {
		return t.bundingCoin(stub, args)
	}else if function == "getSellingList" {
		return t.getSellingList(stub, args)
	}else if function == "getTxByID" {
		return t.getTxByID(stub, args)
	}

	return shim.Error("Invalid invoke function name.")
}

func main(){
	err := shim.Start(new(PayToBitChaincode))
	if err != nil {
		fmt.Printf("Error starting PayToBit chaincode: %s", err)
	}
}

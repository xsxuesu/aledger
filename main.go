package main

import (
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
	"ledger/common"
	"ledger/services"
	"strings"
)
type LedgerChainCode struct {

}

func(lcc *LedgerChainCode)Init(stub shim.ChaincodeStubInterface)pb.Response{
	return shim.Success(nil)
}

func(lcc *LedgerChainCode)Invoke(stub shim.ChaincodeStubInterface)pb.Response{
	funcname,_ := stub.GetFunctionAndParameters()

	switch  strings.ToLower(funcname)  {
	case "version":
		return shim.Success([]byte(common.VERSION))
	case "account_check":
		return services.AccountCheck(stub)
	case "account_register":
		return services.AccountConfirm(stub)
	case "account_generate":
		return services.GeneratePriKey(stub)
	case "account_lock":
		return services.AccountLock(stub)
	case "account_unlock":
		return services.AccountUNLock(stub)
	case "account_get":
		return services.AccountGet(stub)
	case "account_changepri":
		return services.AccountChangePri(stub)
	case "account_list":
		return services.AccountGetAll(stub)
	case "token_create":
		return services.TokenCreate(stub)
	case "token_lock":
		return services.TokenUpdateDisable(stub)
	case "token_unlock":
		return services.TokenUpdateEnable(stub)
	case "token_history":
		return services.TokenGetHistory(stub)
	case "token_list":
		return services.TokenList(stub)
	case "token_get":
		return services.TokenGetByName(stub)
	case "issue":
		return services.LedgerIssue(stub)
	case "transfer":
		return services.LedgerTransfer(stub)
	case "balance":
		return services.LedgerGetBalance(stub)
	case "history":
		return services.LedgerGetHistory(stub)
	case "burn":
		return services.LedgerBurnBalance(stub)
	case "scale":
		return services.LedgerScale(stub)
	case "holdtoken":
		return services.LedgerGetListbyAccount(stub)
	case "signreq":
		return services.SignRequest(stub)
	case "signget":
		return services.SignGetRequest(stub)
	case "signhistory":
		return services.SignHistory(stub)
	case "signresp":
		return services.SignRepsonse(stub)
	default:
		return shim.Error("function not define,please check function")
	}

	return shim.Success(nil)
}

func main(){
	err := shim.Start(new(LedgerChainCode))
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err.Error())
	}
}

package main

import (
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rpc"
	"reflect"
	"os"
	"io/ioutil"
)

const tokenAddress = "0xf4a97a23cd66e7b8bf788d6d6eb4abb4e3b42caf"

////////////////////////////////////////////
//
// base eth rpc data structs
//
////////////////////////////////////////////
type CallArgs struct {
	From 		string
	To   		string
	Gas  		hexutil.Big
	GasPrice 	hexutil.Big
	Value 		hexutil.Big
	Data 		interface{}
}

type Transaction struct {
	From		string
	To 			string
	Gas			hexutil.Big
	GasPrice	hexutil.Big
	Value       hexutil.Big
	Data		string
}

////////////////////////////////////////////
//
// token data structs
//
////////////////////////////////////////////
type Deposit struct {
	hash 		string
	account 	string
	amount 		int
}

type Order struct {
	hash		string
	accountS 	string
	accountB 	string
	amountS		int
	amountB		int
}

type OrderState struct {
	hash		string
	accountS 	string
	accountB 	string
	amountS 	int
	amountB 	int
	ok 			bool
}

////////////////////////////////////////////
//
// token events
//
////////////////////////////////////////////
type DepositEvent struct {
	hash 		string
	account     string
	amount 		int
	ok 			bool
}

type OrderEvent struct {
	hash 		string
	accountS 	string
	accountB 	string
	amountS 	int
	amountB 	int
	ok 			bool
}

type BankToken struct {
	Transfer 		AbiMethod	`methodName:"submitTransfer"`
	Deposit			AbiMethod	`methodName:"submitDeposit"`
	BalanceOf		AbiMethod	`methodName:"balanceOf"`
}

type AbiMethod struct {
	abi.Method
	Abi 		*abi.ABI
	Address 	string
}

var client 		*rpc.Client

func DialChain() {
	var err error
	client, err = rpc.Dial("http://127.0.0.1:8545")
	if err != nil {
		panic(err)
	}
}

func LoadContract() *BankToken {
	tabi := getAbi()

	bank := &BankToken{}
	elem := reflect.ValueOf(bank).Elem()

	for i:=0; i<elem.NumField(); i++ {
		methodName := elem.Type().Field(i).Tag.Get("methodName")

		abiMethod := new(AbiMethod)
		abiMethod.Name = methodName
		abiMethod.Abi = tabi
		abiMethod.Address = tokenAddress

		elem.Field(i).Set(reflect.ValueOf(*abiMethod))
	}

	return bank
}

func getAbi() *abi.ABI {
	tabi := &abi.ABI{}

	dir, _ := os.Getwd()
	abiStr,err := ioutil.ReadFile(dir + "/contracts/transfer/abi.txt")
	if err != nil {
		panic(err)
	}

	if err := tabi.UnmarshalJSON(abiStr); err != nil {
		panic(err)
	}

	return tabi
}

func (method *AbiMethod) Call(result interface{}, tag string, args ...interface{}) error {
	bytes, err := method.Abi.Pack(method.Name, args...)
	if err != nil {
		return err
	}

	c := &CallArgs{}
	c.From = method.Address
	c.To = method.Address
	c.Data = common.ToHex(bytes)

	return client.Call(result, "eth_call", c, tag)
}

func (method *AbiMethod) SendTransaction() (string, error) {

	return "",nil
}
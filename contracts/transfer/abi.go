package main

/*

package eth

import (
"github.com/ethereum/go-ethereum/accounts/abi"
"reflect"
"strings"
"github.com/ethereum/go-ethereum/common"
"github.com/Loopring/ringminer/chainclient"
"math/big"
types "github.com/Loopring/ringminer/types"
ethTypes "github.com/ethereum/go-ethereum/core/types"
"github.com/ethereum/go-ethereum/common/hexutil"
)

//合约的数据结构相关的，其余的不在此处
const Erc20TokenAbiStr = `[{"constant":false,"inputs":[{"name":"_spender","type":"address"},{"name":"_value","type":"uint256"}],"name":"approve","outputs":[{"name":"success","type":"bool"}],"payable":false,"type":"function"},{"constant":true,"inputs":[],"name":"totalSupply","outputs":[{"name":"totalSupply","type":"uint256"}],"payable":false,"type":"function"},{"constant":false,"inputs":[{"name":"_from","type":"address"},{"name":"_to","type":"address"},{"name":"_value","type":"uint256"}],"name":"transferFrom","outputs":[{"name":"success","type":"bool"}],"payable":false,"type":"function"},{"constant":true,"inputs":[{"name":"_owner","type":"address"}],"name":"balanceOf","outputs":[{"name":"balance","type":"uint256"}],"payable":false,"type":"function"},{"constant":false,"inputs":[{"name":"_to","type":"address"},{"name":"_value","type":"uint256"}],"name":"transfer","outputs":[{"name":"success","type":"bool"}],"payable":false,"type":"function"},{"constant":true,"inputs":[{"name":"_owner","type":"address"},{"name":"_spender","type":"address"}],"name":"allowance","outputs":[{"name":"remaining","type":"uint256"}],"payable":false,"type":"function"},{"anonymous":false,"inputs":[{"indexed":true,"name":"_from","type":"address"},{"indexed":true,"name":"_to","type":"address"},{"indexed":false,"name":"_value","type":"uint256"}],"name":"Transfer","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"name":"_owner","type":"address"},{"indexed":true,"name":"_spender","type":"address"},{"indexed":false,"name":"_value","type":"uint256"}],"name":"Approval","type":"event"}]`

type Contract struct {
	Abi *abi.ABI
	Address     string
}

func (c *Contract) GetAbi() interface{} {
	return *c.Abi
}

func (c *Contract) GetAddress() string {
	return c.Address
}

type AbiMethod struct {
	Abi *abi.ABI
	Address	string
	abi.Method
}

func (m *AbiMethod) Call(result interface{}, blockParameter string, args ...interface{}) error {
	dataBytes, err := m.Abi.Pack(m.Name, args...)
	if (nil != err) {
		return err
	}
	data := common.ToHex(dataBytes)
	//when call a contract method，gas,gasPrice and value are not needed.
	arg := &CallArg{}
	arg.From = m.Address	//设置地址，因为rpc.Client.Call不仅给eth_call使用，所以要求地址
	arg.To = m.Address
	arg.Data = data
	return EthClient.Call(result, arg, blockParameter)
}

//contract transaction
func (m *AbiMethod) SendTransaction(from string, gas, gasPrice *big.Int, args ...interface{}) (string,error) {
	dataBytes, err := m.Abi.Pack(m.Name, args...)

	if (nil != err) {
		return "", err
	}

	if nil == gasPrice || gasPrice.Cmp(big.NewInt(0)) <= 0 {
		var gasPriceRes types.HexNumber
		if err = EthClient.GasPrice(&gasPriceRes); nil != err {
			return "", err
		} else {
			gasPrice = gasPriceRes.BigInt()
		}
	}

	if nil == gas || gas.Cmp(big.NewInt(0)) <= 0 {
		var gasRes types.HexNumber
		dataHex := common.ToHex(dataBytes)
		callArg := &CallArg{}
		callArg.From = from
		callArg.To = m.Address
		callArg.Data = dataHex
		callArg.GasPrice = hexutil.Big(*gasPrice)
		if err = EthClient.EstimateGas(&gasRes, callArg); nil != err {
			return "", err
		} else {
			gas = gasRes.BigInt()
		}
	}

	var nonce types.HexNumber
	if err = EthClient.GetTransactionCount(&nonce, from, "pending"); nil != err {
		return "", err
	}

	transaction := ethTypes.NewTransaction(nonce.Uint64(),
		common.HexToAddress(m.Address),
		big.NewInt(0),
		gas,
		gasPrice,
		dataBytes)
	var txHash string

	err = EthClient.SignAndSendTransaction(&txHash, from, transaction)
	return txHash, err
}

func applyAbiMethod(token *chainclient.Erc20Token) {
	e := reflect.ValueOf(token).Elem()
	abi := token.GetAbi().(abi.ABI)

	for _, method := range abi.Methods {
		methodName := strings.ToUpper(method.Name[0:1]) + method.Name[1:]
		abiMethod := &AbiMethod{}
		abiMethod.Name = method.Name
		abiMethod.Abi = &abi
		abiMethod.Address = token.GetAddress()
		e.FieldByName(methodName).Set(reflect.ValueOf(abiMethod))
	}
}

func NewErc20Token(address string) *chainclient.Erc20Token {
	erc20Token := &chainclient.Erc20Token{}

	contract := &Contract{}
	cabi := &abi.ABI{}
	cabi.UnmarshalJSON([]byte(Erc20TokenAbiStr))
	contract.Abi = cabi
	contract.Address = address

	erc20Token.Contract = contract
	applyAbiMethod(erc20Token)

	return erc20Token
}

*/

import "github.com/ethereum/go-ethereum/common/hexutil"

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

}

//
//type AbiMethod interface {
//	Call(result interface{}, methodName string, args ...interface{}) error
//	SendTransaction(result interface{}, methodName string, tx interface{}, blockParameters ...interface{}) (string, error)
//}

type AbiMethod struct {

}

func (method *AbiMethod) Call(result interface{}, methodName string, args ...interface{}) error {
	return nil
}
package types

import (
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/ethereum/go-ethereum/common"
)

// 发送transaction
type Transaction struct {
	From		string
	To 			string
	Gas			hexutil.Big
	GasPrice	hexutil.Big
	Value       hexutil.Big
	Data		string
}

// 使用call方法
type CallArgs struct {
	From 		string
	To   		string
	Gas  		hexutil.Big
	GasPrice 	hexutil.Big
	Value 		hexutil.Big
	Data 		interface{}
}

// 使用newFilter方法生成filter
type FilterReq struct {
	FromBlock string
	ToBlock string
	Address string
	Topics []string
}

// 监听某个filter得到的log数据结构
type FilterLog struct {
	LogIndex 			HexNumber 		`json:"logIndex"`
	BlockNumber 		HexNumber 		`json:"blockNumber"`
	BlockHash 			string 			`json:"blockHash"`
	TransactionHash 	string 			`json:"transactionHash"`
	TransactionIndex 	HexNumber 		`json:"transactionIndex"`
	Address 			string 			`json:"address"`
	Data 				string 			`json:"data"`
	Topics 				[]string 		`json:"topics"`
}

type AbiMethod struct {
	abi.Method
	Abi 		*abi.ABI
	Address 	string
	Client		*rpc.Client
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

	return method.Client.Call(result, "eth_call", c, tag)
}

// sendTransaction是不需要tag的
func (method *AbiMethod) SendTransaction(miner, tokenAddress string,gas, gasPrice int, result interface{}, args ...interface{}) error {
	bytes, err := method.Abi.Pack(method.Name, args...)
	if err != nil {
		return err
	}

	tx := &Transaction{}
	tx.From = miner
	tx.To = tokenAddress
	tx.Gas = Int2HexBigInt(gas)
	tx.GasPrice = Int2HexBigInt(gasPrice)
	tx.Data = common.ToHex(bytes)

	return method.Client.Call(result, "eth_sendTransaction", tx)
}
package abi

import (
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	cm "github.com/dylenfu/eth-libs/common"
	"github.com/ethereum/go-ethereum/rpc"
	"reflect"
	"os"
	"io/ioutil"
	"github.com/dylenfu/eth-libs/types"
	"log"
	//"bytes"
	"github.com/pkg/errors"
)

const (
	miner = "0x4bad3053d574cd54513babe21db3f09bea1d387d"
	tokenAddress = "0x227F88083AE9eE717e39669CB2718E604833fEf9"
	filterId = "0x6b01393d6914ea5118b8ebb123eb07e"
)

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

var (
	client 		*rpc.Client
	tabi 		*abi.ABI
	contract    *BankToken
)

func init() {
	var err error
	client, err = rpc.Dial("http://127.0.0.1:8545")
	if err != nil {
		panic(err)
	}

	filters = make(map[string]string)
	tabi = getAbi()
}

func LoadContract() *BankToken {
	bank := &BankToken{}
	elem := reflect.ValueOf(bank).Elem()

	for i:=0; i < elem.NumField(); i++ {
		methodName := elem.Type().Field(i).Tag.Get("methodName")

		abiMethod := &AbiMethod{}
		abiMethod.Name = methodName
		abiMethod.Abi = tabi
		abiMethod.Address = tokenAddress

		elem.Field(i).Set(reflect.ValueOf(*abiMethod))
	}

	contract = bank

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

type CallArgs struct {
	From 		string
	To   		string
	Gas  		hexutil.Big
	GasPrice 	hexutil.Big
	Value 		hexutil.Big
	Data 		interface{}
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

type Transaction struct {
	From		string
	To 			string
	Gas			hexutil.Big
	GasPrice	hexutil.Big
	Value       hexutil.Big
	Data		string
}

// sendTransaction是不需要tag的
func (method *AbiMethod) SendTransaction(result interface{}, args ...interface{}) error {
	bytes, err := method.Abi.Pack(method.Name, args...)
	if err != nil {
		return err
	}

	tx := &Transaction{}
	tx.From = miner
	tx.To = tokenAddress
	tx.Gas = cm.ToHexBigInt(1200000)
	tx.GasPrice = cm.ToHexBigInt(1)
	tx.Data = common.ToHex(bytes)

	return client.Call(result, "eth_sendTransaction", tx)
}

type FilterReq struct {
	FromBlock string
	ToBlock string
	Address string
	Topics []string
}

var filters map[string]string

// 这里要注意 filterId会变更,
func NewFilter(topic string) error {
	var filterId string

	filter := FilterReq{}
	filter.FromBlock = "latest"
	filter.ToBlock = "latest"
	filter.Address = tokenAddress

	err := client.Call(&filterId, "eth_newFilter", &filter)
	if err != nil {
		return err
	}

	filters[topic] = filterId

	return nil
}

type FilterLog struct {
	LogIndex types.HexNumber `json:"logIndex"`
	BlockNumber types.HexNumber `json:"blockNumber"`
	BlockHash string `json:"blockHash"`
	TransactionHash string `json:"transactionHash"`
	TransactionIndex types.HexNumber `json:"transactionIndex"`
	Address string `json:"address"`
	Data string `json:"data"`
	Topics []string `json:"topics"`
}

type DepositEvent struct {
	hash 		string
	account     string
	amount 		int
	ok 			bool
}

func FilterChanged() error {
	//var logs []FilterLog

	//err := client.Call(&logs, "eth_getFilterChanges", filterId)
	//if err != nil {
	//	return err
	//}

	event := &DepositEvent{}
	name := "DepositFilled"//contract.Deposit.Name

	data := "0x69be7bc7c7c6e216dd9531c88c94769f9f63ce53f47665b5ec7faf55f8094e8100000000000000000000000037303138303536313763303331396139623464640000000000000000000000000000000000000000000000000000000005f5e1000000000000000000000000000000000000000000000000000000000000000001"

	t := hexutil.MustDecode(data)
	abiEvent, ok := tabi.Events[name]
	if !ok {
		err := errors.New("event do not exist")
		panic(err)
		return err
	}

	d := []byte(t)//buff.Bytes()
	if err := cm.UnpackEvent(abiEvent, event, d); err != nil {
		panic(err)
		return err
	} else {
		//log.Println("amount", event.amount)
		//log.Println("account", event.account)
		//log.Println("hash", event.hash)
		//log.Println("isOk", event.ok)
	}

	//log.Println(event.account)
	//log.Println(event.amount)
	//log.Println(event.hash)
	//log.Println(event.ok)

	//for _, v := range logs {
	//	log.Println(v.Data)
	//	bytesArr := get32Bytes([]byte(v.Data))
	//	buff := new(bytes.Buffer)
	//	for _, v := range bytesArr {
	//		buff.Write(common.Hex2Bytes(string(v)))
	//	}
	//
	//	if err := tabi.Unpack(event, name, buff.Bytes()); err != nil {
	//		return err
	//	} else {
	//		log.Println("amount", event.amount)
	//		log.Println("account", event.account)
	//		log.Println("hash", event.hash)
	//		log.Println("isOk", event.ok)
	//	}
	//}

	return nil
}

const step = 64

func get32Bytes(src []byte)[]string {
	var ret []string

	length := len(src)

	log.Println(string(src))
	src = src[2:length]
	log.Println(string(src))

	log.Println("before length", length)
	length = len(src)
	log.Println("after length", length)

	cnt := length / step
	if length % step > 0 {
		cnt++
	}

	log.Println("cnt is ", cnt)
	for i := 0; i < cnt; i++ {
		start := i * step
		end := (i + 1) * step
		if end > length {
			end = length
		}
		sub := string(src[start:end])
		log.Println("sub", i, sub)
		ret = append(ret, sub)
	}

	return ret
}
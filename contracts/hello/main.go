package main

import (
	"io/ioutil"
	abi "github.com/ethereum/go-ethereum/accounts/abi"
	"os"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rpc"
	"log"
	"github.com/ethereum/go-ethereum/common/hexutil"
	cm "github.com/dylenfu/eth-libs/common"
	. "github.com/dylenfu/eth-libs/params"
	"flag"
)

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

var testcase = flag.String("t", "greet", "chose testcase")

func main() {
	flag.Parse()

	habi := newABI()

	client, err := rpc.Dial("http://127.0.0.1:8545")
	if err != nil {
		panic(err)
	}

	var result string

	switch *testcase {

	case "greet":
		args1 := greet(habi)
		if err := client.Call(&result, "eth_call", args1, "latest"); err != nil {
			panic(err)
		}
		log.Println(string(common.FromHex(result)))

	case "set":
		tx := setGreet(habi)
		if err := client.Call(&result, "eth_sendTransaction", tx); err != nil {
			panic(err)
		}
		log.Println(result)

	case "get":
		args2 := getGreet(habi)
		if err := client.Call(&result, "eth_call", args2, "latest"); err != nil {
			panic(err)
		}
		log.Println(string(common.FromHex(result)))
	}

}

// newABI 解析在线编辑器生成的abi字符串,生成对应的智能合约ABI
func newABI() *abi.ABI {
	hABI := abi.ABI{}

	dir, _ := os.Getwd()
	abiStr,err := ioutil.ReadFile(dir + "/contracts/hello/abi.txt")
	if err != nil {
		panic(err)
	}

	if err := hABI.UnmarshalJSON(abiStr); err != nil {
		panic(err)
	}

	return &hABI
}

// 生成eth_call需要的参数,这里主要用到abi的Pack函数,将某个函数的相关参数打包成data
// 注意，在不涉及到修改智能合约变量时，不需要sendTransaction，直接call就可以
func greet(habi *abi.ABI) *CallArgs {
	bytes, err := habi.Pack("greet")
	if err != nil {
		panic(err)
	}

	data := common.ToHex(bytes)
	args := &CallArgs{}
	args.commonArgs(data)

	return args
}

// setGreet函数在智能合约里涉及到修改变量，这时我们不能直接使用call
// 而是要使用sendTransaction的方式调用，然后通过挖矿的方式实现变量变更
// > INFO [09-14|14:10:56] Submitted transaction
// > fullhash=0x1e406b4ef846191ebbf27c405aff6f182e4c6ec0f40ed8fbda60c612abb7b792
// > recipient=0xe5131431f134A961C5Cf3941Ef77182aea203196
func setGreet(habi *abi.ABI) *Transaction {
	bytes, err := habi.Pack("setGreeting", "hahhahammmmmmm")
	if err != nil {
		panic(err)
	}

	tx := &Transaction{}
	tx.From = Miner
	tx.To = TokenAddress
	tx.Gas = cm.ToHexBigInt(100000)
	tx.GasPrice = cm.ToHexBigInt(1)
	tx.Data = common.ToHex(bytes)

	return tx
}

func getGreet(habi *abi.ABI) *CallArgs {
	bytes, err := habi.Pack("getGreeting")
	if err != nil {
		panic(err)
	}

	data := common.ToHex(bytes)
	args := &CallArgs{}
	args.commonArgs(data)

	return args
}

func (args *CallArgs)commonArgs(data string) {
	args.From = TokenAddress
	args.To = TokenAddress
	args.Data = data
}
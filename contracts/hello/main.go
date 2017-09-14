package main

import (
	"io/ioutil"
	abi "github.com/ethereum/go-ethereum/accounts/abi"
	"os"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rpc"
	"log"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

const tokenAddress = "0xe5131431f134a961c5cf3941ef77182aea203196"

type CallArgs struct {
	From 		string
	To   		string
	Gas  		hexutil.Big
	GasPrice 	hexutil.Big
	Value 		hexutil.Big
	Data 		interface{}
}

func main() {
	habi := newABI()

	client, err := rpc.Dial("http://127.0.0.1:8545")
	if err != nil {
		panic(err)
	}

	var result string

	args1 := greet(habi)
	if err := client.Call(&result, "eth_call", args1, "latest"); err != nil {
		panic(err)
	}
	log.Println(string(common.FromHex(result)))

	//args2 := setGreet(habi)
	//if err := client.Call(&result, "eth_call", args2, "latest"); err != nil {
	//	panic(err)
	//}
	//log.Println(result)
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
func setGreet(habi *abi.ABI) *CallArgs {
	bytes, err := habi.Pack("setGreeting", "hahhaha")
	if err != nil {
		panic(err)
	}

	data := common.ToHex(bytes)
	args := &CallArgs{}
	args.commonArgs(data)

	return args
}

func (args *CallArgs)commonArgs(data string) {
	args.From = tokenAddress
	args.To = tokenAddress
	args.Data = data
}
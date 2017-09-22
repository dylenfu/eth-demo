package main

import (
	"log"
	"github.com/ethereum/go-ethereum/common"
	"github.com/dylenfu/eth-libs/types"
	"flag"
	. "github.com/dylenfu/eth-libs/contracts/transfer/abi"
	"math/big"
)

// 该项目实验性地创建一个合约，包含订单信息及转账功能
// 实现合约及部署，通过abi调用方式实现转账操作

var testcase = flag.String("call", "balance", "chose test case")

const (
	account1 = "0x46c5683c754b2eba04b2701805617c0319a9b4dd"
	account2 = "0x56d9620237fff8a6c0f98ec6829c137477887ec4"
)

func main() {
	flag.Parse()

	bank := LoadContract()

	go func() {
		switch *testcase {
		case "balance":
			balance(bank)

		case "deposit":
			deposit(bank)

		case "filter":
			filter()
		}
	}()

	listen()
}

func balance(bank *BankToken) {
	var result types.HexNumber

	addr := common.StringToAddress(account2)
	if err := bank.BalanceOf.Call(&result, "latest", addr); err != nil {
		panic(err)
	}

	log.Println(result.Int())
}

func deposit(bank *BankToken) {
	var result string

	hash := common.FromHex("0x69be7bc7c7c6e216dd9531c88c94769f9f63ce53f47665b5ec7faf55f8094e81")
	account := common.StringToAddress(account1)
	// 这里需要注意一定只能用big.NewInt
	amount := big.NewInt(100000000)

	// 这里一定注意，因为合约里的函数参数是一个个传入的，所以这里不能传一个结构过去
	if err := bank.Deposit.SendTransaction(&result, hash, account, amount); err != nil {
		panic(err)
	}
}

func filter() {
	NewFilter("deposit")
}

func listen() {
	for {
		if err := FilterChanged(); err != nil {
			panic(err)
		}
	}
}

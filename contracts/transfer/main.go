package main

import (
	"log"
	"github.com/ethereum/go-ethereum/common"
	"github.com/dylenfu/eth-libs/types"
	"flag"
	. "github.com/dylenfu/eth-libs/contracts/transfer/abi"
	"math/big"
	. "github.com/dylenfu/eth-libs/params"
)

// 该项目实验性地创建一个合约，包含订单信息及转账功能
// 实现合约及部署，通过abi调用方式实现转账操作
// 注：相关账号密码为101,102,103
var testcase = flag.String("call", "deposit", "chose test case")

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

	addr := common.StringToAddress(Account1)
	if err := bank.BalanceOf.Call(&result, "latest", addr); err != nil {
		panic(err)
	}

	log.Println(result.Int())
}

func deposit(bank *BankToken) {
	var result string

	id := common.FromHex("0x5ad6fe3e08ffa01bb1db674ac8e66c47511e364a4500115dd2feb33dad972d7e")
	account := common.HexToAddress(Account2)
	// 这里需要注意一定只能用big.NewInt
	amount := big.NewInt(200000001)

	// 这里一定注意，因为合约里的函数参数是一个个传入的，所以这里不能传一个结构过去
	if err := bank.Deposit.SendTransaction(&result, id, account, amount); err != nil {
		panic(err)
	}
}

func filter() {
	NewFilter("deposit")
}

func listen() {
	for {
		if err := FilterChanged(); err != nil {
			log.Println(err.Error())
		}
	}
}

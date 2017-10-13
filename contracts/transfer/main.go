package main

import (
	"flag"
	. "github.com/dylenfu/eth-libs/contracts/transfer/abi"
	. "github.com/dylenfu/eth-libs/params"
	"github.com/dylenfu/eth-libs/types"
	"github.com/ethereum/go-ethereum/common"
	"log"
	"math/big"
)

// 该项目实验性地创建一个合约，包含订单信息及转账功能
// 实现合约及部署，通过abi调用方式实现转账操作
// 注：相关账号密码为101,102,103
var testcase = flag.String("call", "deposit", "chose test case")

func main() {
	flag.Parse()

	//v := reflect.ValueOf(*testcase)
	bank := LoadContract()

	switch *testcase {
	case "balance":
		balance(bank)

	case "deposit":
		deposit(bank)

	case "transfer":
		transfer(bank)

	case "listenEvents":
		listenEvents()

	case "listenBlock":
		listenBlock()
	}
}

func balance(bank *BankToken) {
	var result types.HexNumber

	addr := types.Str2Address(Account2)
	if err := bank.BalanceOf.Call(&result, "latest", addr); err != nil {
		panic(err)
	}

	log.Println(result.Int())
}

func deposit(bank *BankToken) {
	var result string

	id := common.FromHex("0x0000000000000000000000000000000000000000000000000000000000000001")

	account := types.Str2Address(Account1)

	// 这里需要注意一定只能用big.NewInt
	amount := big.NewInt(500000000)

	// 这里一定注意，因为合约里的函数参数是一个个传入的，所以这里不能传一个结构过去
	err := bank.Deposit.SendTransaction(Miner, TransferTokenAddress, 1200000, 1, &result, id, account, amount)
	if err != nil {
		panic(err)
	}
}

// 账户1多2少
func transfer(bank *BankToken) {
	var result string
	account1 := types.Str2Address(Account1)
	account2 := types.Str2Address(Account2)

	id := common.FromHex("0x0000000000000000000000000000000000000000000000000000000000000002")
	amount := big.NewInt(100000000)
	err := bank.Transfer.SendTransaction(Miner, TransferTokenAddress, 1200000, 1, &result, id, account2, account1, amount, amount)
	if err != nil {
		panic(err)
	}
}

func listenEvents() {
	filterId, err := NewFilter(BlockNumber)
	if err != nil {
		panic(err)
	}

	for {
		if err := EventChanged(filterId); err != nil {
			log.Println(err.Error())
		}
	}
}

func listenBlock() {
	filterId, err := BlockFilterId()
	if err != nil {
		panic(err)
	}

	for {
		if err := BlockChanged(filterId); err != nil {
			log.Println(err.Error())
		}
	}
}

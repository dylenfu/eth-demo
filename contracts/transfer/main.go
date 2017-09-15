package main

import (
	"log"
	"github.com/ethereum/go-ethereum/common"
	"github.com/dylenfu/eth-libs/types"
)

// 该项目实验性地创建一个合约，包含订单信息及转账功能
// 实现合约及部署，通过abi调用方式实现转账操作
func main() {

	DialChain()
	bank := LoadContract()

	var result types.HexNumber

	addr := common.StringToAddress("1111")
	if err := bank.BalanceOf.Call(&result, "latest", addr); err != nil {
		panic(err)
	}

	log.Println(result.Int())
}

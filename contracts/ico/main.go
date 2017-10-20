package main

import (
	"flag"
	"reflect"
	. "github.com/dylenfu/eth-libs/contracts/ico/contract"
	"github.com/dylenfu/eth-libs/types"
	. "github.com/dylenfu/eth-libs/params"
	"log"
	"math/big"
	"github.com/ethereum/go-ethereum/common"
)

var (
	tokenA = TokenA.Token.(*IcoImpl)
	tokenB = TokenB.Token.(*IcoImpl)
	handle = &Handle{}
	fn = flag.String("fn", "deposit", "chose function")
)

type Handle struct{}

func main() {
	flag.Parse()
	reflect.ValueOf(handle).MethodByName(*fn).Call([]reflect.Value{})
}

func (h *Handle) BalanceOf() {
	var result types.HexNumber

	account := types.Str2Address(Account2)
	err := tokenA.BalanceOf.Call(&result, "latest", account)
	if err != nil {
		panic(err)
	}

	log.Println(result)
}

func (h *Handle) Deposit() {
	var result types.HexNumber
	value := big.NewInt(100000000)

	account := types.Str2Address(Account1)
	err := tokenA.Deposit.SendTransaction(TokenAddressA ,&result, account, value)
	if err != nil {
		panic(err)
	}

	log.Println(result)
}

func (h *Handle) Transfer() {
	var result string

	// 合约里面msg.sender是看谁对transaction签名过即为sender
	to := common.HexToAddress(Account1)
	amount := big.NewInt(100000000)
	if err := tokenA.Transfer.SignAndSendTransaction(TokenAddressA, Account2, &result, to, amount); err != nil {
		panic(err)
	}

	log.Println(result)
}
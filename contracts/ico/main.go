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

	account := types.Str2Address(Account1)
	err := tokenA.BalanceOf.Call(&result, "latest", account)
	if err != nil {
		panic(err)
	}

	log.Println(result)
}

/*
	"0x48ff2269e58a373120FFdBBdEE3FBceA854AC30A":"07ae9ee56203d29171ce3de536d7742e0af4df5b7f62d298a0445d11e466bf9e",
	"0xb5fab0b11776aad5ce60588c16bd59dcfd61a1c2":"11293da8fdfe3898eae7637e429e7e93d17d0d8293a4d1b58819ac0ca102b446",
	*/
func (h *Handle) Deposit() {
	var result types.HexNumber
	value := big.NewInt(1000000000000)

	account := types.Str2Address("0xb5fab0b11776aad5ce60588c16bd59dcfd61a1c2")
	err := tokenB.Deposit.SendTransaction(&result, account, value)
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
	if err := tokenA.Transfer.SignAndSendTransaction(Account2, &result, to, amount); err != nil {
		panic(err)
	}

	log.Println(result)
}
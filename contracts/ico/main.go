package main

import (
	"flag"
	. "github.com/dylenfu/eth-libs/contracts/ico/contract"
	. "github.com/dylenfu/eth-libs/params"
	"github.com/dylenfu/eth-libs/types"
	"github.com/ethereum/go-ethereum/common"
	"log"
	"math/big"
	"reflect"
)

var (
	tokenA = TokenA.Token.(*IcoImpl)
	tokenB = TokenB.Token.(*IcoImpl)
	handle = &Handle{}
	fn     = flag.String("fn", "deposit", "chose function")
)

type Handle struct{}

func main() {
	flag.Parse()
	reflect.ValueOf(handle).MethodByName(*fn).Call([]reflect.Value{})
}

func (h *Handle) BalanceOf() {
	var result types.HexNumber

	account := common.HexToAddress("0xb1018949b241d76a1ab2094f473e9befeabb5ead")
	err := tokenB.BalanceOf.Call(&result, "latest", account)
	if err != nil {
		panic(err)
	}

	log.Println(result.BigInt().String())
}

func (h *Handle) GetBalanceFromMap() {
	var result types.HexNumber

	account := types.Str2Address(Account1)
	if err := tokenA.Balances.Call(&result, "latest", account); err != nil {
		panic(err)
	}
	log.Println(result.Int64())
}

/*
normal user
"0x48ff2269e58a373120FFdBBdEE3FBceA854AC30A"
"0xb5fab0b11776aad5ce60588c16bd59dcfd61a1c2"

loopring test accounts
"0x1b978a1d302335a6f2ebe4b8823b5e17c3c84135"
"0xb1018949b241d76a1ab2094f473e9befeabb5ead"
*/
func (h *Handle) Deposit() {
	var result types.HexNumber
	value, _ := new(big.Int).SetString("20223456789000000000000000000", 0)

	account := types.Str2Address("0xb1018949b241d76a1ab2094f473e9befeabb5ead")
	err := tokenB.Deposit.SendTransaction(&result, account, value)
	if err != nil {
		panic(err)
	}

	log.Println(result)
}

func (h *Handle) Approve() {
	var result bool

	// 合约里面msg.sender是看谁对transaction签名过即为sender
	to := common.HexToAddress(Account1)
	amount := big.NewInt(100000000)
	if err := tokenA.Transfer.SignAndSendTransaction(Account2, &result, to, amount); err != nil {
		panic(err)
	}

	log.Println(result)
}

func (h *Handle) Allowance() {
	var result types.HexNumber

	owner := types.Str2Address(Account1)
	spender := types.Str2Address(Account2)
	err := tokenA.Allowance.Call(&result, "latest", owner, spender)
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

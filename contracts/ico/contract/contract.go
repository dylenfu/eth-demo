package contract

import (
	"github.com/dylenfu/eth-libs/types"
	"github.com/ethereum/go-ethereum/common"
)

const (
	AbiFilePath   = "github.com/dylenfu/eth-libs/contracts/ico/abi.txt"
	EthRpcUrl     = "http://127.0.0.1:8545"
	TokenAddressA = "0xfc2cbce778ddbc4d50bb5b2fc91afe14a8e3953d"
	TokenAddressB = "0x876c8b6ff4a8e87dc6d5e3f64715b58be7d5ab55"
)

var (
	TokenA *types.TokenImpl
	TokenB *types.TokenImpl
)

func init() {
	tokenA := &IcoImpl{}
	tokenB := &IcoImpl{}
	TokenA = types.NewContract(AbiFilePath, TokenAddressA, EthRpcUrl, tokenA)
	TokenB = types.NewContract(AbiFilePath, TokenAddressB, EthRpcUrl, tokenB)
}

type IcoImpl struct {
	Deposit   types.AbiMethod `methodName:"deposit"`
	Transfer  types.AbiMethod `methodName:"transfer"`
	BalanceOf types.AbiMethod `methodName:"balanceOf"`
	Balances  types.AbiMethod `methodName:"balances"`
	Allowance types.AbiMethod `methodName:"allowance"`
	Approve   types.AbiMethod `methodName:"approve"`
	Sign      types.AbiMethod `methodName:"eth_sign"`
}

type Transfer struct {
	From  common.Address `alias:"_from"`
	To    common.Address `alias:"_to"`
	Value common.Address `alias:"_value"`
}

type Approval struct {
	Owner   common.Address `alias:"_owner"`
	Spender common.Address `alias:"_spender"`
	Value   common.Address `alias:"_value"`
}

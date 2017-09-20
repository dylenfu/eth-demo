package abi_test

import (
	"testing"
	"github.com/ethereum/go-ethereum/common/hexutil"
	cm "github.com/dylenfu/eth-libs/common"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"os"
	"io/ioutil"
)

func TestNewFilter(t *testing.T) {
	tabi := newAbi()
	event := &depositEvent{}
	name := "DepositFilled"
	str := "0x69be7bc7c7c6e216dd9531c88c94769f9f63ce53f47665b5ec7faf55f8094e8100000000000000000000000037303138303536313763303331396139623464640000000000000000000000000000000000000000000000000000000005f5e1000000000000000000000000000000000000000000000000000000000000000001"

	data := hexutil.MustDecode(str)
	abiEvent, ok := tabi.Events[name]
	if !ok {
		t.Error("event do not exist")
	}

	if err := cm.UnpackEvent(abiEvent, event, []byte(data)); err != nil {
		panic(err)
	}

	t.Log(event.account)
	t.Log(event.amount)
	t.Log(event.hash)
	t.Log(event.ok)
}

type depositEvent struct {
	hash 		string
	account     string
	amount 		int
	ok 			bool
}

func newAbi() *abi.ABI {
	tabi := &abi.ABI{}

	dir := os.Getenv("GOPATH")
	println(dir)
	abiStr,err := ioutil.ReadFile(dir + "/src/github.com/dylenfu/eth-libs/contracts/transfer/abi.txt")
	if err != nil {
		panic(err)
	}

	if err := tabi.UnmarshalJSON(abiStr); err != nil {
		panic(err)
	}

	return tabi
}

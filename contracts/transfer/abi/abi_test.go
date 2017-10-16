package abi_test

import (
	"bytes"
	cm "github.com/dylenfu/eth-libs/common"
	iabi "github.com/dylenfu/eth-libs/contracts/transfer/abi"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"strings"
	"testing"
	"github.com/dylenfu/eth-libs/types"
)

func TestUnpackMethod(t *testing.T) {
	const definition = `[
	{ "name" : "int", "constant" : false, "outputs": [ { "type": "uint256" } ] },
	{ "name" : "bool", "constant" : false, "outputs": [ { "type": "bool" } ] },
	{ "name" : "bytes", "constant" : false, "outputs": [ { "type": "bytes" } ] },
	{ "name" : "fixed", "constant" : false, "outputs": [ { "type": "bytes32" } ] },
	{ "name" : "multi", "constant" : false, "outputs": [ { "type": "bytes" }, { "type": "bytes" } ] },
	{ "name" : "intArraySingle", "constant" : false, "outputs": [ { "type": "uint256[3]" } ] },
	{ "name" : "addressSliceSingle", "constant" : false, "outputs": [ { "type": "address[]" } ] },
	{ "name" : "addressSliceDouble", "constant" : false, "outputs": [ { "name": "a", "type": "address[]" }, { "name": "b", "type": "address[]" } ] },
	{ "name" : "mixedBytes", "constant" : true, "outputs": [ { "name": "a", "type": "bytes" }, { "name": "b", "type": "bytes32" } ] }]`

	tabi, err := abi.JSON(strings.NewReader(definition))
	if err != nil {
		t.Fatal(err)
	}

	// 64字节
	buff := new(bytes.Buffer)
	buff.Write(common.Hex2Bytes("0000000000000000000000000000000000000000000000000000000000000040")) // offset
	buff.Write(common.Hex2Bytes("0000000000000000000000000000000000000000000000000000000000000080")) // offset
	buff.Write(common.Hex2Bytes("0000000000000000000000000000000000000000000000000000000000000001")) // size
	buff.Write(common.Hex2Bytes("0000000000000000000000000100000000000000000000000000000000000000"))
	buff.Write(common.Hex2Bytes("0000000000000000000000000000000000000000000000000000000000000002")) // size
	buff.Write(common.Hex2Bytes("0000000000000000000000000200000000000000000000000000000000000000"))
	buff.Write(common.Hex2Bytes("0000000000000000000000000300000000000000000000000000000000000000"))

	var outAddrStruct struct {
		A []common.Address
		B []common.Address
	}
	err = cm.UnpackMethod(tabi.Methods["addressSliceDouble"], &outAddrStruct, buff.Bytes())
	if err != nil {
		t.Fatal("didn't expect error:", err)
	}

	if len(outAddrStruct.A) != 1 {
		t.Fatal("expected 1 item, got", len(outAddrStruct.A))
	}

	if outAddrStruct.A[0] != (common.Address{1}) {
		t.Errorf("expected %x, got %x", common.Address{1}, outAddrStruct.A[0])
	}

	if len(outAddrStruct.B) != 2 {
		t.Fatal("expected 1 item, got", len(outAddrStruct.B))
	}

	if outAddrStruct.B[0] != (common.Address{2}) {
		t.Errorf("expected %x, got %x", common.Address{2}, outAddrStruct.B[0])
	}
	if outAddrStruct.B[1] != (common.Address{3}) {
		t.Errorf("expected %x, got %x", common.Address{3}, outAddrStruct.B[1])
	}
}

func TestAddress(t *testing.T) {
	account := "0x56d9620237fff8a6c0f98ec6829c137477887ec4"
	t.Log(account)
	t.Log(common.HexToAddress(account).String())
}

func TestUnpackDepositEvent(t *testing.T) {
	event := iabi.DepositEvent{}

	tabi := types.NewAbi("github.com/dylenfu/eth-libs/contracts/transfer/abi.txt")

	name := "DepositFilled"
	str := "0x5ad6fe3e08ffa01bb1db674ac8e66c47511e364a4500115dd2feb33dad972d7e0000000000000000000000003865633638323963313337343737383837656334000000000000000000000000000000000000000000000000000000000bebc2010000000000000000000000000000000000000000000000000000000000000001"

	data := hexutil.MustDecode(str)

	abiEvent, ok := tabi.Events[name]
	if !ok {
		t.Error("event do not exist")
	}

	if err := cm.UnpackEvent(abiEvent, &event, data, []string{"111"}); err != nil {
		panic(err)
	}

	t.Log(common.BytesToHash(event.Hash).Hex())
	t.Log(event.Account.Hex())
	t.Log(event.Amount)
	t.Log(event.Ok)
}

func TestUnpackTransferEvent(t *testing.T) {
	transfer := iabi.TransferEvent{}

	tabi := types.NewAbi("github.com/dylenfu/eth-libs/contracts/transfer/abi.txt")

	name := "OrderFilled"
	topics := []string{"0xe82b29110155d7f50a67fadb38783bf00fbf992a5c866a55c83f85b7edadd234", "0x0000000000000000000000000000000000000000000000000000000000000002"}
	// str长度为322 包含5个字段，mustDecode后长度为160,但是打印string(str)和common.Bytes2Hex(data)在字面上只差了0x两个字母
	str := "0x00000000000000000000000056d9620237fff8a6c0f98ec6829c137477887ec400000000000000000000000046c5683c754b2eba04b2701805617c0319a9b4dd0000000000000000000000000000000000000000000000000000000005f5e1000000000000000000000000000000000000000000000000000000000005f5e1000000000000000000000000000000000000000000000000000000000000000001"

	data := hexutil.MustDecode(str)
	abiEvent, ok := tabi.Events[name]
	if !ok {
		t.Error("event do not exist")
	}

	if err := cm.UnpackEvent(abiEvent, &transfer, data, topics); err != nil {
		panic(err)
	}

	t.Log("hash", common.BytesToHash(transfer.Hash).Hex())
	t.Log("accounts", transfer.AccountS.Hex())
	t.Log("accountb", transfer.AccountB.Hex())
	t.Log("amounts", transfer.AmountS)
	t.Log("amountb", transfer.AmountB)
	t.Log("ok", transfer.Ok)
}

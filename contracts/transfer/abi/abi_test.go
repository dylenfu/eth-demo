package abi_test

import (
	"testing"
	cm "github.com/dylenfu/eth-libs/common"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"os"
	"io/ioutil"
	"reflect"
	"unsafe"
	"math/big"
	"github.com/ethereum/go-ethereum/common"
	"bytes"
	"strings"
	"github.com/ethereum/go-ethereum/common/hexutil"
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

var testStr = "0x5ad6fe3e08ffa01bb1db674ac8e66c47511e364a4500115dd2feb33dad972d7e0000000000000000000000003865633638323963313337343737383837656334000000000000000000000000000000000000000000000000000000000bebc2010000000000000000000000000000000000000000000000000000000000000001"

func TestAddress(t *testing.T) {
	account := "0x56d9620237fff8a6c0f98ec6829c137477887ec4"
	t.Log(account)
	t.Log(common.HexToAddress(account).String())
}

func TestUnpackEvent(t *testing.T) {
	event := DepositEvent{}

	tabi := newAbi()

	name := "DepositFilled"
	str := testStr

	data := hexutil.MustDecode(str)

	abiEvent, ok := tabi.Events[name]
	if !ok {
		t.Error("event do not exist")
	}

	if err := cm.UnpackEvent(abiEvent, &event, data); err != nil {
		panic(err)
	}

	t.Log(common.BytesToHash(event.Hash).Hex())
	t.Log(event.Account.Hex())
	t.Log(event.Amount)
	t.Log(event.Ok)
}

type DepositEvent struct {
	Hash 		[]byte
	Account     common.Address
	Amount 		*big.Int
	Ok 			bool
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

func BytesToString(b []byte) string {
	bh := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	sh := reflect.StringHeader{bh.Data, bh.Len}
	return *(*string)(unsafe.Pointer(&sh))
}

func StringToBytes(s string) []byte {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := reflect.SliceHeader{sh.Data, sh.Len, 0}
	return *(*[]byte)(unsafe.Pointer(&bh))
}

func TestBytesStringConvert(t *testing.T) {
	b := []byte{'b', 'y', 't', 'e'}
	s := BytesToString(b)
	t.Log(s)
	b = StringToBytes(s)
	t.Log(string(b))
}

// 在传递的时候
func TestBytesStringReflect(t *testing.T) {
	bs := []byte{'h', 'a', 's', 'h'}
	src := reflect.ValueOf(bs)

	str := ""
	dst := reflect.ValueOf(&str).Elem()

	// 这里string(src.Bytes())是字符串内容,而src.String()是类型
	v := string(src.Bytes())
	t.Log("string is ", v)

	dst.SetString(v)
	t.Log(dst.String())
}

// 对于数据结构中的类型:
// slice通过反射转array
func TestBytesStringFieldReflect(t *testing.T) {
	type ts struct {
		srcData []byte
		dstData string
	}

	tsd := &ts{[]byte{'h', 'a','s','h'}, "12"}

	valueOf := reflect.ValueOf(tsd)
	value := valueOf.Elem()

	src := value.Field(0)
	dst := value.Field(1)

	str := string(src.Bytes())
	dst = reflect.ValueOf(str)

	t.Log(str)
	t.Log(dst)
}

func TestArraySliceReflect(t *testing.T) {
	s := [4]byte{'h', 'a', 's', 'h'}
	d := []byte{}

	src := reflect.ValueOf(s)
	dst := reflect.ValueOf(d)
	dst = src

	//reflect.Copy(dst,src)

	t.Log([]byte(dst.Bytes()))
}

func TestArrayStringReflect(t *testing.T) {
	bs := [4]byte{'h', 'a', 's', 'h'}

	src := reflect.ValueOf(bs)
	ts := make([]byte, len(bs))
	for i :=0; i< len(bs); i++{
		ts[i] = src.Index(i).Interface().(byte)
	}

	t.Log(string(ts))
}

// 可以通过reflect直接赋值
func TestBigIntPtrCopy(t *testing.T) {
	bs := big.NewInt(1)
	src := reflect.ValueOf(bs)

	ts := big.NewInt(2)
	dst := reflect.ValueOf(&ts)

	t.Log(dst.CanSet())

	dst = src

	t.Log(dst.Elem().String())
	t.Log(dst)
}
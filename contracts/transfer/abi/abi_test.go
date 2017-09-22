package abi_test

import (
	"testing"
	"github.com/ethereum/go-ethereum/common/hexutil"
	cm "github.com/dylenfu/eth-libs/common"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"os"
	"io/ioutil"
	"reflect"
	"unsafe"
)

func TestNewFilter(t *testing.T) {
	var event depositEvent

	tabi := newAbi()

	name := "DepositFilled"
	str := "0x69be7bc7c7c6e216dd9531c88c94769f9f63ce53f47665b5ec7faf55f8094e8100000000000000000000000037303138303536313763303331396139623464640000000000000000000000000000000000000000000000000000000005f5e1000000000000000000000000000000000000000000000000000000000000000001"

	data := hexutil.MustDecode(str)
	abiEvent, ok := tabi.Events[name]
	if !ok {
		t.Error("event do not exist")
	}

	if err := cm.UnpackEvent(abiEvent, &event, data); err != nil {
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

	t.Log(dst.Kind().String())
	t.Log(str)
	t.Log(dst)
}
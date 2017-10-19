package abi

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"reflect"
	"testing"
	"unsafe"
)

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

	tsd := &ts{[]byte{'h', 'a', 's', 'h'}, "12"}

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
	for i := 0; i < len(bs); i++ {
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

func TestReflectStructTag(t *testing.T) {
	type Deposit struct {
		Id     []byte         `alias:"_id"`
		Owner  common.Address `alias:"_owner"`
		Amount *big.Int       `alias:"_amount"`
	}

	d := Deposit{}
	t.Log(reflect.TypeOf(d).Field(0).Tag.Get("alias"))
}

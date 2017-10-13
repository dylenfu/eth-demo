package abi_test

import (
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
	"reflect"
	"regexp"
	"testing"
)

func TestHexNumber(t *testing.T) {
	i := hexutil.EncodeBig(big.NewInt(32))
	t.Log(reflect.TypeOf(i))
}

func TestRegexp(t *testing.T) {
	reg := regexp.MustCompile(`[a-zA-Z0-9]`)
	list := reg.FindAllString("Address[2][3]", -1)
	for _, v := range list {
		t.Log(v)
	}
}

package abi

import (
	"github.com/ethereum/go-ethereum/common/hexutil"
	"testing"
	"math/big"
	"reflect"
)

func TestHexNumber(t *testing.T) {
	i := hexutil.EncodeBig(big.NewInt(32))
	t.Log(reflect.TypeOf(i))
}

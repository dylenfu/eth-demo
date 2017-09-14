package common

import (
	"github.com/ethereum/go-ethereum/common/hexutil"
	"strconv"
)

func ToHexBigInt(src int) hexutil.Big {
	var ret hexutil.Big

	str := strconv.FormatInt(int64(src), 16)
	// 这里注意，一定要加上"0x"
	des :=  "0x" + str
	ret.UnmarshalText([]byte(des))

	return ret
}
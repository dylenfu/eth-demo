package types

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"math/big"
	"strconv"
)

// 调用jsonrpc call/sendTransaction时字符串转换成以太坊地址
func Str2Address(str string) common.Address{
	return common.HexToAddress(str)
}

// 创建filter时使用int获得FromBlock
func Int2BlockNumHex(height int64) string {
	data := big.NewInt(height)
	return hexutil.EncodeBig(data) //common.Bytes2Hex(height.Bytes())
}

// 使用int类型生成hexutil.big，jsonrpc中call&sendTransaction使用
func Int2HexBigInt(src int) hexutil.Big {
	var ret hexutil.Big

	str := strconv.FormatInt(int64(src), 16)
	// 这里注意，一定要加上"0x"
	des :=  "0x" + str
	ret.UnmarshalText([]byte(des))

	return ret
}

func Uint2HexUintPtr(src uint64) *hexutil.Uint64 {
	var ret *hexutil.Uint64

	str := strconv.FormatUint(src, 16)
	des := "0x" + str
	ret.UnmarshalText([]byte(des))

	common.ToHex([]byte(des))
	return ret
}

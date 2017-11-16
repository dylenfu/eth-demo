package main

import (
	"flag"
	. "github.com/dylenfu/eth-libs/contracts/ringmined/contract"
	"github.com/dylenfu/eth-libs/params"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"qiniupkg.com/x/log.v7"
	"reflect"
	"time"
)

type Handle struct{}

var fn = flag.String("fn", "submitRing", "chose function name")

func main() {
	flag.Parse()
	handle := &Handle{}
	reflect.ValueOf(handle).MethodByName(*fn).Call([]reflect.Value{})
}

func (h *Handle) submitRing() {
	impl := RingMined.Token.(*RingMinedImpl)
	var result string

	ringIndex := big.NewInt(101)
	ringHash := common.HexToHash("0xbf57becafab89ce3f69b1226d32cafb30cfd4db8fcfdc5a6bf9ec607cef14915")
	miner := common.HexToAddress("0x06D534CFA972363ca7D108dBE1E2cAfFef62913B")
	feeRecipient := common.HexToAddress("0x06D534CFA972363ca7D108dBE1E2cAfFef62913B")
	isRinghashReserved := false
	orderHash := "0xbf57becafab89ce3f69b1226d32cafb30cfd4db8fcfdc5a6bf9ec607cef14915"
	amount := big.NewInt(3333)
	reward := big.NewInt(2222)
	fee := big.NewInt(11)

	err := impl.SubmitRing.SendTransaction(&result,
		ringIndex,
		ringHash,
		miner,
		feeRecipient,
		isRinghashReserved,
		orderHash,
		amount,
		reward,
		fee)
	if err != nil {
		panic(err)
	}
}

func (h *Handle) ListenEvent() {
	h.submitRing()
	time.Sleep(5 * time.Second)
	filterId, err := NewFilter(params.BlockNumber)
	if err != nil {
		panic(err)
	}

	for {
		err := GetEvent(filterId)
		if err != nil {
			log.Error(err.Error())
		}
	}
}

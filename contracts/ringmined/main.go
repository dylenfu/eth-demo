package main

import (
	"flag"
	. "github.com/dylenfu/eth-libs/contracts/ringmined/contract"
	"github.com/ethereum/go-ethereum/common"
	"log"
	"math/big"
	"reflect"
)

type Handle struct{}

var (
	fn = flag.String("fn", "SubmitRing", "chose function name")
	tx = flag.String("tx", "0x0", "transaction hash")
)

func main() {
	flag.Parse()
	handle := &Handle{}
	reflect.ValueOf(handle).MethodByName(*fn).Call([]reflect.Value{})
}

func (h *Handle) Min() {
	log.Println("Min Starting......")

	impl := RingMined.Token.(*RingMinedImpl)

	var result string
	miner := common.HexToAddress("0x06D534CFA972363ca7D108dBE1E2cAfFef62913B")
	num := big.NewInt(3000)
	if err := impl.Min.SendTransaction(&result, miner, num); err != nil {
		panic(err)
	}
}

func (h *Handle) SimpleRing() {
	log.Println("Just Ring Starting......")

	impl := RingMined.Token.(*RingMinedImpl)

	var result string

	protocol := common.HexToAddress("0x06D534CFA972363ca7D108dBE1E2cAfFef62913B")
	owner := common.HexToAddress("0xd1ac65fa97a820274b51d92bc46ae08f747e77cg")
	amount := big.NewInt(36)
	res := big.NewInt(1200)

	if err := impl.SimpleRing.SendTransaction(&result, protocol, owner, amount, res); err != nil {
		panic(err)
	}
}

func (h *Handle) JustRing() {
	log.Println("Just Ring Starting......")

	impl := RingMined.Token.(*RingMinedImpl)

	var result string

	ringIndex := big.NewInt(101)
	ringHash := common.HexToHash("0xbf57becafab89ce3f69b1226d32cafb30cfd4db8fcfdc5a6bf9ec607cef14915")
	miner := common.HexToAddress("0x06D534CFA972363ca7D108dBE1E2cAfFef62913B")
	feeRecipient := common.HexToAddress("0x06D534CFA972363ca7D108dBE1E2cAfFef62913B")
	isRinghashReserved := true

	err := impl.JustRing.SendTransaction(&result,
		ringIndex,
		ringHash,
		miner,
		feeRecipient,
		isRinghashReserved)

	if err != nil {
		panic(err)
	}
}

func (h *Handle) SubmitRing() {
	log.Println("Submit Ring Starting......")

	impl := RingMined.Token.(*RingMinedImpl)

	var result string

	ringIndex := big.NewInt(101)
	ringHash := common.HexToHash("0xbf57becafab89ce3f69b1226d32cafb30cfd4db8fcfdc5a6bf9ec607cef14915")
	miner := common.HexToAddress("0x06D534CFA972363ca7D108dBE1E2cAfFef62913B")
	feeRecipient := common.HexToAddress("0xe47f7eb8b08929984713fd0963cd943d79b2b707")
	isRinghashReserved := false
	orderHash := common.HexToHash("0x4ab5dde696d27eb854501a5b03842c5d0b8ab31b3ced8d314a31930976dd85f5")
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

func (h *Handle) RingMinedEvent() {
	txstr := *tx
	if err := GetRingMinedEvent(txstr); err != nil {
		log.Fatal(err)
	}
}

func (h *Handle) RingEvent() {
	txstr := *tx
	if err := GetRingEvent(txstr); err != nil {
		log.Fatal(err)
	}
}

func (h *Handle) MinEvent() {
	txstr := *tx
	if err := GetMinEvent(txstr); err != nil {
		log.Fatal(err)
	}
}

func (h *Handle) SimpleRingEvent() {
	txstr := *tx
	if err := GetSimpleRingEvent(txstr); err != nil {
		panic(err)
	}
}

func (h *Handle) EthLogs() {
	GetLogs()
}
package contract

import (
	"github.com/dylenfu/eth-libs/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"log"
	"math/big"
	"strconv"
)

var RingMined *types.TokenImpl

const (
	AbiFilePath     = "github.com/dylenfu/eth-libs/contracts/ringmined/abi.txt"
	EthRpcUrl       = "http://127.0.0.1:8545"
	ContractAddress = "0xb399bd657c9f03f418e852fe72828ce40339fe85"
)

func init() {
	token := &RingMinedImpl{}
	RingMined = types.NewContract(AbiFilePath, ContractAddress, EthRpcUrl, token)
}

type RingMinedImpl struct {
	SubmitRing types.AbiMethod `methodName:"submitRing"`
	JustRing   types.AbiMethod `methodName:"justRing"`
	Min        types.AbiMethod `methodName:"min"`
	SimpleRing types.AbiMethod `methodName:"simpleRing"`
}

type OrderFilled struct {
	OrderHash     common.Hash `alias:"_orderHash"`
	NextOrderHash common.Hash `alias:"_nextOrderHash"`
	AmountS       *big.Int    `alias:"_amountS"`
	AmountB       *big.Int    `alias:"_amountB"`
	LrcReward     *big.Int    `alias:"_lrcReward"`
	LrcFee        *big.Int    `alias:"_lrcFee"`
}

type SimpleFill struct {
	Owner  common.Address `alias:"owner"`
	Amount *big.Int       `alias:"amount"`
}

type SimpleRingEvent struct {
	Protocol common.Address `alias:"protocol"`
	Res      *big.Int       `alias:"res"`
	Fills    []SimpleFill   `alias:"fills"`
}

type RingEvent struct {
	RingIndex          *big.Int       `alias:"_ringIndex"`
	RingHash           common.Hash    `alias:"_ringhash"`
	Miner              common.Address `alias:"_miner"`
	FeeRecipient       common.Address `alias:"_feeRecipient"`
	IsRingHashReserved bool           `alias:"_isRinghashReserved"`
}

type RingMinedEvent struct {
	RingIndex          *big.Int       `alias:"_ringIndex"`
	RingHash           common.Hash    `alias:"_ringhash"`
	Miner              common.Address `alias:"_miner"`
	FeeRecipient       common.Address `alias:"_feeRecipient"`
	IsRingHashReserved bool           `alias:"_isRinghashReserved"`
	Fills              []OrderFilled  `alias:"_fills"`
}

type MinEvent struct {
	Miner   common.Address `alias:"miner"`
	Amounts []*big.Int     `alias:"amounts"`
}

func NewFilter(height int) (string, error) {
	var filterId string

	// 使用jsonrpc的方式调用newFilter
	filter := types.FilterReq{}
	filter.FromBlock = types.Int2BlockNumHex(int64(height))
	filter.ToBlock = "latest"
	filter.Address = RingMined.TokenAddress

	err := RingMined.Client.Call(&filterId, "eth_newFilter", &filter)
	if err != nil {
		return "", err
	}

	return filterId, nil
}

func GetRingMinedEvent(txhex string) error {
	var (
		recipient types.RTransactionRecipient
		evt       RingMinedEvent
		data      []byte
		err       error
	)

	txhash := common.HexToHash(txhex)
	if err = RingMined.Client.Call(&recipient, "eth_getTransactionReceipt", txhash); err != nil {
		return err
	}

	event := recipient.Logs[0]
	data = hexutil.MustDecode(event.Data)

	log.Printf("recepient block number:%s", recipient.BlockNumber.ToInt().String())
	log.Printf("before hex decord string:%s", event.Data)

	if err = RingMined.Abi.Unpack(&evt, "RingMinedEvent", data, abi.SEL_UNPACK_EVENT); err != nil {
		return err
	}

	log.Printf("evt.ringIndex:%s", evt.RingIndex.String())
	log.Printf("evt.ringHash:%s", evt.RingHash.Hex())
	log.Printf("evt.miner:%s", evt.Miner.Hex())
	log.Printf("evt.feeRecipient:%s", evt.FeeRecipient.Hex())
	log.Printf("evt.isRingHashReserved:%s", strconv.FormatBool(evt.IsRingHashReserved))

	for _, fill := range evt.Fills {
		log.Printf("fill.orderHash:%s", fill.OrderHash.Hex())
		log.Printf("fill.nextOrderHash:%s", fill.NextOrderHash.Hex())
		log.Printf("fill.amountS:%s", fill.AmountS.String())
		log.Printf("fill.amountB:%s", fill.AmountB.String())
		log.Printf("fill.lrcReward:%s", fill.LrcReward.String())
		log.Printf("fill.lrcFee:%s", fill.LrcFee.String())
	}
	return nil
}

func GetRingEvent(txhex string) error {
	var (
		recipient types.RTransactionRecipient
		evt       RingEvent
		data      []byte
		err       error
	)

	txhash := common.HexToHash(txhex)
	if err = RingMined.Client.Call(&recipient, "eth_getTransactionReceipt", txhash); err != nil {
		return err
	}

	event := recipient.Logs[0]
	data = hexutil.MustDecode(event.Data)

	log.Printf("recepient block number:%s", recipient.BlockNumber.ToInt().String())
	log.Printf("before hex decord string:%s", event.Data)

	if err = RingMined.Abi.Unpack(&evt, "RingEvent", data, abi.SEL_UNPACK_EVENT); err != nil {
		return err
	}

	log.Printf("evt.ringIndex:%s", evt.RingIndex.String())
	log.Printf("evt.ringHash:%s", evt.RingHash.Hex())
	log.Printf("evt.miner:%s", evt.Miner.Hex())
	log.Printf("evt.feeRecipient:%s", evt.FeeRecipient.Hex())
	log.Printf("evt.isRingHashReserved:%s", strconv.FormatBool(evt.IsRingHashReserved))

	return nil
}

func GetMinEvent(txhex string) error {
	var (
		recipient types.RTransactionRecipient
		evt       MinEvent
		data      []byte
		err       error
	)

	txhash := common.HexToHash(txhex)
	if err = RingMined.Client.Call(&recipient, "eth_getTransactionReceipt", txhash); err != nil {
		return err
	}

	event := recipient.Logs[0]
	data = hexutil.MustDecode(event.Data)

	log.Printf("recepient block number:%s", recipient.BlockNumber.ToInt().String())
	log.Printf("before hex decord string:%s", event.Data)

	if err = RingMined.Abi.Unpack(&evt, "MinEvent", data, abi.SEL_UNPACK_EVENT); err != nil {
		return err
	}

	log.Printf("evt.miner:%s", evt.Miner.Hex())
	for k, v := range evt.Amounts {
		log.Printf("evt.amount.%d:%s", k, v.String())
	}

	return nil
}

func GetSimpleRingEvent(txhex string) error {
	var (
		recipient types.RTransactionRecipient
		evt       SimpleRingEvent
		data      []byte
		err       error
	)

	txhash := common.HexToHash(txhex)
	if err = RingMined.Client.Call(&recipient, "eth_getTransactionReceipt", txhash); err != nil {
		return err
	}

	event := recipient.Logs[0]
	data = hexutil.MustDecode(event.Data)

	log.Printf("recepient block number:%s", recipient.BlockNumber.ToInt().String())
	log.Printf("before hex decord string:%s", event.Data)

	if err = RingMined.Abi.Unpack(&evt, "SimpleRingEvent", data, abi.SEL_UNPACK_EVENT); err != nil {
		return err
	}

	log.Printf("evt.protocol:%s", evt.Protocol.Hex())
	log.Printf("evt.res:%s", evt.Res.String())

	for k, v := range evt.Fills {
		log.Printf("evt.fills[%d].owner:%s", k, v.Owner.Hex())
		log.Printf("evt.fills[%d].amount:%s", k, v.Amount.String())
	}

	return nil
}

func GetLogs() error {
	var (
		fl  types.FilterLog
		err error
	)

	methodId := ""
	if err = RingMined.Client.Call(&fl, "eth_getFilterChanges", methodId); err != nil {
		return err
	}

	log.Printf("logindex:%d", fl.LogIndex.Int())
	log.Printf("evt.txhash:%s", fl.TransactionHash)

	return nil
}

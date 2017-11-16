package contract

import (
	"github.com/dylenfu/eth-libs/types"
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
	ContractAddress = "0x06D534CFA972363ca7D108dBE1E2cAfFef62913B"
)

func init() {
	token := &RingMinedImpl{}
	RingMined = types.NewContract(AbiFilePath, ContractAddress, EthRpcUrl, token)
}

type RingMinedImpl struct {
	SubmitRing types.AbiMethod `methodName:"submitRing"`
}

type OrderFilled struct {
	OrderHash     common.Hash `alias:"_orderHash"`
	NextOrderHash common.Hash `alias:"_nextOrderHash"`
	AmountS       *big.Int    `alias:"_amountS"`
	AmountB       *big.Int    `alias:"_amountB"`
	LrcReward     *big.Int    `alias:"_lrcReward"`
	LrcFee        *big.Int    `alias:"_lrcFee"`
}

type RingMinedEvent struct {
	RingIndex          *big.Int       `alias:"_ringIndex"`
	RingHash           common.Hash    `alias:"_ringhash"`
	Miner              common.Address `alias:"_miner"`
	FeeRecipient       common.Address `alias:"_feeRecipient"`
	IsRingHashReserved bool           `alias:"_isRinghashReserved"`
	Fills              []OrderFilled  `alias:"_fills"`
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

func GetEvent(filterId string) error {
	var logs []types.FilterLog

	impl := RingMined
	err := RingMined.Client.Call(&logs, "eth_getFilterChanges", filterId)
	if err != nil {
		return err
	}
	evts := impl.Abi.Events

	for _, v := range logs {
		println(v.Data)
		data := hexutil.MustDecode(v.Data)

		switch v.Topics[0] {
		case evts["RingMined"].Id().String():
			if err := showRingMined("SemenEvent", data, v.Topics); err != nil {
				return err
			}
		}
	}

	return nil
}

func showRingMined(eventName string, data []byte, topics []string) error {
	evt := &RingMinedEvent{}
	if err := RingMined.Abi.Unpack(&evt, "RingMined", data); err != nil {
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

package abi

import (
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/common"
	cm "github.com/dylenfu/eth-libs/common"
	"github.com/ethereum/go-ethereum/rpc"
	"reflect"
	"os"
	"io/ioutil"
	"github.com/dylenfu/eth-libs/types"
	"log"
	"github.com/pkg/errors"
	"math/big"
	. "github.com/dylenfu/eth-libs/params"
	"github.com/ethereum/go-ethereum/accounts/abi"
)

var (
	client 		*rpc.Client
	tabi 		*abi.ABI
)

const Topic = "bank"

func init() {
	var err error

	client, err = rpc.Dial("http://127.0.0.1:8545")
	if err != nil {
		panic(err)
	}

	tabi = NewAbi()
}

func NewAbi() *abi.ABI {
	tabi := &abi.ABI{}

	dir := os.Getenv("GOPATH")
	abiStr,err := ioutil.ReadFile(dir + "/src/github.com/dylenfu/eth-libs/contracts/transfer/abi.txt")
	if err != nil {
		panic(err)
	}

	if err := tabi.UnmarshalJSON(abiStr); err != nil {
		panic(err)
	}

	return tabi
}

type BankToken struct {
	Transfer 		types.AbiMethod	`methodName:"submitTransfer"`
	Deposit			types.AbiMethod	`methodName:"submitDeposit"`
	BalanceOf		types.AbiMethod	`methodName:"balanceOf"`
}

func LoadContract() *BankToken {
	contract := &BankToken{}
	elem := reflect.ValueOf(contract).Elem()

	for i:=0; i < elem.NumField(); i++ {
		methodName := elem.Type().Field(i).Tag.Get("methodName")

		abiMethod := &types.AbiMethod{}
		abiMethod.Name = methodName
		abiMethod.Abi = tabi
		abiMethod.Address = TransferTokenAddress
		abiMethod.Client = client

		elem.Field(i).Set(reflect.ValueOf(*abiMethod))
	}

	return contract
}

// filter可以根据blockNumber生成
// 也可以从网络中直接查询eth.filter()
func NewFilter(height *big.Int) (string,error) {
	var filterId string

	// 使用jsonrpc的方式调用newFilter
	filter := types.FilterReq{}
	filter.FromBlock = hexutil.EncodeBig(height)//common.Bytes2Hex(height.Bytes())
	filter.ToBlock = "latest"
	filter.Address = TransferTokenAddress

	err := client.Call(&filterId, "eth_newFilter", &filter)
	if err != nil {
		return "", err
	}

	return filterId, nil
}

type DepositEvent struct {
	Hash 		[]byte
	Account     common.Address
	Amount 		*big.Int
	Ok 			bool
}

type TransferEvent struct {

}

// 监听合约事件并解析
func FilterChanged(filterId string) error {
	var logs []types.FilterLog

	err := client.Call(&logs, "eth_getFilterChanges", filterId)
	if err != nil {
		return err
	}

	depositEventName := "DepositFilled"
	//orderEventName := "OrderFilled"

	for _, v := range logs {
		log.Println(v.Data)
		log.Println(v.TransactionHash)

		// 转换hex
		data := hexutil.MustDecode(v.Data)
		// topics第一个元素就是eventId
		eventId := v.Topics[0]

		switch eventId {
		case tabi.Events[depositEventName].Id().String():
			event, ok := tabi.Events[depositEventName]
			if !ok {
				return errors.New("deposit event do not exsit")
			}
			deposit := &DepositEvent{}

			//
			if err := cm.UnpackEvent(event, deposit, []byte(data)); err != nil {
				return err
			} else {
				log.Println("hash", common.BytesToHash(deposit.Hash).Hex())
				log.Println("account", deposit.Account.Hex())
				log.Println("amount", deposit.Amount)
				log.Println("isOk", deposit.Ok)
			}

		case tabi.Events[""].Id().String():

		}
	}

	return nil
}
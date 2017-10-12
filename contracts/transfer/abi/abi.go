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
func NewFilter(height int) (string,error) {
	var filterId string

	// 使用jsonrpc的方式调用newFilter
	filter := types.FilterReq{}
	filter.FromBlock = types.Int2BlockNumHex(int64(height))
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
	Hash 		[]byte
	AccountS 	common.Address
	AccountB 	common.Address
	AmountS 	*big.Int
	AmountB 	*big.Int
	Ok 			bool
}

// 监听合约事件并解析
func FilterChanged(filterId string) error {
	var logs []types.FilterLog

	// 注意 这里使用filterchanges获得的是以太坊最新的log
	// 使用filterlogs获得的是fromBlock后的所有log
	// 所以，一般而言我们在程序里一般都是启动时使用getFilterLogs
	// 过滤掉已经解析了的logs后，使用getFilterChange继续监听
	err := client.Call(&logs, "eth_getFilterChanges", filterId)
	//err := client.Call(&logs, "eth_getFilterLogs", filterId)
	if err != nil {
		return err
	}

	den := "DepositFilled"
	oen := "OrderFilled"
	denId := tabi.Events[den].Id().String()
	oenId := tabi.Events[oen].Id().String()

	for _, v := range logs {
		// 转换hex
		data := hexutil.MustDecode(v.Data)

		// topics第一个元素就是eventId
		switch v.Topics[0] {
		case denId:
			if err := showDeposit(den, data, v.Topics); err != nil {
				return err
			}
		case oenId:
			if err := showTransfer(oen, data, v.Topics); err != nil {
				return err
			}
		}
	}

	return nil
}

// 解析并显示event数据内容
func showDeposit(eventName string, data []byte, topics []string) error {
	event, ok := tabi.Events[eventName]
	if !ok {
		return errors.New("deposit event do not exsit")
	}
	deposit := &DepositEvent{}

	//
	if err := cm.UnpackEvent(event, deposit, []byte(data), topics); err != nil {
		return err
	}

	log.Println("hash", common.BytesToHash(deposit.Hash).Hex())
	log.Println("account", deposit.Account.Hex())
	log.Println("amount", deposit.Amount)
	log.Println("ok", deposit.Ok)

	return nil
}

// 解析并显示OrderFilledEvent数据内容
func showTransfer(eventName string, data []byte, topics []string) error {
	event, ok := tabi.Events[eventName]
	if !ok {
		return errors.New("orderFilled event do not exist")
	}

	transfer := &TransferEvent{}
	if err := cm.UnpackEvent(event, transfer, []byte(data), topics); err != nil {
		return err
	}

	log.Println("hash", common.BytesToHash(transfer.Hash).Hex())
	log.Println("accounts", transfer.AccountS.Hex())
	log.Println("accountb", transfer.AccountB.Hex())
	log.Println("amounts", transfer.AmountS)
	log.Println("amountb", transfer.AmountB)
	log.Println("ok", transfer.Ok)

	return nil
}
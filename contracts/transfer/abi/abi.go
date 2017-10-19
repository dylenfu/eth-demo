package abi

import (
	cm "github.com/dylenfu/eth-libs/common"
	"github.com/dylenfu/eth-libs/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/pkg/errors"
	"log"
	"math/big"
)

var (
	Bank *types.TokenImpl
)

const (
	Topic        = "bank"
	AbiFilePath  = "github.com/dylenfu/eth-libs/contracts/transfer/abi.txt"
	EthRpcUrl    = "http://127.0.0.1:8545"
	TokenAddress = "0xa221f7c8cd24a7a383d116aa5d7430b48d1e0063"
)

func init() {
	token := &BankToken{}
	Bank = types.NewContract(AbiFilePath, TokenAddress, EthRpcUrl, token)
}

type BankToken struct {
	Transfer  types.AbiMethod `methodName:"submitTransfer"`
	Deposit   types.AbiMethod `methodName:"submitDeposit"`
	BalanceOf types.AbiMethod `methodName:"balanceOf"`
}

type DepositEvent struct {
	Hash    []byte         `alias:"hash"`
	Account common.Address `alias:"account"`
	Amount  *big.Int       `alias:"amount"`
	Ok      bool           `alias:"ok"`
}

type Ha interface{}

// 所有跟event不相关的字段在解析的时候都没有影响
type TransferEvent struct {
	Ha
	Hie      string
	Hash     []byte         `alias:"hash"`
	AccountS common.Address `alias:"accountS"`
	AccountB common.Address `alias:"accountB"`
	AmountS  *big.Int       `alias:"amountS"`
	AmountB  *big.Int       `alias:"amountB"`
	Ok       bool           `alias:"ok"`
}

// filter可以根据blockNumber生成
// 也可以从网络中直接查询eth.filter()
func NewFilter(height int) (string, error) {
	var filterId string

	// 使用jsonrpc的方式调用newFilter
	filter := types.FilterReq{}
	filter.FromBlock = types.Int2BlockNumHex(int64(height))
	filter.ToBlock = "latest"
	filter.Address = Bank.TokenAddress

	err := Bank.Client.Call(&filterId, "eth_newFilter", &filter)
	if err != nil {
		return "", err
	}

	return filterId, nil
}

// 监听合约事件并解析
func EventChanged(filterId string) error {
	var logs []types.FilterLog

	// 注意 这里使用filterchanges获得的是以太坊最新的log
	// 使用filterlogs获得的是fromBlock后的所有log
	// 所以，一般而言我们在程序里一般都是启动时使用getFilterLogs
	// 过滤掉已经解析了的logs后，使用getFilterChange继续监听
	err := Bank.Client.Call(&logs, "eth_getFilterChanges", filterId)
	//err := client.Call(&logs, "eth_getFilterLogs", filterId)
	if err != nil {
		return err
	}

	den := "DepositFilled"
	oen := "OrderFilled"
	denId := Bank.Abi.Events[den].Id().String()
	oenId := Bank.Abi.Events[oen].Id().String()

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
	event, ok := Bank.Abi.Events[eventName]
	if !ok {
		return errors.New("deposit event do not exsit")
	}
	deposit := &DepositEvent{}

	//
	if err := cm.UnpackEvent(event.Inputs, deposit, []byte(data), topics); err != nil {
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
	event, ok := Bank.Abi.Events[eventName]
	if !ok {
		return errors.New("orderFilled event do not exist")
	}

	transfer := &TransferEvent{}
	if err := cm.UnpackEvent(event.Inputs, transfer, []byte(data), topics); err != nil {
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

// 使用jsonrpc eth_newBlockFilter得到一个filterId
// 然后使用jsonrpc eth_getFilterChange得到blockHash数组
// 轮询数组，解析block信息
func BlockFilterId() (string, error) {
	var filterId string
	if err := Bank.Client.Call(&filterId, "eth_newBlockFilter"); err != nil {
		return "", err
	}
	return filterId, nil
}

// 拿到的block一直是最新的
func BlockChanged(filterId string) error {
	var blockHashs []string

	err := Bank.Client.Call(&blockHashs, "eth_getFilterChanges", filterId)
	if err != nil {
		return err
	}

	for _, v := range blockHashs {
		var block types.Block
		// 最后一个参数：true查询整个block信息，false查询block包含的transaction hash
		if err := Bank.Client.Call(&block, "eth_getBlockByHash", v, true); err != nil {
			log.Println(err)
		}
		showBlock(block)
	}

	return nil
}

func showBlock(block types.Block) {
	if len(block.Transactions) > 0 {
		log.Println("number", block.Number.ToInt())
		log.Println("hash", block.Hash)
		log.Println("parentHash", block.ParentHash)
		log.Println("nonce", block.Nonce)
		log.Println("sha3Uncles", block.Sha3Uncles)
		log.Println("logsBloom", block.LogsBloom)
		log.Println("TransactionsRoot", block.TransactionsRoot)
		log.Println("ReceiptsRoot", block.ReceiptsRoot)
		log.Println("Miner", block.Miner)
		log.Println("Difficulty", block.Difficulty.String())
		log.Println("TotalDifficulty", block.TotalDifficulty.String())
		log.Println("ExtraData", block.ExtraData)
		log.Println("Size", block.Size)
		log.Println("GasLimit", block.GasLimit)
		log.Println("GasUsed", block.GasUsed)
		log.Println("Timestamp", block.Timestamp)
		for _, v := range block.Transactions {
			showRecptTransaction(v)
		}
	}
}

func showRecptTransaction(v types.RTransaction) {
	log.Println("transaction.hash", v.Hash)
	log.Println("transaction.nonce", v.Nonce.String())
	log.Println("transaction.blockhash", v.BlockHash)
	log.Println("transaction.blocknumber", v.BlockNumber.String())
	log.Println("transaction.transactionIndex", v.TransactionIndex)
	log.Println("transaction.from", v.From)
	log.Println("transaction.to", v.To)
	log.Println("transaction.gas", v.Gas.String())
	log.Println("transaction.gasPrice", v.GasPrice.String())
	log.Println("transaction.value", v.Value.String())
	log.Println("transaction.data", v.Input)
}

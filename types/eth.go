package types

import (
	"github.com/dylenfu/eth-libs/params"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/rpc"
)

// 监听block时需要解析
type Block struct {
	Number           hexutil.Big
	Hash             string
	ParentHash       string
	Nonce            string
	Sha3Uncles       string
	LogsBloom        string
	TransactionsRoot string
	ReceiptsRoot     string
	Miner            string
	Difficulty       hexutil.Big
	TotalDifficulty  hexutil.Big
	ExtraData        string
	Size             hexutil.Big
	GasLimit         hexutil.Big
	GasUsed          hexutil.Big
	Timestamp        hexutil.Big
	Transactions     []RTransaction
	Uncles           []string
}

// 发送transaction
type Transaction struct {
	From     string
	To       string
	Gas      hexutil.Big
	GasPrice hexutil.Big
	Value    hexutil.Big
	Data     string
}

// 接受到transaction
type RTransaction struct {
	Hash             string
	Nonce            hexutil.Big
	BlockHash        string
	BlockNumber      hexutil.Big
	TransactionIndex hexutil.Big
	From             string
	To               string
	Value            hexutil.Big
	GasPrice         hexutil.Big
	Gas              hexutil.Big
	Input            string
}

type RTransactionRecipient struct {
	TransactionHash   string
	TransactionIndex  hexutil.Big
	BlockNumber       hexutil.Big
	BlockHash         string
	CumulativeGasUsed hexutil.Big
	GasUsed           hexutil.Big
	ContractAddress   string
	Status            hexutil.Big
	Logs              []FilterLog
}

// 使用call方法
type CallArgs struct {
	From     string
	To       string
	Gas      hexutil.Big
	GasPrice hexutil.Big
	Value    hexutil.Big
	Data     interface{}
}

// 使用newFilter方法生成filter
type FilterReq struct {
	FromBlock string
	ToBlock   string
	Address   string
	Topics    []string
}

// 监听某个filter得到的log数据结构
type FilterLog struct {
	Removed 		 bool 	   `json:"removed"`
	LogIndex         HexNumber `json:"logIndex"`
	BlockNumber      HexNumber `json:"blockNumber"`
	BlockHash        string    `json:"blockHash"`
	TransactionHash  string    `json:"transactionHash"`
	TransactionIndex HexNumber `json:"transactionIndex"`
	Address          string    `json:"address"`
	Data             string    `json:"data"`
	Topics           []string  `json:"topics"`
}

type AbiMethod struct {
	abi.Method
	Abi     *abi.ABI
	Address string
	Client  *rpc.Client
}

func (method *AbiMethod) Call(result interface{}, tag string, args ...interface{}) error {
	bytes, err := method.Abi.Pack(method.Name, args...)
	if err != nil {
		return err
	}

	c := &CallArgs{}
	c.From = method.Address
	c.To = method.Address
	c.Data = common.ToHex(bytes)

	return method.Client.Call(result, "eth_call", c, tag)
}

// sendTransaction是不需要tag的
func (method *AbiMethod) SendTransaction(result interface{}, args ...interface{}) error {
	bytes, err := method.Abi.Pack(method.Name, args...)
	if err != nil {
		return err
	}

	tx := &Transaction{}
	tx.From = params.Miner
	tx.To = method.Address
	tx.Gas = Int2HexBigInt(1200000)
	tx.GasPrice = Int2HexBigInt(1)
	tx.Data = common.ToHex(bytes)

	return method.Client.Call(result, "eth_sendTransaction", tx)
}

func (method *AbiMethod) SignAndSendTransaction(from string, result interface{}, args ...interface{}) error {
	bytes, err := method.Abi.Pack(method.Name, args...)
	if err != nil {
		return err
	}

	tx := &Transaction{}
	tx.From = from
	tx.To = method.Address
	tx.Gas = Int2HexBigInt(1200000)
	tx.GasPrice = Int2HexBigInt(1)
	tx.Data = common.ToHex(bytes)

	// todo sign

	//
	//var signRes string
	//txData, err := rlp.EncodeToBytes(tx)
	//if err != nil {
	//	return nil
	//}

	// return method.Client.Call(&signRes, "eth_sign", from, common.ToHex(txData))

	return method.Client.Call(result, "eth_sendTransaction", tx)
}

package main

import (
	"github.com/dylenfu/eth-libs/client/ethsrc/rpc"
	"log"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"flag"
	"reflect"
	"strconv"
)

var call = flag.String("call", "Balance", "chose test case")

type Handle struct {
	client *rpc.Client
}

// 在私链建立三个账户，miner是第一个账户，用于挖矿，account1&account2作为测试账户
// 每一次transaction之前都要解锁相关账户
const(
	Miner = "0x4bad3053d574cd54513babe21db3f09bea1d387d" // pwd 101
	Account1 = "0x46c5683c754b2eba04b2701805617c0319a9b4dd" // pwd 102
	Account2 = "0x56d9620237fff8a6c0f98ec6829c137477887ec4" // pwd 103
)

// transaction数据结构参考 https://github.com/ethereum/wiki/wiki/JSON-RPC#eth_sendtransaction
// 也可以直接使用go-eth本身的数据结构，如果是自己构造的话，需要保证能被正确解析(主要是hexutil.Big)
type transaction struct {
	From		string
	To 			string
	Gas			*hexutil.Big
	GasPrice	*hexutil.Big
	Value       *hexutil.Big
	Data		string
}

// block数据结构参考 https://github.com/ethereum/wiki/wiki/JSON-RPC#eth_getblockbyhash
type block struct {
	Number *hexutil.Big
	Hash string
	ParentHash string
	Nonce string
	Sha3Uncles string
	LogsBloom string
	TransactionRoot string
	StateRoot string
	Miner string
	Difficulty *hexutil.Big
	TotalDifficulty *hexutil.Big
	ExtraData string
	Size *hexutil.Big
	GasLimit *hexutil.Big
	GasUsed *hexutil.Big
	Timestamp string
	Transactions []transaction
	uncles []string
}

// 这里我们使用http的形式连接eth私有链
func main() {
	flag.Parse()
	c, _ := rpc.Dial("http://127.0.0.1:8545")
	handle := &Handle{client: c}
	reflect.ValueOf(handle).MethodByName(*call).Call([]reflect.Value{})
}

// 查询账户余额
func (h *Handle) Balance() {
	var amount hexutil.Big
	h.client.Call(&amount, "eth_getBalance", Account1, "pending")
	log.Println(amount.ToInt().String())
}

// 列出所有账户
func (h *Handle) Accounts() {
	var accounts []string
	h.client.Call(&accounts, "eth_accounts")
	for _, v := range accounts {
		log.Println(v)
	}
}

// 查看区块数量
func (h *Handle) BlockNumber() {
	var num int
	h.client.Call(&num, "eth_blockNumber")
	log.Println(num)
}

// 查看当前客户端是否处于监听状态
func (h *Handle) Listening() {
	var ok bool
	h.client.Call(&ok,"net_listening")
	log.Println(ok)
}

// 查看某个地址code
func (h *Handle) Code() {
	var code string
	h.client.Call(&code, "eth_getCode", Account1, "latest")
	log.Println(code)
}

// 发送transaction转账
func (h *Handle) SendTransaction() {
	var (
		result string
		tx     transaction
	)

	tx.From = Miner
	tx.To   = Account2
	tx.Gas  = toHexBigInt(100000)
	tx.GasPrice = toHexBigInt(1)
	tx.Value = toHexBigInt(1000000000)

	if err := h.client.Call(&result, "eth_sendTransaction", &tx); err != nil {
		panic(err)
	}
	log.Println(result)
}

// 查询transaction数量，注意，这里count一定是string
func (h *Handle) TransactionCount() {
	var count string

	if err := h.client.Call(&count, "eth_getTransactionCount", Miner, "latest"); err != nil {
		panic(err)
	}

	log.Println(count)
}

// 根据number查询block
func (h *Handle) GetBlockByNumber() {
	var b block

	// 这里注意，查询时number中的0x不能少
	if err := h.client.Call(&b, "eth_getBlockByNumber", "0x1", true); err != nil {
		panic(err)
	}

	log.Println(b)
}

// 根据hash查询block
func (h *Handle) GetBlockByHash() {
	var b block

	hash := "0x328075d039e42c5f7b2823bfb2d8334843cc4fbe35c0d4021f6239f03d9d526a"
	if err := h.client.Call(&b, "eth_getBlockByHash", hash, true); err != nil {
		panic(err)
	}

	log.Println(b)
}

// 发送签名后transaction
func (h *Handle) SendRawTransaction() {

}

////////////////////////////////////////////////////////////////////////////
//
// 内部使用方法
//
////////////////////////////////////////////////////////////////////////////
func toHexBigInt(src int) *hexutil.Big {
	var ret hexutil.Big

	str := strconv.FormatInt(int64(src), 16)
	// 这里注意，一定要加上"0x"
	des :=  "0x" + str
	ret.UnmarshalText([]byte(des))

	return &ret
}

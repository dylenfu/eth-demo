package main

import (
	"github.com/dylenfu/eth-libs/client/ethsrc/rpc"
	"log"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"flag"
	"reflect"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"github.com/ethereum/go-ethereum/rlp"
	. "github.com/dylenfu/eth-libs/params"
	tp "github.com/dylenfu/eth-libs/types"
)

var call = flag.String("call", "Balance", "chose test case")

type Handle struct {
	client *rpc.Client
}

// transaction数据结构参考 https://github.com/ethereum/wiki/wiki/JSON-RPC#eth_sendtransaction
// 也可以直接使用go-eth本身的数据结构，如果是自己构造的话，需要保证能被正确解析(主要是hexutil.Big)
type transaction struct {
	From		string
	To 			string
	Gas			hexutil.Big
	GasPrice	hexutil.Big
	Value       hexutil.Big
	Data		string
}

// block数据结构参考 https://github.com/ethereum/wiki/wiki/JSON-RPC#eth_getblockbyhash
type block struct {
	Number 			hexutil.Big
	Hash 			string
	ParentHash 		string
	Nonce 			string
	Sha3Uncles 		string
	LogsBloom 		string
	TransactionRoot string
	StateRoot 		string
	Miner 			string
	Difficulty 		hexutil.Big
	TotalDifficulty hexutil.Big
	ExtraData 		string
	Size 			hexutil.Big
	GasLimit 		hexutil.Big
	GasUsed 		hexutil.Big
	Timestamp 		string
	Transactions 	[]transaction
	uncles 			[]string
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
	var num string
	if err := h.client.Call(&num, "eth_blockNumber"); err != nil {
		panic(err)
	}
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
	tx.Gas  = tp.Int2HexBigInt(100000)
	tx.GasPrice = tp.Int2HexBigInt(1)
	tx.Value = tp.Int2HexBigInt(1000000000)

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
	if err := h.client.Call(&b, "eth_getBlockByNumber", "0x10", true); err != nil {
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

// 签名用户
func (h *Handle) signUser(account,pwd string) (string, error) {
	var result string

	if err := h.client.Call(&result, "eth_sign", account, pwd); err != nil {
		return "", err
	}

	return result, nil
}

// 签名transaction
func (h *Handle) SignTransaction() {
	tx := newTransaction()
	codes, err := rlp.EncodeToBytes(tx)
	if err != nil {
		panic(err)
	}
	hash := common.BytesToHash(codes).String()
	log.Println("rlp hash:", hash)

	var result string
	if err := h.client.Call(&result, "eth_sign", Miner, hash); err != nil {
		panic(err)
	}
	log.Println("sign result:", result)
}

////////////////////////////////////////////////////////////////////////////
//
// 内部使用方法
//
////////////////////////////////////////////////////////////////////////////
func newTransaction() *types.Transaction {
	nonce := 170
	to := common.HexToAddress(Account1)
	amount := big.NewInt(100000000)
	gas := big.NewInt(10000)
	price := big.NewInt(1)
	data := []byte{}
	tx := types.NewTransaction(uint64(nonce), to, amount,gas,price, data)

	return tx
}
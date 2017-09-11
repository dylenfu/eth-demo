package main

import (
	"github.com/dylenfu/eth-libs/client/ethsrc/rpc"
	"log"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"flag"
	//"reflect"
	"reflect"
)

var call = flag.String("call", "Balance", "chose test case")

type Handle struct {
	client *rpc.Client
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
	h.client.Call(&amount, "eth_getBalance", "0xc30ae0bbb7722b1af3ada4db98d86090f8850cdb", "pending")
	log.Println(amount.ToInt().String())
}

func (h *Handle) Accounts() {
	var accounts []string
	h.client.Call(&accounts, "eth_accounts")
	for _, v := range accounts {
		log.Println(v)
	}
}

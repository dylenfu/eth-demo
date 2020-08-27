package main

import (
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/ipfs/go-log"
)

var (
	logger = log.Logger("geth")
	One, _ = new(big.Int).SetString("1000000000000000000", 10)
)

const (
	NetworkID = 10
	GasLimit  = 21000
	GasPrice  = 0
	Nonce     = 5 // increase by 1 for each test

	RpcURL       = "http://localhost:22000"
	Keystore     = `/Users/dylen/software/quorum/node0/data/keystore/`
	KeyFile      = Keystore + `UTC--2020-08-26T09-55-25.770341000Z--57a259e0bcd61dffdd205a5cd046be9068e832dd`
	Passphase    = `111111`
	AdminAddress = `0x57A259e0BcD61dFfdd205a5Cd046be9068E832dd`
	TestAddress  = "0x8409f65FD78a03edd654671a9d15c6E9962C07c9"
)

func init() {
	if err := log.SetLogLevel("*", "DEBUG"); err != nil {
		panic(fmt.Sprintf("failed to initialize logger: [%v]", err))
	}
}

func main() {
	client := getClient()

	srcBalanceBeforeTransfer := client.Balance(AdminAddress)
	logger.Infof("src Balance before transfer %s", srcBalanceBeforeTransfer.String())

	tx := transferETH(Nonce, TestAddress, One)
	signedTx := client.SignTransaction(tx)
	hash := client.SendRawTransaction(signedTx)
	logger.Infof("transfer result %s", hash.Hex())

	time.Sleep(10 * time.Second)

	srcBalanceAfterTransfer := client.Balance(AdminAddress)
	dstBalanceAfterTransfer := client.Balance(TestAddress)
	logger.Infof("src balance after transfer %s", srcBalanceAfterTransfer.String())
	logger.Infof("dest balance after transfer %s", dstBalanceAfterTransfer.String())
}

type wrapClient struct {
	*rpc.Client
	key *keystore.Key
}

func (c *wrapClient) Balance(address string) *big.Int {
	var raw string

	if err := c.Call(
		&raw,
		"eth_getBalance",
		address,
		"latest",
	); err != nil {
		logger.Fatal(err)
	}

	src, _ := new(big.Int).SetString(raw, 0)
	data := new(big.Int).Div(src, One)

	return data
}

func (c *wrapClient) SignTransaction(tx *types.Transaction) string {

	signer := types.HomesteadSigner{}
	signedTx, err := types.SignTx(
		tx,
		signer,
		c.key.PrivateKey,
	)
	if err != nil {
		logger.Fatalf("failed to sign tx: [%v]", err)
	}

	bz, err := rlp.EncodeToBytes(signedTx)
	if err != nil {
		logger.Fatalf("failed to rlp encode bytes: [%v]", err)
	}
	return "0x" + hex.EncodeToString(bz)
}

func (c *wrapClient) SendRawTransaction(signedTx string) common.Hash {
	var result common.Hash
	if err := c.Client.Call(&result, "eth_sendRawTransaction", signedTx); err != nil {
		logger.Fatalf("failed to send raw transaction: [%v]", err)
	}

	return result
}

func transferETH(nonce uint64, toAddress string, value *big.Int) *types.Transaction {
	return types.NewTransaction(
		nonce,
		common.HexToAddress(toAddress),
		value,
		GasLimit,
		big.NewInt(GasPrice),
		nil,
	)
}

func getClient() *wrapClient {
	client, err := rpc.Dial(RpcURL)
	if err != nil {
		panic(err)
	}

	keyJson, err := ioutil.ReadFile(KeyFile)
	if err != nil {
		logger.Fatalf("failed to read file: [%v]", err)
	}

	key, err := keystore.DecryptKey(keyJson, Passphase)
	if err != nil {
		logger.Fatalf("failed to decrypt keyjson: [%v]", err)
	}

	return &wrapClient{
		Client: client,
		key:    key,
	}
}

//func (c *wrapClient) Nonce(address string) uint64 {
//	var raw string
//
//	if err := c.Call(
//		&raw,
//		"eth_getTransactionCount",
//		address,
//		"latest",
//	); err != nil {
//		logger.Fatal(err)
//	}
//
//	data, ok := new(big.Int).SetString(raw, 16)
//	if !ok {
//		return big.NewInt(0).Uint64()
//	} else {
//		return data.Uint64()
//	}
//}


func unlock() {
	// unlock account
	//ks := keystore.NewKeyStore(
	//	Keystore,
	//	keystore.LightScryptN,
	//	keystore.LightScryptP,
	//	)
	//
	//account, err := ks.Find(accounts.Account{
	//	Address: common.HexToAddress(AdminAddress),
	//})
	//if err != nil {
	//	logger.Fatalf("failed to find account in keystore: [%v]", err)
	//}
	//if err := ks.Unlock(account, Passphase); err != nil {
	//	logger.Fatalf("failed to unlock account: [%v]", err)
	//}
}

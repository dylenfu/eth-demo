package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/ipfs/go-log"
	"math/big"
)

var logger = log.Logger("geth")

const (
	Keystore     = `/Users/dylen/software/quorum/node1/keystore/`
	KeyFile      = Keystore + `UTC--2020-08-25T03-33-39.871331000Z--c2e3d3c40b70e6b95feee3f7b10e604ea967db4d`
	Passphase    = `111111`
	AdminAddress = `0xc2E3d3c40B70E6b95fEeE3f7b10E604Ea967Db4d`
	NetworkID    = 10
)

func init() {
	if err := log.SetLogLevel("*", "DEBUG"); err != nil {
		panic(fmt.Sprintf("failed to initialize logger: [%v]", err))
	}
}

func main() {
	client := getClient()

	balance := client.balance(AdminAddress)
	logger.Infof("admin balance %s", balance.String())
}

type wrapClient struct {
	*rpc.Client
}

func (c *wrapClient) balance(address string) *big.Int {
	var amount string

	decimal, _ := new(big.Int).SetString("1000000000000000000", 0)

	if err := c.Call(
		&amount,
		"eth_getBalance",
		address,
		"latest",
	); err != nil {
		logger.Fatal(err)
	}

	src, _ := new(big.Int).SetString(amount, 0)
	data := new(big.Int).Div(src, decimal)

	return data
}

func getClient() *wrapClient {
	client, err := rpc.Dial("http://localhost:22000")
	if err != nil {
		panic(err)
	}

	return &wrapClient{
		Client: client,
	}
}

//
//func sendTx() {
//	// Init a keystore
//	ks := keystore.NewKeyStore(
//		Keystore,
//		keystore.LightScryptN,
//		keystore.LightScryptP)
//	fmt.Println()
//
//	// Create account definitions
//	fromAccDef := accounts.Account{
//		Address: common.HexToAddress(AdminAddress),
//	}
//
//	toAccDef := accounts.Account{
//		Address: common.HexToAddress(AdminAddress),
//	}
//
//	// Find the signing account
//	signAcc, err := ks.Find(fromAccDef)
//	if err != nil {
//		fmt.Println("account keystore find error:")
//		panic(err)
//	}
//	fmt.Printf("account found: signAcc.addr=%s; signAcc.url=%s\n", signAcc.Address.String(), signAcc.URL)
//	fmt.Println()
//
//	// Unlock the signing account
//	errUnlock := ks.Unlock(signAcc, Passphase)
//	if errUnlock != nil {
//		fmt.Println("account unlock error:")
//		panic(err)
//	}
//	fmt.Printf("account unlocked: signAcc.addr=%s; signAcc.url=%s\n", signAcc.Address.String(), signAcc.URL)
//	fmt.Println()
//
//	// Construct the transaction
//	tx := types.NewTransaction(
//		0x0,
//		toAccDef.Address,
//		new(big.Int),
//		0,
//		new(big.Int),
//		[]byte(`cooldatahere`))
//
//	// Open the account key file
//	keyJson, readErr := ioutil.ReadFile(KeyFile)
//	if readErr != nil {
//		fmt.Println("key json read error:")
//		panic(readErr)
//	}
//
//	// Get the private key
//	keyWrapper, keyErr := keystore.DecryptKey(keyJson, Passphase)
//	if keyErr != nil {
//		fmt.Println("key decrypt error:")
//		panic(keyErr)
//	}
//	fmt.Printf("key extracted: addr=%s", keyWrapper.Address.String())
//
//	// Define signer and chain id
//	// chainID := big.NewInt(CHAIN_ID)
//	// signer := types.NewEIP155Signer(chainID)
//	signer := types.HomesteadSigner{}
//
//	// Sign the transaction signature with the private key
//	signature, signatureErr := crypto.Sign(tx.Hash().Bytes(), keyWrapper.PrivateKey)
//	if signatureErr != nil {
//		fmt.Println("signature create error:")
//		panic(signatureErr)
//	}
//
//	signedTx, signErr := tx.WithSignature(signer, signature)
//	if signErr != nil {
//		fmt.Println("signer with signature error:")
//		panic(signErr)
//	}
//
//	// Connect the client
//	client, err := ethclient.Dial("http://localhost:8000") // 8000=geth RPC port
//	if err != nil {
//		fmt.Println("client connection error:")
//		panic(err)
//	}
//	fmt.Println("client connected")
//	fmt.Println()
//
//	// Send the transaction to the network
//	txErr := client.SendTransaction(context.Background(), signedTx)
//	if txErr != nil {
//		fmt.Println("send tx error:")
//		panic(txErr)
//	}
//
//	fmt.Printf("send success tx.hash=%s\n", signedTx.Hash().String())
//}

package types

import (
	"os"
	"io/ioutil"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"reflect"
	"github.com/ethereum/go-ethereum/rpc"
)

type TokenImpl struct {
	Client 	*rpc.Client
	Abi   	*abi.ABI
	Token   interface{}
	ContractFilePath string //ex:"github.com/dylenfu/eth-libs/contracts/transfer/abi.txt"
	TokenAddress string		//ex:""
	Url string    //ex:"http://127.0.0.1:8545"
}

// NewAbi params file
func NewContract(contractFilePath, tokenAddress, rpcUrl string, token interface{}) *TokenImpl {
	impl := &TokenImpl{}
	impl.TokenAddress = tokenAddress
	impl.ContractFilePath = contractFilePath
	impl.Url = rpcUrl
	impl.Token = token

	// connect eth rpc
	impl.Client = DialEthRpc(impl.Url)

	// load contract abi text
	impl.Abi = NewAbi(impl.ContractFilePath)

	// reflect abi method to token impl
	LoadContract(impl.Token, impl.Abi, impl.TokenAddress, impl.Client)

	return impl
}

func DialEthRpc(url string) *rpc.Client {
	client, err := rpc.Dial(url)
	if err != nil {
		panic(err)
	}

	return client
}

func NewAbi(contractFilePath string) *abi.ABI {
	tabi := &abi.ABI{}
	dir := os.Getenv("GOPATH")
	abiStr, err := ioutil.ReadFile(dir + "/src/" + contractFilePath)
	if err != nil {
		panic(err)
	}
	if err := tabi.UnmarshalJSON(abiStr); err != nil {
		panic(err)
	}

	return tabi
}

func LoadContract(token interface{}, abi *abi.ABI, tokenAddress string, client *rpc.Client) {
	elem := reflect.ValueOf(token).Elem()
	for i := 0; i < elem.NumField(); i++ {
		methodName := elem.Type().Field(i).Tag.Get("methodName")

		abiMethod := &AbiMethod{}
		abiMethod.Name = methodName
		abiMethod.Abi = abi
		abiMethod.Client = client
		abiMethod.Address = tokenAddress

		elem.Field(i).Set(reflect.ValueOf(*abiMethod))
	}
}
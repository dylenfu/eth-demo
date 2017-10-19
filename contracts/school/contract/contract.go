package contract

import (
	"github.com/dylenfu/eth-libs/types"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"github.com/ethereum/go-ethereum/common/hexutil"
	cm "github.com/dylenfu/eth-libs/common"
	"errors"
	"log"
)

var School *types.TokenImpl

const (
	AbiFilePath  = "github.com/dylenfu/eth-libs/contracts/school/abi.txt"
	EthRpcUrl    = "http://127.0.0.1:8545"
	TokenAddress = "0x0f24565fe3db5a337373f835b6b3c60865bc7c5c"
)

func init() {
	token := &SchoolImpl{}
	School = types.NewContract(AbiFilePath, TokenAddress, EthRpcUrl, token)
}

type SchoolImpl struct {
	Child   types.AbiMethod `methodName:"setChild"`
	Student types.AbiMethod `methodName:"setStudent"`
	Mates   types.AbiMethod `methodName:"setMates"`
	Class   types.AbiMethod `methodName:"setClass"`
	Grade   types.AbiMethod `methodName:"setGrade"`
}

type ChildEvent struct {
	Addresses []common.Address `alias:"addresses"`
}

type StudentEvent struct {
	AddressList [][2]common.Address `alias:"addressList"`
}

type MatesEvent struct {
	AddressList  [][2]common.Address `alias:"addressList"`
	UintArgsList [][7]*big.Int       `alias:"uintArgsList"`
}

type ClassEvent struct {
	AddressList  [][2]common.Address `alias:"addressList"`
	UintArgsList [][7]*big.Int       `alias:"uintArgsList"`
	VList        []uint8             `alias:"vList"`
}

type GradeEvent struct {
	AddressList  [][2]common.Address `alias:"addressList"`
	UintArgsList [][7]*big.Int       `alias:"uintArgsList"`
	VList        []uint8             `alias:"vList"`
	RList        [][]byte            `alias:"rList"`
}

func NewFilter(height int) (string, error) {
	var filterId string

	// 使用jsonrpc的方式调用newFilter
	filter := types.FilterReq{}
	filter.FromBlock = types.Int2BlockNumHex(int64(height))
	filter.ToBlock = "latest"
	filter.Address = School.TokenAddress

	err := School.Client.Call(&filterId, "eth_newFilter", &filter)
	if err != nil {
		return "", err
	}

	return filterId, nil
}

func GetEvent(filterId string) error {
	var logs []types.FilterLog

	impl := School
	err := School.Client.Call(&logs, "eth_getFilterChanges", filterId)
	if err != nil {
		return err
	}
	evts := impl.Abi.Events

	for _, v := range logs {
		println(v.Data)
		data := hexutil.MustDecode(v.Data)

		switch v.Topics[0] {
		case evts["ChildEvent"].Id().String():
			if err := showChild("ChildEvent", data, v.Topics); err != nil {
				return err
			}
		case evts["StudentEvent"].Id().String():
			if err := showStudent("StudentEvent", data, v.Topics); err != nil {
				return err
			}
		}
	}

	return nil
}

func showChild(eventName string, data []byte, topics []string) error {
	event, ok := School.Abi.Events[eventName]
	if !ok {
		return errors.New("child event do not exist")
	}

	evt := &ChildEvent{}
	if err := cm.UnpackEvent(event.Inputs, evt, data, topics); err != nil {
		return err
	}

	for _, v := range evt.Addresses {
		log.Println(v.Hex())
	}

	return nil
}

func showStudent(eventName string, data []byte, topics []string) error {
	impl := School
	event, ok := impl.Abi.Events[eventName]
	if !ok {
		return errors.New("student event do not exsit")
	}
	evt := &StudentEvent{}

	//
	if err := cm.UnpackEvent(event.Inputs, evt, []byte(data), topics); err != nil {
		return err
	}

	//for _, arr := range evt.AddressList {
	//	for _, v := range arr {
	//		log.Println(v.Hex())
	//	}
	//}
	log.Println("=====end=====")

	return nil
}

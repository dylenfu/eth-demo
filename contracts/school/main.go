package main

import (
	"flag"
	. "github.com/dylenfu/eth-libs/contracts/school/contract"
	"github.com/ethereum/go-ethereum/common"
	"reflect"
	"github.com/dylenfu/eth-libs/params"
	"qiniupkg.com/x/log.v7"
	"time"
)

type Handle struct{}

var fn = flag.String("fn", "Student", "chose function name")

func main() {
	flag.Parse()
	handle := &Handle{}
	reflect.ValueOf(handle).MethodByName(*fn).Call([]reflect.Value{})
}

const (
	addr1 = "0x0c0b638ffccb4bdc4c0d0d5fef062fc512c92511"
	addr2 = "0x96124db0972e3522a9b3910578b3f2e1a50159c7"
	addr3 = "0x86324df0972e3522a9b3910578b3f2e1a50132d5"
)

func (h *Handle) Semen() {
	impl := School.Token.(*SchoolImpl)
	var result string
	addresses := []common.Address{common.HexToAddress(addr1), common.HexToAddress(addr2), common.HexToAddress(addr3), common.HexToAddress(addr1)}
	if err := impl.Semen.SendTransaction(&result, addresses); err != nil {
		panic(err)
	}
}

func (h *Handle) Baby() {
	impl := School.Token.(*SchoolImpl)

	var result string
	addresses := [3]common.Address{common.HexToAddress(addr1), common.HexToAddress(addr2), common.HexToAddress(addr3)}
	if err := impl.Baby.SendTransaction(&result, addresses); err != nil {
		panic(err)
	}
}

func (h *Handle) Child() {
	impl := School.Token.(*SchoolImpl)

	var result string
	addrList := [][3]common.Address{
		[3]common.Address{common.HexToAddress(addr1), common.HexToAddress(addr2), common.HexToAddress(addr3)},
	}
	err := impl.Child.SendTransaction(&result, addrList)

	if err != nil {
		panic(err)
	}

}

func (h *Handle) Student() {
	impl := School.Token.(*SchoolImpl)

	var result string
	addrList := [][2]common.Address{
		[2]common.Address{common.HexToAddress(addr1), common.HexToAddress(addr2)},
		[2]common.Address{common.HexToAddress(addr2), common.HexToAddress(addr1)},
		[2]common.Address{common.HexToAddress(addr3), common.HexToAddress(addr2)},
	}
	err := impl.Student.SendTransaction(&result, addrList)

	if err != nil {
		panic(err)
	}
}

func (h *Handle) Mates() {

}

func (h *Handle) Class() {

}

func (h *Handle) Grade() {

}

func (h *Handle) ListenEvent() {
	h.Semen()
	time.Sleep(3 * time.Second)
	filterId, err := NewFilter(params.BlockNumber)
	if err != nil {
		panic(err)
	}

	for {
		err := GetEvent(filterId)
		if err != nil {
			log.Error(err.Error())
		}
	}
}
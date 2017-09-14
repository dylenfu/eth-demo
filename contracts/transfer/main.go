package main


// 该项目实验性地创建一个合约，包含订单信息及转账功能
// 实现合约及部署，通过abi调用方式实现转账操作
func main() {

}

////////////////////////////////////////////
//
// data structs
//
////////////////////////////////////////////
type Deposit struct {
	hash 		string
	account 	string
	amount 		int
}

type Order struct {
	hash		string
	accountS 	string
	accountB 	string
	amountS		int
	amountB		int
}

type OrderState struct {
	hash		string
	accountS 	string
	accountB 	string
	amountS 	int
	amountB 	int
	ok 			bool
}

////////////////////////////////////////////
//
// events
//
////////////////////////////////////////////
type DepositEvent struct {
	hash 		string
	account     string
	amount 		int
	ok 			bool
}

type OrderEvent struct {
	hash 		string
	accountS 	string
	accountB 	string
	amountS 	int
	amountB 	int
	ok 			bool
}

type BankToken struct {

}

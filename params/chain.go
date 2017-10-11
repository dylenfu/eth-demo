package params

// 在私链建立三个账户，miner是第一个账户，用于挖矿，account1&account2作为测试账户
// 每一次transaction之前都要解锁相关账户
const(
	ChainId 				= 151
	HelloTokenAddress		= "0x227F88083AE9eE717e39669CB2718E604833fEf9"
	TransferTokenAddress	= "0x227F88083AE9eE717e39669CB2718E604833fEf9"
	Miner       			= "0x4bad3053d574cd54513babe21db3f09bea1d387d"    // pwd 101
	Account1    			= "0x46c5683c754b2eba04b2701805617c0319a9b4dd"    // pwd 102
	Account2    			= "0x56d9620237fff8a6c0f98ec6829c137477887ec4"    // pwd 103
	PwdMiner    			= "101"
	PwdAccount1 			= "102"
	PwdAccount2 			= "103"
	TransferFilterId1 		= "0x619362f09bc1e5bb21e6f7850af3cd5f"
)
package main

import (
	"fmt"
	"geth-demo/utils/eth"
	"geth-demo/utils/net"
)

func main() {
	netWorkId, err := net.GetNetWorkId()
	if err != nil {
		panic(err)
	}
	fmt.Println(netWorkId)

	isListening, err := net.GetNetIsListening()
	if err != nil {
		panic(err)
	}
	fmt.Println(isListening)

	peerCount, err := net.GetNetPeerCount()
	if err != nil {
		panic(err)
	}
	fmt.Println(peerCount)

	accounts, err := eth.GetAccountList()
	if err != nil {
		panic(err)
	}
	fmt.Println(accounts)

	// account, err := personal.NewAccount("123")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(account)

	balance, err := eth.GetAccountBalance("0x9bc44f33c7c41970f8d4c6b805952a753fc67e84", "latest")
	if err != nil {
		panic(err)
	}
	fmt.Println(balance)

	gasPrice, err := eth.GetGasPrice()
	if err != nil {
		panic(err)
	}
	fmt.Println(gasPrice)

	coinbase, err := eth.GetCoinbase()
	if err != nil {
		panic(err)
	}
	fmt.Println(coinbase)

	isMining, err := eth.GetIsMining()
	if err != nil {
		panic(err)
	}
	fmt.Println(isMining)

	hashRate, err := eth.GetHashRate()
	if err != nil {
		panic(err)
	}
	fmt.Println(hashRate)

	transactionCount, err := eth.GetTransactionCount("0x9bc44f33c7c41970f8d4c6b805952a753fc67e84", "latest")
	if err != nil {
		panic(err)
	}
	fmt.Println(transactionCount)

	blockNum, err := eth.GetBlockNumber()
	if err != nil {
		panic(err)
	}
	fmt.Println(blockNum)

	// isLock, err := personal.LockAccount("0x9bc44f33c7c41970f8d4c6b805952a753fc67e84")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(isLock)

	// isUnlock, err := personal.UnlockAccount("0x9bc44f33c7c41970f8d4c6b805952a753fc67e84", "123")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(isUnlock)
}

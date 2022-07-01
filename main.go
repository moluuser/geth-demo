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
}

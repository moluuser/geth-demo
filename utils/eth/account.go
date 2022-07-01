package eth

import (
	"errors"
	"geth-demo/utils"
)

// GetAccountList 获取账户地址列表
func GetAccountList() (accounts []string, err error) {
	err = utils.Client.Call(&accounts, "eth_accounts")
	return
}

// GetAccountBalance 获取指定账户余额
func GetAccountBalance(account string, quantity string) (balance string, err error) {
	switch quantity {
	case "latest":
	case "earliest":
	case "pending":
	default:
		err = errors.New("quantity error")
		return
	}
	err = utils.Client.Call(&balance, "eth_getBalance", account, quantity)
	return
}

// GetGasPrice 获取Gas价格，单位wei
func GetGasPrice() (gasPrice string, err error) {
	err = utils.Client.Call(&gasPrice, "eth_gasPrice")
	return
}

// GetCoinbase 获取挖矿账户地址
func GetCoinbase() (coinbase string, err error) {
	err = utils.Client.Call(&coinbase, "eth_coinbase")
	return
}

// Deprecated: 获取当前以太坊协议的版本
func GetProtocolVersion() (protocolVersion string, err error) {
	err = utils.Client.Call(&protocolVersion, "eth_protocolVersion")
	return
}

// GetIsMining 客户端是否在挖矿中
func GetIsMining() (isMining bool, err error) {
	err = utils.Client.Call(&isMining, "eth_mining")
	return
}

// GetHashRate 获取节点挖矿时每秒可算出的哈希数量
func GetHashRate() (hashRate string, err error) {
	err = utils.Client.Call(&hashRate, "eth_hashrate")
	return
}

// GetTransactionCount 获取指定地址发生的交易数量
func GetTransactionCount(account string, quantity string) (transactionCount string, err error) {
	switch quantity {
	case "latest":
	case "earliest":
	case "pending":
	default:
		err = errors.New("quantity error")
		return
	}
	err = utils.Client.Call(&transactionCount, "eth_getTransactionCount", account, quantity)
	return
}

// GetBlockNumber 获取节点当前块编号
func GetBlockNumber() (blockNum string, err error) {
	err = utils.Client.Call(&blockNum, "eth_blockNumber")
	return
}

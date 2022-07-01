package eth

import "geth-demo/utils"

// GetAccountList 获取账户地址列表
func GetAccountList() (accounts []string, err error) {
	err = utils.Client.Call(&accounts, "eth_accounts")
	return
}

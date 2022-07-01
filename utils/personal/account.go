package personal

import "geth-demo/utils"

// GetAccountList 获取账户地址列表
func GetAccountList() (accounts []string, err error) {
	err = utils.Client.Call(&accounts, "personal_listAccounts")
	return
}

func NewAccount(pwd string) (account string, err error) {
	err = utils.Client.Call(&account, "personal_newAccount", pwd)
	return
}

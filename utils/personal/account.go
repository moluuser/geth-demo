package personal

import "geth-demo/utils"

// GetAccountList 获取账户地址列表
func GetAccountList() (accounts []string, err error) {
	err = utils.Client.Call(&accounts, "personal_listAccounts")
	return
}

// NewAccount 创建账户
func NewAccount(pwd string) (account string, err error) {
	err = utils.Client.Call(&account, "personal_newAccount", pwd)
	return
}

// LockAccount 锁定账户
func LockAccount(account string) (isLock bool, err error) {
	err = utils.Client.Call(&isLock, "personal_lockAccount", account)
	return
}

// UnlockAccount 解锁账户
func UnlockAccount(account string, pwd string) (isUnlock bool, err error) {
	err = utils.Client.Call(&isUnlock, "personal_unlockAccount", account, pwd)
	return
}

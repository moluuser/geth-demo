package net

import "geth-demo/utils"

// GetNetWorkId 获取当前连接网络的ID
func GetNetWorkId() (netWorkId string, err error) {
	err = utils.Client.Call(&netWorkId, "net_version")
	return
}

// GetNetIsListening 获取客户端是否处于监听状态
func GetNetIsListening() (isListening bool, err error) {
	err = utils.Client.Call(&isListening, "net_listening")
	return
}

// GetNetPeerCount 获取所连接对端节点旳数量
func GetNetPeerCount() (peerCount string, err error) {
	err = utils.Client.Call(&peerCount, "net_peerCount")
	return
}

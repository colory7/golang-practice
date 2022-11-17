package network3_demo

import "testing"

func TestTcpServer(t *testing.T) {
	startTcpServer()
}

func TestTcpClient(t *testing.T) {
	startTcpClient()
}

func TestUdpServer(t *testing.T) {
	startUdpServer()
}

func TestUdpClient(t *testing.T) {
	startUdpClient()
}




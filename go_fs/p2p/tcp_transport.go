package p2p

import "net"

type TCPtransport struct{
	listenAddress string
	listener net.Listener
	peers map[string]	
}
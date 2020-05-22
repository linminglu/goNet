package main

import (
	"github.com/Quantumoffices/goNet"
	_ "github.com/Quantumoffices/goNet/codec/json"
	_ "github.com/Quantumoffices/goNet/peer/ws"
)

func main() {
	p := goNet.NewPeer(
		goNet.Options{
			Addr:          "ws://192.168.0.125:8080/center/ws",
			PeerType:      goNet.PEERTYPE_SERVER,
			ReadDeadline:  0,
			WriteDeadline: 0,
			PoolSize:      0,
			PanicHandler:  nil,
			AllowMaxConn:  0,
		})
	p.Start()
}

package main

import (
	"github.com/Quantumoffices/goNet"
	_ "github.com/Quantumoffices/goNet/codec/json"
	_ "github.com/Quantumoffices/goNet/peer/tcp"
)

func main() {
	p := goNet.NewPeer(
		goNet.WithPeerType(goNet.PEERTYPE_SERVER),
		goNet.WithAddr(":8087"),
	)
	p.Start()
}

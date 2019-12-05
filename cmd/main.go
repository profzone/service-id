package main

import (
	"github.com/henrylee2cn/erpc/v6"
	"github.com/henrylee2cn/erpc/v6/codec"
	"github.com/profzone/eden-framework/pkg/application"
	"github.com/profzone/service-id/internal/global"
	"github.com/profzone/service-id/internal/routers"
	"time"
)

func main() {
	app := application.NewApplication(runner, &global.Config)
	app.Start()
}

func runner(app *application.Application) error {
	peer := erpc.NewPeer(erpc.PeerConfig{
		ListenPort:       9090,
		DialTimeout:      10 * time.Second,
		RedialTimes:      10,
		RedialInterval:   10 * time.Second,
		DefaultBodyCodec: codec.NAME_PROTOBUF,
		PrintDetail:      true,
		CountTime:        true,
	})
	routers.InitRouters(peer)
	return peer.ListenAndServe()
}

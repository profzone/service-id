package routers

import (
	"github.com/henrylee2cn/erpc/v6"
	"github.com/profzone/service-id/internal/routers/id"
)

var CallRouters = make([]interface{}, 0)
var PushRouters = make([]interface{}, 0)

func init() {
	CallRouters = append(CallRouters, &id.ID{})
}

func InitRouters(peer erpc.Peer) {
	for _, router := range CallRouters {
		peer.RouteCall(router)
	}
	for _, router := range PushRouters {
		peer.RoutePush(router)
	}
}

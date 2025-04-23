package tagged

import (
	"context"

	"github.com/b49nd1n/xray-core/common/net"
	"github.com/b49nd1n/xray-core/features/routing"
)

type DialFunc func(ctx context.Context, dispatcher routing.Dispatcher, dest net.Destination, tag string) (net.Conn, error)

var Dialer DialFunc

package tcp

import (
	"github.com/b49nd1n/xray-core/common"
	"github.com/b49nd1n/xray-core/transport/internet"
)

func init() {
	common.Must(internet.RegisterProtocolConfigCreator(protocolName, func() interface{} {
		return new(Config)
	}))
}

package vmess_test

import (
	"testing"

	"github.com/b49nd1n/xray-core/common"
	"github.com/b49nd1n/xray-core/common/protocol"
	"github.com/b49nd1n/xray-core/common/uuid"
	. "github.com/b49nd1n/xray-core/proxy/vmess"
)

func toAccount(a *Account) protocol.Account {
	account, err := a.AsAccount()
	common.Must(err)
	return account
}

func BenchmarkUserValidator(b *testing.B) {
	for i := 0; i < b.N; i++ {
		v := NewTimedUserValidator()

		for j := 0; j < 1500; j++ {
			id := uuid.New()
			v.Add(&protocol.MemoryUser{
				Email: "test",
				Account: toAccount(&Account{
					Id: id.String(),
				}),
			})
		}

		common.Close(v)
	}
}

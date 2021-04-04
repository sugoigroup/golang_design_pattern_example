package proxy

import (
	"testing"
)

func TestFactory(t *testing.T) {
	proxyPattern := proxyPattern{
		proxy: &proxyService{},
	}
	if proxyPattern.run() != "connected" {
		t.Error("proxy 연결안됨")
	}
}


package ip

import (
	"testing"
)

func TestGetLocalIPv4(t *testing.T) {
	ipv4 := GetLocalIPv4()
	if ipv4 == "" {
		t.Errorf("GetLocalIPv4() = %q", ipv4)
	} else {
		t.Logf("GetLocalIPv4() = %q.", ipv4)
	}
}

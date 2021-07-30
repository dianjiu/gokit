package random

import "testing"

func TestGetLocalIPv4(t *testing.T) {
	str := GenerateRandomString(6)
	if str == "" {
		t.Errorf("GetLocalIPv4() = %q", str)
	} else {
		t.Logf("GetLocalIPv4() = %q.", str)
	}
}

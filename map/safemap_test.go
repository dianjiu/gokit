package _map

import (
	"testing"
)

func Test_beemap(t *testing.T) {
	bm := NewBeeMap()
	if !bm.Set("zhaocloud", 1) {
		t.Error("set Error")
	}
	if !bm.Check("zhaocloud") {
		t.Error("check err")
	}

	if v := bm.Get("zhaocloud"); v.(int) != 1 {
		t.Error("get err")
	}

	bm.Delete("zhaocloud")
	if bm.Check("zhaocloud") {
		t.Error("delete err")
	}
}

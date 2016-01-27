package server

import "testing"

func TestDummy(t *testing.T) {
    var b bool = false
    if b {
        t.Errorf("failed %b", b)
    }
}


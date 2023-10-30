package double_check

import (
	"testing"
)

func TestSafeMap_LoadAndStore(t *testing.T) {
	sm := NewSafeMap()
	t.Log(sm.LoadAndStore("1", 2))
	t.Log(sm.LoadAndStore("1", 2))
	t.Log(sm.LoadAndStore("1", 2))
	t.Log(sm.LoadAndStore("1", 2))
	t.Log(sm.LoadAndStore("1", 2))
	t.Log(sm.LoadAndStore("1", 2))
}

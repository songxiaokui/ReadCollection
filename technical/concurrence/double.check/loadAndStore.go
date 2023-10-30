package double_check

import "sync"

type SafeMap struct {
	mtx sync.RWMutex
	dm  map[string]any
}

func NewSafeMap() *SafeMap {
	return &SafeMap{
		dm: make(map[string]any),
	}
}

// LoadAndStore 存在直接返回 不存在设置新值
func (sm *SafeMap) LoadAndStore(key string, newValue any) (any, bool) {
	sm.mtx.RLock()
	data, ok := sm.dm[key]
	if ok {
		return data, false
	}
	sm.mtx.RUnlock()

	// double check
	// 没拿到 但是在设置之前可能会被设置值 写锁被抢 写之前再次检查
	sm.mtx.Lock()
	defer sm.mtx.Unlock()
	data, ok = sm.dm[key] // 注意: 上写锁后,此处获取读锁将会阻塞,也就是说不可再进行上读锁
	if ok {
		return data, false
	}
	sm.dm[key] = newValue
	return newValue, true

}

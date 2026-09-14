package bilibili

import "sync"

// Storage 是 WBI 密钥缓存的存储抽象。
type Storage interface {
	Set(key string, value any)
	Get(key string) (v any, isSet bool)
}

// MemoryStorage 是基于内存的 Storage 实现。
type MemoryStorage struct {
	data map[string]any
	mu   sync.RWMutex
}

// Set 设置值
func (impl *MemoryStorage) Set(key string, value any) {
	impl.mu.Lock()
	defer impl.mu.Unlock()

	if impl.data == nil {
		impl.data = make(map[string]any)
	}
	impl.data[key] = value
}

// Get 获取值, isSet 表示值是否存在
func (impl *MemoryStorage) Get(key string) (v any, isSet bool) {
	impl.mu.RLock()
	defer impl.mu.RUnlock()

	if v, isSet = impl.data[key]; isSet {
		return v, true
	}
	return nil, false
}

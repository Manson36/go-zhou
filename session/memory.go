package session

import (
	"errors"
	"sync"
)

type MemorySession struct {
	sessionId string
	//存kv
	data   map[string]interface{}
	rwLock sync.RWMutex
}

//构造函数
func NewMemorySession(id string) Session {
	return &MemorySession{
		sessionId: id,
		data:      make(map[string]interface{}, 16),
	}
}

func (m *MemorySession) Set(key string, value interface{}) error {
	//加锁
	m.rwLock.Lock()
	defer m.rwLock.Unlock()
	//设置值
	m.data[key] = value
	return nil
}

func (m *MemorySession) Get(key string) (interface{}, error) {
	//加锁
	m.rwLock.Lock()
	defer m.rwLock.Unlock()
	//获取值
	value, ok := m.data[key]
	if !ok {
		err := errors.New("key not exits in session")
		return value, err
	}
	return value, nil
}

func (m *MemorySession) Del(key string) error {
	//加锁
	m.rwLock.Lock()
	defer m.rwLock.Unlock()
	//删除
	delete(m.data, key)
	return nil
}

func (m *MemorySession) Save() error {
	return nil
}

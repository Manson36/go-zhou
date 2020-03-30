package session

import (
	uuid "github.com/satori/go.uuid"
	"sync"
)

//定义对象
type MemorySessionMgr struct {
	rwLock     sync.RWMutex
	sessionMap map[string]Session
}

func (m *MemorySessionMgr) Init(addr string, options ...interface{}) (err error) {
	return
}

func (m *MemorySessionMgr) CreateSession() (session *MemorySession, err error) {
	m.rwLock.Lock()
	defer m.rwLock.Unlock()
	//go get github.com/satori/go.uuid
	//用uuid作为sessionId
	id := uuid.NewV4()
	sessionId := id.String()
	session = NewMemorySession(sessionId)
	return
}

func (m *MemorySessionMgr) Get(sessionId string) (session Session, err error) {
	panic("implement me")
	return
}

func NewMemorySessionMgr() *MemorySessionMgr {
	return &MemorySessionMgr{
		sessionMap: make(map[string]Session, 1024),
	}
}

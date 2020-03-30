package session

//定义管理者，管理所有Session
type SessionMgr interface {
	//初始化，先想到的是redis地址
	Init(addr string, options ...interface{}) (err error)
	CreateSession() (session Session, err error)
	Get(sessionId string) (session MemorySession, err error)
}

package Youpo

import (
	"github.com/Huabuxiu/Youpo/datastruct"
	"github.com/Jeffail/tunny"
	"sync"
)

//全局对象
var server *Server

type Server struct {
	process *Process

	processPool *tunny.Pool

	dbs []*DB

	clientList *datastruct.LinkedList

	clientLock sync.Mutex

	CloseChan <-chan struct{}
}

//初始化Server
func InitServer() *Server {
	dbs := make([]*DB, 16)
	dbs[0] = MakeDB(0)

	//执行器初始化
	singleProcess := MakeSingleProcess()
	processSinglePool := tunny.NewFunc(1, func(client interface{}) interface{} {
		args := client.(*Client).args
		db := client.(*Client).db
		return singleProcess.Exec(db, args)
	})

	//注册命令表
	StringInit()
	server = &Server{
		process:     singleProcess,
		processPool: processSinglePool,
		dbs:         dbs,
		CloseChan:   make(chan struct{}),
	}
	return server
}

func GetSever() *Server {
	return server
}

func (server *Server) GetProcess() *Process {
	return server.process
}

func (server *Server) GetDB() []*DB {
	return server.dbs
}

func (server *Server) SelectDb(client *Client, dbIndex int) Reply {
	if dbIndex < 0 || dbIndex > len(server.dbs) {
		return MakeErrorReply("dbIndex is out of range")
	}

	client.dbIndex = dbIndex
	client.db = server.dbs[dbIndex]
	return MakeOKReply()
}

func (server *Server) RegisterClient(client *Client) {
	server.clientLock.Lock()
	defer server.clientLock.Unlock()

	server.clientList.Add(client)
}

func (server *Server) RemoveClient(client *Client) {
	server.clientLock.Lock()
	defer server.clientLock.Unlock()

	server.clientList.RemoveNode(client)

}

func (server *Server) Exec(client *Client) Reply {
	//使用协程池来处理执行
	return server.processPool.Process(client).(Reply)
}

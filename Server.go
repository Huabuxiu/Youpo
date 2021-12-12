package Youpo

import (
	"github.com/Huabuxiu/Youpo/datastruct"
	"sync"
)

//全局对象
var server *Server

type Server struct {
	process *Process

	dbs []*DB

	clientList *datastruct.LinkedList

	clientLock sync.Mutex

	CloseChan <-chan struct{}
}

//初始化Server
func InitServer() *Server {
	dbs := make([]*DB, 16)
	dbs[0] = MakeDB(0)
	singleProcess := MakeSingleProcess()

	//注册命令表
	StringInit()
	server = &Server{
		process:   singleProcess,
		dbs:       dbs,
		CloseChan: make(chan struct{}),
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

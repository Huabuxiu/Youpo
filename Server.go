package Youpo

type Server struct {
	process *Process

	dbs []*DB
}

//初始化Server
func InitServer() *Server {
	dbs := make([]*DB, 16)
	dbs[0] = MakeDB(0)
	singleProcess := MakeSingleProcess()

	//注册命令表
	StringInit()
	return &Server{
		process: singleProcess,
		dbs:     dbs,
	}
}

func (receiver *Server) GetProcess() *Process {
	return receiver.process
}

func (receiver *Server) GetDB() []*DB {
	return receiver.dbs
}

func StringInit() {
	RegisterCommand("get", Get, "r", 2)
	RegisterCommand("set", Set, "w", 3)
}

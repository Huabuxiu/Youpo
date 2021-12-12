package networks

import (
	"fmt"
	"github.com/Huabuxiu/Youpo"
	"log"
	"net"
)

type netServer struct {
	ip string

	port int

	listener net.Listener
}

func StartNetServer(ip string, port int) *netServer {
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", ip, port))
	if err != nil {
		//log err
		return nil
	}
	n := &netServer{
		ip:       ip,
		port:     port,
		listener: listener,
	}
	return n
}

func (receiver *netServer) ListenAndProcess() {

	go func() {
		//监听关闭信息
		<-Youpo.GetSever().CloseChan
		//todo log stop tpc
		receiver.listener.Close()
	}()

	defer func() {
		//todo log stop tpc
		receiver.listener.Close()
	}()

	for {
		//建立连接
		con, err := receiver.listener.Accept()

		if err != nil {
			log.Fatal(fmt.Sprintf("accept err: %v", err))
		}

		//异步 处理 连接协议
		go HandleConn(con)
	}
}

func HandleConn(conn net.Conn) {

	//构造client

	client := Youpo.MakeClient(Youpo.GetSever().GetProcess(), conn)
	Youpo.GetSever().SelectDb(client, 0)
	Youpo.GetSever().RegisterClient(client)

	for {
		//阻塞读取

		//	解析协议

		//	提交执行到单线执行器

		//写结果到连接中
	}

}

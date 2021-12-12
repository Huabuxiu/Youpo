package networks

import (
	"fmt"
	"github.com/Huabuxiu/Youpo"
	"io"
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

	defer receiver.listener.Close()

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

	client := Youpo.MakeClient(conn)
	Youpo.GetSever().SelectDb(client, 0)
	Youpo.GetSever().RegisterClient(client)

	//读取来自client 的 Message
	for {
		//阻塞读取
		err := client.ReadMsg()
		if err == io.EOF {
			//移除client
			Youpo.GetSever().RemoveClient(client)
			continue
		}
		// TODO  执行超时

		//	提交执行到单线执行器
		result := Exec(client)

		//写结果到连接中
		_ = client.Write(result)
	}

}

func Exec(client *Youpo.Client) Youpo.Reply {
	return Youpo.GetSever().Exec(client)
}

package Youpo

import (
	"net"
	"time"
)

type Client struct {
	//db 指针
	db *DB

	//选择的db 索引
	dbIndex int

	//执行器
	process *Process

	//tcp 连接
	connection net.Conn

	args []string

	argsNum int

	lastAliveTime time.Time
}

func MakeClient(process *Process, connection net.Conn) *Client {
	client := &Client{
		connection:    connection,
		process:       process,
		lastAliveTime: time.Now(),
	}
	return client
}

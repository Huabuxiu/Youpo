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

	//tcp 连接
	connection net.Conn

	args []string

	argsNum int

	lastAliveTime time.Time
}

func MakeClient(connection net.Conn) *Client {
	client := &Client{
		connection:    connection,
		lastAliveTime: time.Now(),
	}
	return client
}

func (client *Client) ReadMsg() error {

	//	解析协议
	return nil
}

func (client *Client) Write(reply Reply) error {
	return nil
}

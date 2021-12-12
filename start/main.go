package main

import (
	"github.com/Huabuxiu/Youpo"
	"github.com/Huabuxiu/Youpo/networks"
)

var config = &Config{
	Ip:   "127.0.0.1",
	Port: 6379,
}

type Config struct {
	Ip   string
	Port int
}

func main() {

	Youpo.InitServer()

	//todo 建立time ea

	networks.StartNetServer(config.Ip, config.Port).ListenAndProcess()
}

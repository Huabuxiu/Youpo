package Youpo

import (
	"github.com/Huabuxiu/Youpo/networks"
)

func Get(db *DB, args []string) networks.Reply {
	key, exit := db.GetObjectByKey(args[0])
	if !exit {
		return networks.EmptyReply{}
	}
	return networks.MakeStringReply(key.(string))
}

func Set(db *DB, args []string) networks.Reply {
	db.PutObject(args[0], args[1])
	return networks.MakeOKReply()
}

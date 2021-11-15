package Youpo

import (
	"github.com/Huabuxiu/Youpo/networks"
	"strconv"
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

func GetRange(db *DB, args []string) networks.Reply {
	key, exit := db.GetObjectByKey(args[0])
	if !exit {
		return networks.EmptyReply{}
	}
	start, err := strconv.Atoi(args[0])

	if err != nil {
		return networks.MakeErrorReply("ERR value is not an integer or out of range")
	}
	end, err := strconv.Atoi(args[1])

	if err != nil {
		return networks.MakeErrorReply("ERR value is not an integer or out of range")
	}

	runes := []rune(key.(string))

	if end > len(runes) {
		end = len(runes)
	}

	if end < 0 {
		end = len(runes) + end
	}

	if start < 0 {
		start = len(runes) + start
	}

	return networks.MakeStringReply(string(runes[start:end]))

}

func Append(db *DB, args []string) networks.Reply {
	key, exit := db.GetObjectByKey(args[0])
	if exit {
		args[1] = key.(string) + args[1]
	}
	Set(db, args)
	return networks.MakeStringReply(strconv.Itoa(len(args[1])))
}

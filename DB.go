package Youpo

import (
	"github.com/Huabuxiu/Youpo/datastruct"
)

type DB struct {
	index int

	dataMap datastruct.HashMap
}

func MakeDB(index int) *DB {
	return &DB{
		index:   index,
		dataMap: *datastruct.MakeMap(),
	}
}

func (db *DB) GetObjectByKey(key string) (interface{}, bool) {
	return db.dataMap.Get(key)

	//TODO	 过期判断
}

func (db *DB) PutObject(key string, value interface{}) bool {
	db.dataMap.Put(key, value)
	return true
	//	 过期判断
}

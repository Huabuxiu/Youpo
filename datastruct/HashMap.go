package datastruct

import "github.com/Huabuxiu/Youpo/define"

type HashMap struct {
	goMap map[string]interface{}
}

func MakeMap() *HashMap {
	return &HashMap{
		goMap: make(map[string]interface{}),
	}
}

func (hashmap *HashMap) checkInit() {
	if hashmap.goMap == nil {
		panic("map is not init")
	}
}

func (hashmap *HashMap) Size() int {
	hashmap.checkInit()
	return len(hashmap.goMap)
}

func (hashmap *HashMap) IsEmpty() bool {
	return hashmap.goMap == nil || hashmap.Size() == 0
}

func (hashmap *HashMap) IsNotEmpty() bool {
	return !hashmap.IsEmpty()
}

func (hashmap *HashMap) Get(key string) (value interface{}, exist bool) {
	hashmap.checkInit()
	val, ok := hashmap.goMap[key]
	return val, ok
}

func (hashmap *HashMap) ContainsKey(key string) bool {
	hashmap.checkInit()
	_, exist := hashmap.Get(key)

	if exist {
		return true
	}
	return false
}

//put key 相同覆盖
func (hashmap *HashMap) Put(key string, value interface{}) interface{} {
	hashmap.checkInit()
	hashmap.goMap[key] = value
	return value
}

func (hashmap *HashMap) PutAll(argMap *HashMap) {
	hashmap.checkInit()
	if argMap == nil || argMap.goMap == nil {
		return
	}
	for keyItem, valueItem := range argMap.goMap {
		hashmap.Put(keyItem, valueItem)
	}
}

func (hashmap *HashMap) Remove(key string) interface{} {
	hashmap.checkInit()
	get, exist := hashmap.Get(key)
	if !exist {
		return nil
	}
	delete(hashmap.goMap, key)
	return get
}

func (hashmap *HashMap) Keys() []string {
	hashmap.checkInit()
	keys := make([]string, hashmap.Size())
	i := 0
	for key, _ := range hashmap.goMap {
		keys[i] = key
		i++
	}
	return keys
}

func (hashmap *HashMap) Values() []interface{} {
	hashmap.checkInit()
	values := make([]interface{}, hashmap.Size())
	i := 0
	for _, value := range hashmap.goMap {
		values[i] = value
		i++
	}
	return values
}

func (hashmap *HashMap) ForEach(function define.FunctionMultipleInterface) {
	hashmap.checkInit()
	if function == nil {
		return
	}

	for key, value := range hashmap.goMap {
		function(key, value)
	}

}

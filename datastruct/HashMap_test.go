package datastruct

import (
	"fmt"
	"testing"
)

func makeHashMap() *HashMap {
	makeMap := MakeMap()
	for i := 0; i < 10; i++ {
		makeMap.Put(fmt.Sprint(i), i)
	}
	return makeMap
}

func printMap(hashmap *HashMap) {
	hashmap.ForEach(func(key string, value interface{}) interface{} {
		println("key: ", key, ", value: ", value.(int))
		return value
	})
}

func TestHashMap_Size(t *testing.T) {
	hashMap := makeHashMap()
	println(hashMap.Size())
	printMap(hashMap)
}

func TestHashMap_IsEmpty(t *testing.T) {
	hashMap := makeHashMap()
	println(hashMap.IsEmpty())
	println(hashMap.IsNotEmpty())
}

func TestHashMap_Get(t *testing.T) {
	hashMap := makeHashMap()
	for i := 0; i < 10; i++ {
		get, exist := hashMap.Get(fmt.Sprint(i))
		if exist {
			println(get.(int))
		}
	}

	get, exist := hashMap.Get("hjhakjdhsa")
	if exist {
		println(get)
	} else {
		println("key not exit")
	}
}

func TestHashMap_ContainsKey(t *testing.T) {
	hashMap := makeHashMap()
	for i := 0; i < 10; i++ {
		exist := hashMap.ContainsKey(fmt.Sprint(i))
		if exist {
			println("key ", i, " exit")
		} else {
			println("key not exit")
		}
	}

	exist := hashMap.ContainsKey("")
	if exist {
		println("key hjhakjdhsa exit")
	} else {
		println("key not exit")
	}

}

func TestHashMap_Put(t *testing.T) {
	hashMap := makeHashMap()
	hashMap.Put("1", 812739812)
	printMap(hashMap)
}

func TestHashMap_PutAll(t *testing.T) {
	hashMap := makeHashMap()
	hashMap.Put("1", 812739812)

	makeMap := MakeMap()
	makeMap.PutAll(hashMap)
	printMap(makeMap)
}

func TestHashMap_Keys(t *testing.T) {
	hashMap := makeHashMap()
	hashMap.Put("1", 812739812)

	for _, key := range hashMap.Keys() {
		println(key)
	}
}

func TestHashMap_Values(t *testing.T) {
	hashMap := makeHashMap()
	hashMap.Put("1", 812739812)

	for _, key := range hashMap.Values() {
		println(key.(int))
	}
}

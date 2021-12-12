package Youpo

import (
	"github.com/Jeffail/tunny"
	"math/rand"
	"reflect"
	"sync"
	"testing"
	"time"
)

func makeInit() (*Process, []*DB) {
	server := InitServer()

	process := server.GetProcess()
	db := server.GetDB()
	return process, db
}

func TestString_get(t *testing.T) {
	process, db := makeInit()

	argsSet := []string{"set", "test", "1234"}
	argsGet := []string{"get", "test"}

	process.Exec(db[0], argsSet)

	exec := process.Exec(db[0], argsGet)

	println(string(exec.ToBytes()))
}

func TestSet(t *testing.T) {
	type args struct {
		db   *DB
		args []string
	}

	type caseType struct {
		name string
		args args
		want Reply
	}
	var db = MakeDB(0)

	var tests = []caseType{
		// TODO: Add test cases.
		{
			name: "setTest",
			args: args{
				db:   db,
				args: []string{"test", "1234"},
			},
			want: MakeOKReply(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Set(tt.args.db, tt.args.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Set() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGet(t *testing.T) {
	type args struct {
		db   *DB
		args []string
	}
	var db = MakeDB(0)

	tests := []struct {
		name string
		args args
		want Reply
	}{
		// TODO: Add test cases.
		{
			name: "getTest",
			args: args{
				db:   db,
				args: []string{"test"},
			},
			want: MakeStringReply("1234"),
		},
	}
	Set(db, []string{"test", "1234"})

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Get(tt.args.db, tt.args.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAppend(t *testing.T) {
	type args struct {
		db   *DB
		args []string
	}
	var db = MakeDB(0)
	tests := []struct {
		name string
		args args
		want Reply
	}{
		{
			name: "appendNil",
			args: args{
				db:   db,
				args: []string{"test", "123"},
			},
			want: MakeStringReply("3"),
		}, {
			name: "append",
			args: args{
				db:   db,
				args: []string{"test", "0987"},
			},
			want: MakeStringReply("7"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Append(tt.args.db, tt.args.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Append() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAppend_1(t *testing.T) {
	type args struct {
		db   *DB
		args []string
	}
	//var db = MakeDB(0)

	goroutineNum := 2
	tasknums := goroutineNum * 600
	nums := 0

	pool := tunny.NewFunc(goroutineNum, func(payload interface{}) interface{} {
		nums += 1
		return payload
	})

	defer pool.Close()
	now := time.Now()

	var wg sync.WaitGroup
	wg.Add(tasknums)
	for i := 0; i < tasknums; i++ {
		go func(num int) {
			pool.Process(num)
			wg.Done()
		}(i)
	}
	wg.Wait()

	since := time.Since(now)
	println(nums)
	println("cost ", since.Milliseconds(), "avenger ", since.Nanoseconds()/int64(tasknums), "ns")
}

func TestGetP(t *testing.T) {
	type args struct {
		db   *DB
		args []string
	}
	var db = MakeDB(0)

	var qps int64 = 1000000
	var qpsi int = 1000000
	writeRate := 10

	goroutineNum := 2
	tasknums := goroutineNum * qpsi

	pool := tunny.NewFunc(goroutineNum, func(payload interface{}) interface{} {
		if rand.Intn(100) < writeRate {
			Set(db, []string{"test", "1234"})
		} else {
			Get(db, []string{"test"})
		}
		return payload
	})

	var wg sync.WaitGroup
	wg.Add(tasknums)

	now := time.Now()
	for i := 0; i < tasknums; i++ {
		go func(num int) {
			pool.Process(num)
			wg.Done()
		}(i)
	}
	wg.Wait()

	since := time.Since(now)

	println("qps ", qps, " cost ", since.Milliseconds(), "ms", "avenger ", since.Nanoseconds()/(2*qps), "ns")
}

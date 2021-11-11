package Youpo

import (
	"github.com/Huabuxiu/Youpo/networks"
	"reflect"
	"testing"
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
		want networks.Reply
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
			want: networks.MakeOKReply(),
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
		want networks.Reply
	}{
		// TODO: Add test cases.
		{
			name: "getTest",
			args: args{
				db:   db,
				args: []string{"test"},
			},
			want: networks.MakeStringReply("1234"),
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

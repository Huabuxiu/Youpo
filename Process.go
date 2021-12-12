package Youpo

import (
	"github.com/Huabuxiu/Youpo/datastruct"
	"sync"
)

/**
以下是各个 FLAG 的意义：
 *
 * w: write command (may modify the key space).
 *    写入命令，可能会修改 key space
 *
 * r: read command  (will never modify the key space).
 *    读命令，不修改 key space
 * m: may increase memory usage once called. Don't allow if out of memory.
 *    可能会占用大量内存的命令，调用时对内存占用进行检查
 *
 * a: admin command, like SAVE or SHUTDOWN.
 *    管理用途的命令，比如 SAVE 和 SHUTDOWN
 *
 * p: Pub/Sub related command.
 *    发布/订阅相关的命令
 *
 * f: force replication of this command, regardless of server.dirty.
 *    无视 server.dirty ，强制复制这个命令。
 *
 * s: command not allowed in scripts.
 *    不允许在脚本中使用的命令
 *
 * R: random command. Command is not deterministic, that is, the same command
 *    with the same arguments, with the same key space, may have different
 *    results. For instance SPOP and RANDOMKEY are two random commands.
 *    随机命令。
 *    命令是非确定性的：对于同样的命令，同样的参数，同样的键，结果可能不同。
 *    比如 SPOP 和 RANDOMKEY 就是这样的例子。
 *
 * S: Sort command output array if called from script, so that the output
 *    is deterministic.
 *    如果命令在 Lua 脚本中执行，那么对输出进行排序，从而得出确定性的输出。
 *
 * l: Allow command while loading the database.
 *    允许在载入数据库时使用的命令。
 *
 * t: Allow command while a slave has stale data but is not allowed to
 *    server this data. Normally no command is accepted in this condition
 *    but just a few.
 *    允许在附属节点带有过期数据时执行的命令。
 *    这类命令很少有，只有几个。
 *
 * M: Do not automatically propagate the command on MONITOR.
 *    不要在 MONITOR 模式下自动广播的命令。
 *
 * k: Perform an implicit ASKING for this command, so the command will be
 *    accepted in cluster mode if the slot is marked as 'importing'.
 *    为这个命令执行一个显式的 ASKING ，
 *    使得在集群模式下，一个被标示为 importing 的槽可以接收这命令。
*/

//多态 命令函数指针
type Function func(db *DB, args []string) Reply

//命令表
var commandMap = datastruct.MakeMap()

//单例对象
var process *Process
var mutex sync.Mutex

type Process struct {
}

//单例执行器
func MakeSingleProcess() *Process {
	if process == nil {
		mutex.Lock()
		defer mutex.Unlock()

		//双重检查、避免加锁过程中有线程初始化了
		if process == nil {
			process = &Process{}
		}

	}
	return process
}

//注册命令
func RegisterCommand(name string,
	execFunction Function, flag string, arity int) {
	command := Command{
		name:         name,
		execFunction: execFunction,
		flag:         flag,
		arity:        arity,
		microseconds: 0,
		calls:        0,
	}
	commandMap.Put(command.name, &command)
}

type Command struct {
	name string

	execFunction Function

	//命令类型
	flag string

	//期望的参数个数
	arity int

	// 统计信息
	// microseconds 记录了命令执行耗费的总毫微秒数
	// calls 是命令被执行的总次数
	microseconds, calls int64
}

//校验命令参数
func (cmd *Command) checkArgNums(argNum int) bool {
	if cmd.arity >= 0 {
		return cmd.arity == argNum
	}
	return true
}

//执行命令
func (process *Process) Exec(db *DB, args []string) Reply {
	key := string(args[0])

	command, exist := commandMap.Get(key)
	if !exist {
		return EmptyReply{}
	}

	//校验参数个数
	if !command.(*Command).checkArgNums(len(args)) {
		return MakeErrorReply("illegal arg length")
	}

	return command.(*Command).execFunction(db, args[1:])
}

func StringInit() {
	RegisterCommand("get", Get, "r", 2)
	RegisterCommand("set", Set, "w", 3)
}

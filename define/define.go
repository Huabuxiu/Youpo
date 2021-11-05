package define

//定义一个函数接口
type FunctionInterface func(args interface{}) interface{}

type FunctionMultipleInterface func(argOne string, arg interface{}) interface{}

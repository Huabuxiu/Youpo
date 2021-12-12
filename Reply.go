package Youpo

type Reply interface {
	ToBytes() []byte
}

type StringReply struct {
	resp string
}

type EmptyReply struct {
}

type ErrorReply struct {
	errorMsg string
}

func (reply EmptyReply) ToBytes() []byte {
	return nil
}

func (s *StringReply) ToBytes() []byte {
	return []byte(s.resp)
}

func (s *ErrorReply) ToBytes() []byte {
	return []byte(s.errorMsg)
}

func MakeErrorReply(errorMsg string) *ErrorReply {
	return &ErrorReply{errorMsg: errorMsg}
}

func MakeStringReply(resp string) *StringReply {
	return &StringReply{resp: resp}
}

func MakeOKReply() *StringReply {
	return &StringReply{resp: "OK"}
}

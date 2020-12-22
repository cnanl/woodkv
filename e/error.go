package e

import (
	"fmt"
)

//Error 自定义的公共错误码
type Error struct {
	code int
}

var (
	Success  = NewError(0, "成功")
	NotFound = NewError(101, "找不到")
	Unknown  = NewError(102, "未知错误")
)

var codes = map[int]string{}

//NewError 创建一个新的Error
func NewError(code int, msg string) Error {
	if _, ok := codes[code]; ok {
		panic(fmt.Sprintf("错误码 %d 已存在", code))
	}
	codes[code] = msg
	return Error{code}
}

func (e Error) Error() string {
	return fmt.Sprintf("%d: %s", e.code, codes[e.code])
}

func (e *Error) Code() int {
	return e.code
}

func CodeMap() *map[int]string {
	return &codes
}

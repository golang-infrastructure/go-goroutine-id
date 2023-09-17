package goroutine_id

import "errors"

var (
	// ErrGetGoroutineIdFailed 获取goroutine id失败
	ErrGetGoroutineIdFailed = errors.New("get goroutine id failed")
)

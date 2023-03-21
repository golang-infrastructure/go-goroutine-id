package goroutine_id

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
)

// GetGoroutineID 获取当前协程的ID
func GetGoroutineID() (int, error) {
	id, err := strconv.Atoi(GetGoroutineIDAsString())
	if err != nil {
		return 0, fmt.Errorf("cannot get goroutine id: %v", err)
	}
	return id, nil
}

// GetGoroutineIDAsString 获取当前协程的ID，以字符串的形式返回
func GetGoroutineIDAsString() string {
	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	idField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0]
	return idField
}

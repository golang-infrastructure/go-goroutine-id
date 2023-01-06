package goroutine_id

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
)

// GetGoroutineId 获取当前协程的ID
func GetGoroutineId() (int, error) {
	id, err := strconv.Atoi(GetGoroutineIdAsString())
	if err != nil {
		return 0, fmt.Errorf("cannot get goroutine id: %v", err)
	}
	return id, nil
}

// GetGoroutineIdAsString 获取当前协程的ID，以字符串的形式返回
func GetGoroutineIdAsString() string {
	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	idField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0]
	return idField
}

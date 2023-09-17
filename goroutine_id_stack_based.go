package goroutine_id

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
)

// ------------------------------------------------- --------------------------------------------------------------------

// GetGoroutineId 获取当前协程的id
func GetGoroutineId() (int, error) {

	goroutineIdString := GetGoroutineIdAsString()
	if goroutineIdString == "" {
		return 0, ErrGetGoroutineIdFailed
	}

	goroutineId, err := strconv.Atoi(goroutineIdString)
	if err != nil {
		return 0, fmt.Errorf("goroutine goroutineId parse int error, value = %s, msg = %s", goroutineIdString, err.Error())
	}
	return goroutineId, nil
}

// GetGoroutineIdAsString 获取当前协程的id，以字符串的形式返回
func GetGoroutineIdAsString() string {
	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	fields := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))
	if len(fields) < 1 {
		return ""
	}
	return fields[0]
}

// ------------------------------------------------- --------------------------------------------------------------------

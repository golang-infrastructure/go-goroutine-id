package goroutine_id

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
)

// GetGoroutineID 获取当前协程的ID
func GetGoroutineID() (int, error) {
	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	idField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0]
	id, err := strconv.Atoi(idField)
	if err != nil {
		return 0, fmt.Errorf("cannot get goroutine id: %v", err)
	}
	return id, nil
}

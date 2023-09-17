package goroutine_id

import (
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
)

// 只要能成功获取到id就认为是没问题
func TestGetGoroutineID(t *testing.T) {
	var goroutineId int

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		goroutineId, _ = GetGoroutineId()
	}()
	wg.Wait()

	assert.True(t, goroutineId != 0)
}

# go-goroutine-id

# 一、这是啥？解决了啥问题？

在Go中是不推荐使用Goroutine ID的，但是仍然有一些场景有Goroutine ID会更好，这个库就提供了API能够获取协程的ID。 

# 二、安装

```bash
go get -u github.com/golang-infrastructure/go-goroutine-id
```

# 三、API示例

```go
package main

import (
	"fmt"
	goroutine_id "github.com/golang-infrastructure/go-goroutine-id"
	"sync"
)

func main() {

	var goroutineId int
	var err error

	// Start a new coroutine and see what its id is
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		goroutineId, err = goroutine_id.GetGoroutineId()
		if err != nil {
			panic(err)
		}
	}()
	wg.Wait()
	fmt.Println(fmt.Sprintf("Goroutine ID: %d", goroutineId))
	// Output:
	// Goroutine ID: 6

	// Look at the id of the main coroutine
	goroutineId, err = goroutine_id.GetGoroutineId()
	if err != nil {
		panic(err)
	}
	fmt.Println(fmt.Sprintf("Main Goroutine ID: %d", goroutineId))
	// Output:
	// Main Goroutine ID: 1

}
```




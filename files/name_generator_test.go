// Author: FishGoddess
// Email: fishgoddess@qq.com
// Created at 2020/06/24 23:05:51

package files

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// 测试 DefaultNameGenerator 是否会产生重复名字
func TestDefaultNameGenerator(t *testing.T) {

	// 简单看看生成的名字
	nameGenerator1 := DefaultNameGenerator()
	nameGenerator2 := DefaultNameGenerator()
	fmt.Printf("%p %p\n", nameGenerator1, nameGenerator2)
	fmt.Println(nameGenerator1.NextName("", time.Now()), nameGenerator2.NextName("", time.Now()))

	// 并发生成一批名字，并判断唯一性
	nameQueue := make(chan string, 64)
	defer close(nameQueue)

	times := 128
	concurrency := 8
	group := sync.WaitGroup{}
	group.Add(concurrency)
	for i := 0; i < concurrency; i++ {
		go func(id int) {
			if id&1 == 0 {
				for i := 0; i < times; i++ {
					nameQueue <- nameGenerator1.NextName("", time.Now())
				}
			} else {
				for i := 0; i < times; i++ {
					nameQueue <- nameGenerator2.NextName("", time.Now())
				}
			}
			group.Done()
		}(i)
	}

	countMap := map[string]int{}
	done := false
	for !done {
		select {
		case name := <-nameQueue:
			//fmt.Println(name)
			if _, ok := countMap[name]; ok {
				t.Fatalf("name [%s] 重复了！\n", name)
			}
			countMap[name] = 1
		case <-time.After(time.Second):
			done = true
			break
		}
	}

	if len(countMap) != concurrency*times {
		t.Fatal("countMap 数量不对，说明有 name 重复了！")
	}

	group.Wait()
}

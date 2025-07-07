package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

// 示例1：竞态条件
// 展示没有同步的并发访问会导致数据竞争
func raceConditionExample() {
	fmt.Println("=== 竞态条件示例 ===")
	
	var counter int
	var wg sync.WaitGroup
	
	// 启动10个goroutine并发递增counter
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				counter++ // 这里存在竞态条件
			}
		}()
	}
	
	wg.Wait()
	fmt.Printf("无同步的计数器结果: %d (期望: 10000)\n", counter)
}

// 示例2：使用互斥锁解决竞态条件
func mutexExample() {
	fmt.Println("\n=== 互斥锁示例 ===")
	
	var counter int
	var mu sync.Mutex
	var wg sync.WaitGroup
	
	// 启动10个goroutine并发递增counter，使用互斥锁同步
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				mu.Lock()
				counter++
				mu.Unlock()
			}
		}()
	}
	
	wg.Wait()
	fmt.Printf("使用互斥锁的计数器结果: %d (期望: 10000)\n", counter)
}

// 示例3：读写锁
// 读写锁允许多个读操作同时进行，但写操作是互斥的
func rwMutexExample() {
	fmt.Println("\n=== 读写锁示例 ===")
	
	var data = make(map[string]int)
	var rwMu sync.RWMutex
	var wg sync.WaitGroup
	
	// 初始化数据
	data["key1"] = 1
	data["key2"] = 2
	
	// 启动多个读goroutine
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < 3; j++ {
				rwMu.RLock()
				value := data["key1"]
				fmt.Printf("读goroutine %d 读取到值: %d\n", id, value)
				time.Sleep(100 * time.Millisecond)
				rwMu.RUnlock()
			}
		}(i)
	}
	
	// 启动一个写goroutine
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 3; i++ {
			rwMu.Lock()
			data["key1"] = data["key1"] + 1
			fmt.Printf("写goroutine 更新值为: %d\n", data["key1"])
			time.Sleep(200 * time.Millisecond)
			rwMu.Unlock()
		}
	}()
	
	wg.Wait()
}

// 示例4：sync.Once - 确保某个操作只执行一次
var once sync.Once
var config map[string]string

func loadConfig() {
	fmt.Println("正在加载配置...")
	config = make(map[string]string)
	config["host"] = "localhost"
	config["port"] = "8080"
	time.Sleep(500 * time.Millisecond) // 模拟配置加载时间
	fmt.Println("配置加载完成")
}

func getConfig() map[string]string {
	once.Do(loadConfig)
	return config
}

func syncOnceExample() {
	fmt.Println("\n=== sync.Once 示例 ===")
	
	var wg sync.WaitGroup
	
	// 启动多个goroutine尝试加载配置
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			cfg := getConfig()
			fmt.Printf("Goroutine %d 获取配置: %v\n", id, cfg)
		}(i)
	}
	
	wg.Wait()
}

// 示例5：原子操作
// 原子操作提供了一种无锁的并发访问方式
func atomicExample() {
	fmt.Println("\n=== 原子操作示例 ===")
	
	var counter int64
	var wg sync.WaitGroup
	
	// 启动10个goroutine并发递增counter
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				atomic.AddInt64(&counter, 1)
			}
		}()
	}
	
	wg.Wait()
	fmt.Printf("使用原子操作的计数器结果: %d (期望: 10000)\n", counter)
}

// 示例6：银行转账问题 - 演示死锁避免
type Account struct {
	mu      sync.Mutex
	balance int
}

func (a *Account) Deposit(amount int) {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.balance += amount
}

func (a *Account) Withdraw(amount int) bool {
	a.mu.Lock()
	defer a.mu.Unlock()
	if a.balance < amount {
		return false
	}
	a.balance -= amount
	return true
}

func (a *Account) Balance() int {
	a.mu.Lock()
	defer a.mu.Unlock()
	return a.balance
}

// 避免死锁的转账操作
func Transfer(from, to *Account, amount int) {
	// 通过账户地址排序来避免死锁
	first, second := from, to
	if from > to {
		first, second = to, from
	}
	
	first.mu.Lock()
	second.mu.Lock()
	
	defer first.mu.Unlock()
	defer second.mu.Unlock()
	
	if from.balance >= amount {
		from.balance -= amount
		to.balance += amount
		fmt.Printf("转账成功: %d 元\n", amount)
	} else {
		fmt.Printf("转账失败: 余额不足\n")
	}
}

func bankExample() {
	fmt.Println("\n=== 银行转账示例 ===")
	
	account1 := &Account{balance: 1000}
	account2 := &Account{balance: 500}
	
	var wg sync.WaitGroup
	
	// 模拟并发转账
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			amount := rand.Intn(100) + 1
			if id%2 == 0 {
				Transfer(account1, account2, amount)
			} else {
				Transfer(account2, account1, amount)
			}
		}(i)
	}
	
	wg.Wait()
	
	fmt.Printf("账户1余额: %d\n", account1.Balance())
	fmt.Printf("账户2余额: %d\n", account2.Balance())
}

func main() {
	fmt.Println("《Go程序设计语言》第9章：共享变量的并发")
	fmt.Println("============================================")
	
	// 运行所有示例
	raceConditionExample()
	mutexExample()
	rwMutexExample()
	syncOnceExample()
	atomicExample()
	bankExample()
	
	fmt.Println("\n============================================")
	fmt.Println("第9章示例运行完成!")
} 
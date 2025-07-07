// Goroutines 演示Goroutines和通道的使用
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	// 1. 基本Goroutine
	fmt.Printf("基本Goroutine:\n")
	go sayHello("Goroutine")
	sayHello("Main")
	time.Sleep(time.Second) // 等待goroutine完成

	// 2. 多个Goroutines
	fmt.Printf("\n多个Goroutines:\n")
	for i := 0; i < 3; i++ {
		go func(id int) {
			fmt.Printf("Goroutine %d 正在运行\n", id)
		}(i)
	}
	time.Sleep(time.Second)

	// 3. 基本通道
	fmt.Printf("\n基本通道:\n")
	ch := make(chan string)
	go func() {
		ch <- "Hello from goroutine!"
	}()
	message := <-ch
	fmt.Printf("收到消息: %s\n", message)

	// 4. 缓冲通道
	fmt.Printf("\n缓冲通道:\n")
	bufferedCh := make(chan int, 3)
	bufferedCh <- 1
	bufferedCh <- 2
	bufferedCh <- 3
	fmt.Printf("缓冲通道长度: %d, 容量: %d\n", len(bufferedCh), cap(bufferedCh))
	
	fmt.Printf("从缓冲通道读取: %d\n", <-bufferedCh)
	fmt.Printf("从缓冲通道读取: %d\n", <-bufferedCh)

	// 5. 通道的关闭
	fmt.Printf("\n通道的关闭:\n")
	numberCh := make(chan int)
	go func() {
		for i := 1; i <= 3; i++ {
			numberCh <- i
		}
		close(numberCh)
	}()
	
	for num := range numberCh {
		fmt.Printf("收到数字: %d\n", num)
	}

	// 6. select语句
	fmt.Printf("\nselect语句:\n")
	ch1 := make(chan string)
	ch2 := make(chan string)
	
	go func() {
		time.Sleep(time.Millisecond * 500)
		ch1 <- "来自ch1"
	}()
	
	go func() {
		time.Sleep(time.Millisecond * 200)
		ch2 <- "来自ch2"
	}()
	
	select {
	case msg1 := <-ch1:
		fmt.Printf("收到: %s\n", msg1)
	case msg2 := <-ch2:
		fmt.Printf("收到: %s\n", msg2)
	case <-time.After(time.Second):
		fmt.Printf("超时!\n")
	}

	// 7. Worker Pool模式
	fmt.Printf("\nWorker Pool模式:\n")
	jobs := make(chan int, 100)
	results := make(chan int, 100)
	
	// 启动3个worker
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}
	
	// 发送5个任务
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)
	
	// 收集结果
	for a := 1; a <= 5; a++ {
		<-results
	}

	// 8. 生产者-消费者模式
	fmt.Printf("\n生产者-消费者模式:\n")
	producerConsumerDemo()

	// 9. 通道方向
	fmt.Printf("\n通道方向:\n")
	channelDirectionDemo()

	// 10. 同步原语
	fmt.Printf("\n同步原语:\n")
	syncDemo()
}

func sayHello(from string) {
	for i := 0; i < 3; i++ {
		fmt.Printf("Hello from %s\n", from)
		time.Sleep(time.Millisecond * 100)
	}
}

func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Printf("Worker %d 开始任务 %d\n", id, j)
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(1000)))
		fmt.Printf("Worker %d 完成任务 %d\n", id, j)
		results <- j * 2
	}
}

func producerConsumerDemo() {
	ch := make(chan int, 5)
	
	// 生产者
	go func() {
		for i := 1; i <= 10; i++ {
			fmt.Printf("生产: %d\n", i)
			ch <- i
			time.Sleep(time.Millisecond * 100)
		}
		close(ch)
	}()
	
	// 消费者
	go func() {
		for item := range ch {
			fmt.Printf("消费: %d\n", item)
			time.Sleep(time.Millisecond * 200)
		}
	}()
	
	time.Sleep(time.Second * 3)
}

// 通道方向示例
func channelDirectionDemo() {
	ch := make(chan string)
	
	go send(ch)
	go receive(ch)
	
	time.Sleep(time.Second)
}

func send(ch chan<- string) { // 只能发送
	ch <- "数据"
}

func receive(ch <-chan string) { // 只能接收
	msg := <-ch
	fmt.Printf("接收到: %s\n", msg)
}

// 同步原语示例
func syncDemo() {
	var wg sync.WaitGroup
	var mu sync.Mutex
	counter := 0
	
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				mu.Lock()
				counter++
				mu.Unlock()
			}
			fmt.Printf("Goroutine %d 完成\n", id)
		}(i)
	}
	
	wg.Wait()
	fmt.Printf("最终计数: %d\n", counter)
} 
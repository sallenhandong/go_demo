package main

import (
	"fmt"
	"time"
)

//模拟一个耗时任务
func worker(id int, jobs <-chan int, results chan<- int) {

	for j := range jobs {

		fmt.Println("worker", id, "processing job", j)
		time.Sleep(time.Second)
		results <- j * 2
	}

}
func main() {

	//发送工作和接受工作的结果
	//定义两个通道,一个jobs，一个results
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	//启动三个worker协程
	for w := 1; w <= 3; w++ {

		go worker(w, jobs, results)
	}

	//这里发送9个任务，然后关闭通道，告知任务发送完成
	for j := 1; j <= 9; j++ {

		jobs <- j
	}

	close(jobs)

	//然后从results里面获取结果
	for a := 1; a <= 9; a++ {
		<-results
	}

}

package main

import (
	"LearnGo/queue"
	"fmt"
)

func main() {
	initQueue := queue.Queue{}
	fmt.Println(initQueue.IsEmpty())
	initQueue.Push(1)
	initQueue.Push(3)
	initQueue.Push(4)
	initQueue.Push(5)
	fmt.Println(initQueue.IsEmpty())
	fmt.Println(initQueue.Pop())
}

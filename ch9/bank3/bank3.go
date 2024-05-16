package bank3

import "sync"

var mu sync.Mutex
var balance int

func Deposit(amount int) {
	mu.Lock()
	defer mu.Unlock()
	balance += amount
}

func Balance() int{
	mu.Lock()
	defer mu.Unlock()
	return balance // Unlock会在return语句读取完balance的值之后执行
}
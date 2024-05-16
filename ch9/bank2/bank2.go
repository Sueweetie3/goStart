package bank2

var balance int
var sema = make(chan struct{}, 1)

func Deposit(amount int) {
	sema <- struct{}{}
    balance = balance + amount
	<-sema
}
func Balance() int {
	sema <- struct{}{}
    b := balance
	<-sema
    return b
}
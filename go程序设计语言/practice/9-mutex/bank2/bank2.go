package bank2
package bank2
var (
	sema = make(chan struct{}, 1)
	balance int
)

func Deposit(amount int){
	sema <- struct{}{}
	balance += amount
	<- sema
}
func Balance() int{
	sema <- struct{}{}
	defer <-sema
	return balance
}
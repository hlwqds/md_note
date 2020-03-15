package bank

type withdrawInfo struct {
	Amount int
	Ch     chan bool
}

var deposits = make(chan int)
var balances = make(chan int)
var withdraw = make(chan withdrawInfo)

func Deposit(amount int) { deposits <- amount }
func Balance() int       { return <-balances }
func Withdraw(amount int) bool {
	wInfo := withdrawInfo{
		Amount: amount,
		Ch:     make(chan bool),
	}
	withdraw <- wInfo
	return <-wInfo.Ch
}

func teller() {
	var balance int
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		case wInfo := <-withdraw:
			if balance < wInfo.Amount {
				wInfo.Ch <- false
			} else {
				balance -= wInfo.Amount
				wInfo.Ch <- true
			}
		}
	}
}

func init() {
	go teller()
}

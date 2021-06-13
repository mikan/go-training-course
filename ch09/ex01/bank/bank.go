// Package bank は一つの口座を持つ平行的に安全な銀行を提供します。
package bank

type withdrawTransaction struct {
	amount int
	ok     chan bool
}

var deposits = make(chan int)                  // 入金額を送信する
var balances = make(chan int)                  // 残高を受信する
var withdraws = make(chan withdrawTransaction) // 出金額を送信し結果を受信する

func Deposit(amount int) {
	deposits <- amount
}

func Balance() int {
	return <-balances
}

func Withdraw(amount int) bool {
	c := make(chan bool)
	withdraws <- withdrawTransaction{amount: amount, ok: c}
	return <-c
}

func teller() {
	var balance int // balance は teller ゴルーチンに閉じ込められている
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		case transaction := <-withdraws:
			if balance >= transaction.amount {
				balance -= transaction.amount
				transaction.ok <- true
			} else {
				transaction.ok <- false
			}
		}
	}
}

func init() {
	go teller() // モニターゴルーチンを開始する
}

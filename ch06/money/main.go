package main

import "fmt"

type sMoney struct {
	money      float64
	label      string
	amountLeft float64
}

func main() {
	amount := 43.57
	w := sMoney{amountLeft: amount}

	moneyTypes := map[float64]string{20: "£20",
		10: "£10", 5: "£5",
		0.5: "50p", 1.0: "£1",
		0.2: "20p", 0.1: "10p",
		0.05: "5p", 0.02: "2p", 0.01: "1p"}
	for money, label := range moneyTypes {
		w.label = label
		w.money = money
		checkMoney(&w)
	}
}

func checkMoney(w *sMoney) {
	if w.money <= w.amountLeft {
		w.amountLeft = w.amountLeft - w.money
		fmt.Println(w.label)
		checkMoney(w)
	}
}

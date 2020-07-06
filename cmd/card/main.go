package main

import (
	"fmt"
	"time"

	"go_hw6/pkg/card"
)

func main() {
	t1 := &card.Transaction{
		Id:     0,
		Type:   "debit",
		Sum:    735_55,
		Status: "processing",
		MCC:    "5921",
		Date:   time.Date(2020, time.June, 9, 11, 15, 10, 0, time.UTC).Unix(),
	}
	t2 := &card.Transaction{
		Id:     1,
		Type:   "refill",
		Sum:    2000_00,
		Status: "done",
		Date:   time.Date(2020, time.June, 11, 1, 46, 40, 0, time.UTC).Unix(),
	}
	t3 := &card.Transaction{
		Id:     2,
		Type:   "debit",
		Sum:    1203_91,
		Status: "processing",
		MCC:    "5812",
		Date:   time.Date(2020, time.June, 9, 11, 15, 10, 0, time.UTC).Unix(),
	}
	t4 := &card.Transaction{
		Id:     3,
		Type:   "debit",
		Sum:    2560_00,
		Status: "processing",
		MCC:    "5812",
		Date:   time.Date(2020, time.June, 15, 14, 30, 10, 0, time.UTC).Unix(),
	}

	master := &card.Card{
		Id:           0,
		Issuer:       "MasterCard",
		Balance:      45_000_00,
		Currency:     "RUB",
		Number:       "0808",
		Icon:         "https://dlpng.com/png/6794578",
		Transactions: []card.Transaction{*t1, *t2, *t3, *t4},
	}

	fmt.Println("Original: ", master.Transactions)
	newTr := card.SortTransactions(master.Transactions)

	fmt.Println("Sorted: ", newTr)
	fmt.Println("Original after sort: ", master.Transactions)
}
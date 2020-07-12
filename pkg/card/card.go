package card

import (
	"fmt"
	"sort"
	"sync"
	"time"
)

type Transaction struct {
	Id     int64
	Type   string
	Sum    int64
	Status string
	MCC    string
	Date   time.Time
}

type Card struct {
	Id           int64
	Issuer       string
	Balance      int64
	Currency     string
	Number       string
	Icon         string
	Transactions []Transaction
}

func AddTransaction(card *Card, transaction *Transaction) {
	card.Transactions = append(card.Transactions, *transaction)
}

func SortTransactions(transactions []Transaction) []Transaction {
	tr := make([]Transaction, len(transactions))
	copy(tr, transactions)

	sort.SliceStable(tr, func(i, j int) bool {
		return tr[i].Sum > tr[j].Sum
	})
	return tr
}

func (card *Card) makeMapByDate(start, finish time.Time) map[string][]Transaction {
	transByDate := make(map[string][]Transaction)

	for _, trans := range card.Transactions {
		month := trans.Date.Month()
		year := trans.Date.Year()
		if (trans.Date.After(start) && trans.Date.Before(finish)) || (trans.Date.Equal(start) || trans.Date.Equal(finish)) {
			transByDate[fmt.Sprintf("%d-%d", month, year)] = append(transByDate[fmt.Sprintf("%d-%d", month, year)], trans)
		}
	}
	return transByDate
}

func Sum(transitions []Transaction) int64 {
	sum := int64(0)
	for _, t := range transitions {
		sum += t.Sum
	}
	return sum
}

func (card *Card) SumConcurrently(start, finish time.Time) map[string]int64 {
	transByDate := card.makeMapByDate(start, finish)

	wg := sync.WaitGroup{}
	wg.Add(len(transByDate))

	total := make(map[string]int64)
	for key, tr := range transByDate {
		part := tr
		var sum int64
		go func(key string) {
			sum += Sum(part)
			total[key] = sum
			wg.Done()
		}(key)
	}
	wg.Wait()
	return total
}

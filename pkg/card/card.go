package card

import (
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

func (card *Card) SumConcurrently(start, finish time.Time) int64 {
	transByDate := make(map[string][]Transaction)

	for _, trans := range card.Transactions {
		month := trans.Date.Month()
		if (trans.Date.After(start) && trans.Date.Before(finish)) || (trans.Date.Equal(start) || trans.Date.Equal(finish)) {
			transByDate[month.String()] = append(transByDate[month.String()], trans)
		}
	}
	wg := sync.WaitGroup{}
	wg.Add(len(transByDate))

	var total int64
	for _, tr := range transByDate {
		part := tr
		go func() {
			for _, v := range part {
				total += v.Sum
			}
			wg.Done()
		}()
	}
	wg.Wait()
	return total
}

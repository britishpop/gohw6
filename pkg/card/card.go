package card

import "sort"

type Transaction struct {
	Id     int64
	Type   string
	Sum    int64
	Status string
	MCC    string
	Date   int64
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

func SumConcurrently(transactions []int64, goroutines int) int64 {
	wg := sync.WaitGroup{}
	wg.Add(goroutines)

	total := int64(0)
	partSize := // TODO: решаете сами
	for i := 0; i < goroutines; i++ {
	        // ВАЖНО: просто с partSize не прокатит, вам нужно как-то заранее разделить слайс по месяцам
	        // ВАЖНО: этот код - лишь шаблон, который показывает вам как запустить горутину
		part := transactions[i*partSize : (i+1)*partSize]
		go func() {
			sum := Sum(part)
			fmt.Println(sum)
			wg.Done()
		}()
	}

	wg.Wait()
	return total
}

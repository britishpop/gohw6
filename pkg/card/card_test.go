package card

import (
	"reflect"
	"testing"
	"time"
)

func TestSortTransactions(t *testing.T) {
	type args struct {
		transactions []Transaction
	}

	original := []Transaction{{
		Id:     0,
		Type:   "debit",
		Sum:    7800_00,
		Status: "processing",
		MCC:    "5921",
		Date:   time.Date(2020, time.June, 9, 11, 15, 10, 0, time.UTC),
	}, {
		Id:     1,
		Type:   "debit",
		Sum:    90000_00,
		Status: "processing",
		MCC:    "5921",
		Date:   time.Date(2020, time.June, 9, 12, 16, 10, 0, time.UTC),
	}, {
		Id:     2,
		Type:   "debit",
		Sum:    55_00,
		Status: "processing",
		MCC:    "5921",
		Date:   time.Date(2020, time.June, 25, 6, 14, 0, 0, time.UTC),
	}, {
		Id:     3,
		Type:   "debit",
		Sum:    55_00,
		Status: "processing",
		MCC:    "5921",
		Date:   time.Date(2020, time.June, 26, 6, 14, 0, 0, time.UTC),
	}}

	sorted := []Transaction{{
		Id:     1,
		Type:   "debit",
		Sum:    90000_00,
		Status: "processing",
		MCC:    "5921",
		Date:   time.Date(2020, time.June, 9, 12, 16, 10, 0, time.UTC),
	}, {
		Id:     0,
		Type:   "debit",
		Sum:    7800_00,
		Status: "processing",
		MCC:    "5921",
		Date:   time.Date(2020, time.June, 9, 11, 15, 10, 0, time.UTC),
	}, {
		Id:     2,
		Type:   "debit",
		Sum:    55_00,
		Status: "processing",
		MCC:    "5921",
		Date:   time.Date(2020, time.June, 25, 6, 14, 0, 0, time.UTC),
	}, {
		Id:     3,
		Type:   "debit",
		Sum:    55_00,
		Status: "processing",
		MCC:    "5921",
		Date:   time.Date(2020, time.June, 26, 6, 14, 0, 0, time.UTC),
	}}

	tests := []struct {
		name string
		args args
		want []Transaction
	}{
		{
			name: "Сортировка по убыванию",
			args: args{original},
			want: sorted,
		},
	}
	for _, tt := range tests {
		if got := SortTransactions(tt.args.transactions); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("SortTransactions() = %v, want %v", got, tt.want)
		}
	}
}

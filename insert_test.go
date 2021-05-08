package sort_go

import (
	"reflect"
	"testing"
)

// go test -v
func TestInsertSort(t *testing.T) {
	type args struct {
		array []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "case1", args: args{array: []int{1, 3, 4, 2, 345, 23}}, want: []int{1, 2, 3, 4, 23, 345}},
		{name: "case2", args: args{array: []int{1, 4, 22, 345, 23, 234}}, want: []int{1, 4, 22, 23, 234, 345}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InsertSort(tt.args.array); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InsertSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

// go test -bench=InsertSort -benchmem
func BenchmarkInsertSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		InsertSort([]int{1, 4, 22, 345, 23, 234})
	}
}

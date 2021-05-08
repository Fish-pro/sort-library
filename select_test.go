package sort_go

import (
	"reflect"
	"testing"
)

// go test -v
func TestSelectSort(t *testing.T) {
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
			if got := SelectSort(tt.args.array); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SelectSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

//  go test -bench=SelectSort -benchmem
func BenchmarkSelectSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SelectSort([]int{34, 23, 3, 234, 33, 3423, 2343444})
	}
}

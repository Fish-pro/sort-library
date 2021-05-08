package sort_go

import (
	"reflect"
	"testing"
)

// go test -v
func TestQuickSort(t *testing.T) {
	type args struct {
		values []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "case1", args: args{values: []int{1, 3, 4, 2, 345, 23}}, want: []int{1, 2, 3, 4, 23, 345}},
		{name: "case2", args: args{values: []int{1, 4, 22, 345, 23, 234}}, want: []int{1, 4, 22, 23, 234, 345}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := QuickSort(tt.args.values); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("QuickSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

// go test -bench=QuickSort -benchmem
func BenchmarkQuickSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		QuickSort([]int{1, 4, 22, 345, 23, 234})
	}
}

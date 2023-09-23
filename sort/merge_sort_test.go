package sort

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMergeSort(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test",
			args: args{
				arr: []int{14, 12, 15, 13, 11, 16},
			},
			want: []int{11, 12, 13, 14, 15, 16},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MergeSort(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sort(t *testing.T) {
	type args struct {
		arr []int
		tmp []int
		l   int
		r   int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			merge_sort(tt.args.arr, tt.args.tmp, tt.args.l, tt.args.r)
		})
	}
}

func Test_merge(t *testing.T) {
	type args struct {
		arr []int
		tmp []int
		l   int
		m   int
		r   int
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test",
			args: args{
				arr: []int{2, 5, 3, 6},
				tmp: make([]int, 4),
				l:   0,
				m:   1,
				r:   3,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			merge(tt.args.arr, tt.args.tmp, tt.args.l, tt.args.m, tt.args.r)
			fmt.Printf("%+v", tt.args.arr)
		})
	}
}

func Test_merge_sort(t *testing.T) {
	type args struct {
		arr []int
		tmp []int
		l   int
		r   int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			merge_sort(tt.args.arr, tt.args.tmp, tt.args.l, tt.args.r)
		})
	}
}

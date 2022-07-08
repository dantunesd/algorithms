package bubblesort

import (
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	type args struct {
		numbers []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "should sort",
			args: args{
				numbers: []int{9, 5, 1, 4, 2, 8},
			},
			want: []int{1, 2, 4, 5, 8, 9},
		},
		{
			name: "should sort",
			args: args{
				numbers: []int{0, 5, 1, 4},
			},
			want: []int{0, 1, 4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BubbleSort(tt.args.numbers); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BubbleSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

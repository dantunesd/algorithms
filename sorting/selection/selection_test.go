package selection

import (
	"reflect"
	"testing"
)

func TestSort(t *testing.T) {
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
				numbers: []int{2, 3, 1},
			},
			want: []int{1, 2, 3},
		},
		{
			name: "should sort 2",
			args: args{
				numbers: []int{1, 0},
			},
			want: []int{0, 1},
		},
		{
			name: "should sort 3",
			args: args{
				numbers: []int{0, 1},
			},
			want: []int{0, 1},
		},
		{
			name: "should sort 4",
			args: args{
				numbers: []int{1, 1},
			},
			want: []int{1, 1},
		},
		{
			name: "should sort 5",
			args: args{
				numbers: []int{},
			},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sort(tt.args.numbers); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Sort() = %v, want %v", got, tt.want)
			}
		})
	}
}

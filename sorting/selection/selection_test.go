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
				numbers: []int{64, 25, 12, 22, 11},
			},
			want: []int{11, 12, 22, 25, 64},
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

package binary

import "testing"

func TestSearch(t *testing.T) {
	type args struct {
		value int
		list  []int
	}
	tests := []struct {
		name      string
		args      args
		want      int
		wantLoops int
	}{
		{
			name: "should return -1",
			args: args{
				value: 10,
				list:  []int{},
			},
			want:      -1,
			wantLoops: 0,
		},
		{
			name: "should return -1",
			args: args{
				value: 0,
				list:  []int{10, 20, 30, 40, 50},
			},
			want:      -1,
			wantLoops: 2,
		},
		{
			name: "should find index",
			args: args{
				value: 10,
				list:  []int{10, 20, 30, 40, 50},
			},
			want:      0,
			wantLoops: 2,
		},
		{
			name: "should find index",
			args: args{
				value: 20,
				list:  []int{10, 20, 30, 40, 50},
			},
			want:      1,
			wantLoops: 3,
		},
		{
			name: "should find index",
			args: args{
				value: 30,
				list:  []int{10, 20, 30, 40, 50},
			},
			want:      2,
			wantLoops: 1,
		},
		{
			name: "should find index",
			args: args{
				value: 40,
				list:  []int{10, 20, 30, 40, 50},
			},
			want:      3,
			wantLoops: 2,
		},
		{
			name: "should find index",
			args: args{
				value: 50,
				list:  []int{10, 20, 30, 40, 50},
			},
			want:      4,
			wantLoops: 3,
		},
		{
			name: "should find index",
			args: args{
				value: 20,
				list:  []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100},
			},
			want:      1,
			wantLoops: 2,
		},
		{
			name: "should find index",
			args: args{
				value: 30,
				list:  []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100},
			},
			want:      2,
			wantLoops: 3,
		},
		{
			name: "should find index",
			args: args{
				value: 80,
				list:  []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100},
			},
			want:      7,
			wantLoops: 2,
		},
		{
			name: "should find index",
			args: args{
				value: 70,
				list:  []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100},
			},
			want:      6,
			wantLoops: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			index, count := Search(tt.args.value, tt.args.list)
			if index != tt.want {
				t.Errorf("Search() index = %v, want %v", index, tt.want)
			}
			if count != tt.wantLoops {
				t.Errorf("Search() count = %v, want %v", count, tt.wantLoops)
			}
		})
	}
}

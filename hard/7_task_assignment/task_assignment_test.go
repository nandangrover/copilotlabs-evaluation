package main

import (
	"reflect"
	"testing"
)

func TestTaskAssignment(t *testing.T) {
	type args struct {
		k     int
		tasks []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "Test Case #1",
			args: args{
				k:     3,
				tasks: []int{1, 3, 5, 3, 1, 4},
			},
			want: [][]int{
				{4, 2},
				{0, 5},
				{3, 1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := taskAssignment(tt.args.k, tt.args.tasks); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("TaskAssignment() = %v, want %v", got, tt.want)
			}
		})
	}
}

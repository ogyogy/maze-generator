package main

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
)

func TestInitMaze(t *testing.T) {
	type args struct {
		height int
		width  int
	}
	tests := []struct {
		name      string
		args      args
		want_maze [][]int
		want_err  error
	}{
		{
			// 正常系
			name:      "positive",
			args:      args{height: 5, width: 5},
			want_maze: [][]int{{0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}},
			want_err:  nil,
		},
		{
			// 異常系 (高さが5未満)
			name:      "negative",
			args:      args{height: 4, width: 5},
			want_maze: nil,
			want_err:  fmt.Errorf("error: height and width must be at least 5"),
		},
		{
			// 異常系 (幅が5未満)
			name:      "negative",
			args:      args{height: 5, width: 4},
			want_maze: nil,
			want_err:  fmt.Errorf("error: height and width must be at least 5"),
		},
		{
			// 異常系 (高さ、幅ともに5未満)
			name:      "negative",
			args:      args{height: 4, width: 4},
			want_maze: nil,
			want_err:  fmt.Errorf("error: height and width must be at least 5"),
		},
		{
			// 異常系 (高さが偶数)
			name:      "negative",
			args:      args{height: 6, width: 5},
			want_maze: nil,
			want_err:  fmt.Errorf("error: height and width must be odd numbers"),
		},
		{
			// 異常系 (幅が偶数)
			name:      "negative",
			args:      args{height: 5, width: 6},
			want_maze: nil,
			want_err:  fmt.Errorf("error: height and width must be odd numbers"),
		},
		{
			// 異常系 (高さ、幅ともに偶数)
			name:      "negative",
			args:      args{height: 6, width: 6},
			want_maze: nil,
			want_err:  fmt.Errorf("error: height and width must be odd numbers"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got_field, got_err := initMaze(tt.args.height, tt.args.width); !reflect.DeepEqual(got_field, tt.want_maze) || (got_err != nil && got_err.Error() != tt.want_err.Error()) {
				t.Errorf("field = %v, err = %v, want field = %v, err = %v, %v, %v", got_field, got_err, tt.want_maze, tt.want_err, reflect.DeepEqual(got_field, tt.want_maze), errors.Is(got_err, tt.want_err))
			}
		})
	}
}

func TestDigMaze(t *testing.T) {
	type args struct {
		maze [][]int
		x    int
		y    int
	}
	tests := []struct {
		name string
		args args
		want error
	}{
		{name: "positive",
			args: args{
				maze: [][]int{{0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}},
				x:    1,
				y:    1},
			want: nil},
		{name: "negative",
			args: args{maze: [][]int{{0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}},
				x: 0,
				y: 0},
			want: fmt.Errorf("error: x and y must be odd numbers greater than or equal to 1")},
		{name: "negative",
			args: args{maze: [][]int{{0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}},
				x: 0,
				y: 1},
			want: fmt.Errorf("error: x and y must be odd numbers greater than or equal to 1")},
		{name: "negative",
			args: args{maze: [][]int{{0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}},
				x: 1,
				y: 0},
			want: fmt.Errorf("error: x and y must be odd numbers greater than or equal to 1")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := digMaze(tt.args.maze, tt.args.x, tt.args.y); got != nil && got.Error() != tt.want.Error() {
				t.Errorf("got = %v, tt.want = %v", got, tt.want)
			} else if got == nil && tt.args.maze[tt.args.y][tt.args.x] != 1 {
				t.Errorf("Invalid value")
			}
		})
	}
}

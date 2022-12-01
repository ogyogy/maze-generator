package main

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
)

func TestInitField(t *testing.T) {
	type args struct {
		height int
		width  int
	}
	tests := []struct {
		name        string
		args        args
		want_fileld [][]int
		want_err    error
	}{
		{
			// 正常系
			name:        "positive",
			args:        args{height: 5, width: 5},
			want_fileld: [][]int{{0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}},
			want_err:    nil,
		},
		{
			// 異常系 (高さが5未満)
			name:        "negative",
			args:        args{height: 4, width: 5},
			want_fileld: nil,
			want_err:    fmt.Errorf("error: height and width must be at least 5"),
		},
		{
			// 異常系 (幅が5未満)
			name:        "negative",
			args:        args{height: 5, width: 4},
			want_fileld: nil,
			want_err:    fmt.Errorf("error: height and width must be at least 5"),
		},
		{
			// 異常系 (高さ、幅ともに5未満)
			name:        "negative",
			args:        args{height: 4, width: 4},
			want_fileld: nil,
			want_err:    fmt.Errorf("error: height and width must be at least 5"),
		},
		{
			// 異常系 (高さが偶数)
			name:        "negative",
			args:        args{height: 6, width: 5},
			want_fileld: nil,
			want_err:    fmt.Errorf("error: height and width must be odd numbers"),
		},
		{
			// 異常系 (幅が偶数)
			name:        "negative",
			args:        args{height: 5, width: 6},
			want_fileld: nil,
			want_err:    fmt.Errorf("error: height and width must be odd numbers"),
		},
		{
			// 異常系 (高さ、幅ともに偶数)
			name:        "negative",
			args:        args{height: 6, width: 6},
			want_fileld: nil,
			want_err:    fmt.Errorf("error: height and width must be odd numbers"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got_field, got_err := initMaze(tt.args.height, tt.args.width); !reflect.DeepEqual(got_field, tt.want_fileld) || (got_err != nil && got_err.Error() != tt.want_err.Error()) {
				t.Errorf("field = %v, err = %v, want field = %v, err = %v, %v, %v", got_field, got_err, tt.want_fileld, tt.want_err, reflect.DeepEqual(got_field, tt.want_fileld), errors.Is(got_err, tt.want_err))
			}
		})
	}
}

func TestInitCoordinate(t *testing.T) {
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
		{name: "normal", args: args{maze: [][]int{{0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}}, x: 1, y: 1}, want: nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := initCoordinate(tt.args.maze, tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("got = %v, tt.want = %v", got, tt.want)
			} else if got != nil && tt.args.maze[tt.args.y][tt.args.x] != 1 {
				t.Errorf("error")
			}
		})
	}
}

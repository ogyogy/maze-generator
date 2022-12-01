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
			m := Maze{tt.args.height, tt.args.width, nil}
			if got := m.InitMaze(tt.args.height, tt.args.width); !reflect.DeepEqual(m.Maze, tt.want_maze) || (got != nil && got.Error() != tt.want_err.Error()) {
				// 初期生成した迷路の値またはエラーメッセージが想定と異なる場合テスト失敗
				t.Errorf("field = %v, err = %v, want field = %v, err = %v, %v, %v", m.Maze, got, tt.want_maze, tt.want_err, reflect.DeepEqual(m.Maze, tt.want_maze), errors.Is(got, tt.want_err))
			}
		})
	}
}

func TestGenerateMaze(t *testing.T) {
	type args struct {
		height int
		width  int
		x      int
		y      int
	}
	tests := []struct {
		name string
		args args
		want error
	}{
		// 正常系
		{name: "positive",
			args: args{
				height: 5,
				width:  5,
				x:      1,
				y:      1},
			want: nil},
		// 異常系 (高さが5未満)
		{name: "negative",
			args: args{
				height: 4,
				width:  5,
				x:      1,
				y:      1},
			want: fmt.Errorf("error: height and width must be at least 5")},
		// 異常系 (幅が5未満)
		{name: "negative",
			args: args{
				height: 5,
				width:  4,
				x:      1,
				y:      1},
			want: fmt.Errorf("error: height and width must be at least 5")},
		// 異常系 (高さが偶数)
		{name: "negative",
			args: args{
				height: 6,
				width:  5,
				x:      1,
				y:      1},
			want: fmt.Errorf("error: height and width must be odd numbers")},
		// 異常系 (幅が偶数)
		{name: "negative",
			args: args{
				height: 5,
				width:  6,
				x:      1,
				y:      1},
			want: fmt.Errorf("error: height and width must be odd numbers")},
		// 異常系 (x座標が偶数)
		{name: "negative",
			args: args{
				height: 5,
				width:  5,
				x:      0,
				y:      1},
			want: fmt.Errorf("error: x and y must be odd numbers greater than or equal to 1")},
		// 異常系 (y座標が偶数)
		{name: "negative",
			args: args{
				height: 5,
				width:  5,
				x:      1,
				y:      0},
			want: fmt.Errorf("error: x and y must be odd numbers greater than or equal to 1")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := Maze{tt.args.height, tt.args.width, nil}
			if got := m.GenerateMaze(tt.args.height, tt.args.width, tt.args.x, tt.args.y); got != nil && got.Error() != tt.want.Error() {
				// エラーメッセージが想定と異なる場合テスト失敗
				t.Errorf("got = %v, tt.want = %v", got, tt.want)
			} else if got == nil && m.Maze[tt.args.y][tt.args.x] == WALL {
				// 初期座標が壁の場合テスト失敗
				t.Errorf("initialization failed")
			}
		})
	}
}

func TestDigMaze(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want error
	}{
		// 正常系
		{name: "positive",
			args: args{
				x: 1,
				y: 1},
			want: nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := Maze{5, 5, nil}
			if got := m.DigMaze(tt.args.x, tt.args.y); got != nil {
				t.Errorf("got != nil")
			}
		})
	}
}

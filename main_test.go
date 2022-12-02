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
			if got := m.InitMaze(); !reflect.DeepEqual(m.Maze, tt.want_maze) || (got != nil && got.Error() != tt.want_err.Error()) {
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
			if got := m.GenerateMaze(tt.args.x, tt.args.y); got != nil && got.Error() != tt.want.Error() {
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
			m.InitMaze()
			if got := m.DigMaze(tt.args.x, tt.args.y); got != nil {
				t.Errorf("got != nil")
			}
		})
	}
}

func TestMovePlayer(t *testing.T) {
	type args struct {
		x         int
		y         int
		direction int
	}
	tests := []struct {
		name     string
		args     args
		want_err error
		want_x   int
		want_y   int
	}{
		// 正常系
		{name: "positive", args: args{x: 2, y: 2, direction: UP}, want_err: nil, want_x: 2, want_y: 1},
		{name: "positive", args: args{x: 2, y: 2, direction: DOWN}, want_err: nil, want_x: 2, want_y: 3},
		{name: "positive", args: args{x: 2, y: 2, direction: LEFT}, want_err: nil, want_x: 1, want_y: 2},
		{name: "positive", args: args{x: 2, y: 2, direction: RIGHT}, want_err: nil, want_x: 3, want_y: 2},
		// 異常系 (移動先が壁)
		{name: "negative", args: args{x: 1, y: 1, direction: UP}, want_err: fmt.Errorf("you cannot walk through walls"), want_x: 1, want_y: 1},
		{name: "negative", args: args{x: 3, y: 3, direction: DOWN}, want_err: fmt.Errorf("you cannot walk through walls"), want_x: 3, want_y: 3},
		{name: "negative", args: args{x: 1, y: 1, direction: LEFT}, want_err: fmt.Errorf("you cannot walk through walls"), want_x: 1, want_y: 1},
		{name: "negative", args: args{x: 3, y: 3, direction: RIGHT}, want_err: fmt.Errorf("you cannot walk through walls"), want_x: 3, want_y: 3},
		// 異常系 (移動先が範囲外)
		{name: "negative", args: args{x: 0, y: 0, direction: UP}, want_err: fmt.Errorf("index out of range"), want_x: 0, want_y: 0},
		{name: "negative", args: args{x: 4, y: 4, direction: DOWN}, want_err: fmt.Errorf("index out of range"), want_x: 4, want_y: 4},
		{name: "negative", args: args{x: 0, y: 0, direction: LEFT}, want_err: fmt.Errorf("index out of range"), want_x: 0, want_y: 0},
		{name: "negative", args: args{x: 4, y: 4, direction: RIGHT}, want_err: fmt.Errorf("index out of range"), want_x: 4, want_y: 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := Maze{5, 5, nil}
			m.InitMaze()
			height, width := m.Height, m.Width
			p := Player{tt.args.x, tt.args.y}
			// 外周以外を通路にする
			for i := 0; i < height; i++ {
				for j := 0; j < width; j++ {
					if i > 0 && j > 0 && i < width-1 && j < height-1 {
						m.Maze[i][j] = PATH
					}
				}
			}
			got := p.MovePlayer(m, tt.args.direction)
			if got != nil && got.Error() != tt.want_err.Error() {
				t.Errorf("got != nil && got.Error() != tt.want_err.Error()")
			}
			if p.X != tt.want_x || p.Y != tt.want_y {
				t.Errorf("p.X = %v, tt.want_x = %v, p.Y = %v, tt.want_y = %v", p.X, tt.want_x, p.Y, tt.want_y)
			}
		})
	}
}

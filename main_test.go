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
		name     string
		args     args
		wantMaze [][]int
		wantErr  error
	}{
		{
			// 正常系
			name:     "positive",
			args:     args{height: 5, width: 5},
			wantMaze: [][]int{{0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}, {0, 0, 0, 0, 0}},
			wantErr:  nil,
		},
		{
			// 異常系 (高さが5未満)
			name:     "negative",
			args:     args{height: 4, width: 5},
			wantMaze: nil,
			wantErr:  fmt.Errorf("error: height and width must be at least 5"),
		},
		{
			// 異常系 (幅が5未満)
			name:     "negative",
			args:     args{height: 5, width: 4},
			wantMaze: nil,
			wantErr:  fmt.Errorf("error: height and width must be at least 5"),
		},
		{
			// 異常系 (高さ、幅ともに5未満)
			name:     "negative",
			args:     args{height: 4, width: 4},
			wantMaze: nil,
			wantErr:  fmt.Errorf("error: height and width must be at least 5"),
		},
		{
			// 異常系 (高さが偶数)
			name:     "negative",
			args:     args{height: 6, width: 5},
			wantMaze: nil,
			wantErr:  fmt.Errorf("error: height and width must be odd numbers"),
		},
		{
			// 異常系 (幅が偶数)
			name:     "negative",
			args:     args{height: 5, width: 6},
			wantMaze: nil,
			wantErr:  fmt.Errorf("error: height and width must be odd numbers"),
		},
		{
			// 異常系 (高さ、幅ともに偶数)
			name:     "negative",
			args:     args{height: 6, width: 6},
			wantMaze: nil,
			wantErr:  fmt.Errorf("error: height and width must be odd numbers"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := Maze{tt.args.height, tt.args.width, nil}
			if got := m.Init(); !reflect.DeepEqual(m.Maze, tt.wantMaze) || (got != nil && got.Error() != tt.wantErr.Error()) {
				// 初期生成した迷路の値またはエラーメッセージが想定と異なる場合テスト失敗
				t.Errorf("field = %v, err = %v, want field = %v, err = %v, %v, %v", m.Maze, got, tt.wantMaze, tt.wantErr, reflect.DeepEqual(m.Maze, tt.wantMaze), errors.Is(got, tt.wantErr))
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
			m.Init()
			if got := m.Dig(tt.args.x, tt.args.y); got != nil {
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
		name    string
		args    args
		wantErr error
		wantX   int
		wantY   int
	}{
		// 正常系
		{name: "positive", args: args{x: 2, y: 2, direction: UP}, wantErr: nil, wantX: 2, wantY: 1},
		{name: "positive", args: args{x: 2, y: 2, direction: DOWN}, wantErr: nil, wantX: 2, wantY: 3},
		{name: "positive", args: args{x: 2, y: 2, direction: LEFT}, wantErr: nil, wantX: 1, wantY: 2},
		{name: "positive", args: args{x: 2, y: 2, direction: RIGHT}, wantErr: nil, wantX: 3, wantY: 2},
		// 異常系 (移動先が壁)
		{name: "negative", args: args{x: 1, y: 1, direction: UP}, wantErr: fmt.Errorf("you cannot walk through walls"), wantX: 1, wantY: 1},
		{name: "negative", args: args{x: 3, y: 3, direction: DOWN}, wantErr: fmt.Errorf("you cannot walk through walls"), wantX: 3, wantY: 3},
		{name: "negative", args: args{x: 1, y: 1, direction: LEFT}, wantErr: fmt.Errorf("you cannot walk through walls"), wantX: 1, wantY: 1},
		{name: "negative", args: args{x: 3, y: 3, direction: RIGHT}, wantErr: fmt.Errorf("you cannot walk through walls"), wantX: 3, wantY: 3},
		// 異常系 (移動先が範囲外)
		{name: "negative", args: args{x: 0, y: 0, direction: UP}, wantErr: fmt.Errorf("index out of range"), wantX: 0, wantY: 0},
		{name: "negative", args: args{x: 4, y: 4, direction: DOWN}, wantErr: fmt.Errorf("index out of range"), wantX: 4, wantY: 4},
		{name: "negative", args: args{x: 0, y: 0, direction: LEFT}, wantErr: fmt.Errorf("index out of range"), wantX: 0, wantY: 0},
		{name: "negative", args: args{x: 4, y: 4, direction: RIGHT}, wantErr: fmt.Errorf("index out of range"), wantX: 4, wantY: 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := Maze{5, 5, nil}
			m.Init()
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
			got := p.Move(m, tt.args.direction)
			if got != nil && got.Error() != tt.wantErr.Error() {
				t.Errorf("got != nil && got.Error() != tt.want_err.Error()")
			}
			if p.X != tt.wantX || p.Y != tt.wantY {
				t.Errorf("p.X = %v, tt.want_x = %v, p.Y = %v, tt.want_y = %v", p.X, tt.wantX, p.Y, tt.wantY)
			}
		})
	}
}

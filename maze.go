package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Maze struct {
	Height int
	Width  int
	Maze   [][]int
	GoalX  int
	GoalY  int
}

// 迷路を初期化
// 高さ、幅はそれぞれ5以上の奇数を指定する必要がある
func (m *Maze) Init() error {
	height, width := m.Height, m.Width
	if height < 5 || width < 5 {
		// 高さまたは幅が5未満の場合エラー
		// 空のスライスとエラーメッセージを返却
		// スライスのゼロ値はnil
		// nilスライスは0の長さと容量を持つ
		return fmt.Errorf("error: height and width must be at least 5")
	} else if height%2 == WALL || width%2 == WALL {
		return fmt.Errorf("error: height and width must be odd numbers")
	}
	m.Maze = make([][]int, height)
	for i := range m.Maze {
		// field[i] = 0で初期化
		m.Maze[i] = make([]int, width)
	}
	return nil
}

// 座標(x, y)を起点に迷路に穴を掘る
func (m *Maze) Dig(x, y int) error {
	rand.Seed(time.Now().UnixNano())
	// 上下左右に進めるか
	var up, down, left, right = true, true, true, true
	// 掘り進められなくなるまでループ
	for up || down || left || right {
		// ランダムに方向を選択
		direction := rand.Intn(4)
		// 上方向
		if direction == UP {
			if y-2 > 0 && m.Maze[y-2][x] == WALL {
				// 2マス先が迷路の範囲内かつ壁の場合掘り進める
				m.Maze[y-1][x] = PATH
				m.Maze[y-2][x] = PATH
				m.Dig(x, y-2)
			} else {
				up = false
			}
		}
		// 下方向
		if direction == DOWN {
			if y+2 < m.Height && m.Maze[y+2][x] == WALL {
				// 2マス先が迷路の範囲内かつ壁の場合掘り進める
				m.Maze[y+1][x] = PATH
				m.Maze[y+2][x] = PATH
				m.Dig(x, y+2)
			} else {
				down = false
			}
		}
		// 左方向
		if direction == LEFT {
			if x-2 > 0 && m.Maze[y][x-2] == WALL {
				// 2マス先が迷路の範囲内かつ壁の場合掘り進める
				m.Maze[y][x-1] = PATH
				m.Maze[y][x-2] = PATH
				m.Dig(x-2, y)
			} else {
				left = false
			}
		}
		// 上方向
		if direction == RIGHT {
			if x+2 < m.Width && m.Maze[y][x+2] == WALL {
				// 2マス先が迷路の範囲内かつ壁の場合掘り進める
				m.Maze[y][x+1] = PATH
				m.Maze[y][x+2] = PATH
				m.Dig(x+2, y)
			} else {
				right = false
			}
		}
	}
	return nil
}

// 穴掘り法で迷路を生成
// 初期化済みの迷路maze、初期座標(x, y)を指定する
// 初期座標は奇数である必要がある
func (m *Maze) GenerateMaze(x, y int) error {
	// 初期生成
	err := m.Init()
	if err != nil {
		// 初期生成で異常が発生した場合はエラー
		return err
	} else if x < 1 || y < 1 || x%2 == 0 || y%2 == 0 {
		// 初期座標に偶数が指定された場合エラー
		return fmt.Errorf("error: x and y must be odd numbers greater than or equal to 1")
	} else {
		m.Maze[y][x] = PATH
	}
	m.Dig(x, y)
	return nil
}

// ゴールの座標をランダムに設定
func (m *Maze) SetRandGoal() {
	rand.Seed(time.Now().UnixNano())
	var gx, gy int
	for m.Maze[gy][gx] != PATH {
		gx, gy = rand.Intn(m.Width), rand.Intn(m.Height)
	}
	m.Maze[gy][gx] = GOAL
	m.GoalX, m.GoalY = gx, gy
}

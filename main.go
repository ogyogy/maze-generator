package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	WALL = iota
	PATH
	GOAL
)

const (
	UP = iota
	DOWN
	LEFT
	RIGHT
)

type Maze struct {
	Height int
	Width  int
	Maze   [][]int
}

type Player struct {
	X int
	Y int
}

type Game struct {
	Maze   Maze
	Player Player
}

// 迷路を初期化
// 高さ、幅はそれぞれ5以上の奇数を指定する必要がある
func (m *Maze) InitMaze() error {
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
func (m *Maze) DigMaze(x, y int) error {
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
				m.DigMaze(x, y-2)
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
				m.DigMaze(x, y+2)
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
				m.DigMaze(x-2, y)
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
				m.DigMaze(x+2, y)
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
	err := m.InitMaze()
	if err != nil {
		// 初期生成で異常が発生した場合はエラー
		return err
	} else if x < 1 || y < 1 || x%2 == 0 || y%2 == 0 {
		// 初期座標に偶数が指定された場合エラー
		return fmt.Errorf("error: x and y must be odd numbers greater than or equal to 1")
	} else {
		m.Maze[y][x] = PATH
	}
	m.DigMaze(x, y)
	return nil
}

// 迷路を標準出力
func (g *Game) DisplayMaze() {
	fmt.Println(g.Player)
	for i, v := range g.Maze.Maze {
		for j, vv := range v {
			if vv == WALL {
				fmt.Print("#")
			} else if vv == PATH {
				if g.Player.X == j && g.Player.Y == i {
					fmt.Print("@")
				} else {
					fmt.Print(".")
				}
			} else if vv == GOAL {
				fmt.Print("G")
			}
		}
		fmt.Println()
	}
}

// プレイヤーを移動する
func (p *Player) MovePlayer(m Maze, direction int) error {
	newX, newY := p.X, p.Y
	// 引数directionに基づき移動先の座標を計算
	if direction == UP {
		newY--
	} else if direction == DOWN {
		newY++
	} else if direction == LEFT {
		newX--
	} else if direction == RIGHT {
		newX++
	}
	if newX < 0 || newY < 0 || newX > m.Width-1 || newY > m.Height-1 {
		// 移動先の座標が範囲外の場合エラー
		return fmt.Errorf("index out of range")
	} else if m.Maze[newY][newX] == WALL {
		// 移動先の座標が壁の場合エラー
		return fmt.Errorf("you cannot walk through walls")
	} else {
		// エラーに該当しない場合はプレイヤーの座標を更新
		p.X = newX
		p.Y = newY
	}
	return nil
}

// プレイヤーの座標を初期化
func (p *Player) SetRandCoord(m Maze) {
	rand.Seed(time.Now().UnixNano())
	for {
		sx, sy := rand.Intn(m.Width), rand.Intn(m.Height)
		if m.Maze[sx][sy] == PATH {
			p.X, p.Y = sx, sy
			break
		}
	}
}

// ゴールの座標を設定
func (m *Maze) setGoal() {
	rand.Seed(time.Now().UnixNano())
	for {
		gx, gy := rand.Intn(m.Width), rand.Intn(m.Height)
		if m.Maze[gx][gy] == PATH {
			m.Maze[gx][gy] = GOAL
			break
		}
	}
}

// ゲームの実行
func (g *Game) Run() {
	rand.Seed(time.Now().UnixNano())
	// 迷路の生成
	g.Maze.InitMaze()
	// 穴掘り法の開始座標をランダム生成
	sx, sy := rand.Intn(g.Maze.Width-1), rand.Intn(g.Maze.Height-1)
	if sx%2 == 0 {
		sx++
	}
	if sy%2 == 0 {
		sy++
	}
	err := g.Maze.GenerateMaze(sx, sy)
	// プレイヤーの生成
	g.Player.SetRandCoord(g.Maze)
	// ゴールの生成
	g.Maze.setGoal()
	if err != nil {
		fmt.Println(err)
	} else {
		g.DisplayMaze()
	}
}

func main() {
	// 迷路の初期化
	m := Maze{7, 7, nil}
	// プレイヤーの初期化
	p := Player{0, 0}
	// ゲームの実行
	g := Game{m, p}
	g.Run()
}

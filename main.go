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

type Game struct {
	Maze   Maze
	Player Player
}

// 迷路を標準出力
func (g *Game) DisplayMaze() {
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

// ゲームの実行
func (g *Game) Run() error {
	rand.Seed(time.Now().UnixNano())
	// 迷路の初期化
	err := g.Maze.Init()
	if err != nil {
		return err
	}
	// 穴掘り法の開始座標をランダム生成
	sx, sy := rand.Intn(g.Maze.Width-1), rand.Intn(g.Maze.Height-1)
	if sx%2 == 0 {
		sx++
	}
	if sy%2 == 0 {
		sy++
	}
	// 穴掘り法による迷路の生成
	err = g.Maze.GenerateMaze(sx, sy)
	if err != nil {
		return err
	}
	// プレイヤーの生成
	g.Player.SetRandCoord(g.Maze)
	// ゴールの生成
	gx, gy := g.Maze.SetRandGoal()
	// 迷路を表示
	g.DisplayMaze()
	// 移動
	var s string
	for {
		fmt.Print("move (hjkl), quit (q): ")
		_, err := fmt.Scan(&s)
		if err != nil {
			return err
		}
		// プレイヤーの移動量
		var d int
		if s == "h" {
			d = LEFT
		} else if s == "j" {
			d = DOWN
		} else if s == "k" {
			d = UP
		} else if s == "l" {
			d = RIGHT
		} else if s == "q" {
			break
		} else {
			fmt.Println("try again")
			continue
		}
		err = g.Player.Move(g.Maze, d)
		if err != nil {
			fmt.Println(err)
		}
		// ゴール判定
		if gx == g.Player.X && gy == g.Player.Y {
			fmt.Println("finish")
			return nil
		}
		g.DisplayMaze()
	}
	return nil
}

func main() {
	// 迷路の初期化
	m := Maze{7, 7, nil}
	// プレイヤーの初期化
	p := Player{0, 0}
	// ゲームの実行
	g := Game{m, p}
	err := g.Run()
	if err != nil {
		fmt.Println(err)
	}
}

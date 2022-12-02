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

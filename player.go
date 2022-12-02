package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Player struct {
	X int
	Y int
}

// プレイヤーを移動する
func (p *Player) Move(m Maze, direction int) error {
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

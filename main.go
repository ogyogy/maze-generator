package main

import "fmt"

// 0は壁、1は通路を表す
const (
	WALL = 0
	PATH = 1
)

type Maze struct {
	Height int
	Width  int
	Maze   [][]int
}

// 迷路を高さheight * 幅widthマスで初期化
// 高さ、幅はそれぞれ5以上の奇数を指定する必要がある
func (m *Maze) InitMaze(height, width int) error {
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
	return nil
}

// 穴掘り法で迷路を生成
// 初期化済みの迷路maze、初期座標(x, y)を指定する
// 初期座標は奇数である必要がある
func (m *Maze) GenerateMaze(height, width, x, y int) error {
	// 初期生成
	err := m.InitMaze(height, width)
	if err != nil {
		// 初期生成で異常が発生した場合はエラー
		return err
	} else if x < 1 || y < 1 || x%2 == 0 || y%2 == 0 {
		// 初期座標に偶数が指定された場合エラー
		return fmt.Errorf("error: x and y must be odd numbers greater than or equal to 1")
	} else {
		m.Maze[y][x] = PATH
	}
	return nil
}

// func main() {
// 	err := generateMaze(5, 5, 1, 1)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// }

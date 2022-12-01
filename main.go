package main

import "fmt"

// 迷路を高さheight * 幅widthマスで初期化
// 0は壁、1は通路を表す
func initMaze(height, width int) ([][]int, error) {
	// 高さ、幅はそれぞれ5以上の奇数を指定する必要がある
	if height < 5 || width < 5 {
		// 空のスライスとエラーメッセージを返却
		// スライスのゼロ値はnil
		// nilスライスは0の長さと容量を持つ
		return nil, fmt.Errorf("error: height and width must be at least 5")
	} else if height%2 == 0 || width%2 == 0 {
		return nil, fmt.Errorf("error: height and width must be odd numbers")
	}
	field := make([][]int, height)
	for i := range field {
		// field[i] = 0で初期化
		field[i] = make([]int, width)
	}
	return field, nil
}

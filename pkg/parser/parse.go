package parser

import (
	"bufio"
	"io"
	"os"
	"sudoku/pkg/logger"
	"sudoku/ranks"
)

// Parse 从文件读取数独
func Parse(sudokuFile string) *ranks.Ranks {
	var sudoku [9][9]int
	f, err := os.OpenFile(sudokuFile, os.O_RDONLY, 0644)
	if err != nil {
		panic("无法读取文件")
	}
	r := bufio.NewReader(f)
	i := 0
	for {
		b, err := r.ReadByte()
		if err == io.EOF {
			break
		}
		// 非0-9的数字，继续往后写
		v := int(b - '0')
		if v < 0 || v > 9 {
			continue
		}
		sudoku[i/9][i%9] = v
		i++
	}
	if i != 81 {
		panic("错误的数据源")
	}
	logger.Infof("读取到的数独表%v", sudoku)
	return ranks.New(sudoku)
}

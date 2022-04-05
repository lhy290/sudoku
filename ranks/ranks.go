package ranks

import (
	"fmt"
	"os"
	"sudoku/pkg/logger"
)

// 9*9
type Ranks struct {
	sudoku [9][9]*box
	// 3*3 Box
	baskets [3][3]basket
	// 9*row
	rows [9]row
	// 9*column
	columns [9]column

	sureNumCount   int
	eliminateCount int
}

func New(input [9][9]int) *Ranks {
	var rows [9]row
	var columns [9]column
	var baskets [3][3]basket
	var sureNumCount int

	ret := &Ranks{}
	for i, line := range input {
		for j, v := range line {
			b := newBox(v, ret, i, j)
			b.row = &rows[i]
			b.column = &columns[j]
			b.basket = &baskets[i/3][j/3]
			ret.sudoku[i][j] = b
			rows[i][j] = b
			columns[j][i] = b
			baskets[i/3][j/3][i%3][j%3] = b
			if b.isPerfect() {
				sureNumCount++
			}
		}
	}

	ret.rows = rows
	ret.columns = columns
	ret.baskets = baskets
	ret.sureNumCount = sureNumCount
	return ret
}

// Copy 相当于从old数据中重新new一个
func Copy(old *Ranks) *Ranks {
	var oldData [9][9]int
	for i := range old.sudoku {
		for j := range old.sudoku[i] {
			oldData[i][j] = old.sudoku[i][j].number
		}
	}
	ret := New(oldData)
	ret.eliminateCount = old.eliminateCount
	return ret
}

// Eliminate 简单的行列排查和栏目排查
func (r *Ranks) Eliminate() int {
	for _, v := range r.rows {
		v.eliminate()
	}
	for _, v := range r.columns {
		v.eliminate()
	}
	for _, line := range r.baskets {
		for _, v := range line {
			v.eliminate()
		}
	}
	r.eliminateCount++
	return r.sureNumCount
}

// Abandon 对每个格子使用摒除法
func (r *Ranks) Abandon() int {
	for _, row := range r.rows {
		for _, v := range row {
			if !v.isPerfect() {
				if v.ranksAbandon() {
					return r.sureNumCount
				}
			}
		}
	}
	return r.sureNumCount
}

// Trail 试错
func (r *Ranks) Trail() *Ranks {
	fmt.Println("===========================================================================")
	fmt.Printf("=============================== TRIAL START %d ===============================\n", r.eliminateCount)
	defer func() {
		fmt.Printf("================================ END TRIAL %d ================================\n", r.eliminateCount)
		fmt.Println("===========================================================================")
	}()

	// 在新的数独上试错，避免对正确的数独产生影响
	newSudoku := Copy(r)
	newSudoku.Eliminate()
	// 找到可能数字最少的格子试错
	minBox := newSudoku.getMinumBox()

	// 对可能的每个数字
	for k := range minBox.possible {
		// 把尝试的数字放入新数独
		trialSudoku := Copy(newSudoku)
		newBox := newBox(k, trialSudoku, minBox.x, minBox.y)
		trialSudoku.replaceBox(newBox)
		newBox.complete(k)

		// 做简单的排查
		before := trialSudoku.Eliminate()
		trialSudoku.Print()
		for {
			// 完成则退出
			if trialSudoku.Perfect() {
				trialSudoku.Print()
				return trialSudoku
			}
			// 非法则尝试下一个数字
			if !trialSudoku.IsValid() {
				break
			}

			// 第二次排查
			after := trialSudoku.Eliminate()
			trialSudoku.Print()
			// fmt.Printf("sureNum count: %d\n", trialSudoku.sureNumCount)
			// 试错没有进展，迭代试错
			if before == after {
				return trialSudoku.Trail()
			}
			before = after
			// if trialSudoku.eliminateCount > 1000 {
			// 	break
			// }
		}
	}

	return r
}

// Perfect 数独是否计算完成
func (r *Ranks) Perfect() bool {
	return r.sureNumCount == 81
}

// IsValid 判断当前数独数据是否合法
func (r *Ranks) IsValid() bool {
	for _, v := range r.rows {
		if !v.isValid() {
			return false
		}
	}
	for _, v := range r.columns {
		if !v.isValid() {
			return false
		}
	}
	for _, line := range r.baskets {
		for _, v := range line {
			if !v.isValid() {
				return false
			}
		}
	}
	return true
}

func (r *Ranks) Save(file string) {
	f, err := os.OpenFile(file, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		panic("无法写入文件 ")
	}
	f.WriteString(r.String())
}

func (r *Ranks) Print() {
	logger.Debugf("###############第%d次排除###############", r.eliminateCount)
	fmt.Printf("%v", r)
}

func (r *Ranks) String() string {
	var ret string
	for _, line := range r.sudoku {
		var row string
		for _, v := range line {
			row += fmt.Sprint(v.number)
		}
		ret += row + "\n"
	}
	return ret
}

// 未知数字最少的第一个box
func (r *Ranks) getMinumBox() *box {
	min := 9
	x, y := 0, 0
	for i, line := range r.sudoku {
		for j, v := range line {
			if !v.isPerfect() && len(v.possible) < min {
				min = len(v.possible)
				x, y = i, j
			}
		}
	}
	return r.sudoku[x][y]
}

func (r *Ranks) replaceBox(b *box) {
	x, y := b.x, b.y
	r.sudoku[x][y] = b
	r.rows[x][y] = b
	r.columns[y][x] = b
	r.baskets[x/3][y/3][x%3][y%3] = b
	if b.perfect {
		r.sureNumCount++
	}
}

package ranks

// 1*1
type box struct {
	// completed if number is not 0
	number int
	// possible nums
	possible map[int]struct{}
	// is perfect
	perfect bool
	ranks   *Ranks
	row     *row
	column  *column
	basket  *basket
	x       int
	y       int
}

func newBox(num int, r *Ranks, x int, y int) *box {
	possible := make(map[int]struct{})
	if num == 0 {
		possible = map[int]struct{}{
			1: {},
			2: {},
			3: {},
			4: {},
			5: {},
			6: {},
			7: {},
			8: {},
			9: {},
		}
	}
	return &box{
		number:   num,
		possible: possible,
		perfect:  len(possible) == 0,
		ranks:    r,
		x:        x,
		y:        y,
	}
}

func (b *box) isPerfect() bool {
	return b.perfect
}

// 从格子的可能数字中移除确定的数
func (b *box) rmPossibles(sureBoxes []int) {
	for _, v := range sureBoxes {
		delete(b.possible, v)
	}
	if len(b.possible) == 1 {
		b.complete(getSureNumber(b.possible))
	}
}

// 摒除法
func (b *box) ranksAbandon() bool {
	// 行、列、栏目分别做摒除法
	return b.rowAbandon() || b.columnAbandon() || b.basketAbandon()
	// 这个算法待验证，是否需要
	// otherPossibles := make(map[int]struct{})
	// for _, v := range b.row {
	// 	if !v.isSame(b) && !v.isPerfect() {
	// 		for p, _ := range v.possible {
	// 			otherPossibles[p] = struct{}{}
	// 		}
	// 	}
	// }
	// for _, v := range b.column {
	// 	if !v.isSame(b) && !v.isPerfect() {
	// 		for p, _ := range v.possible {
	// 			otherPossibles[p] = struct{}{}
	// 		}
	// 	}
	// }
	// for _, line := range b.basket {
	// 	for _, v := range line {
	// 		if !v.isSame(b) && !v.isPerfect() {
	// 			for p, _ := range v.possible {
	// 				otherPossibles[p] = struct{}{}
	// 			}
	// 		}
	// 	}
	// }

	// possibles := make(map[int]struct{})
	// for k, _ := range b.possible {
	// 	possibles[k] = struct{}{}
	// }
	// for k, _ := range otherPossibles {
	// 	if _, ok := possibles[k]; ok {
	// 		delete(possibles, k)
	// 	}
	// }
	// if len(possibles) == 1 {
	// 	b.complete(getSureNumber(possibles))
	// }
}

// 一行中，某一个数仅在这个格子中出现，则必定属于这个格子
func (b *box) rowAbandon() bool {
	otherPossibles := make(map[int]struct{})
	for _, v := range b.row {
		if !v.isSame(b) && !v.isPerfect() {
			for p := range v.possible {
				otherPossibles[p] = struct{}{}
			}
		}
	}
	possibles := make(map[int]struct{})
	for k := range b.possible {
		possibles[k] = struct{}{}
	}
	for k := range otherPossibles {
		delete(possibles, k)
	}
	if len(possibles) == 1 {
		b.complete(getSureNumber(possibles))
		return true
	}
	return false
}

// 一列中，某一个数仅在这个格子中出现，则必定属于这个格子
func (b *box) columnAbandon() bool {
	otherPossibles := make(map[int]struct{})
	for _, v := range b.column {
		if !v.isSame(b) && !v.isPerfect() {
			for p := range v.possible {
				otherPossibles[p] = struct{}{}
			}
		}
	}
	possibles := make(map[int]struct{})
	for k := range b.possible {
		possibles[k] = struct{}{}
	}
	for k := range otherPossibles {
		delete(possibles, k)
	}
	if len(possibles) == 1 {
		b.complete(getSureNumber(possibles))
		return true
	}
	return false
}

// 一个栏目中，某一个数仅在这个格子中出现，则必定属于这个格子
func (b *box) basketAbandon() bool {
	otherPossibles := make(map[int]struct{})
	for _, line := range b.basket {
		for _, v := range line {
			if !v.isSame(b) && !v.isPerfect() {
				for p := range v.possible {
					otherPossibles[p] = struct{}{}
				}
			}
		}
	}
	possibles := make(map[int]struct{})
	for k := range b.possible {
		possibles[k] = struct{}{}
	}
	for k := range otherPossibles {
		delete(possibles, k)
	}
	if len(possibles) == 1 {
		b.complete(getSureNumber(possibles))
		return true
	}
	return false
}

// b和o是否是同一个坐标
func (b *box) isSame(o *box) bool {
	return b.x == o.x && b.y == o.y
}

func (b *box) complete(num int) {
	if num == 0 {
		panic("不能用0完成一个box")
	}
	if !b.perfect {
		b.perfect = true
		b.number = num
		b.ranks.sureNumCount++
	}
}

func getSureNumber(nums map[int]struct{}) int {
	if len(nums) != 1 {
		panic("无法从多个数中返回确定的一个数")
	}
	ret := 0
	for k := range nums {
		ret = k
	}
	return ret
}

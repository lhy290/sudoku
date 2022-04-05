package ranks

// 3*3
type basket [3][3]*box

func (b *basket) eliminate() {
	var sureNums []int
	for _, line := range b {
		for _, v := range line {
			if v.isPerfect() {
				sureNums = append(sureNums, v.number)
			}
		}
	}
	for _, line := range b {
		for _, v := range line {
			if !v.isPerfect() {
				v.rmPossibles(sureNums)
			}
		}
	}
}

func (b *basket) isValid() bool {
	boxes := make(map[int]struct{})
	for _, line := range b {
		for _, v := range line {
			if !v.isPerfect() {
				continue
			}
			if _, ok := boxes[v.number]; ok {
				return false
			}
			boxes[v.number] = struct{}{}
		}
	}
	return true
}

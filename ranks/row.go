package ranks

// 1*9
type row [9]*box

func (r *row) eliminate() {
	var sureNums []int
	for _, v := range r {
		if v.isPerfect() {
			sureNums = append(sureNums, v.number)
		}
	}
	for _, v := range r {
		if !v.isPerfect() {
			v.rmPossibles(sureNums)
		}
	}
}

func (r *row) isValid() bool {
	boxes := make(map[int]struct{})
	for _, b := range r {
		if !b.isPerfect() {
			continue
		}
		if _, ok := boxes[b.number]; ok {
			return false
		}
		boxes[b.number] = struct{}{}
	}
	return true
}

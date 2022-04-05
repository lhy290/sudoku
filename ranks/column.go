package ranks

// 9*1
type column [9]*box

func (c *column) eliminate() {
	var sureNums []int
	for _, v := range c {
		if v.isPerfect() {
			sureNums = append(sureNums, v.number)
		}
	}
	for _, v := range c {
		if !v.isPerfect() {
			v.rmPossibles(sureNums)
		}
	}
}

func (c *column) isValid() bool {
	boxes := make(map[int]struct{})
	for _, b := range c {
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

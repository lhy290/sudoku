package eliminater

import "sudoku/pkg/parser"

type Eliminater struct {
	in  string
	out string
}

func New(in, out string) *Eliminater {
	return &Eliminater{
		in:  in,
		out: out,
	}
}

// Run 计算和排除工作
func (e *Eliminater) Run() {
	r := parser.Parse(e.in)
	if !r.IsValid() {
		panic("读取的数独不合法，请检查确认")
	}
	// 先计算一次
	before := r.Eliminate()
	r.Print()
	for {
		// 再计算第二次
		after := r.Eliminate()
		r.Print()
		if r.Perfect() {
			r.Save(e.out)
			break
		}
		if before == after {
			// 如果不能增加确定数字的格子，则是用摒除法
			after = r.Abandon()
			r.Print()
		}
		if before == after {
			// 连使用摒除法都解决不了，则开始试错
			r = r.Trail()
		}
		before = after
	}
}

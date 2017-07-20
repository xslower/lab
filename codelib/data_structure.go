package main

import (
	"fmt"
)

type leaf struct {
	d string
	l *leaf
	r *leaf
}

func NewStack() Stack {
	return Stack{0, make([]*leaf, 100)}
}

type Stack struct {
	pos int
	stk []*leaf
}

func (this *Stack) Push(l *leaf) {
	this.stk[this.pos] = l
	this.pos++
}
func (this *Stack) Pop() *leaf {
	if this.pos == 0 {
		return nil
	}
	this.pos--
	return this.stk[this.pos]
}
func (this *Stack) Len() int {
	return this.pos
}
func init() {

}

func visit(p *leaf) {
	fmt.Print(p.d)
}

/**
 * 二叉树遍历
 */
func traverseTree(p *leaf) {
	var stack = NewStack()
	stack.Push(p)
	//前序
	// for stack.Len() > 0 {
	// 	p = stack.Pop()
	// 	visit(p)
	// 	if p.r != nil {
	// 		stack.Push(p.r)
	// 	}
	// 	if p.l != nil {
	// 		stack.Push(p.l)
	// 	}
	// }
	//中序
	// for p != nil || stack.Len() > 0 {
	// 	if p != nil {
	// 		stack.Push(p)
	// 		p = p.l
	// 	} else {
	// 		p = stack.Pop()
	// 		visit(p)
	// 		p = p.r
	// 	}
	// }
	// 后序
	var stat = map[*leaf]bool{}
	stat[nil] = true
	for stack.Len() > 0 {
		p = stack.Pop()
		//后序
		// if stat[p.l] && stat[p.r] {
		// 	visit(p)
		// 	stat[p] = true
		// } else {
		// 	stack.Push(p)
		// 	if p.r != nil {
		// 		stack.Push(p.r)
		// 	}
		// 	if p.l != nil {
		// 		stack.Push(p.l)
		// 	}
		// }
		// 中序
		if stat[p.l] {
			visit(p)
			stat[p] = true
			if p.r != nil {
				stack.Push(p.r)
			}
		} else {
			stack.Push(p)
			if p.l != nil {
				stack.Push(p.l)
			}
		}
	}
}

func main() {
	var a = &leaf{`a`, nil, nil}
	var b = &leaf{`b`, nil, nil}
	var c = &leaf{`c`, a, b}
	var d = &leaf{`d`, nil, c}
	var e = &leaf{`e`, d, nil}
	traverseTree(e)
}

package main

import "fmt"

type stack struct{ s []*node }

func (s *stack) Top() *node   { return s.s[len(s.s)-1] }
func (s *stack) Pop()         { s.s = s.s[:len(s.s)-1] }
func (s *stack) Push(v *node) { s.s = append(s.s, v) }
func (s *stack) Empty() bool  { return len(s.s) == 0 }

type node struct {
	v int
	l *node
	r *node
}

func (n *node) String() string {
	if n == nil {
		return "x"
	}
	return fmt.Sprint(n.v, n.l, n.r)
}

func findLCA(n *node, a, b int) *node {
	for n.v != a && n.v != b {
		if (a < n.v && b > n.v) || (a > n.v && b < n.v) {
			return n
		}
		if a < n.v {
			n = n.l
		} else {
			n = n.r
		}
	}
	return n
}

func makePreOrderTree(vals ...int) *node {
	root := &node{v: vals[0]}
	s := stack{}
	s.Push(root)
	for _, v := range vals[1:] {
		curr := s.Top()
		if v < curr.v {
			curr.l = &node{v: v}
			s.Push(curr.l)
		} else {
			for !s.Empty() && v > s.Top().v {
				curr = s.Top()
				s.Pop()
			}
			curr.r = &node{v: v}
			s.Push(curr.r)
		}
	}
	return root
}

func main() {
	var N int
	fmt.Scan(&N)

	arr := make([]int, N)
	for i := range arr {
		fmt.Scan(&arr[i])
	}

	tree := makePreOrderTree(arr...)
	// fmt.Println("tree: ", tree)

	for i := 0; i < N; i++ {
		var a, b int
		fmt.Scan(&a, &b)
		n := findLCA(tree, a, b)

		fmt.Printf("%t %d\n", n.v == a, n.v)
	}
}

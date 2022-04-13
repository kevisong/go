package binarytree

type Node struct {
	value int
	left  *Node
	right *Node
}

func Insert(node *Node, value int) *Node {
	if node == nil {
		node = &Node{value: value}
	}
	if value < node.value {
		node.left = Insert(node.left, value)
	} else {
		node.right = Insert(node.right, value)
	}
	return node
}

func InOrderTraverseRecurive(node *Node) []int {
	if node == nil {
		return []int{}
	}
	arr := make([]int, 0)
	arr = append(arr, InOrderTraverseRecurive(node.left)...)
	arr = append(arr, node.value)
	arr = append(arr, InOrderTraverseRecurive(node.right)...)
	return arr
}

func InOrderTraverseIterative(node *Node) []int {
	arr := make([]int, 0)
	stack := make([]*Node, 0)
	for len(stack) > 0 || node != nil {
		for node != nil {
			stack = append(stack, node) // push
			node = node.left            // as long as there is a left node
		}
		last := stack[len(stack)-1] // pop
		stack = stack[:len(stack)-1]
		arr = append(arr, last.value) // get the result
		node = last.right             // as long as there is a right node
	}
	return arr
}

func PreOrderTraverseRecurive(node *Node) []int {
	if node == nil {
		return []int{}
	}
	arr := make([]int, 0)
	arr = append(arr, node.value)
	arr = append(arr, PreOrderTraverseRecurive(node.left)...)
	arr = append(arr, PreOrderTraverseRecurive(node.right)...)
	return arr
}

func PreOrderTraverseIterative(node *Node) []int {
	arr := make([]int, 0)
	stack := make([]*Node, 0)
	for len(stack) > 0 || node != nil {
		for node != nil {
			arr = append(arr, node.value)
			stack = append(stack, node)
			node = node.left
		}
		last := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		node = last.right
	}
	return arr
}

func PostOrderTraverseRecurive(node *Node) []int {
	if node == nil {
		return []int{}
	}
	arr := make([]int, 0)
	arr = append(arr, PostOrderTraverseRecurive(node.left)...)
	arr = append(arr, PostOrderTraverseRecurive(node.right)...)
	arr = append(arr, node.value)
	return arr
}

func InOrderTraverseIterativeClean(node *Node) []int {
	// TODO: impl
	arr := make([]int, 0)
	stack := make([]*Node, 0)
	for len(stack) > 0 {
		if node != nil {
			stack = append(stack, node)

		}
	}
	return arr
}

func PostOrderTraverseIterative(node *Node) []int {
	// TODO: impl
	arr := make([]int, 0)
	return arr
}

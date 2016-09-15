package main

type Node struct {
	Right *Node
	Left  *Node
}

func (tree *Node) Height (seed int) int {
	lSeed, rSeed := seed, seed

	if tree.Left != nil {
		lSeed = tree.Left.Height(seed + 1)
	}

	if tree.Right != nil {
		rSeed = tree.Right.Height(seed + 1)
	}

	if lSeed > rSeed {
		return lSeed
	}
	return rSeed
}

func main() {

	// example tree
	var tree = Node{
		Left: &Node{
			Left: &Node{
				Right: &Node{},
			},
		},
		Right: &Node{
			Left: &Node{
				Right: &Node{},
				Left: &Node{
					Left: &Node{
						Right: &Node{},
					},
				},
			},
		},
	}

	println(tree.Height(0))
}
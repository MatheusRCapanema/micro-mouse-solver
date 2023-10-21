package solvers

import "micro-mouse-solver/pkg/api"

type Node struct {
	Pos    int
	Parent *Node
}

func constructPath(n *Node) []int {
	var path []int
	for n != nil {
		path = append([]int{n.Pos}, path...)
		n = n.Parent
	}
	return path
}

func BFS(client *api.Client, mazeID, labirintoName string, start int) []int {
	visited := make(map[int]bool)
	queue := []*Node{{Pos: start, Parent: nil}}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		respMove, err := client.Move(mazeID, labirintoName, current.Pos)
		if err != nil {
			return []int{}
		}

		if respMove.Final {
			return constructPath(current)
		}

		for _, nextPos := range respMove.Movimentos {
			if !visited[nextPos] {
				visited[nextPos] = true
				queue = append(queue, &Node{Pos: nextPos, Parent: current})
			}
		}
	}

	return []int{}
}

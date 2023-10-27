package solvers

import (
	"fmt"
	"micro-mouse-solver/pkg/api"
	"micro-mouse-solver/pkg/models"
)

type Node struct {
	Pos    int
	Parent *Node
	Final  bool
}

func constructPath(n *Node) []int {
	var path []int
	for n != nil {
		path = append([]int{n.Pos}, path...)
		n = n.Parent
	}
	return path
}

func BFS(client *api.Client, mazeID, labirintoName string, start int, respStart models.StartResponse) []int {
	visited := make(map[int]bool)

	// Use a resposta inicial para povoar a queue com os primeiros movimentos válidos.
	queue := make([]*Node, 0)
	for _, move := range respStart.Movimentos {
		queue = append(queue, &Node{Pos: move, Parent: &Node{Pos: start, Parent: nil}})
	}

	visited[start] = true

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current.Final {
			return constructPath(current)
		}

		fmt.Println("Explorando posição:", current.Pos) // Log da posição atual

		respMove, err := client.Move(mazeID, labirintoName, current.Pos)
		if err != nil {
			fmt.Println("Erro ao mover para a posição", current.Pos, ":", err) // Log de erro
			continue                                                           // Continue para o próximo movimento em caso de erro.
		}

		if respMove.Final {
			return constructPath(current)
		}

		fmt.Println("Movimentos possíveis a partir de", current.Pos, ":", respMove.Movimentos) // Log dos movimentos possíveis

		for _, nextPos := range respMove.Movimentos {
			if !visited[nextPos] {
				// Verifique se este movimento é o final
				testMoveResp, err := client.Move(mazeID, labirintoName, nextPos)
				if err == nil && testMoveResp.Final {
					// Se for o final, retorne o caminho até este ponto
					return constructPath(&Node{Pos: nextPos, Parent: current, Final: true})
				}

				visited[nextPos] = true
				queue = append(queue, &Node{Pos: nextPos, Parent: current})
			}
		}
	}

	// Se a função saiu do loop sem retornar um caminho, isso significa que não encontrou uma solução.
	fmt.Println("Não foi encontrado um caminho válido.") // Log de mensagem de erro
	return []int{}
}

func DFS(client *api.Client, mazeID, labirintoName string, currentNode *Node, visited map[int]bool) []int {

	fmt.Println("Explorando posição:", currentNode.Pos) // Log da posição atual

	respMove, err := client.Move(mazeID, labirintoName, currentNode.Pos)
	if err != nil {
		fmt.Println("Erro ao mover para a posição", currentNode.Pos, ":", err) // Log de erro
		return nil
	}

	if respMove.Final {
		return constructPath(currentNode)
	}

	fmt.Println("Movimentos possíveis a partir de", currentNode.Pos, ":", respMove.Movimentos) // Log dos movimentos possíveis

	for _, nextPos := range respMove.Movimentos {
		if !visited[nextPos] {
			visited[nextPos] = true
			result := DFS(client, mazeID, labirintoName, &Node{Pos: nextPos, Parent: currentNode}, visited)
			if result != nil { // Se encontrarmos um caminho válido, retornamos imediatamente
				return result
			}
		}
	}

	return nil
}

func DFSStart(client *api.Client, mazeID, labirintoName string, start int, respStart models.StartResponse) []int {
	visited := make(map[int]bool)
	visited[start] = true

	for _, move := range respStart.Movimentos {
		result := DFS(client, mazeID, labirintoName, &Node{Pos: move, Parent: &Node{Pos: start, Parent: nil}}, visited)
		if result != nil { // Se encontrarmos um caminho válido, retornamos imediatamente
			return result
		}
	}

	// Se chegarmos aqui, não encontramos um caminho válido.
	fmt.Println("Não foi encontrado um caminho válido.") // Log de mensagem de erro
	return []int{}
}

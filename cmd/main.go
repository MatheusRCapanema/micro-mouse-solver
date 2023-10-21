package cmd

import (
	"fmt"
	"micro-mouse-solver/pkg/api"
	"micro-mouse-solver/pkg/solvers"
)

func main() {
	// Suposições iniciais:
	// - 'mazeID' é uma string constante representando o ID do labirinto.
	// - 'labirintoName' é uma string constante representando o nome do labirinto.
	const mazeID = "exampleID"
	const labirintoName = "sample-maze"

	// 1. Iniciar o labirinto
	client := api.NewClient()
	respStart, err := client.StartMaze(mazeID, labirintoName)
	if err != nil {
		fmt.Println("Erro ao iniciar o labirinto:", err)
		return
	}

	// Capturando a posição inicial da resposta
	start := respStart.PosAtual

	//// 2. Obter a representação do labirinto da API
	//maze, err := client.GetMazeMap() // Esta chamada ainda é necessária?
	//if err != nil {
	//	fmt.Println("Erro ao obter o labirinto:", err)
	//	return
	//}

	// 3. Resolver o labirinto usando BFS
	// Suponho que sua função BFS agora só precisa da posição inicial e do próprio labirinto
	path := solvers.BFS(client, mazeID, labirintoName, start)

	// 4. Movimentar-se pelo labirinto usando o caminho encontrado
	for _, move := range path {
		respMove, err := client.Move(mazeID, labirintoName, move)
		if err != nil {
			fmt.Println("Erro ao movimentar:", err)
			return
		}

		// Verificar se respMove indica que o fim foi alcançado e, nesse caso, interromper
		if respMove.Final {
			fmt.Println("Chegamos ao final do labirinto!")
			break
		}
	}

	// 5. Opcionalmente, fazer uma requisição para /validar_caminho para validar o caminho encontrado
	respValidate, err := client.ValidatePath(mazeID, labirintoName, path)
	if err != nil {
		fmt.Println("Erro ao validar o caminho:", err)
		return
	}
	if respValidate.CaminhoValido {
		fmt.Println("O caminho é válido!")
	} else {
		fmt.Println("O caminho não é válido.")
	}
}

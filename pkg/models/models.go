package models

type MazeMap map[int][]int

type StartRequest struct {
	ID        string `json:"id"`
	Labirinto string `json:"labirinto"`
}

type StartResponse struct {
	PosAtual   int   `json:"pos_atual"`
	Inicio     bool  `json:"inicio"`
	Final      bool  `json:"final"`
	Movimentos []int `json:"movimentos"`
}

type MoveRequest struct {
	ID          string `json:"id"`
	Labirinto   string `json:"labirinto"`
	NovaPosicao int    `json:"nova_posicao"`
}

type MoveResponse struct {
	PosAtual   int   `json:"pos_atual"`
	Inicio     bool  `json:"inicio"`
	Final      bool  `json:"final"`
	Movimentos []int `json:"movimentos"`
}

type ValidateRequest struct {
	ID              string `json:"id"`
	Labirinto       string `json:"labirinto"`
	TodosMovimentos []int  `json:"todos_movimentos"`
}

type ValidateResponse struct {
	CaminhoValido        bool `json:"caminho_valido"`
	QuantidadeMovimentos int  `json:"quantidade_movimentos"`
}

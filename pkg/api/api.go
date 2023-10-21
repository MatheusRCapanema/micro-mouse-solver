package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"micro-mouse-solver/pkg/models"
	"net/http"
)

const baseURL = "https://gtm.delary.dev"

type Client struct {
	httpClient *http.Client
}

func NewClient() *Client {
	return &Client{
		httpClient: &http.Client{},
	}
}

// Endpoint /iniciar
func (c *Client) StartMaze(id, labirinto string) (response models.StartResponse, err error) {
	body := models.StartRequest{
		ID:        id,
		Labirinto: labirinto,
	}

	payload, _ := json.Marshal(body)
	resp, err := c.httpClient.Post(baseURL+"/iniciar", "application/json", bytes.NewBuffer(payload))
	if err != nil {
		return response, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return response, errors.New("failed to start the maze")
	}

	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return models.StartResponse{}, err
	}
	return response, nil
}

// Endpoint /movimentar
func (c *Client) Move(id, labirinto string, novaPosicao int) (response models.MoveResponse, err error) {
	body := models.MoveRequest{
		ID:          id,
		Labirinto:   labirinto,
		NovaPosicao: novaPosicao,
	}

	payload, _ := json.Marshal(body)
	resp, err := c.httpClient.Post(baseURL+"/movimentar", "application/json", bytes.NewBuffer(payload))
	if err != nil {
		return response, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return response, errors.New("failed to move in the maze")
	}

	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return models.MoveResponse{}, err
	}
	return response, nil
}

// Endpoint /validar_caminho
func (c *Client) ValidatePath(id, labirinto string, todosMovimentos []int) (response models.ValidateResponse, err error) {
	body := models.ValidateRequest{
		ID:              id,
		Labirinto:       labirinto,
		TodosMovimentos: todosMovimentos,
	}

	payload, _ := json.Marshal(body)
	resp, err := c.httpClient.Post(baseURL+"/validar_caminho", "application/json", bytes.NewBuffer(payload))
	if err != nil {
		return response, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return response, errors.New("failed to validate the path")
	}

	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return models.ValidateResponse{}, err
	}
	return response, nil
}

// Endpoint /labirintos
func (c *Client) ListMazes() ([]string, error) {
	resp, err := c.httpClient.Get(baseURL + "/labirintos")
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	var mazes []string
	err = json.NewDecoder(resp.Body).Decode(&mazes)
	if err != nil {
		return nil, err
	}

	return mazes, nil
}

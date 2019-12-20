package utils

import (
	"encoding/json"
	"math/rand"
	"net/http"
	_ "strconv"
	"time"
)


type TextResponse struct {
	Status    int    `json:"status"`
	ImagePath string `json:"image_path"`
}

type TaskIdResponse struct {
   TaskID string `json:"task_id"`
}

type Error struct {
	Status     int           `json:"status"`
	Message    interface{}   `json:"message"`
}

type Success struct {
	text TextResponse
	taskId TaskIdResponse
}

func GenerateRandomWord(n int) string {
	var wordRunes = []rune("abc123defg456hijkl789mno101112pqrs131415tuvwxyz") //ABCDEFGHIJKLMNOPQRSTUVWXYZ
	b := make([]rune, n)
	for i := range b {
		rand.Seed(time.Now().UnixNano())
		b[i] = wordRunes[rand.Intn(len(wordRunes))]
	}

	return string(b)
}


func (resp Error) RespondWithErrorJSON(w http.ResponseWriter, status int, error interface{}) {
	w.Header().Set("Content-Type", "application/json")

	resp.Status = status
	resp.Message = error

	json.NewEncoder(w).Encode(resp)
}

func (resp Success) RespondWithSuccessTextJSON(w http.ResponseWriter, text string) {
	w.Header().Set("Content-Type", "application/json")
	var response TextResponse

	response.Status = 200
	response.ImagePath = text

	resp.text = response

	json.NewEncoder(w).Encode(resp.text)
}

func (resp Success) RespondWithSuccessTaskIdJSON(w http.ResponseWriter, taskId string) {
	w.Header().Set("Content-Type", "application/json")
	var response TaskIdResponse
	response.TaskID = taskId

	resp.taskId = response

	json.NewEncoder(w).Encode(resp.taskId)
}

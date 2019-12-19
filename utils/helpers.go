package utils

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	_ "strconv"
	"time"

	"github.com/davecgh/go-spew/spew"
)


type TextResponse struct {
	Text string `json:"text"`
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

//os and file upload
func CreateDirectory(foldername string) (string, error) {
	_, err := os.Stat(foldername)
	if os.IsNotExist(err) {
		errDir := os.MkdirAll(foldername, 0755)
		if errDir != nil {
			return "", err
		}
	}

	return foldername, nil
}

func SpewResultForDebugging(description string, v interface{}) {
	fmt.Println("**** Start Result ******")
	fmt.Println(description)
	spew.Dump(v)
	fmt.Println("**** End Result ******")
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
	response.Text = text

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

func EncodeStringToBase64(imageName string) string {
	return  base64.StdEncoding.EncodeToString([]byte(imageName))
}

func DecodeStringToBase64(taskId string) (string, error) {
	s, err :=  base64.StdEncoding.DecodeString(taskId)
	if err != nil {
		return "", err
	}
	return string(s), nil
}
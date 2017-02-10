package apiaigo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type APIAI struct {
	Version   string
	Language  string
	AuthToken string
	SessionID string
}

const apiEndpoint string = "https://api.api.ai/v1/%s?v=%s"

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func (api APIAI) SendText(text string) ResponseStruct {
	query := QueryStruct{
		Query:     text,
		Language:  api.Language,
		SessionID: api.SessionID,
	}
	response := Response(query, api)
	return response
}

func (api APIAI) TTS(text string, filepath string) {
	req, err := http.NewRequest("GET", fmt.Sprintf(apiEndpoint, "tts", api.Version), nil)
	checkError(err)
	req.Header.Add("Authorization", "Bearer "+api.AuthToken)
	req.Header.Add("Accept-Language", "en-US")
	client := http.Client{}
	resp, err := client.Do(req)
	checkError(err)
	defer resp.Body.Close()
	file, err := os.Open(filepath)
	checkError(err)
	io.Copy(file, resp.Body)
}

func Response(query QueryStruct, api APIAI) ResponseStruct {
	result := ResponseStruct{}
	url := fmt.Sprintf(apiEndpoint, "query", api.Version)
	log.Println(url)
	jsonQuery, err := json.Marshal(&query)
	checkError(err)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonQuery))
	checkError(err)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+api.AuthToken)
	client := &http.Client{}
	resp, err := client.Do(req)
	checkError(err)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	checkError(err)
	log.Println(string(body))
	err = json.Unmarshal(body, &result)
	checkError(err)
	return result
}
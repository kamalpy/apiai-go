package apiaigo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

// APIAI Base APIAI struct. All fields are required.
type APIAI struct {
	Version   string
	Language  string
	AuthToken string
	SessionID string
}

const apiEndpoint string = "https://api.api.ai/v1/%s?v=%s"

// SendText is useful to send a simple query to API.ai
func (api APIAI) SendText(text string) (ResponseStruct, error) {
	query := QueryStruct{
		Query:     text,
		Language:  api.Language,
		SessionID: api.SessionID,
	}
	return Response(query, api)
}

// TTS sends query to API.ai and saves the speech response to specified filepath in wav format
func (api APIAI) TTS(text string, filepath string) error {
	req, err := http.NewRequest("GET", fmt.Sprintf(apiEndpoint, "tts", api.Version), nil)
	if err != nil {
		return err
	}
	req.Header.Add("Authorization", "Bearer "+api.AuthToken)
	req.Header.Add("Accept-Language", "en-US")
	client := http.Client{}
	q := req.URL.Query()
	q.Add("text", text)
	req.URL.RawQuery = q.Encode()
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	file, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer file.Close()
	io.Copy(file, resp.Body)
	return nil
}

func Response(query QueryStruct, api APIAI) (ResponseStruct, error) {
	result := ResponseStruct{}
	url := fmt.Sprintf(apiEndpoint, "query", api.Version)
	jsonQuery, err := json.Marshal(&query)
	if err != nil {
		return ResponseStruct{}, err
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonQuery))
	if err != nil {
		return ResponseStruct{}, err
	}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("Authorization", "Bearer "+api.AuthToken)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return ResponseStruct{}, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ResponseStruct{}, err
	}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return ResponseStruct{}, err
	}
	return result, nil
}

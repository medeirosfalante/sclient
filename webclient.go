package sclient

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

//WebClient object request client
type WebClient struct {
	Uri string
}

// QueryString is a object a request client
type QueryString struct {
}

func request(method string, uri string, command *string, payloadJSON []byte, queryStringByte []byte) (*Response, error) {
	log.Printf(uri)
	req, err := http.NewRequest(method, uri, bytes.NewBuffer(payloadJSON))
	if queryStringByte != nil {
		q := req.URL.Query()
		queryStringParse, _ := ParseJSON(queryStringByte)
		children, _ := queryStringParse.ChildrenMap()
		for key, child := range children {
			q.Add(key, child.Data().(string))
		}
		req.URL.RawQuery = q.Encode()
	}

	req.Header.Add("Content-Type", "application/json")
	if os.Getenv("USER_API_TOKEN") != "" {
		req.Header.Add("Authorization", fmt.Sprintf("ApiToken %s", os.Getenv("USER_API_TOKEN")))
	}
	if err != nil {
		log.Printf("erro create request: %s", err)
		return nil, err
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Erro fetch request: %s ", err)
		return nil, err
	}
	defer resp.Body.Close()
	bodyString, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("erro parse bodyString response date: %s", err)
		return nil, err
	}
	jsonParsed, err := ParseJSON(bodyString)
	if err != nil {
		log.Println("erro parse Json body", err)
	}

	return jsonParsed, nil

}

// GET func in WebClient
func (web WebClient) GET(command string, queryString []byte) (*Response, error) {
	response, err := request("GET", web.Uri, &command, nil, queryString)
	return response, err
}

// POST func in WebClient
func (web WebClient) POST(command string, payload []byte) (*Response, error) {
	response, err := request("POST", web.Uri, &command, payload, nil)
	return response, err
}

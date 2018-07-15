package sclient_test

import (
	"encoding/json"
	"sclient"
	"testing"
)

type TestPost struct {
	title  string `json:"title"`
	body   string `json:"body"`
	userId int64  `json:"userId"`
}

func TestGET(t *testing.T) {
	webclient := sclient.WebClient{
		Uri: "https://jsonplaceholder.typicode.com/",
	}
	_, err := webclient.GET("post", nil)
	if err != nil {
		t.Errorf("erro test GET: %s", err)
	}

}

func TestPOST(t *testing.T) {
	webclient := sclient.WebClient{
		Uri: "https://jsonplaceholder.typicode.com",
	}
	testPost := TestPost{
		title:  "foo",
		body:   "bar",
		userId: 1,
	}
	testPostByte, _ := json.Marshal(&testPost)
	_, err := webclient.POST("posts", testPostByte)
	if err != nil {
		t.Errorf("erro test POST: %s", err)
	}

}

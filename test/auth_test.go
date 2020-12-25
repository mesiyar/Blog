package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"testing"
)

func TestAuth(t *testing.T)  {

	req, err :=http.NewRequest("POST", "http://localhost:8000/auth", strings.NewReader("username=eddie13&password=123456"))

	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	byts, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	t.Log(string(byts))
}
package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
)

func PostJsonAsText(url string, jsonAsText string) {
	var jsonStr = []byte(jsonAsText)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	panicIfError(err, "couldn't post")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	client.Do(req)
}

func GetHttpGetAsText(url string) (responsebody string) {
	response, err := http.Get(url)
	panicIfError(err, "couldn't make request")
	defer response.Body.Close()
	bodyAsBytes, err := ioutil.ReadAll(response.Body)
	panicIfError(err, "couldn't read body")
	return string(bodyAsBytes)
}

func panicIfError(err error, msg string) {
	if err != nil {
		errorText := msg + err.Error()
		log.Fatal(errorText)
		panic(errorText)
	}
}

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	username := "d6b41e685927de8f"
	password := "cI9Bqta2X7tOZQNqOyEwbK729BrQohaK0ijqxu9CY5LdRC"

	url := "http://localhost:18083/api/v5/nodes"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}
	req.SetBasicAuth(username, password)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(resp.Body)
	if err != nil {
		panic(err)
	}

	var data interface{}
	json.Unmarshal(buf.Bytes(), &data)
	fmt.Println(data)
}

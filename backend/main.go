package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Response struct {
	Data string `json:"data"`
}

func handlePost(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 Not Found", http.StatusNotFound)
		return
	}

	switch r.Method {
	case "POST":
		requestData, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "400 Bad Request", http.StatusBadRequest)
			return
		}
		requestString := string(requestData)
		fmt.Printf("Received: %s\n", requestString)
		Data := fmt.Sprintf("Modified: %s", requestString)
		fmt.Printf("Data: %s\n", Data)
		response := Response{Data}
		responseJson, err := json.Marshal(response)
		if err != nil {
			http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
			return
		}
		responseJsonString := string(responseJson)
		fmt.Printf("Responding: %s\n", responseJsonString)
		w.Header().Add("Access-Control-Allow-Origin", "*")
		w.Write(responseJson)
	default:
		http.Error(w, "405 Method Not Allowed.", http.StatusMethodNotAllowed)
	}
}

func main() {
	http.HandleFunc("/", handlePost)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

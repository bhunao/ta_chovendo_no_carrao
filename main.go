package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
)

var tpl = template.Must(template.ParseFiles("index.html"))

type ViewData struct {
	Message string
}

func parseJSON(jsonString string) (map[string]interface{}, error) {
	// Define a map to store the JSON data
	var data map[string]interface{}

	// Unmarshal the JSON string into the map
	err := json.Unmarshal([]byte(jsonString), &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func get_weather(url string) (string, error) {
	// Make GET request
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error making GET request:", err)
		return "", err
	}
	defer resp.Body.Close()

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return "", err
	}

	// Print response body
	coiso, err := parseJSON(string(body))
	if err != nil {
		fmt.Println("Error parsing json body:", err)
		return "", err
	}
	fmt.Println(coiso)
	fmt.Println(string(body))
	return string(body), nil
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	data := ViewData{Message: "TA SOL"}
	tpl.Execute(w, data)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	mux := http.NewServeMux()
	mux.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("js"))))
	mux.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))
	mux.Handle("/assets/img/", http.StripPrefix("/assets/img/", http.FileServer(http.Dir("assets/img"))))


	mux.HandleFunc("/", indexHandler)
	http.ListenAndServe(":"+port, mux)
}

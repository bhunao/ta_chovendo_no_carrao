package main

import (
	"encoding/json"
	"fmt"
	"io"
	"html/template"
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

func get_weather() (string, error) {
    // Define the URL
    url := "http://dataservice.accuweather.com/currentconditions/v1/36302?apikey=RCwn2dVmR7RtmGGeRKfjIrzAAozQp70w&language=pt-br"

    // Make GET request
    resp, err := http.Get(url)
    if err != nil {
        fmt.Println("Error making GET request:", err)
        return "", err
    }
    defer resp.Body.Close()

    // Read response body
    body, err  := io.ReadAll(resp.Body)
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
	err, d := get_weather()
	// data := "talvez"
	data := map[string]string{"algo": "talvez"}
	tpl.Execute(w, data)
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/", indexHandler)
	http.ListenAndServe(":"+port, mux)
}

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
)

// interface for JSON data
type JSONData interface{}

var prettyFlag = flag.Bool("pretty", true, "Enable/disable pretty JSON formatting")
var headersFlag = flag.Bool("headers", true, "Enable/disable technical information of request")
var portFlag = flag.Int("port", 8080, "Set listening port of webhook catcher")
var feedbackFlag = flag.Bool("answer", true, "Enable/disable answer from service (Status: OK, "+
	"Data received)")

func main() {
	flag.Parse()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error reading request body", http.StatusInternalServerError)
			return
		}

		var requestData JSONData
		err = json.Unmarshal(body, &requestData)
		if err != nil {
			http.Error(w, fmt.Sprintf("Error parsing JSON: %v", err), http.StatusBadRequest)
			return
		}
		// parse JSON to empty interface
		var formattedJSON []byte
		if *prettyFlag {
			formattedJSON, err = json.MarshalIndent(requestData, "", "  ")
		} else {
			formattedJSON, err = json.Marshal(requestData)
		}
		if err != nil {
			http.Error(w, fmt.Sprintf("Error formatting JSON: %v", err), http.StatusInternalServerError)
			return
		}

		// Send data to console
		if *headersFlag {
			fmt.Print("\n\nStart technical info\n")
			fmt.Print("====================\n")
			fmt.Println("Received request:")
			fmt.Printf("  Method: %s\n", r.Method)
			fmt.Printf("  URL: %s\n", r.URL)
			fmt.Println("  Headers:")
			for key, values := range r.Header {
				for _, value := range values {
					fmt.Printf("    %s: %s\n", key, value)
				}
			}
			fmt.Print("=====================\n")
			fmt.Print("End of technical info\n\n\n")
		}
		fmt.Printf("  Parsed JSON data:\n%s\n\n", formattedJSON)

		if *feedbackFlag {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Data received successfully\n"))
		}
	})

	// Run service
	address := fmt.Sprintf("localhost:%d", *portFlag)
	fmt.Printf("Server is listening on http://%s\n", address)
	fmt.Printf("Pretty JSON formatting: %t\n", *prettyFlag)
	fmt.Printf("Headers: %t\n", *headersFlag)
	fmt.Printf("Answer: %t\n\n", *feedbackFlag)
	err := http.ListenAndServe(address, nil)
	if err != nil {
		fmt.Printf("Error starting the server: %v\n", err)
	}
}

package main

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"strings"
)

// Run with
//		go run .
// Send request with:
//		curl -F 'file=@/path/matrix.csv' "localhost:8080/echo"

func catchError(err error, w http.ResponseWriter) bool {
	if err != nil {
		w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
		return true
	}
	return false
}

func computeFile(routeName string, responseWriter http.ResponseWriter, request *http.Request){

	switch routeName {
		case 'echo':
			
		case 'invert':
			statement(s)
		case 'flatten':
			statement(s)
		case 'sum':
			statement(s)	
		case 'multiply':
			statement(s)
		
		default:
			statement(s)    
	}
}

func main() {

	http.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {

		file, _, err := r.FormFile("file")
		if catchError(err, w) == true {
			return
		}

		defer file.Close()
		records, err := csv.NewReader(file).ReadAll()
		if catchError(err, w) == true {
			return
		}
		var response string
		for _, row := range records {
			response = fmt.Sprintf("%s%s\n", response, strings.Join(row, ","))
		}
		fmt.Fprint(w, response)
	})
	http.ListenAndServe(":8080", nil)
}

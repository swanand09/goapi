package main

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"os"
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

func processFile(routeName string, responseWriter http.ResponseWriter, request *http.Request) {

	//file, _, err := request.FormFile("file")
	file, err := os.Open("matrix.csv")
	if catchError(err, responseWriter) == true {
		return
	}

	defer file.Close()
	records, err := csv.NewReader(file).ReadAll()
	if catchError(err, responseWriter) == true {
		return
	}

	//count := 0
	var response string

	switch routeName {

	case "echo":

		for _, row := range records {
			response = fmt.Sprintf("%s%s\n", response, strings.Join(row, ","))
		}

	case "invert":

		strEle := ""
		count := 0
		for count < len(records) {

			for _, row := range records {

				for key, value := range row {

					if key == count {
						strEle += string(value)
					}
				}
			}
			count++

		}

		response = fmt.Sprintf("%s%s", response, strEle)

	case "flatten":

		for _, row := range records {
			response = fmt.Sprintf("%s%s", response, strings.Join(row, ","))
		}

	case "sum":

		totalSum := 0
		for _, row := range records {

			for val := range row {
				totalSum += int(val)
			}
		}

		response = fmt.Sprintf("%s%s\n", response, string(totalSum))

	case "multiply":

		multiply := 0
		for _, row := range records {

			for val := range row {
				multiply *= int(val)
			}
		}

		response = fmt.Sprintf("%s%s\n", response, string(multiply))

	default:

		for _, row := range records {
			response = fmt.Sprintf("%s%s\n", response, strings.Join(row, ","))
		}
	}

	fmt.Fprint(responseWriter, response)
}

func invert(records []string) string {

	strEle := ""
	count := 0
	for count < len(records) {

		for _, row := range records {

			for key, value := range row {

				if key == count {
					strEle += string(value)
				}
			}
		}
		count++

	}

	return strEle

}

func main() {

	http.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {

		processFile("echo", w, r)
	})

	http.HandleFunc("/invert", func(w http.ResponseWriter, r *http.Request) {

		processFile("invert", w, r)
	})

	http.HandleFunc("/flatten", func(w http.ResponseWriter, r *http.Request) {

		processFile("flatten", w, r)
	})

	http.HandleFunc("/sum", func(w http.ResponseWriter, r *http.Request) {

		processFile("sum", w, r)
	})

	http.HandleFunc("/multiply", func(w http.ResponseWriter, r *http.Request) {

		processFile("multiply", w, r)
	})

	http.ListenAndServe(":8080", nil)
}

package main

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

// Run with
//		go run .
// Send request with:
//		curl -F 'file=@matrix.csv' "localhost:8080/echo"

//catch error
func catchError(err error, w http.ResponseWriter) bool {
	if err != nil {
		w.Write([]byte(fmt.Sprintf("error %s", err.Error())))
		return true
	}
	return false
}

//execute command based from curl
//curl -F 'file=@matrix.csv' "localhost:8080/echo"
//curl -F 'file=@matrix.csv' "localhost:8080/invert"
//curl -F 'file=@matrix.csv' "localhost:8080/flatten"
func processFile(routeName string, responseWriter http.ResponseWriter, request *http.Request) {

	file, _, err := request.FormFile("file")
	//file, err := os.Open("matrix.csv")
	if catchError(err, responseWriter) {
		return
	}

	defer file.Close()
	records, err := csv.NewReader(file).ReadAll()
	if catchError(err, responseWriter) {
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

						if len(strEle)%(4+(count*6)) != 0 || len(strEle) == 0 {

							strEle += string(value) + ","
						} else {

							strEle += string(value) + "\n"
						}

					}
				}

			}
			count++
		}
		response = fmt.Sprintf("%s%s", response, strEle)

	case "flatten":

		strElem := ""
		for _, row := range records {
			strElem += strings.Join(row, ",") + ","

		}
		response = fmt.Sprintf("%s%s\n", response, strings.TrimRight(strElem, ","))

	case "sum":

		totalSum := 0
		for _, row := range records {

			for _, val := range row {
				tempInt, err := strconv.ParseInt(val, 10, 64)

				if err == nil {
					totalSum += int(tempInt)
				}
			}
		}

		response = fmt.Sprint(totalSum)
		response += "\n"

	case "multiply":

		multiply := 0
		for _, row := range records {

			for _, val := range row {

				tempInt, err := strconv.ParseInt(val, 10, 64)
				if err == nil && multiply != 0 {

					multiply *= int(tempInt)
				}
				if err == nil && multiply == 0 {

					multiply = int(tempInt)
				}

			}
		}
		response = fmt.Sprint(multiply)
		response += "\n"

	default:

		for _, row := range records {
			response = fmt.Sprintf("%s%s\n", response, strings.Join(row, ","))
		}
	}

	fmt.Fprint(responseWriter, response)
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

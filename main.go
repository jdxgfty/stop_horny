package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

var ALLOWED_HEADERS = []string{"GET", "HEAD"}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if !contains(ALLOWED_HEADERS, r.Method) {
			w.Header().Set("Allow", "GET, HEAD")
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		img, err := os.ReadFile("stop_horny.jpg")
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("an unknown error occured"))
			return
		}
		headers := map[string]string{
			"Content-Disposition": `inline; filename="stop_horny.jpg"`,
			"Content-Length":      fmt.Sprint(len(img)),
		}
		for k, v := range headers {
			w.Header().Set(k, v)
		}
		w.Write(img)
	})

	log.Println("Starting server")
	log.Fatalln(http.ListenAndServe(":8080", mux))
}

func contains(array []string, value string) bool {
	for _, v := range array {
		if v == value {
			return true
		}
	}
	return false
}

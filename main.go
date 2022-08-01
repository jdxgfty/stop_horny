package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" && r.Method != "HEAD" {
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

	log.Fatalln(http.ListenAndServe(":8080", mux))
}

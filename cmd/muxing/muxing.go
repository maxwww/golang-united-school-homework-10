package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

/**
Please note Start functions is a placeholder for you to start your own solution.
Feel free to drop gorilla.mux if you want and use any other solution available.

main function reads host/port from env just for an example, flavor it following your taste
*/

// Start /** Starts the web server listener on given host and port.
func Start(host string, port int) {
	router := mux.NewRouter()

	router.HandleFunc("/name/{name}", NameHandler).Methods(http.MethodGet)
	router.HandleFunc("/bad", BadHandler).Methods(http.MethodGet)
	router.HandleFunc("/data", DataHandler).Methods(http.MethodPost)
	router.HandleFunc("/headers", HeadersHandler).Methods(http.MethodPost)

	log.Println(fmt.Printf("Starting API server on %s:%d\n", host, port))
	if err := http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), router); err != nil {
		log.Fatal(err)
	}
}

func NameHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	nameValue := vars["name"]
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello, %s!", nameValue)
}

func BadHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
}

func DataHandler(w http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	r.Body.Close()
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "I got message:\n%s", string(data))
}

func HeadersHandler(w http.ResponseWriter, r *http.Request) {
	aHeader := r.Header.Get("a")
	bHeader := r.Header.Get("b")
	a, err := strconv.Atoi(aHeader)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	b, err := strconv.Atoi(bHeader)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	result := strconv.Itoa(a + b)
	w.Header().Set("a+b", result)
	log.Println(result)
	w.WriteHeader(http.StatusOK)

}

//main /** starts program, gets HOST:PORT param and calls Start func.
func main() {
	host := os.Getenv("HOST")
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		port = 8081
	}
	Start(host, port)
}

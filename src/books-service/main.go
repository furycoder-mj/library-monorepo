package main

import (
	"log"
	"net/http"
	"os"
)

func handleHello(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		w.Write([]byte("Hello Books!"))
		w.WriteHeader(http.StatusOK)
		return
	}
	w.WriteHeader(http.StatusBadRequest)
}
func main() {
	helloHandler := http.HandlerFunc(handleHello)
	http.Handle(os.Getenv("BasePath")+"/", helloHandler)
	err := http.ListenAndServe(":"+os.Getenv("Port"), nil)
	if err != nil {
		log.Fatal(err)
	}
}

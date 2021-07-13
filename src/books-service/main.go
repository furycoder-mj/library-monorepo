package main

import (
	"log"
	"net/http"
	"os"
	"sync"
)

func handleHello(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		w.Write([]byte("Hello Books!"))
		w.WriteHeader(http.StatusOK)
		return
	}
	w.WriteHeader(http.StatusBadRequest)
}
func startServer(wg *sync.WaitGroup) *http.Server {
	nonTlsSrv := &http.Server{Addr: ":" + os.Getenv("Port")}
	http.HandleFunc(os.Getenv("BasePath")+"/", handleHello)
	go func() {
		defer wg.Done() // let main know we are done cleaning up

		// always returns error. ErrServerClosed on graceful close
		if err := nonTlsSrv.ListenAndServe(); err != http.ErrServerClosed {
			// unexpected error. port in use?
			log.Fatalf("ListenAndServe(): %v", err)
		}
	}()
	return nonTlsSrv
}
func main() {
	httpServerExitDone := &sync.WaitGroup{}
	httpServerExitDone.Add(1)
	_ = startServer(httpServerExitDone)
	httpServerExitDone.Wait()
}

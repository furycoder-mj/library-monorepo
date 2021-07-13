package main

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"sync"
	"testing"
	"time"
)

func TestMain(m *testing.M) {
	// init server {setup code}
	httpServerExitDone := &sync.WaitGroup{}
	httpServerExitDone.Add(1)
	nonTlsServer := startServer(httpServerExitDone)
	time.Sleep(5 * time.Second)
	// init complete

	exitCode := m.Run()

	// shutdown server {teardown code}
	nonTlsServer.Shutdown(context.TODO())

	//exit
	os.Exit(exitCode)
}

func TestStartServer(t *testing.T) {
	newreq := func(method, url string, body io.Reader) *http.Request {
		r := httptest.NewRequest(method, url, body)
		return r
	}

	tests := []struct {
		name string
		r    *http.Request
	}{
		{name: "Pass", r: newreq("GET", "http://localhost:5000/books", nil)},
		{name: "Fail", r: newreq("POST", "http://localhost:5000/books", nil)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := httptest.NewRecorder()

			handleHello(res, tt.r)

			if res.Result().StatusCode != http.StatusOK {
				t.Errorf("Expected status code:%d, received status code:%d", http.StatusOK, res.Result().StatusCode)
			}
		})
	}

}

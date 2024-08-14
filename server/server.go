package server

import (
	"fmt"
	"net/http"
)

func StartServer(host, port string, ready chan<- bool) *http.Server {
	srv := &http.Server{Addr: fmt.Sprintf("%s:%s", host, port)}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	go func() {
		ready <- true
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("ListenAndServe(): %v", err)
		}
	}()

	return srv
}

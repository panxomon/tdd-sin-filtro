package server_test

import (
	"fmt"
	"net/http"
	"tdd/server"
	"testing"
)

func Test_Server(t *testing.T) {
	type test struct {
		name string
		host string
		port string
	}

	cases := []test{{
		name: "When Server Response OK",
		host: "127.0.0.1",
		port: "8080",
	}}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			ready := make(chan bool)

			// Inicia el servidor
			srv := server.StartServer(tc.host, tc.port, ready)

			// Espera hasta que el servidor esté listo
			<-ready

			// Ahora realiza la solicitud HTTP
			url := fmt.Sprintf("http://%s:%s/", tc.host, tc.port)
			resp, err := http.Get(url)
			if err != nil {
				t.Fatalf("Expected no error, got %v", err)
			}
			defer resp.Body.Close()

			if resp.StatusCode != http.StatusOK {
				t.Errorf("Expected status code %v, got %v", http.StatusOK, resp.StatusCode)
			}

			// Cierra el servidor
			if err := srv.Close(); err != nil {
				t.Fatalf("Server close failed: %v", err)
			}
		})
	}
}

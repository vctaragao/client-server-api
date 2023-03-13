package main

import (
	"fmt"
	"net/http"

	"github.com/vctaragao/client-server-api/server/service"
)

func main() {
	http.HandleFunc("/cotacao", service.GetDolarQuotation)

	fmt.Println("Server listening on: localhost:8080")
	http.ListenAndServe(":8080", nil)
}

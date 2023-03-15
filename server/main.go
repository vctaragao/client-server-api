package main

import (
	"fmt"
	"net/http"

	"github.com/vctaragao/client-server-api/server/application/service"
	"github.com/vctaragao/client-server-api/server/storage"
)

func main() {
	repo, err := storage.NewSqlite()
	if err != nil {
		fmt.Printf("unable to create repo: %v", err)
		return
	}

	getDolarQuotation := service.NewGetDolarQuotationService(repo)

	http.HandleFunc("/cotacao", getDolarQuotation.Execute)

	fmt.Println("Server listening on: localhost:8080")
	http.ListenAndServe(":8080", nil)
}

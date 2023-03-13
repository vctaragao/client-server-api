package service

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
)

func GetDolarQuotation(w http.ResponseWriter, r *http.Request) {
	resp, err := makeQuotationRequest()
	if err != nil {
		fmt.Println(err)
		return
	}

	body, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		fmt.Printf("Unable to read Body %v", err)
		return
	}

	_, err = w.Write(body)
	if err != nil {
		fmt.Printf("Unable to write to response writer %v", err)
		return
	}
}

func makeQuotationRequest() (*http.Response, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*300)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "https://economia.awesomeapi.com.br/json/last/USD-BRL", nil)
	if err != nil {
		return &http.Response{}, fmt.Errorf("unable to create request %v", err)

	}

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return &http.Response{}, fmt.Errorf("unable to make request %v", err)
	}

	return resp, nil
}

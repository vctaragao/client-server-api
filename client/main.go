package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

type QuotationDto struct {
	Dolar float64 `json:"dolar"`
}

func main() {
	resp, err := makeRequest()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	dto, err := parseResponse(resp)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	err = saveQuotation(dto)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
}

func makeRequest() (*http.Response, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*300)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, "GET", "http://localhost:8080/cotacao", nil)
	if err != nil {
		return &http.Response{}, fmt.Errorf("unable to create request: %v", err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return &http.Response{}, fmt.Errorf("unable to do request: %v", err)
	}

	return resp, nil
}

func parseResponse(resp *http.Response) (*QuotationDto, error) {
	if resp.StatusCode != http.StatusOK {
		return &QuotationDto{}, fmt.Errorf("an unespected error was received from the server")
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return &QuotationDto{}, fmt.Errorf("unable to read response body: %v", err)
	}

	fmt.Println(string(data))

	var dto QuotationDto
	err = json.Unmarshal(data, &dto)
	if err != nil {
		return &QuotationDto{}, fmt.Errorf("unable to parse response: %v", err)
	}

	return &dto, nil
}

func saveQuotation(dto *QuotationDto) error {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	err = os.WriteFile(dir+"/context.txt", []byte(fmt.Sprintf("Dolar: %.5f", dto.Dolar)), 0644)
	if err != nil {
		return fmt.Errorf("unable to write to context file: %v", err)
	}

	return nil
}

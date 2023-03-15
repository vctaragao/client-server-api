package service

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/vctaragao/client-server-api/server/application"
	"github.com/vctaragao/client-server-api/server/application/dto"
)

type GetDolarQuotationService struct {
	repo application.Repository
}

func NewGetDolarQuotationService(repo application.Repository) *GetDolarQuotationService {
	return &GetDolarQuotationService{
		repo: repo,
	}
}

func (s *GetDolarQuotationService) Execute(w http.ResponseWriter, r *http.Request) {
	quotationDto, err := makeQuotationRequest()
	if err != nil {
		fmt.Println(err)
		returnServerError(w)
		return
	}

	err = addQuotation(quotationDto, s.repo)
	if err != nil {
		fmt.Println(err)
		returnServerError(w)
		return
	}

	err = writeResponse(quotationDto, w)
	if err != nil {
		fmt.Println(err)
		returnServerError(w)
		return
	}
}

func makeQuotationRequest() (*dto.QuotationDto, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*200)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, "GET", "https://economia.awesomeapi.com.br/json/last/USD-BRL", nil)
	if err != nil {
		return &dto.QuotationDto{}, fmt.Errorf("unable to create request %v", err)
	}

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return &dto.QuotationDto{}, fmt.Errorf("unable to make request %v", err)
	}

	quotationDto, err := parseResponse(resp)
	if err != nil {
		return &dto.QuotationDto{}, fmt.Errorf("unable to parse response: %v", err)
	}

	return quotationDto, nil
}

func parseResponse(resp *http.Response) (*dto.QuotationDto, error) {
	body, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return &dto.QuotationDto{}, fmt.Errorf("unable to read body %v", err)
	}

	var quotationDto dto.QuotationDto
	err = json.Unmarshal(body, &quotationDto)
	if err != nil {
		return &dto.QuotationDto{}, fmt.Errorf("unable to parse body: %v", err)
	}

	return &quotationDto, nil
}

func addQuotation(dto *dto.QuotationDto, repo application.Repository) error {
	dolar, err := parseDolar(dto)
	if err != nil {
		return err
	}

	_, err = repo.AddQuotation(dolar)
	if err != nil {
		return fmt.Errorf("unable to add quotation: %v", err)
	}

	return nil
}

func writeResponse(quotationDto *dto.QuotationDto, w http.ResponseWriter) error {
	dolar, err := parseDolar(quotationDto)
	if err != nil {
		return err
	}

	dto := dto.NewResponseDto(dolar)

	response, err := json.Marshal(dto)
	if err != nil {
		return fmt.Errorf("unable to parse response %v", err)
	}

	_, err = w.Write(response)
	if err != nil {
		return fmt.Errorf("unable to write to response writer %v", err)
	}

	return nil
}

func returnServerError(w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte("unespected error"))
}

func parseDolar(dto *dto.QuotationDto) (float64, error) {
	dolar, err := strconv.ParseFloat(dto.Usdbrl.Bid, 64)
	if err != nil {
		return 0.0, fmt.Errorf("unable for format dolar: %v", err)
	}

	return dolar, nil
}

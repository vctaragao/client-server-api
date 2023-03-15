package dto

type ResponseDto struct {
	Dolar float64 `json:"dolar"`
}

func NewResponseDto(dolar float64) *ResponseDto {
	return &ResponseDto{
		Dolar: dolar,
	}
}

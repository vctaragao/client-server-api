package application

type Repository interface {
	AddQuotation(value float64) (int64, error)
}

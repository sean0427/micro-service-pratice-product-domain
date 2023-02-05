package model

type GetProductsParams struct {
	ManufacturerID *string
	Name           *string
}

func StringToPointer(s string) *string {
	return &s
}

package entity

type ProductRequest struct {
	Name  string `json:"name"`
	Price int    `json:"price"`
}

type CustomerRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type OrderRequest struct {
	Product  ProductRequest  `json:"product"`
	Customer CustomerRequest `json:"customer"`
}

type Invoice struct {
	ID         string `json:"id"`
	InvoiceUrl string `json:"invoice_url"`
}

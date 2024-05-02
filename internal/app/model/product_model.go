package model

type Product struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Name     string `json:"name"`
	Price    int    `json:"price"`
	Quantity int    `json:"quantity"`
}
type ProductResponse struct {
    Products []Product `json:"products"`
    Total    int          `json:"total"`
    Message  string       `json:"message"`
    Status   int          `json:"status"`
    Success  bool         `json:"success"`
}